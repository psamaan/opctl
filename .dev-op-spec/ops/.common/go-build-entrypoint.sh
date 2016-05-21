#!/bin/sh

go get -t ./... && \
go build -a -o ${PATH_TO_TMP_DIR}/opctl-${GOOS}-x86_64${OUTPUT_FILE_EXTENSION}
