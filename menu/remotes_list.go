package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
	"strings"
)

func getRemotes(r *git.Repo) string {
	remotes, err := r.GetRemotesWithURLs()
	if err != nil {
		return styled.Warn("error: " + err.Error())
	}
	if len(remotes) == 0 {
		return styled.Hint("no remotes configured")
	}
	lines := make([]string, len(remotes))
	for i, remote := range remotes {
		lines[i] = styled.Hint(remote)
	}
	return "Remotes:\n" + strings.Join(lines, "\n")
}

var RemotesListItem = MenuItem{
	Label: "List",
	Info:  getRemotes,
}
