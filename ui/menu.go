package ui

import (
	"strings"

	git "git-tui/git-ops"
)

// MenuItem is a single entry in a navigable menu.
// Set Children for a submenu, Action for a leaf that runs a git operation.
type MenuItem struct {
	Label    string
	Action   func(*git.Repo) string // nil if submenu
	Children []MenuItem             // nil if leaf action
}

// rootMenu is the top-level menu tree shown on startup.
var rootMenu = []MenuItem{
	{
		Label: "Check current branch",
		Action: func(r *git.Repo) string {
			branch, err := r.Branch()
			if err != nil {
				return WarnStyle.Render("error: " + err.Error())
			}
			return "Current branch: " + HintStyle.Render(branch)
		},
	},
	{
		Label: "Remotes",
		Action: func(r *git.Repo) string {
			remotes, err := r.Remotes()
			if err != nil {
				return WarnStyle.Render("error: " + err.Error())
			}
			if len(remotes) == 0 {
				return HintStyle.Render("no remotes configured")
			}
			lines := make([]string, len(remotes))
			for i, remote := range remotes {
				lines[i] = HintStyle.Render(remote)
			}
			return "Remotes:\n" + strings.Join(lines, "\n")
		},
	},
}
