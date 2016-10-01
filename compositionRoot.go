package main

//go:generate counterfeiter -o ./fakeCompositionRoot.go --fake-name fakeCompositionRoot ./ compositionRoot

import (
  "github.com/opspec-io/cli/core"
  "github.com/opspec-io/sdk-golang/adapters"
)

type compositionRoot interface {
  CoreApi() core.Api
}

func newCompositionRoot(
engineHost adapters.EngineHost,
) compositionRoot {

  return &_compositionRoot{
    coreApi:core.New(engineHost),
  }

}

type _compositionRoot struct {
  coreApi core.Api
}

func (_compositionRoot _compositionRoot) CoreApi() core.Api {
  return _compositionRoot.coreApi
}
