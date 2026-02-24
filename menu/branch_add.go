package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func createBranch(r *git.Repo, values []string) string {
	name := values[0]
	out, err := r.CheckoutToNew(name)
	if err != nil {
		return styled.Warn("error: " + err.Error())
	}
	return styled.Success("✓ " + out)
}

var BranchAddItem = MenuItem{
	Label: "Checkout to new branch",
	Flow: func(_ *git.Repo) *InputFlow {
		return &InputFlow{
			Title: "Checkout to new branch",
			Steps: []InputStep{
				{Label: "Name"},
			},
			OnSubmit: createBranch,
		}
	},
}
