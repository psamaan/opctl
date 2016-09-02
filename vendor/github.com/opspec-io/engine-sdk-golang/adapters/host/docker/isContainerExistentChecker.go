package docker

import (
  "fmt"
  "os/exec"
)

type isContainerExistentChecker interface {
  IsContainerExistentCheck(
  image string,
  ) (isContainerExistent bool, err error)
}

func newIsContainerExistentChecker(
) (isContainerExistentChecker isContainerExistentChecker) {

  isContainerExistentChecker = &_isContainerExistentChecker{}

  return

}

type _isContainerExistentChecker struct{}

func (this _isContainerExistentChecker) IsContainerExistentCheck(
image string,
) (isContainerExistent bool, err error) {

  dockerPsCmd :=
    exec.Command(
      "docker",
      "ps",
      "-a",
      "-q",
      "-f",
      fmt.Sprintf("name=%v", containerName),
      "-f",
      fmt.Sprintf("ancestor=%v", image),
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
