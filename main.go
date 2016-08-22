package main

import (
  "os"
  "fmt"
)

func main() {

  defer func() {
    if panicArg := recover(); panicArg != nil {
      switch err := panicArg.(type) {
      case exitReq:
        fmt.Fprintln(os.Stderr, err.Message)
        os.Exit(err.Code)
      default:
        os.Exit(1)
      }
    }
  }()

  newCli(newCompositionRoot()).Run(os.Args)

}
