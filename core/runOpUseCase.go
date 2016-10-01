package core

//go:generate counterfeiter -o ./fakeRunOpUseCase.go --fake-name fakeRunOpUseCase ./ runOpUseCase

import (
  "fmt"
  "os"
  "syscall"
  "os/signal"
  "strings"
  "github.com/opspec-io/sdk-golang"
  "path"
  "github.com/peterh/liner"
  "github.com/opspec-io/sdk-golang/models"
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
workDirPathGetter workDirPathGetter,
) runOpUseCase {
  return _runOpUseCase{
    exiter:exiter,
    opspecSdk:opspecSdk,
    workDirPathGetter:workDirPathGetter,
  }
}

type _runOpUseCase struct {
  exiter            exiter
  opspecSdk         opspec.Sdk
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
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }

  line := liner.NewLiner()
  line.SetCtrlCAborts(true)
  defer line.Close()

  // [only] pass defined params
  argsMap := make(map[string]string)
  for _, opParam := range opView.Inputs {

    if providedArg, ok := providedArgMap[opParam.Name]; ok {
      argsMap[opParam.Name] = providedArg
    } else if ("" != os.Getenv(opParam.Name)) {
      argsMap[opParam.Name] = os.Getenv(opParam.Name)
    } else {
      var argValue string
      argPrompt := fmt.Sprintf("%v: ", opParam.Name)
      if (opParam.IsSecret) {
        argValue, err = line.PasswordPrompt(argPrompt)
      } else {
        argValue, err = line.Prompt(argPrompt)
      }

      if (nil != err) {
        this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
        return // support fake exiter
      }

      argsMap[opParam.Name] = argValue
    }
  }

  // init signal channel
  intSignalsReceived := 0
  signalChannel := make(chan os.Signal, 1)
  defer close(signalChannel)

  signal.Notify(
    signalChannel,
    syscall.SIGINT, //handle SIGINTs
  )

  // init event channel
  eventChannel, err := this.opspecSdk.GetEventStream()
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }

  rootOpRunId, err := this.opspecSdk.StartOpRun(
    models.StartOpRunReq{
      Args:argsMap,
      OpUrl:opPath,
    },
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }

  for {
    select {

    case <-signalChannel:
      if (intSignalsReceived == 0) {

        intSignalsReceived++
        fmt.Println()
        fmt.Println("Gracefully stopping... (signal Control-C again to force)")

        this.opspecSdk.KillOpRun(
          models.KillOpRunReq{
            OpRunId:rootOpRunId,
          },
        )
      } else {
        this.exiter.Exit(ExitReq{Message:"Terminated by Control-C", Code:130})
        return // support fake exiter
      }

    case event, isEventChannelOpen := <-eventChannel:
      if (!isEventChannelOpen) {
        this.exiter.Exit(ExitReq{Message:"Event channel closed unexpectedly", Code:1})
        return // support fake exiter
      }

      if (nil != event.ContainerStdOutWrittenTo && event.ContainerStdOutWrittenTo.RootOpRunId == rootOpRunId) {
        fmt.Fprintf(os.Stdout, "%v \n", string(event.ContainerStdOutWrittenTo.Data))
      } else if (nil != event.ContainerStdErrWrittenTo && event.ContainerStdErrWrittenTo.RootOpRunId == rootOpRunId) {
        fmt.Fprintf(os.Stderr, "%v \n", string(event.ContainerStdErrWrittenTo.Data))
      } else if (nil != event.OpRunStarted && event.OpRunStarted.RootOpRunId == rootOpRunId) {
        fmt.Printf(
          "OpRunStarted: Id=%v OpRef=%v Timestamp=%v \n",
          event.OpRunStarted.OpRunId,
          event.OpRunStarted.OpRef,
          event.Timestamp,
        )
      } else if (nil != event.OpRunEnded && event.OpRunEnded.RootOpRunId == rootOpRunId) {
        fmt.Printf(
          "OpRunEnded: Id=%v Outcome=%v Timestamp=%v \n",
          event.OpRunEnded.OpRunId,
          event.OpRunEnded.Outcome,
          event.Timestamp,
        )
        if (event.OpRunEnded.OpRunId == rootOpRunId) {
          switch event.OpRunEnded.Outcome{
          case models.OpRunOutcomeSucceeded:
            this.exiter.Exit(ExitReq{Message:"", Code:0})
          case models.OpRunOutcomeKilled:
            this.exiter.Exit(ExitReq{Message:"", Code:137})
          case models.OpRunOutcomeFailed:
            this.exiter.Exit(ExitReq{Message:"", Code:1})
          default:
            // fallback to general error
            this.exiter.Exit(ExitReq{Message:fmt.Sprintf("Received unknown outcome `%v`", event.OpRunEnded.Outcome), Code:1})
          }
          return // support fake exiter
        }
      }
    }

  }

}
