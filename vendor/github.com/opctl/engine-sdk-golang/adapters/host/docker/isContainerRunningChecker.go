package docker

import (
  "fmt"
  "os/exec"
)

type isContainerRunningChecker interface {
  IsContainerRunningCheck(
  ) (isContainerRunning bool, err error)
}

func newIsContainerRunningChecker(
) (isContainerRunningChecker isContainerRunningChecker) {

  isContainerRunningChecker = &_isContainerRunningChecker{}

  return

}

type _isContainerRunningChecker struct{}

func (this _isContainerRunningChecker) IsContainerRunningCheck(
) (isContainerRunning bool, err error) {

  dockerPsCmd :=
  exec.Command(
    "docker",
    "ps",
    "-q",
    "-f",
    fmt.Sprintf("name=%v", containerName),
  )

  dockerPsCmdOutput, err := dockerPsCmd.Output()
  if (nil != err) {
    return
  }

  if (len(dockerPsCmdOutput) > 0 ) {
    isContainerRunning = true
  }

  return
}
