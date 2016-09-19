package core

//go:generate counterfeiter -o ./fakeCompositionRoot.go --fake-name fakeCompositionRoot ./ compositionRoot

import (
  "github.com/opspec-io/sdk-golang"
  "os"
)

type compositionRoot interface {
  CreateCollectionUseCase() createCollectionUseCase
  CreateOpUseCase() createOpUseCase
  KillOpRunUseCase() killOpRunUseCase
  ListOpsInCollectionUseCase() listOpsInCollectionUseCase
  RunOpUseCase() runOpUseCase
  SetCollectionDescriptionUseCase() setCollectionDescriptionUseCase
  SetOpDescriptionUseCase() setOpDescriptionUseCase
  StreamEventsUseCase() streamEventsUseCase
}

func newCompositionRoot(
) (compositionRoot compositionRoot) {

  exiter := newExiter()

  opspecSdk := opspec.New()
  workDirPathGetter := newWorkDirPathGetter()

  compositionRoot = &_compositionRoot{
    createCollectionUseCase:newCreateCollectionUseCase(opspecSdk, workDirPathGetter),
    createOpUseCase:newCreateOpUseCase(opspecSdk, workDirPathGetter),
    killOpRunUseCase:newKillOpRunUseCase(opspecSdk),
    listOpsInCollectionUseCase:newListOpsInCollectionUseCase(opspecSdk, workDirPathGetter, os.Stdout),
    runOpUseCase:newRunOpUseCase(exiter, opspecSdk, workDirPathGetter),
    setCollectionDescriptionUseCase:newSetCollectionDescriptionUseCase(opspecSdk, workDirPathGetter),
    setOpDescriptionUseCase:newSetOpDescriptionUseCase(opspecSdk, workDirPathGetter),
    streamEventsUseCase:newStreamEventsUseCase(exiter, opspecSdk),
  }

  return

}

type _compositionRoot struct {
  createCollectionUseCase         createCollectionUseCase
  createOpUseCase                 createOpUseCase
  killOpRunUseCase                killOpRunUseCase
  listOpsInCollectionUseCase      listOpsInCollectionUseCase
  runOpUseCase                    runOpUseCase
  setCollectionDescriptionUseCase setCollectionDescriptionUseCase
  setOpDescriptionUseCase         setOpDescriptionUseCase
  streamEventsUseCase             streamEventsUseCase
}

func (this _compositionRoot) CreateCollectionUseCase() createCollectionUseCase {
  return this.createCollectionUseCase
}

func (this _compositionRoot) CreateOpUseCase() createOpUseCase {
  return this.createOpUseCase
}

func (this _compositionRoot) KillOpRunUseCase() killOpRunUseCase {
  return this.killOpRunUseCase
}

func (this _compositionRoot) ListOpsInCollectionUseCase() listOpsInCollectionUseCase {
  return this.listOpsInCollectionUseCase
}

func (this _compositionRoot) RunOpUseCase() runOpUseCase {
  return this.runOpUseCase
}

func (this _compositionRoot) SetCollectionDescriptionUseCase() setCollectionDescriptionUseCase {
  return this.setCollectionDescriptionUseCase
}

func (this _compositionRoot) SetOpDescriptionUseCase() setOpDescriptionUseCase {
  return this.setOpDescriptionUseCase
}

func (this _compositionRoot) StreamEventsUseCase() streamEventsUseCase {
  return this.streamEventsUseCase
}
