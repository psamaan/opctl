package main

//go:generate counterfeiter -o ./fakeCompositionRoot.go --fake-name fakeCompositionRoot ./ compositionRoot

import (
  "github.com/opspec-io/cli/core"
)

type compositionRoot interface {
  CoreApi() core.Api
}

func newCompositionRoot(
) compositionRoot {

  return &_compositionRoot{
    coreApi:core.New(),
  }

}

type _compositionRoot struct {
  coreApi core.Api
}

func (_compositionRoot _compositionRoot) CoreApi() core.Api {
  return _compositionRoot.coreApi
}
