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

func branchFetch(all []MenuItem) func(offset int) []MenuItem {
	return func(offset int) []MenuItem {
		if offset >= len(all) {
			return nil
		}
		end := offset + logPageSize
		if end > len(all) {
			end = len(all)
		}
		return all[offset:end]
	}
}

func BranchListMenu(r *git.Repo) MenuItem {
	return MenuItem{
		Label: "Checkout",
		LevelSubmenu: func(r *git.Repo) MenuLevel {
			names, err := r.GetInactiveBranchList()
			if err != nil {
				return MenuLevel{Items: []MenuItem{GitErrorMenu(err)}}
			}

			all := make([]MenuItem, len(names))
			for i, name := range names {
				all[i] = singleBranchItem(name)
			}

			fetch := branchFetch(all)
			level := MenuLevel{Items: fetch(0), Cursor: 0}
			if len(all) > logPageSize {
				level.Scroll = &ScrollState{Offset: 0, Fetch: fetch}
			}
			return level
		},
	}
}
