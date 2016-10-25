package core

//go:generate counterfeiter -o ./fakeListOpsInCollectionUseCase.go --fake-name fakeListOpsInCollectionUseCase ./ listOpsInCollectionUseCase

import (
  "path"
  "text/tabwriter"
  "fmt"
)

func (this _api) ListOpsInCollection(
) {
  _tabWriter := new(tabwriter.Writer)
  defer _tabWriter.Flush()
  _tabWriter.Init(this.writer, 0, 8, 1, '\t', 0)

  fmt.Fprintln(_tabWriter, "NAME\tDESCRIPTION")

  ops, err := this.bundle.GetCollection(
    path.Join(this.workDirPathGetter.Get(), ".opspec"),
  )
  if (nil != err) {
    this.exiter.Exit(ExitReq{Message:err.Error(), Code:1})
    return // support fake exiter
  }

  for _, op := range ops.Ops {

    fmt.Fprintf(_tabWriter, "%v\t%v", op.Name, op.Description)
    fmt.Fprintln(_tabWriter)

  }
}
