package models

import "encoding/json"

type EventMsg struct {
  Type string `json:"type"`
  Data json.RawMessage `json:"data"`
}
