# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`go-cmdhelping` is a Go library providing opinionated helpers for building CLIs with [cobra](https://github.com/spf13/cobra) and [viper](https://github.com/spf13/viper). It unifies configuration from command-line arguments, environment variables, and config files through a single `ContextVariable` abstraction.

Part of [Senzing Garage](https://github.com/senzing-garage) - experimental projects, not production-ready.

## Build & Development Commands

```bash
make build                      # Build binaries (output: target/<os>-<arch>/)
make test                       # Run all tests with gotestfmt
make lint                       # Run golangci-lint, govulncheck, cspell
make coverage                   # Run tests with coverage report
make clean                      # Clean build artifacts
make dependencies               # Update Go dependencies
make dependencies-for-development  # Install dev tools (golangci-lint, gotestfmt, etc.)
make fix                        # Auto-fix linter issues (gofumpt, wsl, etc.)
```

Run a single test:

```bash
go test -v -run TestFunctionName ./package/...
```

Run tests for a specific package:

```bash
go test -v ./cmdhelper/...
```

## Architecture

### Core Abstraction: ContextVariable

The `option.ContextVariable` struct (defined in [option/main.go](option/main.go)) ties together:

- `Arg`: CLI flag name (kebab-case, e.g., `database-url`)
- `Envar`: Environment variable name (e.g., `SENZING_TOOLS_DATABASE_URL`)
- `Default`: Default value
- `Help`: Help text template (uses `%s` for env var name)
- `Type`: One of `optiontype.Bool`, `Int`, `String`, `StringSlice`, `Uint`, `Uint32`, `Uint64`

### Package Structure

- **option/**: `ContextVariable` definition and pre-defined Senzing-specific options
  - [option.go](option/option.go): Pre-built options (DatabaseURL, HTTPPort, LogLevel, etc.)
  - [optiontype/](option/optiontype/): Type enum for supported flag types

- **cmdhelper/**: Cobra/Viper integration functions
  - `Init(cobraCommand, contextVariables)`: Registers cobra flags for all variables
  - `PreRun(cobraCommand, args, configName, contextVariables)`: Loads config files, binds viper to flags
  - `Version(buildTag, buildIteration)`: Creates version string

- **settings/**: Builds Senzing engine configuration JSON from viper values

- **constant/**: Common constants (`SetEnvPrefix = "SENZING_TOOLS"`, version template)

### Usage Pattern

```go
// Define which options your CLI uses
contextVariables := []option.ContextVariable{option.DatabaseURL, option.HTTPPort, option.LogLevel}

// Initialize cobra flags
cmdhelper.Init(cobraCommand, contextVariables)

// In PreRun, load config and bind viper
cmdhelper.PreRun(cobraCommand, os.Args, "myapp", contextVariables)

// Access values via viper
url := viper.GetString(option.DatabaseURL.Arg)
```

### Configuration Resolution Order

1. Command-line flags (highest priority)
2. Environment variables (SENZING*TOOLS*\* prefix)
3. Config file (searched in: `~/.senzing-tools/`, `~/`, `/etc/senzing-tools/`)
4. Default values

## Code Style

- Uses extensive golangci-lint configuration ([.github/linters/.golangci.yaml](.github/linters/.golangci.yaml))
- Max line length: 120 characters (enforced by golines)
- Formatting: gofumpt (stricter gofmt)
- Run `make fix` to auto-fix most style issues
