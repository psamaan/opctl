package main

//go:generate counterfeiter -o ./fakeCli.go --fake-name FakeCli ./ cli

import (
  mow "github.com/jawher/mow.cli"
)

const cliVersion = "0.1.9"

type cli interface {
  Run(args []string) error
}

func newCli(
compositionRoot compositionRoot,
) cli {

  cli := mow.App("opctl", "control http://opspec.io compliant ops")
  cli.Version("v version", cliVersion)

  cli.Command("collection", "Collection related actions", func(collectionCmd *mow.Cmd) {

    collectionCmd.Command("create", "Create a collection", func(createCmd *mow.Cmd) {

      name := createCmd.StringArg("NAME", "", "name of the collection")
      description := createCmd.StringOpt("d description", "", "description of the collection")

      createCmd.Action = func() {
        compositionRoot.CreateCollectionUseCase().Execute(*description, *name)
      }

    })

    collectionCmd.Command("set", "Set collection attributes", func(setCmd *mow.Cmd) {
      setCmd.Command("description", "Set the description of a collection", func(descriptionCmd *mow.Cmd) {
        description := descriptionCmd.StringArg("DESCRIPTION", "", "description of the collection")

        descriptionCmd.Action = func() {
          compositionRoot.SetCollectionDescriptionUseCase().Execute(*description)
        }
      })
    })
  })

  cli.Command("events", "Stream events from the engine", func(eventsCmd *mow.Cmd) {
    eventsCmd.Action = func() {
      compositionRoot.StreamEventsUseCase().Execute()
    }
  })

  cli.Command("kill", "Kill an op run", func(killCmd *mow.Cmd) {
    opRunId := killCmd.StringArg("OP_RUN_ID", "", "the id of the op run to kill (must be a root op run)")

    killCmd.Action = func() {
      compositionRoot.KillOpRunUseCase().Execute(*opRunId)
    }
  })

  cli.Command("ls", "List ops in a collection", func(lsCmd *mow.Cmd) {
    lsCmd.Action = func() {
      compositionRoot.ListOpsInCollectionUseCase().Execute()
    }
  })

  cli.Command("op", "Op related actions", func(opCmd *mow.Cmd) {

    opCmd.Command("create", "Create an op", func(createCmd *mow.Cmd) {

      description := createCmd.StringOpt("d description", "", "description of the op")
      name := createCmd.StringArg("NAME", "", "name of the op")

      createCmd.Action = func() {
        compositionRoot.CreateOpUseCase().Execute(*description, *name)
      }

    })

    opCmd.Command("set", "Set op attributes", func(setCmd *mow.Cmd) {
      setCmd.Command("description", "Set the description of an op", func(descriptionCmd *mow.Cmd) {
        description := descriptionCmd.StringArg("DESCRIPTION", "", "description of the op")
        name := descriptionCmd.StringArg("NAME", "", "name of the op")

        descriptionCmd.Action = func() {
          compositionRoot.SetOpDescriptionUseCase().Execute(*description, *name)
        }
      })
    })
  })

  cli.Command("run", "Run an op", func(runCmd *mow.Cmd) {
    args := runCmd.StringsOpt("a", []string{}, "Pass args to op in format: NAME[=VALUE] (gets from env if not provided)")
    name := runCmd.StringArg("OP_URL", "", "url of the op (op name if in collection)")

    runCmd.Action = func() {
      compositionRoot.RunOpUseCase().Execute(*args, *name)
    }
  })

  return cli

}
