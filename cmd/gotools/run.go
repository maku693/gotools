package main

import (
	"fmt"
	"os"
	"os/exec"
)

type RunCommand struct{}

func (c *RunCommand) Execute(args []string) error {
	if err := Setenv(); err != nil {
		return err
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if options.Verbose {
		fmt.Fprintf(os.Stderr, "* running command: %s\n", cmd)
	}

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

var runCommand RunCommand

func init() {
	parser.AddCommand("run",
		"Run a tool",
		"",
		&runCommand)
}
