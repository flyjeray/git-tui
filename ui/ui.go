package ui

import (
	git "git-tui/git-ops"

	tea "github.com/charmbracelet/bubbletea"
)

// menuLevel is one frame in the navigation stack.
type menuLevel struct {
	items  []MenuItem
	cursor int
}

// Model holds the state for the main TUI.
type Model struct {
	repo        *git.Repo
	repoWarning string
	stack       []menuLevel // navigation stack; top = current menu
	result      string      // non-empty while a result view is shown
}

func InitialModel(repo *git.Repo, repoWarning string) Model {
	return Model{
		repo:        repo,
		repoWarning: repoWarning,
		stack:       []menuLevel{{items: rootMenu}},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// top returns the current (deepest) menu level.
func (m Model) top() menuLevel {
	return m.stack[len(m.stack)-1]
}
