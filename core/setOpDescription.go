package core

//go:generate counterfeiter -o ./fakeSetOpDescriptionUseCase.go --fake-name fakeSetOpDescriptionUseCase ./ setOpDescriptionUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang/models"
)

func (this _api) SetOpDescription(
description string,
name string,
) {
  err := this.bundle.SetOpDescription(
    models.SetOpDescriptionReq{
      PathToOp:path.Join(this.workDirPathGetter.Get(), ".opspec", name),
      Description:description,
    },
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }
}
