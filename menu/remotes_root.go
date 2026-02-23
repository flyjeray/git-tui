package menu

import git "git-tui/git-ops"

var RemotesMenuItem = MenuItem{
	Label: "Remotes",
	Submenu: func(r *git.Repo) []MenuItem {
		return []MenuItem{RemotesListItem, RemotesAddItem, RemoteDeleteMenu(r)}
	},
}
