package opctlengine

//go:generate counterfeiter -o ./fakeGetLivenessUseCase.go --fake-name fakeGetLivenessUseCase ./ getLivenessUseCase

type getLivenessUseCase interface {
  Execute(
  ) (err error)
}

func newGetLivenessUseCase(
httpClient httpClient,
reqFactory reqFactory,
) getLivenessUseCase {

  return &_getLivenessUseCase{
    httpClient:httpClient,
    reqFactory:reqFactory,
  }

}

type _getLivenessUseCase struct {
  httpClient httpClient
  reqFactory reqFactory
}

func (this _getLivenessUseCase) Execute(
) (err error) {

  httpReq, err := this.reqFactory.Construct(
    "GET",
    "liveness",
    nil,
  )
  if (nil != err) {
    return
  }

  _, err = this.httpClient.Do(
    httpReq,
    nil,
  )

  return

}
