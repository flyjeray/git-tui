//go:build windows

package main

import (
	"fmt"
	"os/exec"
)

// ensureInPath adds dir to the current user's persistent PATH via PowerShell,
// if it isn't already present.
func ensureInPath(_ string, dir string) error {
	// Check if dir is already in the user PATH.
	checkScript := fmt.Sprintf(
		`if (([Environment]::GetEnvironmentVariable("Path","User") -split ";") -contains "%s") { exit 0 } else { exit 1 }`,
		dir,
	)
	if err := exec.Command("powershell", "-NoProfile", "-Command", checkScript).Run(); err == nil {
		return nil
	}

	// Append dir to the user PATH in the registry.
	addScript := fmt.Sprintf(
		`$p = [Environment]::GetEnvironmentVariable("Path","User"); [Environment]::SetEnvironmentVariable("Path", ($p.TrimEnd(";") + ";%s"), "User")`,
		dir,
	)
	return exec.Command("powershell", "-NoProfile", "-Command", addScript).Run()
}

func pathExportHint(dir string) string {
	return `add "` + dir + `" to your user PATH via System Properties > Environment Variables`
}
