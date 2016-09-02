package docker

type compositionRoot interface {
  EnsureRunningUseCase() ensureRunningUseCase
  GetHostnameUseCase() getHostnameUseCase
}

func newCompositionRoot(
) (compositionRoot compositionRoot) {

  ensureRunningUseCase := newEnsureRunningUseCase(
    newContainerRemover(),
    newContainerStarter(newPathNormalizer()),
    newContainerChecker(),
  )

  compositionRoot = &_compositionRoot{
    ensureRunningUseCase:ensureRunningUseCase,
    getHostnameUseCase:newGetHostnameUseCase(),
  }

  return

}

type _compositionRoot struct {
  ensureRunningUseCase ensureRunningUseCase
  getHostnameUseCase   getHostnameUseCase
}

func (this _compositionRoot) EnsureRunningUseCase(
) ensureRunningUseCase {
  return this.ensureRunningUseCase
}

func (this _compositionRoot) GetHostnameUseCase(
) getHostnameUseCase {
  return this.getHostnameUseCase
}
