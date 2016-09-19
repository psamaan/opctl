package core

//go:generate counterfeiter -o ./fakeCreateCollectionUseCase.go --fake-name fakeCreateCollectionUseCase ./ createCollectionUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang"
  models "github.com/opspec-io/sdk-golang/models"
)

type createCollectionUseCase interface {
  Execute(
  description string,
  name string,
  ) error
}

func newCreateCollectionUseCase(
opspecSdk opspec.Sdk,
workDirPathGetter workDirPathGetter,
) createCollectionUseCase {
  return _createCollectionUseCase{
    opspecSdk:opspecSdk,
    workDirPathGetter:workDirPathGetter,
  }
}

type _createCollectionUseCase struct {
  opspecSdk         opspec.Sdk
  workDirPathGetter workDirPathGetter
}

func (this _createCollectionUseCase) Execute(
description string,
name string,
) error {
  err := this.opspecSdk.CreateCollection(
    *models.NewCreateCollectionReq(
      path.Join(this.workDirPathGetter.Get(), name),
      name,
      description,
    ),
  )
  if (nil != err) {
    return err
  }

  return nil
}
