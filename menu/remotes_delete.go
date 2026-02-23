package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func noRemotesItem(prefix string) MenuItem {
	return MenuItem{
		Label:  prefix + " (no remotes)",
		Result: func(_ *git.Repo) string { return styled.Hint("no remotes configured") },
	}
}

func createRemoteItem(name string) MenuItem {
	return MenuItem{
		Label: name,
		Confirm: &ConfirmPrompt{
			Prompt: fmt.Sprintf("Delete remote %q?", name),
			OnYes: func(r *git.Repo) string {
				if err := r.RemoveRemote(name); err != nil {
					return styled.Warn("error: " + err.Error())
				}
				return styled.Success("✓ removed remote " + name)
			},
		},
	}
}

func remoteToDeleteSelectMenu(remotes []string) MenuItem {
	return MenuItem{
		Label: "Delete",
		Submenu: func(r *git.Repo) []MenuItem {
			items := make([]MenuItem, len(remotes))
			for i, name := range remotes {
				items[i] = createRemoteItem(name)
			}
			return items
		},
	}
}

func RemoteDeleteMenu(r *git.Repo) MenuItem {
	remotes, err := r.GetRemoteNames()

	if err != nil {
		return GitErrorMenu(err)
	}

	if len(remotes) == 0 {
		return noRemotesItem("Delete")
	}

	return remoteToDeleteSelectMenu(remotes)
}
