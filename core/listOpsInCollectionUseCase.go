package core

//go:generate counterfeiter -o ./fakeListOpsInCollectionUseCase.go --fake-name fakeListOpsInCollectionUseCase ./ listOpsInCollectionUseCase

import (
  "path"
  "github.com/opspec-io/sdk-golang"
  "text/tabwriter"
  "fmt"
  "io"
)

type listOpsInCollectionUseCase interface {
  Execute(
  ) error
}

func newListOpsInCollectionUseCase(
opspecSdk opspec.Sdk,
workDirPathGetter workDirPathGetter,
writer io.Writer,
) listOpsInCollectionUseCase {

  return _listOpsInCollectionUseCase{
    opspecSdk:opspecSdk,
    workDirPathGetter:workDirPathGetter,
    writer:writer,
  }
}

type _listOpsInCollectionUseCase struct {
  opspecSdk         opspec.Sdk
  workDirPathGetter workDirPathGetter
  writer            io.Writer
}

func (this _listOpsInCollectionUseCase) Execute(
) error {
  _tabWriter := new(tabwriter.Writer)
  defer _tabWriter.Flush()
  _tabWriter.Init(this.writer, 0, 8, 1, '\t', 0)

  fmt.Fprintln(_tabWriter, "NAME\tDESCRIPTION")

  ops, err := this.opspecSdk.GetCollection(
    path.Join(this.workDirPathGetter.Get(), ".opspec"),
  )
  if (nil != err) {
    return err
  }

  for _, op := range ops.Ops {

    fmt.Fprintf(_tabWriter, "%v\t%v", op.Name, op.Description)
    fmt.Fprintln(_tabWriter)

  }

  return nil
}
