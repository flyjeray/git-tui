package menu

import (
	"fmt"
	git "git-tui/git-ops"
	styles "git-tui/styles"
	urltools "net/url"
)

func createRemote(r *git.Repo, values []string) string {
	name, url := values[0], values[1]
	if _, err := urltools.ParseRequestURI(url); err != nil {
		return styles.WarnStyle.Render("error: URL is not valid url")
	}
	if err := r.AddRemote(name, url); err != nil {
		return styles.WarnStyle.Render("error: " + err.Error())
	}
	return styles.SuccessStyle.Render(fmt.Sprintf("✓ added remote %q → %s", name, url))

}

var RemotesAddItem = MenuItem{
	Label: "Add",
	Operation: func(_ *git.Repo) Operation {
		return Operation{Flow: &InputFlow{
			Title: "Add remote",
			Steps: []inputStep{
				{Label: "Name"},
				{Label: "URL"},
			},
			OnSubmit: createRemote,
		}}
	},
}
