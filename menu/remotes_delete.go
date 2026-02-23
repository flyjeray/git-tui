package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

var noRemotes = []MenuItem{{
	Label:  "(no remotes)",
	Result: func(_ *git.Repo) string { return styled.Hint("no remotes configured") },
}}

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

var RemotesDeleteItem = MenuItem{
	Label: "Delete",
	Submenu: func(r *git.Repo) []MenuItem {
		if r == nil {
			return noRemotes
		}
		names, err := r.GetRemoteNames()
		if err != nil || len(names) == 0 {
			return noRemotes
		}
		items := make([]MenuItem, len(names))
		for i, name := range names {
			items[i] = createRemoteItem(name)
		}
		return items
	},
}
