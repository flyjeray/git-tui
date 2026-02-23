package menu

import git "git-tui/git-ops"

func RemotesMenuItem(r *git.Repo) MenuItem {
	return MenuItem{
		Label: "Remotes",
		Submenu: func(_ *git.Repo) []MenuItem {
			return []MenuItem{RemotesListItem, RemotesAddItem, RemoteDeleteMenu(r)}
		},
	}
}
