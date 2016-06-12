package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/engine-sdk-golang"
  "github.com/opctl/engine-sdk-golang/models"
  "fmt"
  "os"
  "net/url"
)

func addOpCmd(
opctlCli *cli.Cli,
opctlEngineSdk opctlengine.Sdk,
) {

  opctlCli.Command("add-op", "Add an op", func(opAddCmd *cli.Cmd) {

    opAddCmd.Spec = "[--description] OP_NAME"

    var (
      name = opAddCmd.StringArg("OP_NAME", "", "name of the op")
      description = opAddCmd.StringOpt("description", "", "description of the op")
    )

    opAddCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      var projectUrl *url.URL
      projectUrl, err = url.Parse(currentWorkDir)
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      opctlEngineSdk.AddOp(
        *models.NewAddOpReq(
          projectUrl,
          *name,
          *description,
        ),
      )

    }

  })

}
