package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opspec-io/sdk-golang"
)

func opCmd(
opctlCli *cli.Cli,
opspecSdk opspec.Sdk,
) {

  opctlCli.Command(
    "op",
    "Op related actions",
    func(opCmd *cli.Cmd) {
      opCreateCmd(opCmd, opspecSdk)
      opSetCmd(opCmd, opspecSdk)
    },
  )

}
