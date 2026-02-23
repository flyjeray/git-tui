package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styled "git-tui/styles"
	urltools "net/url"
)

func createRemote(r *git.Repo, values []string) string {
	name, url := values[0], values[1]
	if _, err := urltools.ParseRequestURI(url); err != nil {
		return styled.Warn("error: URL is not valid url")
	}
	if err := r.AddRemote(name, url); err != nil {
		return styled.Warn("error: " + err.Error())
	}
	return styled.Success(fmt.Sprintf("✓ added remote %q → %s", name, url))
}

var RemotesAddItem = MenuItem{
	Label: "Add",
	Flow: func(_ *git.Repo) *InputFlow {
		return &InputFlow{
			Title: "Add remote",
			Steps: []InputStep{
				{Label: "Name"},
				{Label: "URL"},
			},
			OnSubmit: createRemote,
		}
	},
}
