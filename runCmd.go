package main

import (
  "fmt"
  "github.com/jawher/mow.cli"
  "github.com/opctl/sdk-for-golang/sdk"
  "github.com/opctl/sdk-for-golang/sdk/models"
  "net/url"
  "os"
  "syscall"
  "os/signal"
  "strings"
)

func runCmd(
dosCli *cli.Cli,
sdk sdk.Client,
) {

  dosCli.Command("run", "Run an op", func(runCmd *cli.Cmd) {

    runCmd.Spec = "[-a...] OP_NAME"

    var (
      argsSlice = runCmd.StringsOpt("a", []string{}, "Pass args to op in format: NAME[=VALUE] (gets VALUE from env if not provided)")
      name = runCmd.StringArg("OP_NAME", "", "the name of the op")
    )

    runCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      argsMap := make(map[string]string)
      for _, arg := range *argsSlice {

        argParts := strings.Split(arg, "=")

        argName := argParts[0]
        var argValue string
        if (len(argParts) > 1) {
          argValue = argParts[1]
        } else {
          argValue = os.Getenv(arg)
        }

        argsMap[argName] = argValue
      }

      var opUrl *url.URL
      opUrl, err = url.Parse(
        fmt.Sprintf(
          "%v/.opspec/%v",
          currentWorkDir,
          *name,
        ),
      )
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      // init signal channel
      intSignalsReceived := 0
      signalChannel := make(chan os.Signal, 1)
      signal.Notify(
        signalChannel,
        syscall.SIGINT, //handle SIGINTs
      )

      // init event channel
      eventChannel, err := sdk.GetEventStream()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      opRunId, correlationId, err := sdk.RunOp(
        *models.NewRunOpReq(
          argsMap,
          opUrl,
        ),
      )
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      for {
        select {

        case <-signalChannel:
          if (intSignalsReceived == 0) {

            sdk.KillOpRun(
              *models.NewKillOpRunReq(
                opRunId,
              ),
            )

            intSignalsReceived++
            fmt.Println()
            fmt.Println("Gracefully stopping... (press Ctrl+C again to force)")
          } else {
            os.Exit(1)
          }

        case event, isEventChannelOpen := <-eventChannel:
          if (!isEventChannelOpen) {
            os.Exit(1)
          }

          switch event := event.(type) {
          case models.LogEntryEmittedEvent:
            // @TODO: this doesn't catch log entries for the same tree but triggered from different actions (such as kills)
            if (event.CorrelationId() == correlationId) {
              fmt.Printf(
                "%v \n",
                event.LogEntryMsg(),
              )
            }
          case models.OpRunFinishedEvent:
            if (event.RootOpRunId() == opRunId) {
              fmt.Printf(
                "OpRunFinished: Id=%v ExitCode=%v Timestamp=%v \n",
                event.OpRunId(),
                event.OpRunExitCode(),
                event.Timestamp(),
              )
              if (event.OpRunId() == opRunId) {
                os.Exit(event.OpRunExitCode())
              }
            }
          case models.OpRunStartedEvent:
            if (event.RootOpRunId() == opRunId) {
              opUrl := event.OpRunOpUrl()
              fmt.Printf(
                "OpRunStarted: Id=%v OpUrl=%v Timestamp=%v \n",
                event.OpRunId(),
                opUrl.String(),
                event.Timestamp(),
              )

            }
          case models.OpRunKilledEvent:
            if (event.RootOpRunId() == opRunId) {
              fmt.Printf(
                "OpRunKilled: Id=%v Timestamp=%v \n",
                event.OpRunId(),
                event.Timestamp(),
              )
            }
          default: // no op
          }
        }

      }

    }

  })

}
