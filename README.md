[![Build Status](https://travis-ci.org/opspec-io/opctl.svg?branch=master)](https://travis-ci.org/opspec-io/opctl)
[![Coverage](https://codecov.io/gh/opspec-io/opctl/branch/master/graph/badge.svg)](https://codecov.io/gh/opspec-io/opctl)

CLI for controlling http://opspec.io compliant ops.

> *Be advised: this project is currently at Major version zero. Per the
> semantic versioning spec: "Major version zero (0.y.z) is for initial
> development. Anything may change at any time. The public API should
> not be considered stable."*

# Usage

for usage guidance simply execute without any arguments:

```SHELL
opctl

Usage: opctl [OPTIONS] COMMAND [arg...]

control http://opspec.io compliant ops

Options:
  -v, --version    Show the version and exit

Commands:
  collection   Collection related actions
  events       Get real time events from the server
  kill         Kill an op run
  ls           List ops
  op           Op related actions
  run          Run an op

Run 'opctl COMMAND --help' for more information on a command.
```

# Supported Use Cases

- get cli version
- create collection
- create op
- kill op run
- list ops in collection
- run op
- set collection description
- set op description
- stream events

# Releases

All releases will be [tagged](https://github.com/opspec-io/opctl/tags) and
made available on the
[releases](https://github.com/opspec-io/opctl/releases)
[releases](https://github.com/opspec-io/sdk-golang/releases) page with
links to docs.

# Versioning

This project adheres to the [Semantic Versioning](http://semver.org/)
specification

# Installation

see [INSTALLATION.md](INSTALLATION.md)

# Contributing

see [CONTRIBUTING.md](CONTRIBUTING.md)

# Changelog

see [CHANGELOG.md](CHANGELOG.md)
