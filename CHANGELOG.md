# Change Log
All notable changes to this project will be documented in this file.
## 0.1.8 - 2016-08-08
### Added
- support for private docker images

## 0.1.7 - 2016-08-02
### Added
- support for [opspec 0.1.1](https://opspec.io)

### Removed
- support for [opspec 0.1.0](https://opspec.io)

## 0.1.6 - 2016-07-20
### Fixed
- windows path handling

## 0.1.5 - 2016-07-15
### Fixed
- [regression] exit code not representative of operation outcome

## 0.1.4 - 2016-07-09
### Added
- [Support new opspec subop `isParallel` flag](https://github.com/opctl/engine/issues/11)

## 0.1.3 - 2016-06-22
### Changed
- updated to opctl/engine:0.1.2 (see [CHANGELOG.md](https://github.com/opctl/engine/tree/0.1.2/CHANGELOG.md))

## 0.1.2 - 2016-06-22
### Changed
- updated to opctl/engine:0.1.1 (see [CHANGELOG.md](https://github.com/opctl/engine/tree/0.1.1/CHANGELOG.md))

### Removed
- `add-sub-op` sub command

## 0.1.1 - 2016-06-16
### Added
- `collection` sub command
- `collection set description` sub command
- `op set description` sub command
- `op create` sub command
- made `opctl run` autopopulate op params from env
- better error for docker not running
- refactor to use [opspec sdk](https://github.com/opspec-io/sdk-golang) for use cases it supports (rather than re-invent the wheel)

### Removed
- `add-op` sub command
- `op set-description` sub command

### Fixed
- engine API errors not handled

## 0.1.0 - 2016-05-27
