package main

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	ToolsFile string `short:"f" description:"Path to the tools.go file"`
	Verbose   bool   `short:"v" description:"Verbose output"`
}

var options Options

var parser = flags.NewParser(&options, flags.Default)

func main() {
	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}
}
