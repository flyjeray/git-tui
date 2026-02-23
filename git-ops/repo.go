package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

type Repo struct {
	Root string
}

// Find returns the Repo for the git repository containing dir.
// Returns an error if dir is not inside a git repository.
func Find(dir string) (*Repo, error) {
	out, err := exec.Command("git", "-C", dir, "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return nil, fmt.Errorf("not a git repository")
	}
	return &Repo{Root: strings.TrimSpace(string(out))}, nil
}
