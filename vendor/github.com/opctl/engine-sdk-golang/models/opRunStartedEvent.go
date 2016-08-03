package models

import (
  "time"
)

type OpRunStartedEvent interface {
  CorrelationId() string
  OpRunId() string
  OpRunOpUrl() string
  ParentOpRunId() string
  RootOpRunId() string
  Timestamp() time.Time
}
