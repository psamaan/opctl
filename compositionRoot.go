package main

import (
  "github.com/opctl/sdk-for-golang/sdk"
  sdkDosEngine "github.com/opctl/sdk-for-golang/sdk/adapters/dosruntime/dosengine"
  sdkNetHttp "github.com/opctl/sdk-for-golang/sdk/adapters/http/net"
  "net/url"
)

type compositionRoot interface {
  DevOpSpecSdk() sdk.Client
}

func newCompositionRoot(
) (compositionRoot compositionRoot, err error) {

  sdkDosRuntime := sdkDosEngine.New()

  baseUrl, err := url.Parse("http://192.168.99.100:42224/")
  if (nil != err) {
    return
  }

  sdkHttpAdapter, err := sdkNetHttp.New(
    sdkNetHttp.NewConfig(*baseUrl),
  )
  if (nil != err) {
    return
  }

  devOpSpecSdk, err := sdk.New(
    sdkDosRuntime,
    sdkHttpAdapter,
  )
  if (nil != err) {
    return
  }

  compositionRoot = &_compositionRoot{
    devOpSpecSdk:devOpSpecSdk,
  }

  return

}

type _compositionRoot struct {
  devOpSpecSdk sdk.Client
}

func (this _compositionRoot) DevOpSpecSdk() sdk.Client {
  return this.devOpSpecSdk
}
