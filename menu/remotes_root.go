package menu

import git "git-tui/git-ops"

var RemotesMenuItem = MenuItem{
	Label: "Remotes",
	Submenu: func(_ *git.Repo) []MenuItem {
		return []MenuItem{RemotesListItem, RemotesAddItem, RemotesDeleteItem}
	},
}
