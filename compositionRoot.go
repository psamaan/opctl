package main

//go:generate counterfeiter -o ./fakeCompositionRoot.go --fake-name fakeCompositionRoot ./ compositionRoot

import (
  "github.com/opctl/engine-sdk-golang"
  dockerSdkHostAdapter "github.com/opctl/engine-sdk-golang/adapters/host/docker"
  "github.com/opspec-io/sdk-golang"
)

type compositionRoot interface {
  OpSpecSdk() opspec.Sdk
  OpCtlEngineSdk() opctlengine.Sdk
}

func newCompositionRoot(
) (compositionRoot compositionRoot, err error) {

  opctlEngineSdk, err := opctlengine.New(
    dockerSdkHostAdapter.New(),
  )
  if (nil != err) {
    return
  }

  compositionRoot = &_compositionRoot{
    opspecSdk:opspec.New(),
    opctlEngineSdk:opctlEngineSdk,
  }

  return

}

type _compositionRoot struct {
  opspecSdk      opspec.Sdk
  opctlEngineSdk opctlengine.Sdk
}

func (this _compositionRoot) OpSpecSdk() opspec.Sdk {
  return this.opspecSdk
}

func (this _compositionRoot) OpCtlEngineSdk() opctlengine.Sdk {
  return this.opctlEngineSdk
}
