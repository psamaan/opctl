#!/bin/sh -e

go get -t ./...
go get github.com/onsi/ginkgo/ginkgo
/golang/bin/ginkgo -r -cover
