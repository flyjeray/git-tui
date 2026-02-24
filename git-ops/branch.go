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

func (r *Repo) GetInactiveBranchList() ([]string, error) {
	out, err := exec.Command("git", "branch").Output()
	if err != nil {
		return []string{}, fmt.Errorf("could not get list of branches: %w", err)
	}

	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return []string{}, nil
	}

	branches := []string{}

	for _, line := range strings.Split(raw, "\n") {
		fields := strings.Fields(line) // ["*" if active, "name"]
		if len(fields) > 1 {
			continue
		}
		if fields[0] != "" {
			branches = append(branches, fields[0])
		}
	}

	return branches, nil
}

func (r *Repo) Checkout(branch string) (string, error) {
	out, err := exec.Command("git", "checkout", branch).Output()
	if err != nil {
		return "", fmt.Errorf("could not checkout: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func (r *Repo) CheckoutToNew(branch string) (string, error) {
	out, err := exec.Command("git", "checkout", "-b", branch).Output()
	if err != nil {
		return "", fmt.Errorf("could not checkout: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

func (r *Repo) DeleteBranch(branch string) (string, error) {
	out, err := exec.Command("git", "branch", "-d", branch).Output()
	if err != nil {
		return "", fmt.Errorf("could not delete branch: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}
