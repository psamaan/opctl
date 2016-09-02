package core

//go:generate counterfeiter -o ./fakeKillOpRunUseCase.go --fake-name fakeKillOpRunUseCase ./ killOpRunUseCase

import (
  "github.com/opspec-io/engine-sdk-golang/models"
  "github.com/opspec-io/engine-sdk-golang"
)

type killOpRunUseCase interface {
  Execute(
  opRunId string,
  ) error
}

func newKillOpRunUseCase(
opctlEngineSdk opctlengine.Sdk,
) killOpRunUseCase {
  return _killOpRunUseCase{
    opctlEngineSdk:opctlEngineSdk,
  }
}

type _killOpRunUseCase struct {
  opctlEngineSdk opctlengine.Sdk
}

func (this _killOpRunUseCase) Execute(
opRunId string,
) error {

  _, err := this.opctlEngineSdk.KillOpRun(
    *models.NewKillOpRunReq(
      opRunId,
    ),
  )

  return err

}
