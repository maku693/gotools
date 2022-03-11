package main

import (
	"fmt"
	"os"
	"path"

	"github.com/maku693/gotools"
)

func Setenv() error {
	toolsFile, err := gotools.ToolsFile(options.ToolsFile)
	if err != nil {
		return err
	}

	if options.Verbose {
		fmt.Fprintf(os.Stderr, "* tools file: %s\n", toolsFile)
	}

	toolsDir := path.Dir(toolsFile)
	gobin := path.Join(toolsDir, "bin")

	if err := os.Setenv("GOBIN", gobin); err != nil {
		return err
	}

	if err := os.Setenv("PATH", os.ExpandEnv("$PATH:$GOBIN")); err != nil {
		return err
	}

	return nil
}
