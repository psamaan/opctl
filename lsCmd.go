package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/engine-sdk-golang"
  "fmt"
  "os"
  "text/tabwriter"
  "net/url"
)

func lsCmd(
opctlCli *cli.Cli,
opctlEngineSdk opctlengine.Sdk,
) {

  opctlCli.Command("ls", "List ops", func(opLsCmd *cli.Cmd) {

    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 0, 8, 1, '\t', 0)

    opLsCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      var projectUrl *url.URL
      projectUrl, err = url.Parse(currentWorkDir)
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      fmt.Fprintln(w, "NAME\tDESCRIPTION")

      ops, err := opctlEngineSdk.ListOps(
        projectUrl,
      )
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      for _, op := range ops {

        fmt.Fprintf(w, "%v\t%v", op.Name, op.Description)
        fmt.Fprintln(w)

      }

      w.Flush()
    }

  })

}
