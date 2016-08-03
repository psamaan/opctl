package models

import (
  "time"
)

type OpRunEndedEvent interface {
  CorrelationId() string
  OpRunId() string
  Outcome() string
  RootOpRunId() string
  Timestamp() time.Time
}

