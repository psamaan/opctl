package opctlengine

//go:generate counterfeiter -o ./fakeKillOpRunUseCase.go --fake-name fakeKillOpRunUseCase ./ killOpRunUseCase

import (
  "github.com/opctl/engine-sdk-golang/models"
)

type killOpRunUseCase interface {
  Execute(
  req models.KillOpRunReq,
  ) (
  correlationId string,
  err error,
  )
}

func newKillOpRunUseCase(
httpClient httpClient,
reqFactory reqFactory,
) killOpRunUseCase {

  return &_killOpRunUseCase{
    httpClient:httpClient,
    reqFactory:reqFactory,
  }

}

type _killOpRunUseCase struct {
  httpClient httpClient
  reqFactory reqFactory
}

func (this _killOpRunUseCase) Execute(
req models.KillOpRunReq,
) (
correlationId string,
err error,
) {

  httpReq, err := this.reqFactory.Construct(
    "POST",
    "op-run-kills",
    req,
  )
  if (nil != err) {
    return
  }

  httpResp, err := this.httpClient.Do(
    httpReq,
    nil,
  )
  if (nil != err) {
    return
  }

  correlationId = httpResp.Header.Get("Correlation-Id")

  return

}
