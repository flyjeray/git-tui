package gitops

import (
	"os/exec"
	"strings"
)

func (r *Repo) Commit(commit string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", commit)
	out, err := cmd.CombinedOutput()
	msg := strings.TrimSpace(string(out))
	return msg, err
}
