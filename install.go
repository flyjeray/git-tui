package main

import (
	"fmt"
	styles "git-tui/styles"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

func selfInstall() {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}

	var installDir, binaryName string
	if runtime.GOOS == "windows" {
		installDir = filepath.Join(home, "AppData", "Local", "Programs", "gt")
		binaryName = "gt.exe"
	} else {
		installDir = filepath.Join(home, ".local", "bin")
		binaryName = "gt"
	}

	installPath := filepath.Join(installDir, binaryName)

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
		fmt.Println(styles.Warn("⚠ installed, but could not update PATH: " + err.Error()))
		fmt.Println(styles.Hint("  Add this manually: " + pathExportHint(installDir)))
		fmt.Println()
		return
	}

	fmt.Println(styles.Success("✓ git-tui installed to " + installPath))
	fmt.Println(styles.Hint("  Open a new terminal and run \"gt\" to use it from anywhere."))
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
