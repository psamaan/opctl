package models

import "time"

type Event interface {
  CorrelationId() string
  Timestamp() time.Time
}

