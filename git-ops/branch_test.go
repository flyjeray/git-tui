package gitops_test

import (
	"slices"
	"testing"
)

func TestGetCurrentBranch(t *testing.T) {
	r := setupRepo(t)
	makeCommit(t, r, "initial")

	got, err := r.GetCurrentBranch()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "main" {
		t.Errorf("got %q, want %q", got, "main")
	}
}

func TestGetInactiveBranchList(t *testing.T) {
	r := setupRepo(t)
	makeCommit(t, r, "initial")

	gitCmd(t, r.Root, "branch", "feature-a")
	gitCmd(t, r.Root, "branch", "feature-b")

	branches, err := r.GetInactiveBranchList()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// the active branch must not appear
	if slices.Contains(branches, "main") {
		t.Error("active branch 'main' should not be in inactive list")
	}

	for _, want := range []string{"feature-a", "feature-b"} {
		if !slices.Contains(branches, want) {
			t.Errorf("expected %q in inactive list", want)
		}
	}
}
