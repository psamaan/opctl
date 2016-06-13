package main

import (
  "github.com/jawher/mow.cli"
  "fmt"
  "os"
  "github.com/opspec-io/sdk-golang"
  "github.com/opspec-io/sdk-golang/models"
  "path"
)

func opCreateCmd(
opCmd *cli.Cmd,
opspecSdk opspec.Sdk,
) {

  opCmd.Command("create", "Create an op", func(opCreateCmd *cli.Cmd) {

    opCreateCmd.Spec = "[--description] OP_NAME"

    var (
      name = opCreateCmd.StringArg("OP_NAME", "", "name of the op")
      description = opCreateCmd.StringOpt("description", "", "description of the op")
    )

    opCreateCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      opspecSdk.CreateOp(
        *models.NewCreateOpReq(
          path.Join(currentWorkDir, ".opspec", *name),
          *name,
          *description,
        ),
      )

    }

  })

}
