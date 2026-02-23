package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styles "git-tui/styles"
)

var noRemotes = []MenuItem{{
	Label: "(no remotes)",
	Operation: func(_ *git.Repo) Operation {
		return Operation{Title: styles.HintStyle.Render("no remotes configured")}
	},
}}

func createRemoteItem(name string) MenuItem {
	return MenuItem{
		Label: name,
		Operation: func(_ *git.Repo) Operation {
			return Operation{ConfirmPrompt: &ConfirmPrompt{
				Prompt: fmt.Sprintf("Delete remote %q?", name),
				OnYes: func(r *git.Repo) string {
					if err := r.RemoveRemote(name); err != nil {
						return styles.WarnStyle.Render("error: " + err.Error())
					}
					return styles.SuccessStyle.Render("✓ removed remote " + name)
				},
			}}
		}}
}

var RemotesDeleteItem = MenuItem{
	Label: "Delete",
	Submenu: func(r *git.Repo) []MenuItem {
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
