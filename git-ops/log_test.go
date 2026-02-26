package gitops_test

import (
	"testing"
)

func TestGetLog(t *testing.T) {
	r := setupRepo(t)
	makeCommit(t, r, "first commit")
	makeCommit(t, r, "second commit")
	makeCommit(t, r, "third commit")

	tests := []struct {
		name         string
		skip, count  int
		wantSubjects []string
	}{
		{
			name:         "all commits, newest first",
			skip:         0,
			count:        10,
			wantSubjects: []string{"third commit", "second commit", "first commit"},
		},
		{
			name:         "limited to one",
			skip:         0,
			count:        1,
			wantSubjects: []string{"third commit"},
		},
		{
			name:         "skip the newest",
			skip:         1,
			count:        10,
			wantSubjects: []string{"second commit", "first commit"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			entries, err := r.GetLog(tt.skip, tt.count)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(entries) != len(tt.wantSubjects) {
				t.Fatalf("got %d entries, want %d", len(entries), len(tt.wantSubjects))
			}
			for i, want := range tt.wantSubjects {
				if entries[i].Subject != want {
					t.Errorf("entries[%d].Subject = %q, want %q", i, entries[i].Subject, want)
				}
				if entries[i].Hash == "" {
					t.Errorf("entries[%d].Hash is empty", i)
				}
				if entries[i].Author != "Test User" {
					t.Errorf("entries[%d].Author = %q, want %q", i, entries[i].Author, "Test User")
				}
			}
		})
	}
}
