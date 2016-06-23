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
  collectionCmd(cli, compositionRoot.OpSpecSdk())
  eventsCmd(cli, compositionRoot.OpCtlEngineSdk())
  killCmd(cli, compositionRoot.OpCtlEngineSdk())
  lsCmd(cli, compositionRoot.OpSpecSdk())
  opCmd(cli, compositionRoot.OpSpecSdk())
  runCmd(cli, compositionRoot.OpCtlEngineSdk(), compositionRoot.OpSpecSdk())
  cli.Version("v version", "opctl version 0.1.3")
  cli.Run(os.Args)

}
