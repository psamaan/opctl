[![Build Status](https://travis-ci.org/opctl/engine-sdk-golang.svg?branch=master)](https://travis-ci.org/opctl/engine-sdk-golang)
[![Coverage](https://codecov.io/gh/opctl/engine-sdk-golang/branch/master/graph/badge.svg)](https://codecov.io/gh/opctl/engine-sdk-golang)

Golang SDK for https://github.com/opctl/engine

*Be advised: this project is currently at Major version zero. Per the semantic versioning spec: 
"Major version zero (0.y.z) is for initial development. Anything may change at any time. The public API should not be considered stable."*

# Supported Use Cases
- get event stream
- get liveness
- kill op run
- run op

# Runtime Dependencies

The environment in which the sdk executes, must have the following available on its $Path for 
full functionality:

- [docker](https://github.com/docker/docker) >= 1.10

Note: if using Windows or OSX, you need to update your docker-machine to use NFS instead of vboxfs 
(or suffer painfully slow performance). One recommended way to achieve this is via 
[docker-machine-nfs](https://github.com/adlogix/docker-machine-nfs). 
Your mileage may vary.

# Versioning
This project adheres to the [Semantic Versioning](http://semver.org/) specification

# Contributing

see [CONTRIBUTING.md](CONTRIBUTING.md)

# Changelog
see [CHANGELOG.md](CHANGELOG.md)