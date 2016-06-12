package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opspec-io/sdk-golang"
)

func opSetCmd(
opCmd *cli.Cmd,
opspecSdk opspec.Sdk,
) {

  opCmd.Command(
    "set",
    "Set op attributes",
    func(opSetCmd *cli.Cmd) {
      opSetDescriptionCmd(opSetCmd, opspecSdk)
    },
  )

}
