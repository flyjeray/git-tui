package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func GitErrorMenu(err error) MenuItem {
	return MenuItem{
		Label:  "Pull",
		Result: func(_ *git.Repo) string { return styled.Warn("git error: " + err.Error()) },
	}
}

// rootMenu is the top-level menu tree shown on startup.
func rootMenu(r *git.Repo) []MenuItem {
	return []MenuItem{
		PullMenuItem(r),
		CommitMenuItem(r),
		PushMenuItem(r),
		BranchRootMenu,
		RemotesMenuItem,
		LogMenuItem,
		CustomCommandMenuItem,
	}
}

func GetStartMenu(r *git.Repo) []MenuLevel {
	return []MenuLevel{{
		Items:  rootMenu(r),
		Cursor: 0,
	}}
}
