package opctlengine

//go:generate counterfeiter -o ./fakeCompositionRoot.go --fake-name fakeCompositionRoot ./ compositionRoot

import (
  "github.com/opctl/engine-sdk-golang/ports"
  "net/http"
)

type compositionRoot interface {
  GetLivenessUseCase() getLivenessUseCase
  GetEventStreamUseCase() getEventStreamUseCase
  KillOpRunUseCase() killOpRunUseCase
  RunOpUseCase() runOpUseCase
}

func newCompositionRoot(
host ports.Host,
) (compositionRoot compositionRoot) {

  httpClient := newHttpClient(http.DefaultClient)
  reqFactory := newReqFactory(host)

  compositionRoot = &_compositionRoot{
    getLivenessUseCase:newGetLivenessUseCase(httpClient, reqFactory),
    getEventStreamUseCase:newGetEventStreamUseCase(host),
    killOpRunUseCase:newKillOpRunUseCase(httpClient, reqFactory),
    runOpUseCase: newRunOpUseCase(httpClient, reqFactory),
  }

  return

}

type _compositionRoot struct {
  getLivenessUseCase        getLivenessUseCase
  getEventStreamUseCase     getEventStreamUseCase
  killOpRunUseCase          killOpRunUseCase
  runOpUseCase              runOpUseCase
}

func (this _compositionRoot) GetLivenessUseCase() getLivenessUseCase {
  return this.getLivenessUseCase
}

func (this _compositionRoot) GetEventStreamUseCase() getEventStreamUseCase {
  return this.getEventStreamUseCase
}

func (this _compositionRoot) KillOpRunUseCase() killOpRunUseCase {
  return this.killOpRunUseCase
}

func (this _compositionRoot) RunOpUseCase() runOpUseCase {
  return this.runOpUseCase
}
