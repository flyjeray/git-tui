package gitops

import (
	"os/exec"
	"strings"
)

func (r *Repo) RunCommand(args ...string) (string, error) {
	cmdArgs := append([]string{"-C", r.Root}, args...)
	out, err := exec.Command("git", cmdArgs...).CombinedOutput()
	return strings.TrimSpace(string(out)), err
}
