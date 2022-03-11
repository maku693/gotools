package gotools

import (
	"errors"
	"go/parser"
	"go/token"
	"os"
	"path"
	"reflect"
	"strings"
)

// ParseTools parses tools.go file of provided filename.
//
// If filename is empty string, it tries to infer default file path.
func ParseTools(filename string) ([]*Tool, error) {
	toolsFile, err := ToolsFile(filename)
	if err != nil {
		return nil, err
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, toolsFile, nil, parser.ImportsOnly|parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var tools []*Tool
	for _, importSpec := range f.Imports {
		packagePath := strings.TrimFunc(importSpec.Path.Value, func(r rune) bool { return r == '"' })
		var tags []string

		// Parse tag specification if exists
		if comment := importSpec.Comment; comment != nil {
			tagsValue, ok := reflect.StructTag(comment.Text()).Lookup("tags")
			if ok {
				tags = strings.Split(tagsValue, ",")
			}
		}

		tools = append(tools, &Tool{
			PackagePath: packagePath,
			Tags:        tags,
		})
	}

	return tools, nil
}

type Tool struct {
	PackagePath string
	Tags        []string
}

// ToolsFile returns the path of tools.go file.
//
// If filename is empty string, it search for go.mod file in parent directories to
// infer the default tools.go file path.
func ToolsFile(filename string) (string, error) {
	if filename != "" {
		return filename, nil
	}

	modfile := ""
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		candidate := path.Join(dir, "go.mod")
		info, err := os.Stat(candidate)
		// If go.mod is not found on current `dir` then walk directory upwards
		if errors.Is(err, os.ErrNotExist) {
			parent := path.Dir(dir)
			// Stop walking if there's no more parent directory
			if parent == dir {
				break
			}
			dir = parent
			continue
		}
		if err != nil {
			return "", err
		}
		if !info.IsDir() {
			modfile = candidate
			break
		}
	}

	if modfile == "" {
		return "", errors.New("failed to infer module root")
	}

	pkgRoot := path.Dir(modfile)
	defaultPath := path.Join(pkgRoot, "tools/tools.go")

	return defaultPath, nil
}
