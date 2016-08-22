package docker

import (
  "fmt"
  "os/exec"
)

type isContainerExistentChecker interface {
  IsContainerExistentCheck(
  ) (isContainerExistent bool, err error)
}

func newIsContainerExistentChecker(
) (isContainerExistentChecker isContainerExistentChecker) {

  isContainerExistentChecker = &_isContainerExistentChecker{}

  return

}

type _isContainerExistentChecker struct{}

func (this _isContainerExistentChecker) IsContainerExistentCheck(
) (isContainerExistent bool, err error) {

  dockerPsCmd :=
    exec.Command(
      "docker",
      "ps",
      "-a",
      "-q",
      "-f",
      fmt.Sprintf("name=%v", containerName),
    )

  dockerPsCmdOutput, err := dockerPsCmd.Output()
  if (nil != err) {
    return
  }

  if (len(dockerPsCmdOutput) > 0 ) {
    isContainerExistent = true
  }

  return
}
