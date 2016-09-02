package opctlengine

//go:generate counterfeiter -o ./fakeSdk.go --fake-name FakeSdk ./ Sdk

import (
  "github.com/opspec-io/engine-sdk-golang/models"
  "github.com/opspec-io/engine-sdk-golang/ports"
  "time"
  "errors"
)

type Sdk interface {
  GetLiveness(
  ) (err error)

  GetEventStream(
  ) (
  stream chan models.Event,
  err error,
  )

  KillOpRun(
  req models.KillOpRunReq,
  ) (
  correlationId string,
  err error,
  )

  RunOp(
  req models.RunOpReq,
  ) (
  opRunId string,
  correlationId string,
  err error,
  )
}

func New(
host ports.Host,
) (sdk Sdk, err error) {

  err = host.EnsureRunning("opspec/engine:0.1.7")
  if (nil != err) {
    return
  }

  compositionRoot := newCompositionRoot(
    host,
  )

  timeout := time.After(15 * time.Second)
  tick := time.Tick(500 * time.Millisecond)
  // Keep trying until we're timed out or engine is alive
  forLoop: for {
    select {
    // Got a timeout! fail with a timeout error
    case <-timeout:
      err = errors.New("Timeout exceeded while liveness checking engine. \n")
      return
    // Got a tick, we should check on doSomething()
    case <-tick:
      err := compositionRoot.GetLivenessUseCase().Execute()
      if err == nil {
        break forLoop
      }
    }
  }

  sdk = &_sdk{
    compositionRoot:compositionRoot,
  }

  return
}

type _sdk struct {
  compositionRoot compositionRoot
}

func (this _sdk) GetLiveness(
) (err error) {
  return this.
  compositionRoot.
    GetLivenessUseCase().
    Execute()
}

func (this _sdk) GetEventStream(
) (stream chan models.Event, err error) {
  return this.
  compositionRoot.
    GetEventStreamUseCase().
    Execute()
}

func (this _sdk) KillOpRun(
req models.KillOpRunReq,
) (
correlationId string,
err error,
) {
  return this.
  compositionRoot.
    KillOpRunUseCase().
    Execute(req)
}

func (this _sdk) RunOp(
req models.RunOpReq,
) (
opRunId string,
correlationId string,
err error,
) {
  return this.
  compositionRoot.
    RunOpUseCase().
    Execute(
    req,
  )
}
