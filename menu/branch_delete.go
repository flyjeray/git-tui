package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func singleBranchDeleteItem(branch string) MenuItem {
	return MenuItem{
		Label: branch,
		Confirm: &ConfirmPrompt{
			Prompt: "Delete branch " + branch + "?",
			OnYes: func(r *git.Repo) string {
				if _, err := r.DeleteBranch(branch); err != nil {
					return styled.Warn("error: " + err.Error())
				}
				return styled.Success("✓ deleted " + branch)
			},
		},
	}
}

var BranchDeleteMenu = MenuItem{
	Label: "Delete",
	LevelSubmenu: func(r *git.Repo) MenuLevel {
		names, err := r.GetInactiveBranchList()
		if err != nil {
			return MenuLevel{Items: []MenuItem{GitErrorMenu(err)}}
		}

		all := make([]MenuItem, len(names))
		for i, name := range names {
			all[i] = singleBranchDeleteItem(name)
		}

		fetch := branchFetch(all)
		level := MenuLevel{Items: fetch(0), Cursor: 0}
		if len(all) > logPageSize {
			level.Scroll = &ScrollState{Offset: 0, Fetch: fetch}
		}
		return level
	},
}
