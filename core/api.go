package core

import (
  "github.com/opspec-io/sdk-golang/pkg/engineprovider"
  "github.com/opspec-io/sdk-golang/pkg/engineclient"
  "github.com/opspec-io/sdk-golang/pkg/bundle"
  "io"
  "os"
)

//go:generate counterfeiter -o ./fakeApi.go --fake-name FakeApi ./ Api

type Api interface {
  CreateCollection(
  description string,
  name string,
  )

  CreateOp(
  description string,
  name string,
  )

  KillOpRun(
  opRunId string,
  )

  ListOpsInCollection(
  )

  RunOp(
  args []string,
  name string,
  )

  SetCollectionDescription(
  description string,
  )

  SetOpDescription(
  description string,
  name string,
  )

  StreamEvents(
  )
}

func New(
engineProvider engineprovider.EngineProvider,
) Api {

  return &_api{
    bundle:bundle.New(),
    exiter:newExiter(),
    engineClient:engineclient.New(engineProvider),
    workDirPathGetter:newWorkDirPathGetter(),
    writer:os.Stdout,
  }

}

type _api struct {
  bundle bundle.Bundle
  exiter exiter
  engineClient engineclient.EngineClient
  workDirPathGetter workDirPathGetter
  writer io.Writer
}
