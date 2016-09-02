package main

import (
  "os"
  "fmt"
  "github.com/opspec-io/cli/core"
)

func main() {

  defer func() {
    if panicArg := recover(); panicArg != nil {
      switch err := panicArg.(type) {
      case core.ExitReq:
        fmt.Fprintln(os.Stderr, err.Message)
        os.Exit(err.Code)
      default:
        os.Exit(1)
      }
    }
  }()

  newCli(newCompositionRoot()).Run(os.Args)

}
