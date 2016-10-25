package core

//go:generate counterfeiter -o ./fakeCreateOpUseCase.go --fake-name fakeCreateOpUseCase ./ createOpUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang/models"
)

func (this _api) CreateOp(
description string,
name string,
) {
  err := this.bundle.CreateOp(
    *models.NewCreateOpReq(
      path.Join(this.workDirPathGetter.Get(), ".opspec", name),
      name,
      description,
    ),
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }
}
