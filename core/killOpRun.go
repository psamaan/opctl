package core

//go:generate counterfeiter -o ./fakeKillOpRunUseCase.go --fake-name fakeKillOpRunUseCase ./ killOpRunUseCase

import (
  "github.com/opspec-io/sdk-golang/models"
)

func (this _api) KillOpRun(
opRunId string,
) {

  err := this.engineClient.KillOpRun(
    models.KillOpRunReq{
      OpRunId:opRunId,
    },
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }

}
