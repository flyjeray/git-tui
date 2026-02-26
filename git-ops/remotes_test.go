package gitops_test

import (
	"strings"
	"testing"
)

func TestGetRemoteNames(t *testing.T) {
	r := setupRepo(t)

	// no remotes yet
	names, err := r.GetRemoteNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(names) != 0 {
		t.Errorf("expected no remotes, got %v", names)
	}

	// add one and check it appears
	gitCmd(t, r.Root, "remote", "add", "origin", "https://example.com/repo.git")

	names, err = r.GetRemoteNames()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(names) != 1 || names[0] != "origin" {
		t.Errorf("got %v, want [origin]", names)
	}
}

func TestGetRemotesWithURLs(t *testing.T) {
	r := setupRepo(t)
	gitCmd(t, r.Root, "remote", "add", "origin", "https://example.com/repo.git")

	// git remote -v lists each remote twice (fetch + push); the function must deduplicate
	remotes, err := r.GetRemotesWithURLs()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(remotes) != 1 {
		t.Fatalf("expected 1 entry after dedup, got %d: %v", len(remotes), remotes)
	}
	if !strings.Contains(remotes[0], "origin") || !strings.Contains(remotes[0], "https://example.com/repo.git") {
		t.Errorf("unexpected format: %q", remotes[0])
	}
}
