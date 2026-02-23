package menu

import (
	git "git-tui/git-ops"
	styles "git-tui/styles"
)

var BranchMenuItem = MenuItem{
	Label: "Branch",
	Operation: func(r *git.Repo) Operation {
		branch, err := r.GetCurrentBranch()
		if err != nil {
			return Operation{Title: styles.WarnStyle.Render("error: " + err.Error())}
		}
		return Operation{Title: "Current branch: " + styles.HintStyle.Render(branch)}
	},
}
