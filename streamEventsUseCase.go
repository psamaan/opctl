package main

//go:generate counterfeiter -o ./fakeStreamEventsUseCase.go --fake-name fakeStreamEventsUseCase ./ streamEventsUseCase

import (
  "github.com/opctl/engine-sdk-golang/models"
  "fmt"
  "github.com/opctl/engine-sdk-golang"
)

type streamEventsUseCase interface {
  Execute()
}

func newStreamEventsUseCase(
exiter exiter,
opctlEngineSdk opctlengine.Sdk,
) streamEventsUseCase {
  return _streamEventsUseCase{
    exiter:exiter,
    opctlEngineSdk:opctlEngineSdk,
  }
}

type _streamEventsUseCase struct {
  exiter         exiter
  opctlEngineSdk opctlengine.Sdk
}

func (this _streamEventsUseCase) Execute(
) {

  eventChannel, err := this.opctlEngineSdk.GetEventStream()
  if (nil != err) {
    this.exiter.Exit(1)
    return
  }

  for {

    event, isEventChannelOpen := <-eventChannel
    if (!isEventChannelOpen) {
      this.exiter.Exit(1)
      return
    }

    switch event := event.(type) {
    case models.LogEntryEmittedEvent:
      fmt.Printf(
        "%v \n",
        event.LogEntryMsg(),
      )
    case models.OpRunStartedEvent:
      fmt.Printf(
        "OpRunStarted: Id=%v OpUrl=%v Timestamp=%v \n",
        event.OpRunId(),
        event.OpRunOpUrl(),
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
}
