# scgen

Generate shellcode from the command line

## Install
scgen is written in Go and has no run-time dependencies. If you have Go installed and configured (i.e. with `$GOPATH/bin` in your `$PATH`):
```bash
go get -u github.com/peterstrongg/scgen
```
Otherwise, [download](https://github.com/peterstrongg/scgen/releases) a release for your platform. To make it easier to execute you can put the binary in your `$PATH`.

## Usage
```bash
scgen <binary_executable>
```
