package main

import (
  "os"
  "fmt"
  "github.com/opspec-io/cli/core"
  "github.com/opspec-io/sdk-golang/docker"
)

func main() {

  defer func() {
    if panicArg := recover(); panicArg != nil {
      switch err := panicArg.(type) {
      case core.ExitReq:
        fmt.Fprintln(os.Stderr, err.Message)
        os.Exit(err.Code)
      default:
        fmt.Fprintf(os.Stderr,"%v", err)
        os.Exit(1)
      }
    }
  }()

  newCli(newCompositionRoot(docker.New())).Run(os.Args)

}
