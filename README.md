# Random

[![Go Reference](https://pkg.go.dev/badge/github.com/sg0hsmt/random.svg)](https://pkg.go.dev/github.com/sg0hsmt/random)
[![Test](https://github.com/sg0hsmt/random/actions/workflows/test.yaml/badge.svg)](https://github.com/sg0hsmt/random/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sg0hsmt/random)](https://goreportcard.com/report/github.com/sg0hsmt/random)
[![License](https://img.shields.io/github/license/sg0hsmt/random.svg)](https://github.com/sg0hsmt/random/blob/master/LICENSE)
[![Release](https://img.shields.io/github/release/sg0hsmt/random.svg)](https://github.com/sg0hsmt/random/releases/latest)

Package random is generics based random choice utility.
Similar of Python's standard package.

Note: This was created to help me learn about generics.

## Requirements

Go 1.18 or later

## Usage

See example test code.

## Test

```console
go1.18beta1 test .
```

You can use GitHub Actions locally by [act](https://github.com/nektos/act).

```console
act -j test
```

## Develop Setting (VSCode)

Download and install go1.18beta1.

https://pkg.go.dev/golang.org/dl/go1.18beta1

```console
go install golang.org/dl/go1.18beta1@latest
go1.18beta1 download
```

Download and install gopls by go1.18beta1.

https://github.com/golang/tools/blob/master/gopls/doc/advanced.md

```console
go1.18beta1 get golang.org/x/tools/gopls@master golang.org/x/tools@master
go1.18beta1 install golang.org/x/tools/gopls
```

Change vscode settings.

https://github.com/golang/vscode-go/blob/master/docs/settings.md#goalternatetools

```json
{
  "go.alternateTools": {
    "go": "go1.18beta1"
  }
}
```
