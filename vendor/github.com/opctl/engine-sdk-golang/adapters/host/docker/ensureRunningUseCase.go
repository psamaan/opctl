package docker

import (
  "fmt"
  "errors"
)

type ensureRunningUseCase interface {
  Execute(
  image string,
  ) (err error)
}

func newEnsureRunningUseCase(
containerRemover           containerRemover,
containerStarter           containerStarter,
isContainerExistentChecker isContainerExistentChecker,
isContainerRunningChecker  isContainerRunningChecker,
) (ensureRunningUseCase ensureRunningUseCase) {

  ensureRunningUseCase = &_ensureRunningUseCase{
    containerRemover:containerRemover,
    containerStarter:containerStarter,
    isContainerExistentChecker:isContainerExistentChecker,
    isContainerRunningChecker:isContainerRunningChecker,
  }

  return

}

type _ensureRunningUseCase struct {
  containerRemover           containerRemover
  containerStarter           containerStarter
  isContainerExistentChecker isContainerExistentChecker
  isContainerRunningChecker  isContainerRunningChecker
}

func (this _ensureRunningUseCase) Execute(
image string,
) (err error) {

  // if already running we're done
  isContainerRunning, err := this.isContainerRunningChecker.IsContainerRunningCheck(image)
  if (nil != err) {
    err = errors.New(
      fmt.Sprintf("Unable to connect to docker engine\n error was: %v \n", err),
    )
    return
  } else if (isContainerRunning) {
    return
  }

  // if not running but exists we need to kill it.
  isContainerExistent, err := this.isContainerExistentChecker.IsContainerExistentCheck()
  if (nil != err) {
    return
  }
  if (isContainerExistent) {
    err = this.containerRemover.ContainerRemove()
    if (nil != err) {
      return
    }
  }

  err = this.containerStarter.ContainerStart(image)

  return
}
