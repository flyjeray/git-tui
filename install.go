package main

import (
	"fmt"
	styles "git-tui/styles"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func selfInstall() {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	installDir := filepath.Join(home, ".local", "bin")
	installPath := filepath.Join(installDir, "gt")

	execPath, err := os.Executable()
	if err != nil {
		return
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		return
	}

	// Already running from the install location — nothing to do.
	// Use os.SameFile (inode comparison) instead of string comparison so that
	// hard links and any path-normalisation differences are handled correctly.
	execInfo, execStatErr := os.Lstat(execPath)
	if execStatErr == nil {
		if installInfo, err := os.Lstat(installPath); err == nil {
			if os.SameFile(execInfo, installInfo) {
				return
			}
		}
	}

	// 0755: rwxr-xr-x — owner can write, everyone can read/execute
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return
	}

	// Write to a temp file first, then atomically rename into place.
	// A direct O_TRUNC open on the install path would truncate the file before
	// writing; if that path is currently executing, macOS sends SIGKILL.
	// os.Rename replaces the directory entry without touching the old inode,
	// so any running process mapped from the old file is unaffected.
	tmpPath := installPath + ".tmp"
	if err := copyBinary(execPath, tmpPath); err != nil {
		os.Remove(tmpPath)
		return
	}
	if err := os.Rename(tmpPath, installPath); err != nil {
		os.Remove(tmpPath)
		return
	}

	if err := ensureInPath(home, installDir); err != nil {
		fmt.Println(styles.Warn("⚠ installed, but could not update PATH in shell RC: " + err.Error()))
		fmt.Println(styles.Hint("  Add this manually: export PATH=\"" + installDir + ":$PATH\""))
		fmt.Println()
		return
	}

	fmt.Println(styles.Success("✓ git-tui installed to " + installPath))
	fmt.Println(styles.Hint("  Restart your terminal and call \"gt\" command to use it from anywhere."))
	fmt.Println()
}

func copyBinary(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

// ensureInPath appends an export PATH line for dir to the user's shell RC file,
// if it isn't already present. Supports zsh and bash; skips other shells silently.
func ensureInPath(home, dir string) error {
	// test error message to check logging
	// return fmt.Errorf("intentional test error")

	// Detect shell from $SHELL env var (e.g. /bin/zsh, /usr/bin/bash)
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

	// Avoid duplicate entries
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
