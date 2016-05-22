package main

import (
  "fmt"
  "github.com/jawher/mow.cli"
  "github.com/opctl/sdk-for-golang/sdk"
  "github.com/opctl/sdk-for-golang/sdk/models"
  "os"
)

func eventsCmd(
dosCli *cli.Cli,
sdk sdk.Client,
) {

  dosCli.Command("events", "Get real time events from the server", func(eventsCmd *cli.Cmd) {

    eventsCmd.Action = func() {

      eventStream, err := sdk.GetEventStream()
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
          fmt.Printf(
            "%v \n",
            event.LogEntryMsg(),
          )
        case models.OpRunFinishedEvent:
          fmt.Printf(
            "OpEventsFinished: Id=%v ExitCode=%v Timestamp=%v \n",
            event.OpRunId(),
            event.OpRunExitCode(),
            event.Timestamp(),
          )
        case models.OpRunStartedEvent:
          opUrl := event.OpRunOpUrl()
          fmt.Printf(
            "OpEventsStarted: Id=%v OpUrl=%v Timestamp=%v \n",
            event.OpRunId(),
            opUrl.String(),
            event.Timestamp(),
          )
        case models.OpRunKilledEvent:
          fmt.Printf(
            "OpEventsKilled: Id=%v Timestamp=%v \n",
            event.OpRunId(),
            event.Timestamp(),
          )
        default: // no op
        }

      }

      fmt.Println()

    }

  })

}
