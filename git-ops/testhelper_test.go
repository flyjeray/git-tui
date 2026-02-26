package gitops_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	gitops "git-tui/git-ops"
)

// setupRepo creates a temporary git repository and returns a Repo pointing at
// it. The directory is automatically removed when the test ends.
func setupRepo(t *testing.T) *gitops.Repo {
	t.Helper()

	dir, err := os.MkdirTemp("", "git-tui-test-*")
	if err != nil {
		t.Fatal(err)
	}
	// On macOS /var is a symlink to /private/var; git resolves the real path,
	// so we normalise here to keep comparisons consistent.
	dir, err = filepath.EvalSymlinks(dir)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })

	gitCmd(t, dir, "init", "-b", "main")
	gitCmd(t, dir, "config", "user.email", "test@example.com")
	gitCmd(t, dir, "config", "user.name", "Test User")

	return &gitops.Repo{Root: dir}
}

// gitCmd runs a git command in dir and fails the test if it exits non-zero.
func gitCmd(t *testing.T, dir string, args ...string) {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", dir}, args...)...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git %v: %s", args, out)
	}
}

// makeCommit creates an empty commit with the given message in r.
func makeCommit(t *testing.T, r *gitops.Repo, message string) {
	t.Helper()
	gitCmd(t, r.Root, "commit", "--allow-empty", "-m", message)
}
