package main

import (
  "github.com/jawher/mow.cli"
  "fmt"
  "os"
  "github.com/opspec-io/sdk-golang"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "path"
)

func opSetDescriptionCmd(
opSetCmd *cli.Cmd,
opspecSdk opspec.Sdk,
) {

  opSetCmd.Command(
    "description",
    "Set the description of an op",
    func(opSetDescriptionCmd *cli.Cmd) {

      opSetDescriptionCmd.Spec = "OP_DESCRIPTION OP_NAME"

      var (
        name = opSetDescriptionCmd.StringArg("OP_NAME", "", "the name of the op")
        description = opSetDescriptionCmd.StringArg("OP_DESCRIPTION", "", "description of the op")
      )

      opSetDescriptionCmd.Action = func() {

        currentWorkDir, err := os.Getwd()
        if (nil != err) {
          fmt.Fprintln(os.Stderr, err)
          os.Exit(1)
        }

        err = opspecSdk.SetOpDescription(
          *opspecModels.NewSetOpDescriptionReq(
            path.Join(currentWorkDir, ".opspec", *name),
            *description,
          ),
        )
        if (nil != err) {
          fmt.Fprintln(os.Stderr, err)
          os.Exit(1)
        }

      }

    },
  )
}
