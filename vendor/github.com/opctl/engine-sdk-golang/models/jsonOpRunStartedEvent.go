package models

import (
  "time"
  "encoding/json"
)

type JsonOpRunStartedEvent struct {
  data struct {
         CorrelationId string `json:"correlationId"`
         OpRunId       string `json:"opRunId"`
         OpRunOpUrl    string `json:"opRunOpUrl"`
         ParentOpRunId string `json:"parentOpRunId"`
         RootOpRunId   string `json:"rootOpRunId"`
         Timestamp     time.Time `json:"timestamp"`
       }
}

func (this *JsonOpRunStartedEvent) UnmarshalJSON(data []byte) error {

  return json.Unmarshal(data, &this.data)

}

func (this JsonOpRunStartedEvent) CorrelationId() string {
  return this.data.CorrelationId
}

func (this JsonOpRunStartedEvent) OpRunId() string {
  return this.data.OpRunId
}

func (this JsonOpRunStartedEvent) OpRunOpUrl() string {
  return this.data.OpRunOpUrl
}

func (this JsonOpRunStartedEvent) ParentOpRunId() string {
  return this.data.ParentOpRunId
}

func (this JsonOpRunStartedEvent) RootOpRunId() string {
  return this.data.RootOpRunId
}

func (this JsonOpRunStartedEvent) Timestamp() time.Time {
  return this.data.Timestamp
}
