package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opspec-io/sdk-golang"
)

func collectionSetCmd(
collectionCmd *cli.Cmd,
opspecSdk opspec.Sdk,
) {

  collectionCmd.Command(
    "set",
    "Set collection attributes",
    func(collectionSetCmd *cli.Cmd) {
      collectionSetDescriptionCmd(collectionSetCmd, opspecSdk)
    },
  )

}
