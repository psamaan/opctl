package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/sdk-for-golang/sdk"
  "github.com/opctl/sdk-for-golang/sdk/models"
)

func killCmd(
dosCli *cli.Cli,
sdk sdk.Client,
) {

  dosCli.Command("kill", "Kill an op run", func(runCmd *cli.Cmd) {

    runCmd.Spec = "OP_RUN_ID"

    var (
      opRunId = runCmd.StringArg("OP_RUN_ID", "", "the id of the op run to kill (must be a root op run)")
    )

    runCmd.Action = func() {

      sdk.KillOpRun(
        *models.NewKillOpRunReq(
          *opRunId,
        ),
      )

    }

  })

}
