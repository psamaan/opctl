package core

//go:generate counterfeiter -o ./fakeKillOpRunUseCase.go --fake-name fakeKillOpRunUseCase ./ killOpRunUseCase

import ()
import (
  "github.com/opspec-io/sdk-golang"
  "github.com/opspec-io/sdk-golang/models"
)

type killOpRunUseCase interface {
  Execute(
  opRunId string,
  ) error
}

func newKillOpRunUseCase(
opspecSdk opspec.Sdk,
) killOpRunUseCase {
  return _killOpRunUseCase{
    opspecSdk:opspecSdk,
  }
}

type _killOpRunUseCase struct {
  opspecSdk opspec.Sdk
}

func (this _killOpRunUseCase) Execute(
opRunId string,
) error {

  err := this.opspecSdk.KillOpRun(
    models.KillOpRunReq{
      OpRunId:opRunId,
    },
  )

  return err

}
