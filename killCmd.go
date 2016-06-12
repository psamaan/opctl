package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/engine-sdk-golang"
  "github.com/opctl/engine-sdk-golang/models"
)

func killCmd(
opctlCli *cli.Cli,
opctlEngineSdk opctlengine.Sdk,
) {

  opctlCli.Command("kill", "Kill an op run", func(runCmd *cli.Cmd) {

    runCmd.Spec = "OP_RUN_ID"

    var (
      opRunId = runCmd.StringArg("OP_RUN_ID", "", "the id of the op run to kill (must be a root op run)")
    )

    runCmd.Action = func() {

      opctlEngineSdk.KillOpRun(
        *models.NewKillOpRunReq(
          *opRunId,
        ),
      )

    }

  })

}
