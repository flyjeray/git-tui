package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

const logPageSize = 5

func logFetch(r *git.Repo) func(offset int) []MenuItem {
	return func(offset int) []MenuItem {
		entries, err := r.GetLog(offset, logPageSize)
		if err != nil {
			return []MenuItem{{
				Label:  "Error",
				Result: func(_ *git.Repo) string { return styled.Warn("error: " + err.Error()) },
			}}
		}
		items := make([]MenuItem, 0, len(entries))
		for _, e := range entries {
			e := e
			items = append(items, MenuItem{
				Label: fmt.Sprintf("%s - %s (by %s)", e.Hash, e.Subject, e.Author),
				Info:  func(_ *git.Repo) string { return styled.Hint(e.Hash) },
			})
		}
		return items
	}
}

var LogMenuItem = MenuItem{
	Label: "Log",
	LevelSubmenu: func(r *git.Repo) MenuLevel {
		fetch := logFetch(r)
		return MenuLevel{
			Items:  fetch(0),
			Cursor: 0,
			Scroll: &ScrollState{Offset: 0, Fetch: fetch},
		}
	},
}
