package core

//go:generate counterfeiter -o ./fakeStreamEventsUseCase.go --fake-name fakeStreamEventsUseCase ./ streamEventsUseCase

import (
  "fmt"
  "os"
  "time"
)

func (this _api) StreamEvents(
) {

  eventChannel, err := this.engineClient.GetEventStream()
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }

  for {

    event, isEventChannelOpen := <-eventChannel
    if (!isEventChannelOpen) {
      this.exiter.Exit(ExitReq{Message:"Event channel closed unexpectedly", Code:1})
      return // support fake exiter
    }

    if (nil != event.ContainerStdOutWrittenTo) {
      fmt.Fprintf(os.Stdout, "%v \n", string(event.ContainerStdOutWrittenTo.Data))
    } else if (nil != event.ContainerStdErrWrittenTo) {
      fmt.Fprintf(os.Stderr, "%v \n", string(event.ContainerStdErrWrittenTo.Data))
    } else if (nil != event.OpRunStarted) {
      fmt.Printf(
        "OpRunStarted: Id=%v OpRef=%v Timestamp=%v \n",
        event.OpRunStarted.OpRunId,
        event.OpRunStarted.OpRef,
        event.Timestamp.Format(time.RFC3339),
      )
    } else if (nil != event.OpRunEnded) {
      fmt.Printf(
        "OpRunEnded: Id=%v Outcome=%v Timestamp=%v \n",
        event.OpRunEnded.OpRunId,
        event.OpRunEnded.Outcome,
        event.Timestamp.Format(time.RFC3339),
      )
    }

  }
}
