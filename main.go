package main

import (
  "github.com/jawher/mow.cli"
  "os"
  "fmt"
)

func main() {

  compositionRoot, err := newCompositionRoot()
  if (nil != err) {
    fmt.Fprint(os.Stderr, err)
    os.Exit(1)
  }

  dosCli := cli.App("opctl", "control http://opendevops.io/ compliant ops")
  addOpCmd(dosCli, compositionRoot.Sdk())
  addSubOpCmd(dosCli, compositionRoot.Sdk())
  eventsCmd(dosCli, compositionRoot.Sdk())
  killCmd(dosCli, compositionRoot.Sdk())
  lsCmd(dosCli, compositionRoot.Sdk())
  opCmd(dosCli, compositionRoot.Sdk())
  runCmd(dosCli, compositionRoot.Sdk())
  dosCli.Version("v version", "opctl version 0.9.0")
  dosCli.Run(os.Args)

}
