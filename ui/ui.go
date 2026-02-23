package ui

import (
	git "git-tui/git-ops"

	tea "github.com/charmbracelet/bubbletea"
)

// Model holds the state for the main TUI.
type Model struct {
	repo        *git.Repo
	repoWarning string
	cursor      int
	result      string
}

var menuItems = []string{
	"Check current branch",
	"Check remotes",
}

func InitialModel(repo *git.Repo, repoWarning string) Model {
	return Model{
		repo:        repo,
		repoWarning: repoWarning,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
