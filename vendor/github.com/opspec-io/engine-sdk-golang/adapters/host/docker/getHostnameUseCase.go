package docker

import (
  "os"
  "strings"
)

type getHostnameUseCase interface {
  Execute(
  ) (hostname string, err error)
}

func newGetHostnameUseCase(
) (getHostnameUseCase getHostnameUseCase) {

  getHostnameUseCase = &_getHostnameUseCase{}

  return

}

type _getHostnameUseCase struct{}

func (this _getHostnameUseCase) Execute(
) (hostname string, err error) {

  dockerComposeHost, isDockerMachine := os.LookupEnv("DOCKER_HOST")
  if (isDockerMachine) {
    addrParts := strings.Split(dockerComposeHost, ":")
    hostname = strings.TrimPrefix(addrParts[1], "//")
  } else {
    hostname = "localhost"
  }

  return
}
