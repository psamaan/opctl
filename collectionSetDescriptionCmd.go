package main

import (
  "github.com/jawher/mow.cli"
  opspecModels "github.com/opspec-io/sdk-golang/models"
  "github.com/opspec-io/sdk-golang"
  "fmt"
  "os"
  "path"
)

func collectionSetDescriptionCmd(
collectionSetCmd *cli.Cmd,
opspecSdk opspec.Sdk,
) {

  collectionSetCmd.Command(
    "description",
    "Set the description of a collection",
    func(collectionSetDescriptionCmd *cli.Cmd) {

      collectionSetDescriptionCmd.Spec = "COLLECTION_DESCRIPTION"

      var (
        description = collectionSetDescriptionCmd.StringArg("COLLECTION_DESCRIPTION", "", "description of the collection")
      )

      collectionSetDescriptionCmd.Action = func() {

        currentWorkDir, err := os.Getwd()
        if (nil != err) {
          fmt.Fprintln(os.Stderr, err)
          os.Exit(1)
        }

        err = opspecSdk.SetCollectionDescription(
          *opspecModels.NewSetCollectionDescriptionReq(
            path.Join(currentWorkDir, ".opspec"),
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
