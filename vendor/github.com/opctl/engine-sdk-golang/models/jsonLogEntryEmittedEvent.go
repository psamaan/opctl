package models

import (
  "time"
  "encoding/json"
)

type JsonLogEntryEmittedEvent struct {
  data struct {
         CorrelationId        string `json:"correlationId"`
         LogEntryMsg          string `json:"logEntryMsg"`
         LogEntryOutputStream string `json:"logEntryOutputStream"`
         Timestamp            time.Time `json:"timestamp"`
       }
}

func (this *JsonLogEntryEmittedEvent) UnmarshalJSON(data []byte) error {

  return json.Unmarshal(data, &this.data)

}

func (this JsonLogEntryEmittedEvent) CorrelationId() string {
  return this.data.CorrelationId
}

func (this JsonLogEntryEmittedEvent) LogEntryMsg() string {
  return this.data.LogEntryMsg
}

func (this JsonLogEntryEmittedEvent) LogEntryOutputStream() string {
  return this.data.LogEntryOutputStream
}

func (this JsonLogEntryEmittedEvent) Timestamp() time.Time {
  return this.data.Timestamp
}
