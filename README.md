local-go
========

local-go provides equivalents of go command line tools that automatically set up a GOPATH structure for packages that are outside of your default GOPATH so that commands like `godoc` can work properly.

Currently supported:

* go (lgo)
* godoc (lgodoc)
* gofmt (lgofmt)
* goimports (lgoimports)

Usage:

`lgo <path> [<arg>...]`

## Building and Installing

To build and install, `cd` into one of the directories and run `go install`.