package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

func (r *Repo) GetCurrentBranch() (string, error) {
	out, err := exec.Command("git", "branch", "--show-current").Output()
	if err != nil {
		return "", fmt.Errorf("could not get branch: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
