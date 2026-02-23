package gitops

import (
	"os/exec"
	"strings"
)

func (r *Repo) StageAll() (string, error) {
	cmd := exec.Command("git", "-C", r.Root, "add", "-A")
	out, err := cmd.CombinedOutput()
	msg := strings.TrimSpace(string(out))
	return msg, err
}
