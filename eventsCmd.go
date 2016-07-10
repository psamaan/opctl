package main

import (
  "fmt"
  "github.com/jawher/mow.cli"
  "github.com/opctl/engine-sdk-golang"
  "github.com/opctl/engine-sdk-golang/models"
  "os"
)

func eventsCmd(
opctlCli *cli.Cli,
opctlEngineSdk opctlengine.Sdk,
) {

  opctlCli.Command("events", "Get real time events from the server", func(eventsCmd *cli.Cmd) {

    eventsCmd.Action = func() {

      eventStream, err := opctlEngineSdk.GetEventStream()
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
        case models.OpRunStartedEvent:
          opUrl := event.OpRunOpUrl()
          fmt.Printf(
            "OpRunStarted: Id=%v OpUrl=%v Timestamp=%v \n",
            event.OpRunId(),
            opUrl.String(),
            event.Timestamp(),
          )
        case models.OpRunEndedEvent:
          fmt.Printf(
            "OpRunEnded: Outcome=%v Id=%v Timestamp=%v \n",
            event.Outcome(),
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
