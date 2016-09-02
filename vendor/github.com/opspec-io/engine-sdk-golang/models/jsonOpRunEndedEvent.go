package models

import (
  "time"
  "encoding/json"
)

type JsonOpRunEndedEvent struct {
  data struct {
         CorrelationId string `json:"correlationId"`
         OpRunId       string `json:"opRunId"`
         Outcome       string `json:"outcome"`
         RootOpRunId   string `json:"rootOpRunId"`
         Timestamp     time.Time `json:"timestamp"`
       }
}

func (this *JsonOpRunEndedEvent) UnmarshalJSON(data []byte) error {

  return json.Unmarshal(data, &this.data)

}

func (this JsonOpRunEndedEvent) CorrelationId() string {
  return this.data.CorrelationId
}

func (this JsonOpRunEndedEvent) OpRunId() string {
  return this.data.OpRunId
}

func (this JsonOpRunEndedEvent) Outcome() string {
  return this.data.Outcome
}

func (this JsonOpRunEndedEvent) RootOpRunId() string {
  return this.data.RootOpRunId
}

func (this JsonOpRunEndedEvent) Timestamp() time.Time {
  return this.data.Timestamp
}
