package menu

import (
	"errors"

	git "git-tui/git-ops"
)

// fetchRemotesAndBranch runs GetRemoteNames and GetCurrentBranch concurrently
// and returns both results. The two git commands are independent so there is no
// reason to wait for one before starting the other.
func fetchRemotesAndBranch(r *git.Repo) ([]string, string, error) {
	type remotesResult struct {
		val []string
		err error
	}
	type branchResult struct {
		val string
		err error
	}

	remotesCh := make(chan remotesResult, 1)
	branchCh := make(chan branchResult, 1)

	go func() {
		val, err := r.GetRemoteNames()
		remotesCh <- remotesResult{val, err}
	}()
	go func() {
		val, err := r.GetCurrentBranch()
		branchCh <- branchResult{val, err}
	}()

	rr := <-remotesCh
	br := <-branchCh

	if err := errors.Join(rr.err, br.err); err != nil {
		return nil, "", err
	}
	return rr.val, br.val, nil
}
