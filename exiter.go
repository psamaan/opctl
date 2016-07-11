package main

//go:generate counterfeiter -o ./fakeExiter.go --fake-name fakeExiter ./ exiter

import "os"

type exiter interface {
  Exit(code int)
}

func newExiter() exiter {
  return _exiter{}
}

type _exiter struct{}

func (this _exiter) Exit(code int) {
  os.Exit(code)
}
