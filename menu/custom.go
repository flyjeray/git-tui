package menu

import (
	"strings"

	git "git-tui/git-ops"
	styled "git-tui/styles"
)

var CustomCommandMenuItem = MenuItem{
	Label: "Run custom command",
	Flow: func(_ *git.Repo) *InputFlow {
		return &InputFlow{
			Title: "Custom git command",
			Steps: []InputStep{
				{Label: "git", Hint: "e.g. status, log --oneline -5, diff HEAD~1"},
			},
			OnSubmit: func(r *git.Repo, values []string) string {
				args := strings.Fields(values[0])
				if len(args) == 0 {
					return styled.Warn("no command entered")
				}
				out, err := r.RunCommand(args...)
				if err != nil {
					return styled.Warn(out)
				}
				if out == "" {
					return styled.Hint("(no output)")
				}
				return out
			},
		}
	},
}
