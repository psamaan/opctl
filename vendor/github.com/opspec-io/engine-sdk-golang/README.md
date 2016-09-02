[![Build Status](https://travis-ci.org/opspec-io/engine-sdk-golang.svg?branch=master)](https://travis-ci.org/opspec-io/engine-sdk-golang)
[![Coverage](https://codecov.io/gh/opspec-io/engine-sdk-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/opspec-io/engine-sdk-golang)

Golang SDK for https://github.com/opspec-io/engine

*Be advised: this project is currently at Major version zero. Per the
semantic versioning spec: "Major version zero (0.y.z) is for initial
development. Anything may change at any time. The public API should not
be considered stable."*

# Supported Use Cases

- get event stream
- get liveness
- kill op run
- run op

# Runtime Dependencies

The environment in which the sdk executes, must have the following
available on its $Path for full functionality:

- [docker](https://github.com/docker/docker) >= 1.10

Note: if using Windows or OSX, you need to update your docker-machine to
use NFS instead of vboxfs (or suffer painfully slow performance). One
recommended way to achieve this is via
[docker-machine-nfs](https://github.com/adlogix/docker-machine-nfs).
Your mileage may vary.

# Releases

All releases will be
[tagged](https://github.com/opspec-io/engine-sdk-golang/tags) and made
available on the
[releases](https://github.com/opspec-io/engine-sdk-golang/releases) page
with links to the corresponding version of the
[CHANGELOG.md](CHANGELOG.md) doc.

# Versioning

This project adheres to the [Semantic Versioning](http://semver.org/)
specification

# Contributing

see [CONTRIBUTING.md](CONTRIBUTING.md)

# Changelog

see [CHANGELOG.md](CHANGELOG.md)
