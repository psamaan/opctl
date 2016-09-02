package core

//go:generate counterfeiter -o ./fakeApi.go --fake-name FakeApi ./ Api

type Api interface {
  CreateCollection(
  description string,
  name string,
  )

  CreateOp(
  description string,
  name string,
  )

  KillOpRun(
  opRunId string,
  )

  ListOpsInCollection(
  )

  RunOp(
  args []string,
  name string,
  )

  SetCollectionDescription(
  description string,
  )

  SetOpDescription(
  description string,
  name string,
  )

  StreamEvents(
  )
}

func New(
) Api {

  return &_api{
    compositionRoot:
    newCompositionRoot(),
  }

}

type _api struct {
  compositionRoot compositionRoot
}

func (this _api) CreateCollection(
description string,
name string,
) {
  this.
  compositionRoot.
    CreateCollectionUseCase().
    Execute(description, name)
}

func (this _api) CreateOp(
description string,
name string,
) {
  this.
  compositionRoot.
    CreateOpUseCase().
    Execute(description, name)
}

func (this _api) KillOpRun(
opRunId string,
) {
  this.
  compositionRoot.
    KillOpRunUseCase().
    Execute(opRunId)
}

func (this _api) ListOpsInCollection(
) {
  this.
  compositionRoot.
    ListOpsInCollectionUseCase().
    Execute()
}

func (this _api) RunOp(
args []string,
name string,
) {
  this.
  compositionRoot.
    RunOpUseCase().
    Execute(args, name)
}

func (this _api) SetCollectionDescription(
description string,
) {
  this.
  compositionRoot.
    SetCollectionDescriptionUseCase().
    Execute(description)
}

func (this _api) SetOpDescription(
description string,
name string,
) {
  this.
  compositionRoot.
    SetOpDescriptionUseCase().
    Execute(description, name)
}

func (this _api) StreamEvents(
) {
  this.
  compositionRoot.
    StreamEventsUseCase().
    Execute()
}
