package menu

import (
	git "git-tui/git-ops"
)

var BranchRootMenu = MenuItem{
	Label: "Branches",
	Submenu: func(r *git.Repo) []MenuItem {
		return []MenuItem{BranchCurrentItem, BranchListMenu(r)}
	},
}
