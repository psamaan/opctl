# Change Log

All notable changes to this project will be documented in this file.

## 0.1.12 - 2016-10-03

### Changed

- bumped [opspec sdk](https://github.com/opspec-io/sdk-golang) version
- event timestamps now output in RFC3339 format

### Fixed

- [Emitted ContainerStd*WrittenToEvent.Data Incomplete](https://github.com/opspec-io/engine/issues/32)

## 0.1.11 - 2016-09-09

### Added

- support for [opspec 0.1.2](https://opspec.io)

### Fixed

- [failure of serial operation run does not immediately fail all following operations](https://github.com/opspec-io/cli/issues/5)

### Removed

- support for < [opspec 0.1.2](https://opspec.io)

## 0.1.10 - 2016-09-02

### Fixed

- [opctl does not wait for parallel op containers to die before returning](https://github.com/opspec-io/cli/issues/8)
- [Many parallel ops crash engine](https://github.com/opspec-io/engine/issues/17)

## 0.1.9 - 2016-08-21

### Added

- prompt for op args if not provided or in environment

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

- [Support new opspec subop `isParallel` flag](https://github.com/opspec-io/engine/issues/11)

## 0.1.3 - 2016-06-22

### Changed

- updated to opspec/engine:0.1.2 (see
  [CHANGELOG.md](https://github.com/opspec-io/engine/tree/0.1.2/CHANGELOG.md))

## 0.1.2 - 2016-06-22

### Changed

- updated to opspec/engine:0.1.1 (see
  [CHANGELOG.md](https://github.com/opspec-io/engine/tree/0.1.1/CHANGELOG.md))

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
- refactor to use [opspec sdk](https://github.com/opspec-io/sdk-golang)
  for use cases it supports (rather than re-invent the wheel)

### Removed

- `add-op` sub command
- `op set-description` sub command

### Fixed

- engine API errors not handled

## 0.1.0 - 2016-05-27

