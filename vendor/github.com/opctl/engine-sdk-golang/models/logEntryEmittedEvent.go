package models

import (
  "time"
)

type LogEntryEmittedEvent interface {
  CorrelationId() string
  LogEntryMsg() string
  LogEntryOutputStream() string
  Timestamp() time.Time
}
