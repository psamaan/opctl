package docker

import (
  "os/exec"
)

type containerRemover interface {
  ContainerRemove(
  ) (err error)
}

func newContainerRemover(
) (containerRemover containerRemover) {

  containerRemover = &_containerRemover{}

  return

}

type _containerRemover struct{}

func (this _containerRemover) ContainerRemove(
) (err error) {

  dockerRmCmd :=
  exec.Command(
    "docker",
    "rm",
    "-v",
    containerName,
  )

  _, err = dockerRmCmd.Output()

  return
}
