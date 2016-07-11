package main

//go:generate counterfeiter -o ./fakeWorkDirPathGetter.go --fake-name fakeWorkDirPathGetter ./ workDirPathGetter

import "os"

type workDirPathGetter interface {
  Get() (workDirPath string)
}

func newWorkDirPathGetter() workDirPathGetter {
  return _workDirPathGetter{}
}

type _workDirPathGetter struct{}

func (this _workDirPathGetter) Get() (workDirPath string) {
  workDirPath, err := os.Getwd()
  if (err != nil) {
    panic(err)
  }
  return
}
