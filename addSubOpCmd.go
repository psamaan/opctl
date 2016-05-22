package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/sdk-for-golang/sdk"
  "github.com/opctl/sdk-for-golang/sdk/models"
  "fmt"
  "os"
  "net/url"
)

func addSubOpCmd(
dosCli *cli.Cli,
sdk sdk.Client,
) {

  dosCli.Command("add-sub-op", "Add a sub op", func(opAddSubOpCmd *cli.Cmd) {

    opAddSubOpCmd.Spec = "[--preceding-sub-op-url] SUB_OP_URL OP_NAME"

    precedingSubOpUrl :=
    opAddSubOpCmd.String(
      cli.StringOpt{
        Name: "preceding-sub-op-url",
        Value: "",
        Desc: "the url of an existing sub-op which will precede this one",
        HideValue: true,
      })

    subOpUrl :=
    opAddSubOpCmd.StringArg(
      "SUB_OP_URL",
      "",
      "the url of the sub-op (must match an op name)",
    )

    opName :=
    opAddSubOpCmd.StringArg(
      "OP_NAME",
      "",
      "the name of the op to add the sub-op to",
    )

    opAddSubOpCmd.Action = func() {

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

      sdk.AddSubOp(
        *models.NewAddSubOpReq(
          projectUrl,
          *opName,
          *subOpUrl,
          *precedingSubOpUrl,
        ),
      )
    }

  })
}
