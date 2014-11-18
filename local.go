package lgo

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	EXIT_USAGE = 1 + iota // Improper arguments
	EXIT_IO               // IO error
	EXIT_EXEC             // Error running subcommand
)

// args must be os.Args (assumes same invariants)
func Main(bin string, args []string) int {
	if len(args) < 2 {
		return usage(args)
	}
	path, err := filepath.Abs(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not resolve path: %v\n", err)
		return EXIT_IO
	}
	dir, err := ioutil.TempDir("", args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not create temp directory: %v\n", err)
		return EXIT_IO
	}
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not remove temp directory: %v\n", err)
		}
	}()
	if err := os.Symlink(path, dir+"/src"); err != nil {
		fmt.Fprintf(os.Stderr, "could not symlink src: %v\n", err)
		return EXIT_IO
	}
	cmd := exec.Command(bin, args[2:]...)
	// Important that GOPATH comes before GOROOT
	// (otherwise GOROOT will be used instead)
	cmd.Env = append([]string{"GOPATH=" + dir}, os.Environ()...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	err = cmd.Run()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			fmt.Fprintf(os.Stderr, "error running \"%s\": %v\n", bin, err)
		} else {
			fmt.Fprintf(os.Stderr, "could not run \"%s\": %v\n", bin, err)
		}
		return EXIT_EXEC
	}
	return 0
}

func usage(args []string) int {
	fmt.Fprintf(os.Stderr, "Usage: %v <path> [<args>...]\n", args[0])
	return EXIT_USAGE
}
