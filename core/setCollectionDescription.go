package core

//go:generate counterfeiter -o ./fakeSetCollectionDescriptionUseCase.go --fake-name fakeSetCollectionDescriptionUseCase ./ setCollectionDescriptionUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang/models"
)

func (this _api) SetCollectionDescription(
description string,
) {
  err := this.bundle.SetCollectionDescription(
    models.SetCollectionDescriptionReq{
      PathToCollection:path.Join(this.workDirPathGetter.Get(), ".opspec"),
      Description:description,
    },
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }
}
