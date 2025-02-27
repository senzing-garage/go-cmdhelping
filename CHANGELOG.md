# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], [markdownlint],
and this project adheres to [Semantic Versioning].

## [Unreleased]

## [0.3.2] - 2025-02-27

- Update dependencies

## [0.3.1] - 2024-09-09

- Update dependencies

## [0.3.0] - 2024-08-20

- Change from `g2` to `sz`/`er`

### Added in 0.3.0

## [0.2.3] - 2024-07-11

### Added in 0.2.3

- New Options:
  - AvoidServe
  - EngineInstanceName
  - EngineSettings
  - LicenseDaysLeft
  - LicenseRecordsPercent

### Deprecated in 0.2.3

- Deprecated Options:
  - EngineConfigurationJSON
  - EngineModuleName

## [0.2.2] - 2024-06-11

### Changed in 0.2.2

- From `engineconfigurationjson` to `options`

## [0.2.1] - 2024-04-19

### Changed in 0.2.1

- Migrated from `G2` to `Sz` prefixes
- Update dependencies
  - github.com/senzing-garage/go-helpers v0.5.1
  - github.com/stretchr/testify v1.9.0

## [0.2.0] - 2024-01-02

- Migrated from `github.com/senzing-garage/go-common` to `github.com/senzing-garage/go-helpers`

### Changed in 0.2.0

- Renamed module to `github.com/senzing-garage/go-cmdhelping`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-common v0.4.0
  - github.com/spf13/cobra v1.8.0
  - github.com/spf13/viper v1.18.2

## [0.1.9] - 2023-10-17

### Changed in 0.1.9

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-common v0.3.1
  - github.com/spf13/viper v1.17.0

## [0.1.8] - 2023-08-29

### Changed in 0.1.8

- Updated monitoring period default

## [0.1.7] - 2023-08-18

### Changed in 0.1.7

- In `go.mod` update to `go 1.21`
- Improved `engineconfiguration`

## [0.1.6] - 2023-08-16

### Changed in 0.1.6

- Updated option names InputURL, OutputURL, JSONOutput

## [0.1.5] - 2023-08-05

### Changed in 0.1.5

- Refactored to `template-go`
- Updated dependencies
  - github.com/senzing-garage/go-common v0.2.11

## [0.1.4] - 2023-08-02

### Changed in 0.1.4

- Fixed default value of `SENZING_TOOLS_SUPPORT_PATH`

## [0.1.3] - 2023-07-27

### Added to 0.1.3

- Swapped port numbers for serve-http, serve-grpc, and observe.

## [0.1.2] - 2023-07-26

### Added to 0.1.2

- Added options: `JsonOutput`

## [0.1.1] - 2023-07-21

### Added to 0.1.1

- Added options: `ErrorId`, `TtyOnly`.

## [0.1.0] - 2023-07-21

### Added to 0.1.0

- Initial functionality: cmdhelper, constant, and option packages.

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[markdownlint]: https://dlaa.me/markdownlint/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html
