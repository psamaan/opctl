package main

//go:generate counterfeiter -o ./fakeSetOpDescriptionUseCase.go --fake-name fakeSetOpDescriptionUseCase ./ setOpDescriptionUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
)

type setOpDescriptionUseCase interface {
  Execute(
  description string,
  name string,
  ) error
}

func newSetOpDescriptionUseCase(
opspecSdk opspec.Sdk,
workDirPathGetter workDirPathGetter,
) setOpDescriptionUseCase {
  return _setOpDescriptionUseCase{
    opspecSdk:opspecSdk,
    workDirPathGetter:workDirPathGetter,
  }
}

type _setOpDescriptionUseCase struct {
  opspecSdk         opspec.Sdk
  workDirPathGetter workDirPathGetter
}

func (this _setOpDescriptionUseCase) Execute(
description string,
name string,
) error {
  err := this.opspecSdk.SetOpDescription(
    *opspecModels.NewSetOpDescriptionReq(
      path.Join(this.workDirPathGetter.Get(), ".opspec", name),
      description,
    ),
  )
  if (nil != err) {
    return err
  }

  return nil
}
