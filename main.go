package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/mod/modfile"
)

// ResolveModulePath resolves the Go module name by parsing the go.mod file of the target project.
// If no go.mod is present, it falls back to the directory-based path resolution.
func ResolveModulePath(dir string) (string, error) {
	goModPath := filepath.Join(dir, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		// Fallback to directory-based path resolution
		return filepath.Base(dir), nil
	}

	content, err := ioutil.ReadFile(goModPath)
	if err != nil {
		return "", err
	}

	file, err := modfile.Parse(goModPath, content, nil)
	if err != nil {
		return "", err
	}

	if file.Module != nil {
		return file.Module.Mod.Path, nil
	}

	return filepath.Base(dir), nil
}

func main() {
	fmt.Println("Buffalo Generator Module Path Resolver")
}
