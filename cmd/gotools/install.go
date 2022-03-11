package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/maku693/gotools"
)

type InstallCommand struct {
	DryRun bool `long:"dry-run" description:"Do not actually install tools"`
}

func (c *InstallCommand) Execute([]string) error {
	if err := Setenv(); err != nil {
		return err
	}

	tools, err := gotools.ParseTools(options.ToolsFile)
	if err != nil {
		return fmt.Errorf("failed to parse tools.go: %w", err)
	}

	for _, tool := range tools {
		args := []string{"install"}
		if len(tool.Tags) > 0 {
			args = append(args, "-tags")
			args = append(args, tool.Tags...)
		}
		args = append(args, tool.PackagePath)

		cmd := exec.Command("go", args...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if options.Verbose {
			fmt.Fprintf(os.Stderr, "* installing tool: %s\n", tool.PackagePath)
		}

		if !c.DryRun {
			if err := cmd.Run(); err != nil {
				return err
			}
		}
	}

	return nil
}

var installCommand InstallCommand

func init() {
	parser.AddCommand("install",
		"Install tools",
		"",
		&installCommand)
}
