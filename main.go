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
  addOpCmd(dosCli, compositionRoot.DevOpSpecSdk())
  addSubOpCmd(dosCli, compositionRoot.DevOpSpecSdk())
  eventsCmd(dosCli, compositionRoot.DevOpSpecSdk())
  killCmd(dosCli, compositionRoot.DevOpSpecSdk())
  lsCmd(dosCli, compositionRoot.DevOpSpecSdk())
  opCmd(dosCli, compositionRoot.DevOpSpecSdk())
  runCmd(dosCli, compositionRoot.DevOpSpecSdk())
  dosCli.Version("v version", "opctl version 0.9.0")
  dosCli.Run(os.Args)

}
