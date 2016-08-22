package main

//go:generate counterfeiter -o ./fakeExiter.go --fake-name fakeExiter ./ exiter

type exitReq struct {
  Message string
  Code    int
}

type exiter interface {
  Exit(req exitReq)
}

func newExiter() exiter {
  return _exiter{}
}

type _exiter struct{}

func (this _exiter) Exit(req exitReq) {
  panic(req)
}
