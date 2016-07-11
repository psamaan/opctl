package main

//go:generate counterfeiter -o ./fakeRunOpUseCase.go --fake-name fakeRunOpUseCase ./ runOpUseCase

import (
  "fmt"
  "github.com/opctl/engine-sdk-golang"
  "github.com/opctl/engine-sdk-golang/models"
  engineModels "github.com/opctl/engine/core/models"
  "net/url"
  "os"
  "syscall"
  "os/signal"
  "strings"
  "github.com/opspec-io/sdk-golang"
  "path"
  "errors"
)

type runOpUseCase interface {
  Execute(
  args []string,
  name string,
  )
}

func newRunOpUseCase(
exiter exiter,
opspecSdk opspec.Sdk,
opctlEngineSdk opctlengine.Sdk,
workDirPathGetter workDirPathGetter,
) runOpUseCase {
  return _runOpUseCase{
    exiter:exiter,
    opspecSdk:opspecSdk,
    opctlEngineSdk:opctlEngineSdk,
    workDirPathGetter:workDirPathGetter,
  }
}

type _runOpUseCase struct {
  exiter            exiter
  opspecSdk         opspec.Sdk
  opctlEngineSdk    opctlengine.Sdk
  workDirPathGetter workDirPathGetter
}

func (this _runOpUseCase) Execute(
args []string,
name string,
) {
  providedArgMap := make(map[string]string)
  for _, arg := range args {

    argParts := strings.Split(arg, "=")

    argName := argParts[0]
    var argValue string
    if (len(argParts) > 1) {
      argValue = argParts[1]
    } else {
      argValue = os.Getenv(arg)
    }

    providedArgMap[argName] = argValue
  }

  opPath := path.Join(this.workDirPathGetter.Get(), ".opspec", name)
  opView, err := this.opspecSdk.GetOp(opPath)
  if (nil != err) {
    this.exiter.Exit(1)
    return
  }

  // [only] pass defined params
  argsMap := make(map[string]string)
  for _, opParam := range opView.Params {
    if providedArg, ok := providedArgMap[opParam.Name]; ok {
      argsMap[opParam.Name] = providedArg
    } else {
      argsMap[opParam.Name] = os.Getenv(opParam.Name)
    }
  }

  // init signal channel
  intSignalsReceived := 0
  signalChannel := make(chan os.Signal, 1)
  signal.Notify(
    signalChannel,
    syscall.SIGINT, //handle SIGINTs
  )

  // init event channel
  eventChannel, err := this.opctlEngineSdk.GetEventStream()
  if (nil != err) {
    this.exiter.Exit(1)
    return
  }

  rootOpRunId, correlationId, err := this.opctlEngineSdk.RunOp(
    *models.NewRunOpReq(
      argsMap,
      &url.URL{Path:opPath},
    ),
  )
  if (nil != err) {
    this.exiter.Exit(1)
    return
  }

  for {
    select {

    case <-signalChannel:
      if (intSignalsReceived == 0) {

        this.opctlEngineSdk.KillOpRun(
          *models.NewKillOpRunReq(
            rootOpRunId,
          ),
        )

        intSignalsReceived++
        fmt.Println()
        fmt.Println("Gracefully stopping... (press Ctrl+C again to force)")
      } else {
        this.exiter.Exit(130)
        return
      }

    case event, isEventChannelOpen := <-eventChannel:
      if (!isEventChannelOpen) {
        err = errors.New("event channel closed unexpectedly")
        this.exiter.Exit(1)
        return
      }

      switch event := event.(type) {
      case models.LogEntryEmittedEvent:
        // @TODO: this doesn't catch log entries for the same tree but triggered from different actions (such as kills) see https://github.com/opctl/engine/issues/2
        if (event.CorrelationId() == correlationId) {
          fmt.Printf(
            "%v \n",
            event.LogEntryMsg(),
          )
        }
      case models.OpRunStartedEvent:
        if (event.RootOpRunId() == rootOpRunId) {
          opUrl := event.OpRunOpUrl()
          fmt.Printf(
            "OpRunStarted: Id=%v OpUrl=%v Timestamp=%v \n",
            event.OpRunId(),
            opUrl.String(),
            event.Timestamp(),
          )
        }
      case models.OpRunEndedEvent:
        if (event.RootOpRunId() == rootOpRunId) {
          fmt.Printf(
            "OpRunEnded: Outcome:%v Id=%v Timestamp=%v \n",
            event.Outcome(),
            event.OpRunId(),
            event.Timestamp(),
          )
          if (event.OpRunId() == rootOpRunId) {
            switch event.Outcome(){
            case engineModels.OpRunOutcomeSucceeded:
              this.exiter.Exit(0)
            case engineModels.OpRunOutcomeKilled:
              this.exiter.Exit(130)
            default:
              // fallback to general error
              this.exiter.Exit(1)
            }
            return
          }
        }
      default: // no op
      }
    }

  }

}
