package main

//go:generate counterfeiter -o ./fakeCreateOpUseCase.go --fake-name fakeCreateOpUseCase ./ createOpUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
)

type createOpUseCase interface {
  Execute(
  description string,
  name string,
  ) error
}

func newCreateOpUseCase(
opspecSdk opspec.Sdk,
workDirPathGetter workDirPathGetter,
) createOpUseCase {
  return _createOpUseCase{
    opspecSdk:opspecSdk,
    workDirPathGetter:workDirPathGetter,
  }
}

type _createOpUseCase struct {
  opspecSdk         opspec.Sdk
  workDirPathGetter workDirPathGetter
}

func (this _createOpUseCase) Execute(
description string,
name string,
) error {
  err := this.opspecSdk.CreateOp(
    *opspecModels.NewCreateOpReq(
      path.Join(this.workDirPathGetter.Get(), ".opspec", name),
      name,
      description,
    ),
  )
  if (nil != err) {
    return err
  }

  return nil
}
