package main

import (
  mowCli "github.com/jawher/mow.cli"
  "os"
  "fmt"
)

func main() {

  compositionRoot, err := newCompositionRoot()
  if (nil != err) {
    fmt.Fprint(os.Stderr, err)
    os.Exit(1)
  }

  cli := mowCli.App("opctl", "control http://opspec.io compliant ops")
  addOpCmd(cli, compositionRoot.Sdk())
  addSubOpCmd(cli, compositionRoot.Sdk())
  eventsCmd(cli, compositionRoot.Sdk())
  killCmd(cli, compositionRoot.Sdk())
  lsCmd(cli, compositionRoot.Sdk())
  opCmd(cli, compositionRoot.Sdk())
  runCmd(cli, compositionRoot.Sdk())
  cli.Version("v version", "opctl version 0.1.0")
  cli.Run(os.Args)

}
