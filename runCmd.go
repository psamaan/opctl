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
)

func runCmd(
dosCli *cli.Cli,
devOpSpecSdk sdk.Client,
) {

  dosCli.Command("run", "Run an op", func(runCmd *cli.Cmd) {

    runCmd.Spec = "OP_NAME"

    var (
      name = runCmd.StringArg("OP_NAME", "", "the name of the op")
    )

    runCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        cli.Exit(1)
      }

      var opUrl *url.URL
      opUrl, err = url.Parse(
        fmt.Sprintf(
          "%v/.dev-op-spec/ops/%v",
          currentWorkDir,
          *name,
        ),
      )
      if (nil != err) {
        fmt.Errorf(err.Error())
        os.Exit(1)
      }

      opRunId, correlationId, err := devOpSpecSdk.RunOp(
        *models.NewRunOpReq(
          opUrl,
        ),
      )
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
      }

      // handle SIGINT
      signalChannel := make(chan os.Signal, 1)
      signal.Notify(
        signalChannel,
        syscall.SIGINT,
      )

      go func() {

        <-signalChannel

        devOpSpecSdk.KillOpRun(
          *models.NewKillOpRunReq(
            opRunId,
          ),
        )

        return

      }()

      eventStream, err := devOpSpecSdk.GetEventStream()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
      }

      for {

        event, isOpen := <-eventStream
        if (!isOpen) {
          return
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
              return
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

      if (nil != err) {

        fmt.Fprintln(os.Stderr, *name + " unsuccessful")
        os.Exit(2)

      } else {

        fmt.Printf("%v successful \n", *name)

      }

      fmt.Println()

    }

  })

}
