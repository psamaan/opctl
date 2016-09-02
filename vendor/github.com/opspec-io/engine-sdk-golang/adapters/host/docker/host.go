package docker

import (
  "github.com/opspec-io/engine-sdk-golang/ports"
)

func New(
) (host ports.Host) {

  host = &_host{
    compositionRoot:newCompositionRoot(),
  }

  return

}

type _host struct {
  compositionRoot compositionRoot
}

func (this _host) EnsureRunning(
image string,
) (err error) {

  return this.
  compositionRoot.
    EnsureRunningUseCase().
    Execute(image)

}

func (this _host) GetHostname(
) (
hostname string,
err error,
) {

  return this.
  compositionRoot.
    GetHostnameUseCase().
    Execute()

}
