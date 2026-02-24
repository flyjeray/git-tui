package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func singleBranchMergeItem(branch string) MenuItem {
	return MenuItem{
		Label: branch,
		Confirm: &ConfirmPrompt{
			Prompt: "Delete branch " + branch + "?",
			OnYes: func(r *git.Repo) string {
				if _, err := r.MergeBranch(branch); err != nil {
					return styled.Warn("error: " + err.Error())
				}
				return styled.Success("✓ merged " + branch)
			},
		},
	}
}

var BranchMergeMenu = MenuItem{
	Label: "Merge",
	LevelSubmenu: func(r *git.Repo) MenuLevel {
		names, err := r.GetInactiveBranchList()
		if err != nil {
			return MenuLevel{Items: []MenuItem{GitErrorMenu(err)}}
		}

		all := make([]MenuItem, len(names))
		for i, name := range names {
			all[i] = singleBranchMergeItem(name)
		}

		fetch := branchFetch(all)
		level := MenuLevel{Items: fetch(0), Cursor: 0}
		if len(all) > logPageSize {
			level.Scroll = &ScrollState{Offset: 0, Fetch: fetch}
		}
		return level
	},
}
