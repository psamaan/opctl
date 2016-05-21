#!/bin/sh

cat cli.coverprofile > coverage.txt && \
curl -s https://codecov.io/bash | bash -s
