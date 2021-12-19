# Random

Package random is generics based random choice utility.
Similar of Python's standard package.

Note: This was created to help me learn about generics.

## Requirements

- Go 1.18 or later

## Usage

See example test code.

## Test

```console
$ go1.18beta1 test .
```

## Develop Setting (VSCode)

Download and install go1.18beta1.

https://pkg.go.dev/golang.org/dl/go1.18beta1

```console
$ go install golang.org/dl/go1.18beta1@latest
$ go1.18beta1 download
```

Download and install gopls by go1.18beta1.

https://github.com/golang/tools/blob/master/gopls/doc/advanced.md

```console
$ go1.18beta1 get golang.org/x/tools/gopls@master golang.org/x/tools@master
$ go1.18beta1 install golang.org/x/tools/gopls
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
