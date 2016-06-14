package main

import (
  "github.com/jawher/mow.cli"
  "fmt"
  "os"
  "text/tabwriter"
  "github.com/opspec-io/sdk-golang"
  "path"
)

func lsCmd(
opctlCli *cli.Cli,
opspecSdk opspec.Sdk,
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

      fmt.Fprintln(w, "NAME\tDESCRIPTION")

      ops, err := opspecSdk.GetCollection(
        path.Join(currentWorkDir, ".opspec"),
      )
      if (nil != err) {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
      }

      for _, op := range ops.Ops {

        fmt.Fprintf(w, "%v\t%v", op.Name, op.Description)
        fmt.Fprintln(w)

      }

      w.Flush()
    }

  })

}
