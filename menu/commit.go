package menu

import (
	git "git-tui/git-ops"
	styled "git-tui/styles"
)

func commitProcess(r *git.Repo, values []string) string {
	var commit = values[0]

	if commit == "" {
		commit = "Update"
	}

	var _, stageErr = r.StageAll()

	if stageErr != nil {
		return styled.Warn("staging error: " + stageErr.Error())
	}

	var _, commitErr = r.Commit(commit)

	if commitErr != nil {
		return styled.Warn("commit error: " + commitErr.Error())
	}

	return styled.Success("Updates are committed!")
}

func singleOriginCommit() MenuItem {
	return MenuItem{
		Label: "Commit",
		Flow: func(_ *git.Repo) *InputFlow {
			return &InputFlow{
				Title: "Commit details",
				Steps: []InputStep{
					{Label: "Title"},
				},
				OnSubmit: func(r *git.Repo, values []string) string {
					return commitProcess(r, values)
				},
			}
		},
	}
}

func CommitMenuItem(r *git.Repo) MenuItem {
	return singleOriginCommit()
}
