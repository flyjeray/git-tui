package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func singleOriginPull(remote, branch string) MenuItem {
	return MenuItem{
		Label: "Pull from " + remote,
		Confirm: &ConfirmPrompt{
			Prompt: fmt.Sprintf("Pull latest changes from %q?", remote),
			OnYes: func(r *git.Repo) string {
				msg, err := r.PullFromRemote(remote, branch)
				if err != nil {
					return styled.Warn("error: " + msg)
				}
				return styled.Success("✓ pull completed") + "\n" + styled.Hint(msg)
			},
		},
	}
}

func multipleOriginsPull(origins []string, branch string) MenuItem {
	return MenuItem{
		Label: "Pull",
		Submenu: func(r *git.Repo) []MenuItem {
			items := make([]MenuItem, len(origins))
			for i, name := range origins {
				items[i] = singleOriginPull(name, branch)
			}
			return items
		},
	}
}

func PullMenuItem(r *git.Repo) MenuItem {
	names, remotesErr := r.GetRemoteNames()

	if remotesErr != nil {
		return GitErrorMenu(remotesErr)
	}

	branch, branchErr := r.GetCurrentBranch()

	if branchErr != nil {
		return GitErrorMenu(branchErr)
	}

	if len(names) == 0 {
		return noRemotesItem("Pull")
	}

	if len(names) == 1 {
		return singleOriginPull(names[0], branch)
	}

	return multipleOriginsPull(names, branch)
}
