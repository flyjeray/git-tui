package menu

import (
	git "git-tui/git-ops"
)

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

var BranchRootMenu = MenuItem{
	Label: "Branches",
	Submenu: func(r *git.Repo) []MenuItem {
		return []MenuItem{BranchCurrentItem, BranchListMenu(r), BranchAddItem, BranchDeleteMenu}
	},
}
