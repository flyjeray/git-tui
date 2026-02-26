package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

func (r *Repo) GetRemotesWithURLs() ([]string, error) {
	out, err := exec.Command("git", "-C", r.Root, "remote", "-v").Output()
	if err != nil {
		return nil, fmt.Errorf("could not get remotes: %w", err)
	}
	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return []string{}, nil
	}

	urls := map[string]string{}

	for _, line := range strings.Split(raw, "\n") {
		fields := strings.Fields(line) // ["name", "url", "(operations)"]
		if len(fields) < 2 {
			continue
		}
		name, url := fields[0], fields[1]
		if _, exists := urls[name]; !exists {
			urls[name] = url
		}
	}

	result := make([]string, 0, len(urls))
	for key, value := range urls {
		result = append(result, fmt.Sprintf("%s - %s", key, value))
	}
	return result, nil
}

func (r *Repo) GetRemoteNames() ([]string, error) {
	out, err := exec.Command("git", "-C", r.Root, "remote").Output()
	if err != nil {
		return nil, fmt.Errorf("could not list remotes: %w", err)
	}
	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return []string{}, nil
	}
	return strings.Split(raw, "\n"), nil
}

func (r *Repo) AddRemote(name, url string) error {
	out, err := exec.Command("git", "-C", r.Root, "remote", "add", name, url).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", strings.TrimSpace(string(out)))
	}
	return nil
}

func (r *Repo) RemoveRemote(name string) error {
	out, err := exec.Command("git", "-C", r.Root, "remote", "remove", name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s", strings.TrimSpace(string(out)))
	}
	return nil
}
