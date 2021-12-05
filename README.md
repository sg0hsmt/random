# Random

Package random is generics based random choice utility.
Similar of Python's standard package.

Note: This was created to help me learn about generics.

## Requirements

- Go 1.18 or later

## Usage

See example test code.

## Test

```bash
$ gotip test .
```

## Develop Setting (VSCode)

Download and install gotip.

https://pkg.go.dev/golang.org/dl/gotip

```bash
$ go install golang.org/dl/gotip@latest
$ gotip download
```

Download and install gopls by gotip.

https://github.com/golang/tools/blob/master/gopls/doc/advanced.md

```bash
$ gotip get golang.org/x/tools/gopls@master golang.org/x/tools@master
$ gotip install golang.org/x/tools/gopls
```

Change vscode settings.

https://github.com/golang/vscode-go/blob/master/docs/settings.md#goalternatetools

```json
{
  "go.alternateTools": {
    "go": "gotip"
  }
}
```
