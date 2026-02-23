package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

var BranchCurrentItem = MenuItem{
	Label: "Current branch",
	Info: func(r *git.Repo) string {
		branch, err := r.GetCurrentBranch()
		if err != nil {
			return styled.Warn("error: " + err.Error())
		}
		return "Current branch: " + styled.Hint(branch)
	},
}
