package main

import (
  "github.com/jawher/mow.cli"
  "github.com/opctl/sdk-for-golang/sdk"
  "fmt"
  "os"
  "text/tabwriter"
  "net/url"
)

func lsCmd(
dosCli *cli.Cli,
sdk sdk.Client,
) {

  dosCli.Command("ls", "List ops", func(opLsCmd *cli.Cmd) {

    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 0, 8, 1, '\t', 0)

    opLsCmd.Action = func() {

      currentWorkDir, err := os.Getwd()
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        cli.Exit(1)
      }

      var projectUrl *url.URL
      projectUrl, err = url.Parse(currentWorkDir)
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        cli.Exit(1)
      }

      fmt.Fprintln(w, "NAME\tDESCRIPTION")

      ops, err := sdk.ListOps(
        projectUrl,
      )
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        cli.Exit(1)
      }

      for _, op := range ops {

        fmt.Fprintf(w, "%v\t%v", op.Name, op.Description)
        fmt.Fprintln(w)

      }

      w.Flush()
    }

  })

}
