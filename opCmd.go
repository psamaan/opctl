package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/sdk-for-golang/sdk"
  "github.com/opctl/sdk-for-golang/sdk/models"
  "fmt"
  "os"
  "net/url"
)

func opCmd(
dosCli *cli.Cli,
devOpSpecSdk sdk.Client,
) {

  dosCli.Command("op", "Op-related operations", func(opCmd *cli.Cmd) {
    opSetDescriptionCmd(opCmd, devOpSpecSdk)
  })

}

func opSetDescriptionCmd(
opCmd *cli.Cmd,
devOpSpecSdk sdk.Client,
) {

  opCmd.Command("set-description", "Set the description of an op", func(opSetDescriptionCmd *cli.Cmd) {

    opSetDescriptionCmd.Spec = "OP_DESCRIPTION OP_NAME"

    var (
      name = opSetDescriptionCmd.StringArg("OP_NAME", "", "the name of the op")
      description = opSetDescriptionCmd.StringArg("OP_DESCRIPTION", "", "description of the op")
    )

    opSetDescriptionCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        cli.Exit(1)
      }

      var projectUrl *url.URL
      projectUrl, err = url.Parse(currentWorkDir)
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        cli.Exit(1)
      }

      devOpSpecSdk.SetDescriptionOfOp(
        *models.NewSetDescriptionOfOpReq(
          projectUrl,
          *description,
          *name,
        ),
      )
    }

  })
}
