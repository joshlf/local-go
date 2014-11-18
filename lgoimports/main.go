package main

import (
	"github.com/synful/local-go"
	"os"
)

func main() {
	os.Exit(lgo.Main("goimports", os.Args))
}
