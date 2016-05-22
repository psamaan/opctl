package main

import (
  "github.com/opctl/sdk-for-golang/sdk"
  dockerSdkHostAdapter "github.com/opctl/sdk-for-golang/sdk/adapters/host/docker"
)

type compositionRoot interface {
  Sdk() sdk.Client
}

func newCompositionRoot(
) (compositionRoot compositionRoot, err error) {

  sdk, err := sdk.New(
    dockerSdkHostAdapter.New(),
  )
  if (nil != err) {
    return
  }

  compositionRoot = &_compositionRoot{
    sdk:sdk,
  }

  return

}

type _compositionRoot struct {
  sdk sdk.Client
}

func (this _compositionRoot) Sdk() sdk.Client {
  return this.sdk
}
