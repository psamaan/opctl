# Change Log
All notable changes to this project will be documented in this file.

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
