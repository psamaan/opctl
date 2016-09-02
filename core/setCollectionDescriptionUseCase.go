package core

//go:generate counterfeiter -o ./fakeSetCollectionDescriptionUseCase.go --fake-name fakeSetCollectionDescriptionUseCase ./ setCollectionDescriptionUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
)

type setCollectionDescriptionUseCase interface {
  Execute(
  description string,
  ) error
}

func newSetCollectionDescriptionUseCase(
opspecSdk opspec.Sdk,
workDirPathGetter workDirPathGetter,
) setCollectionDescriptionUseCase {
  return _setCollectionDescriptionUseCase{
    opspecSdk:opspecSdk,
    workDirPathGetter:workDirPathGetter,
  }
}

type _setCollectionDescriptionUseCase struct {
  opspecSdk         opspec.Sdk
  workDirPathGetter workDirPathGetter
}

func (this _setCollectionDescriptionUseCase) Execute(
description string,
) error {
  err := this.opspecSdk.SetCollectionDescription(
    *opspecModels.NewSetCollectionDescriptionReq(
      path.Join(this.workDirPathGetter.Get(), ".opspec"),
      description,
    ),
  )
  if (nil != err) {
    return err
  }

  return nil
}
