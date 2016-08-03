package opspec

//go:generate counterfeiter -o ./fakeOpViewFactory.go --fake-name fakeOpViewFactory ./ opViewFactory

import (
  "github.com/opspec-io/sdk-golang/models"
  "path"
  "errors"
)

type opViewFactory interface {
  Construct(
  opBundlePath string,
  ) (
  opView models.OpView,
  err error,
  )
}

func newOpViewFactory(
filesystem Filesystem,
yamlCodec yamlCodec,
) opViewFactory {

  return &_opViewFactory{
    filesystem:filesystem,
    yamlCodec:yamlCodec,
  }

}

type _opViewFactory struct {
  filesystem Filesystem
  yamlCodec  yamlCodec
}

func (this _opViewFactory) Construct(
opBundlePath string,
) (
opView models.OpView,
err error,
) {

  opFilePath := path.Join(opBundlePath, NameOfOpFile)

  opFileBytes, err := this.filesystem.GetBytesOfFile(
    opFilePath,
  )
  if (nil != err) {
    return
  }

  opFile := models.OpFile{}
  err = this.yamlCodec.FromYaml(
    opFileBytes,
    &opFile,
  )
  if (nil != err) {
    return
  }

  var run models.RunInstruction
  if (nil != opFile.Run.Container) {
    run = models.NewContainerRunInstruction(opFile.Run.Container)
  } else if (len(opFile.Run.SubOps) > 0) {
    run = models.NewSubOpsRunInstruction(opFile.Run.SubOps)
  } else {
    err = errors.New("run not set or invalid")
    return
  }

  opView = *models.NewOpView(
    opFile.Description,
    opFile.Inputs,
    opFile.Name,
    opFile.Outputs,
    run,
    opFile.Version,
  )

  return

}

