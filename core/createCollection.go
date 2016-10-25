package core

//go:generate counterfeiter -o ./fakeCreateCollectionUseCase.go --fake-name fakeCreateCollectionUseCase ./ createCollectionUseCase

import (
  "path"
   "github.com/opspec-io/sdk-golang/models"
)

func (this _api) CreateCollection(
description string,
name string,
) {
  err := this.bundle.CreateCollection(
    *models.NewCreateCollectionReq(
      path.Join(this.workDirPathGetter.Get(), name),
      name,
      description,
    ),
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }
}
