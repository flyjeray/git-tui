//go:build !windows

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ensureInPath appends an export PATH line for dir to the user's shell RC file,
// if it isn't already present. Supports zsh and bash; skips other shells silently.
func ensureInPath(home, dir string) error {
	shell := os.Getenv("SHELL")
	var rcFile string
	switch {
	case strings.HasSuffix(shell, "zsh"):
		rcFile = filepath.Join(home, ".zshrc")
	case strings.HasSuffix(shell, "bash"):
		rcFile = filepath.Join(home, ".bashrc")
	default:
		return nil // unsupported shell, skip
	}

	data, _ := os.ReadFile(rcFile)
	if strings.Contains(string(data), dir) {
		return nil
	}

	f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintf(f, "\nexport PATH=\"%s:$PATH\"\n", dir)
	return err
}

func pathExportHint(dir string) string {
	return `export PATH="` + dir + `:$PATH"`
}
