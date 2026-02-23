package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func singleBranchItem(branch string) MenuItem {
	return MenuItem{
		Label: branch,
		Confirm: &ConfirmPrompt{
			Prompt: "Checkout to " + branch + "?",
			OnYes: func(r *git.Repo) string {
				if _, err := r.Checkout(branch); err != nil {
					return styled.Warn("error: " + err.Error())
				}
				return styled.Success("✓ checkout to " + branch)
			},
		},
	}
}

func branches(r *git.Repo) []MenuItem {
	branches, err := r.GetInactiveBranchList()

	if err != nil {
		return []MenuItem{GitErrorMenu(err)}
	}

	items := []MenuItem{}

	for _, branch := range branches {
		items = append(items, singleBranchItem(branch))
	}

	return items
}

func BranchListMenu(r *git.Repo) MenuItem {
	return MenuItem{
		Label:   "Checkout",
		Submenu: branches,
	}
}
