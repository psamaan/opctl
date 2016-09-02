package opctlengine

//go:generate counterfeiter -o ./fakeGetEventStreamUseCase.go --fake-name fakeGetEventStreamUseCase ./ getEventStreamUseCase

import (
  "github.com/opspec-io/engine-sdk-golang/ports"
  "github.com/opspec-io/engine-sdk-golang/models"
  "fmt"
  "log"
  "encoding/json"
  "github.com/gorilla/websocket"
)

type getEventStreamUseCase interface {
  Execute(
  ) (stream chan models.Event, err error)
}

func newGetEventStreamUseCase(
host ports.Host,
) getEventStreamUseCase {

  return &_getEventStreamUseCase{
    host:host,
  }

}

type _getEventStreamUseCase struct {
  host ports.Host
}

func (this _getEventStreamUseCase) Execute(
) (eventStream chan models.Event, err error) {

  eventStream = make(chan models.Event, 1000)

  hostname, err := this.host.GetHostname()
  if (nil != err) {
    return
  }

  c, _, err := websocket.DefaultDialer.Dial(
    fmt.Sprintf("ws://%v:42224/event-stream", hostname),
    nil,
  )
  if (err != nil) {
    fmt.Println(err)
  }

  go func() {
    defer c.Close()
    for {

      _, bytes, err := c.ReadMessage()
      if err != nil {
        log.Println("read:", err)
        return
      }

      var eventMsg models.EventMsg
      err = json.Unmarshal(bytes, &eventMsg)
      if (nil != err) {
        fmt.Printf("json.Unmarshal err: %v \n", err)
      }

      var event models.Event

      switch eventMsg.Type {
      case "LogEntryEmitted":
        var logEntryEmittedEvent models.JsonLogEntryEmittedEvent
        err = json.Unmarshal(eventMsg.Data, &logEntryEmittedEvent)
        event = logEntryEmittedEvent
      case "OpRunEnded":
        var opRunEndedEvent models.JsonOpRunEndedEvent
        err = json.Unmarshal(eventMsg.Data, &opRunEndedEvent)
        event = opRunEndedEvent
      case "OpRunStarted":
        var opRunStartedEvent models.JsonOpRunStartedEvent
        err = json.Unmarshal(eventMsg.Data, &opRunStartedEvent)
        event = opRunStartedEvent
      default:
        err = fmt.Errorf("received unexpected eventMsg type: %v \n", eventMsg.Type)
      }

      if (nil != err) {
        log.Fatalln(err)
      } else {
        eventStream <- event
      }

    }
  }()

  return

}
