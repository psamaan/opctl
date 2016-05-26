#!/bin/sh

go get -t ./... && \
go build -a -o .tmp/opctl-${GOOS}-x86_64${OUTPUT_FILE_EXTENSION}
