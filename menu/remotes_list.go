package menu

import (
	git "git-tui/git-ops"
	styles "git-tui/styles"
	"strings"
)

func getRemotes(r *git.Repo) Operation {
	remotes, err := r.GetRemotesWithURLs()

	if err != nil {
		return Operation{Title: styles.WarnStyle.Render("error: " + err.Error())}
	}
	if len(remotes) == 0 {
		return Operation{Title: styles.HintStyle.Render("no remotes configured")}
	}

	lines := make([]string, len(remotes))
	for i, remote := range remotes {
		lines[i] = styles.HintStyle.Render(remote)
	}

	return Operation{Title: "Remotes:\n" + strings.Join(lines, "\n")}
}

var RemotesListItem = MenuItem{
	Label:     "List",
	Operation: getRemotes,
}
