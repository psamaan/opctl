package opctlengine

//go:generate counterfeiter -o ./fakeRunOpUseCase.go --fake-name fakeRunOpUseCase ./ runOpUseCase

import (
  "github.com/opspec-io/engine-sdk-golang/models"
  "bytes"
)

type runOpUseCase interface {
  Execute(
  req models.RunOpReq,
  ) (
  opRunId string,
  correlationId string,
  err error,
  )
}

func newRunOpUseCase(
httpClient httpClient,
reqFactory reqFactory,
) runOpUseCase {

  return &_runOpUseCase{
    httpClient:httpClient,
    reqFactory:reqFactory,
  }

}

type _runOpUseCase struct {
  httpClient httpClient
  reqFactory reqFactory
}

func (this _runOpUseCase) Execute(
req models.RunOpReq,
) (
opRunId string,
correlationId string,
err error,
) {

  httpReq, err := this.reqFactory.Construct(
    "POST",
    "op-runs",
    req,
  )
  if (nil != err) {
    return
  }

  opRunIdBuffer := bytes.Buffer{}

  httpResp, err := this.httpClient.Do(
    httpReq,
    &opRunIdBuffer,
  )

  correlationId = httpResp.Header.Get("Correlation-Id")

  opRunId = opRunIdBuffer.String()

  return

}
