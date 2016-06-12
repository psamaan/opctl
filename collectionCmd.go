package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opspec-io/sdk-golang"
)

func collectionCmd(
opctlCli *cli.Cli,
opspecSdk opspec.Sdk,
) {

  opctlCli.Command(
    "collection",
    "Collection related actions",
    func(collectionCmd *cli.Cmd) {
      collectionSetCmd(collectionCmd, opspecSdk)
    },
  )

}
