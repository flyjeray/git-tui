package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func pushProcess(r *git.Repo, remote, branch string) string {
	var _, pushErr = r.PushToRemote(remote, branch)

	if pushErr != nil {
		return styled.Warn("push error: " + pushErr.Error())
	}

	return styled.Success("Updates are pushed to remote!")
}

func singleOriginPush(remote, branch string) MenuItem {
	return MenuItem{
		Label: "Push to " + remote,
		Confirm: &ConfirmPrompt{
			Prompt: fmt.Sprintf("Push latest changes to %q?", remote),
			OnYes: func(r *git.Repo) string {
				return pushProcess(r, remote, branch)
			},
		},
	}
}

func multipleOriginsPush(origins []string, branch string) MenuItem {
	return MenuItem{
		Label: "Push",
		Submenu: func(r *git.Repo) []MenuItem {
			items := make([]MenuItem, len(origins))
			for i, name := range origins {
				items[i] = singleOriginPush(name, branch)
			}
			return items
		},
	}
}

func PushMenuItem(r *git.Repo) MenuItem {
	remotes, remotesErr := r.GetRemoteNames()

	if remotesErr != nil {
		return GitErrorMenu(remotesErr)
	}

	branch, branchErr := r.GetCurrentBranch()

	if branchErr != nil {
		return GitErrorMenu(branchErr)
	}

	if len(remotes) == 0 {
		return noRemotesItem("Push")
	}

	if len(remotes) == 1 {
		return singleOriginPush(remotes[0], branch)
	}

	return multipleOriginsPush(remotes, branch)
}
