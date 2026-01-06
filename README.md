# go-cmdhelping

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage] where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: go-cmdhelping is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`go-cmdhelping` contains utility packages for working with [viper] and [cobra].

[![Go Reference Badge]][Package reference]
[![Go Report Card Badge]][Go Report Card]
[![License Badge]][License]
[![go-test-linux.yaml Badge]][go-test-linux.yaml]
[![go-test-darwin.yaml Badge]][go-test-darwin.yaml]
[![go-test-windows.yaml Badge]][go-test-windows.yaml]

[![golangci-lint.yaml Badge]][golangci-lint.yaml]

## Overview

`go-cmdhelping` is an opinionated use of [viper] and [cobra].

Context variables are specified in a list of [option.ContextVariable].

The `[]option.ContextVariable` list is processed by functions
in the [cmdhelper] package.

## Use

See [main.go] for an example of use.

## References

1. [SDK documentation]
1. [Development]
1. [Errors]
1. [Examples]

[SDK documentation]: https://pkg.go.dev/github.com/senzing-garage/go-cmdhelping
[cmdhelper]: cmdhelper
[cobra]: https://github.com/spf13/cobra
[Development]: docs/development.md
[Errors]: docs/errors.md
[Examples]: docs/examples.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/go-cmdhelping.svg
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/senzing-garage/go-cmdhelping
[Go Report Card]: https://goreportcard.com/report/github.com/senzing-garage/go-cmdhelping
[go-test-darwin.yaml Badge]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/go-test-darwin.yaml/badge.svg
[go-test-darwin.yaml]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/go-test-darwin.yaml
[go-test-linux.yaml Badge]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/go-test-linux.yaml/badge.svg
[go-test-linux.yaml]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/go-test-linux.yaml
[go-test-windows.yaml Badge]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/go-test-windows.yaml/badge.svg
[go-test-windows.yaml]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/go-test-windows.yaml
[golangci-lint.yaml Badge]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/golangci-lint.yaml/badge.svg
[golangci-lint.yaml]: https://github.com/senzing-garage/go-cmdhelping/actions/workflows/golangci-lint.yaml
[License Badge]: https://img.shields.io/badge/License-Apache2-brightgreen.svg
[License]: https://github.com/senzing-garage/go-cmdhelping/blob/main/LICENSE
[main.go]: main.go
[option.ContextVariable]: option/option.go
[Senzing Garage]: https://github.com/senzing-garage-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[Senzing]: https://senzing.com/
[viper]: https://github.com/spf13/viper
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/go-cmdhelping
