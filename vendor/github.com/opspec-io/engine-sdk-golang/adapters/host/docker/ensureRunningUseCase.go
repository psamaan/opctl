package docker

type ensureRunningUseCase interface {
  Execute(
  image string,
  ) (err error)
}

func newEnsureRunningUseCase(
containerRemover           containerRemover,
containerStarter           containerStarter,
containerChecker  containerChecker,
) (ensureRunningUseCase ensureRunningUseCase) {

  ensureRunningUseCase = &_ensureRunningUseCase{
    containerRemover:containerRemover,
    containerStarter:containerStarter,
    containerChecker:containerChecker,
  }

  return

}

type _ensureRunningUseCase struct {
  containerRemover containerRemover
  containerStarter containerStarter
  containerChecker containerChecker
}

func (this _ensureRunningUseCase) Execute(
image string,
) (err error) {

  // handle obsolete container
  this.containerRemover.RemoveIfExists(obsoleteContainerName)

  // if valid container running or error checking, return
  isValidContainerRunning, err := this.containerChecker.IsValidContainerRunning(image)
  if (nil != err || isValidContainerRunning) {
    return
  }

  // cleanup invalid container
  this.containerRemover.RemoveIfExists(containerName)

  // start fresh container
  err = this.containerStarter.Start(image)

  return
}
