package gitops

import (
	"os/exec"
	"strings"
)

func (r *Repo) PushToRemote(remote, branch string) (string, error) {
	cmd := exec.Command("git", "-C", r.Root, "push", "--set-upstream", remote, branch)
	out, err := cmd.CombinedOutput()
	msg := strings.TrimSpace(string(out))

	return msg, err
}
