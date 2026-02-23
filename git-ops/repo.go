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

// Branch returns the name of the currently checked-out branch.
func (r *Repo) Branch() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		return "", fmt.Errorf("could not get branch: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// Remotes returns remotes formatted as "{name} - {url} ({operations})".
// fetch and push are merged into one entry when the URL is the same.
func (r *Repo) Remotes() ([]string, error) {
	out, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		return nil, fmt.Errorf("could not get remotes: %w", err)
	}
	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return []string{}, nil
	}

	// git remote -v outputs: "name\turl (op)" — two lines per remote (fetch + push)
	ops := map[string][]string{} // "name\turl" -> operations
	var order []string

	for _, line := range strings.Split(raw, "\n") {
		fields := strings.Fields(line) // ["name", "url", "(op)"]
		if len(fields) < 3 {
			continue
		}
		name, url, op := fields[0], fields[1], strings.Trim(fields[2], "()")
		k := name + "\t" + url
		if len(ops[k]) == 0 {
			order = append(order, k)
		}
		ops[k] = append(ops[k], op)
	}

	result := make([]string, 0, len(order))
	for _, k := range order {
		name, url, _ := strings.Cut(k, "\t")
		result = append(result, fmt.Sprintf("%s - %s (%s)", name, url, strings.Join(ops[k], ", ")))
	}
	return result, nil
}
