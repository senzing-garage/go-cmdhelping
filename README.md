# go-cmdhelping

## :warning: WARNING: go-cmdhelping is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`go-cmdhelping` contains utility packages for working with
[viper](https://github.com/spf13/viper) and
[cobra](https://github.com/spf13/cobra)

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/go-cmdhelping.svg)](https://pkg.go.dev/github.com/senzing/go-cmdhelping)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/go-cmdhelping)](https://goreportcard.com/report/github.com/senzing/go-cmdhelping)
[![go-test.yaml](https://github.com/Senzing/go-cmdhelping/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/go-cmdhelping/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/go-cmdhelping/blob/main/LICENSE)

## Overview

`go-cmdhelping` is an opinionated use of
[viper](https://github.com/spf13/viper) and
[cobra](https://github.com/spf13/cobra).

Context variables are specified in a list of
[option.ContextVariable](option/option.go).

The `[]option.ContextVariable` list is processed by functions in the
[cmdhelper](cmdhelper) package.

## Use

See
[main.go](main.go)
for an example of use.

## References

- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
