package gitops_test

import (
	"os"
	"testing"

	gitops "git-tui/git-ops"
)

func TestFind(t *testing.T) {
	repo := setupRepo(t)

	notGitDir, err := os.MkdirTemp("", "git-tui-nongit-*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(notGitDir) })

	tests := []struct {
		name    string
		dir     string
		wantErr bool
	}{
		{"valid git repo", repo.Root, false},
		{"non-git directory", notGitDir, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := gitops.Find(tt.dir)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Find(%q) error = %v, wantErr %v", tt.dir, err, tt.wantErr)
			}
			if !tt.wantErr && got.Root != repo.Root {
				t.Errorf("Root = %q, want %q", got.Root, repo.Root)
			}
		})
	}
}
