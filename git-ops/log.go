package gitops

import (
	"fmt"
	"os/exec"
	"strings"
)

type LogEntry struct {
	Hash    string
	Subject string
	Author  string
}

func (r *Repo) GetLog(skip, count int) ([]LogEntry, error) {
	out, err := exec.Command(
		"git", "-C", r.Root,
		"log",
		"--format=%h\x1f%s\x1f%an",
		fmt.Sprintf("-n%d", count),
		fmt.Sprintf("--skip=%d", skip),
	).Output()
	if err != nil {
		return nil, fmt.Errorf("could not get log: %w", err)
	}
	raw := strings.TrimSpace(string(out))
	if raw == "" {
		return nil, nil
	}
	lines := strings.Split(raw, "\n")
	entries := make([]LogEntry, 0, len(lines))
	for _, line := range lines {
		parts := strings.SplitN(line, "\x1f", 3)
		if len(parts) != 3 {
			continue
		}
		entries = append(entries, LogEntry{
			Hash:    parts[0],
			Subject: parts[1],
			Author:  parts[2],
		})
	}
	return entries, nil
}
