package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

// Pull runs "git pull" in the repository root and returns the command output
// (stdout and stderr combined). On error, the returned string contains the
// trimmed git output for easier display in the UI.
func (r *Repo) Pull() (string, error) {
	cmd := exec.Command("git", "-C", r.Root, "pull")
	out, err := cmd.CombinedOutput()
	msg := strings.TrimSpace(string(out))
	if err != nil {
		if msg == "" {
			msg = "git pull failed"
		}
		return msg, fmt.Errorf("%s", msg)
	}
	if msg == "" {
		msg = "Already up to date."
	}
	return msg, nil
}

// PullFromRemote runs "git pull <remote>" targeting a specific named remote.
func (r *Repo) PullFromRemote(remote, branch string) (string, error) {
	cmd := exec.Command("git", "-C", r.Root, "pull", remote, branch)
	out, err := cmd.CombinedOutput()
	msg := strings.TrimSpace(string(out))
	if err != nil {
		if msg == "" {
			msg = "git pull failed"
		}
		return msg, fmt.Errorf("%s", msg)
	}
	if msg == "" {
		msg = "Already up to date."
	}
	return msg, nil
}
