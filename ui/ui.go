package ui

import (
	git "git-tui/git-ops"
	menu "git-tui/menu"

	tea "github.com/charmbracelet/bubbletea"
)

// Model holds the state for the main TUI.
type Model struct {
	repo        *git.Repo
	repoWarning string
	stack       []menu.MenuLevel    // navigation stack; top = current menu
	result      string              // non-empty = result view
	confirm     *menu.ConfirmPrompt // non-nil = confirmation view
	input       *menu.InputFlow     // non-nil = input form view
}

func InitialModel(repo *git.Repo, repoWarning string) Model {
	return Model{
		repo:        repo,
		repoWarning: repoWarning,
		stack:       menu.StartMenu,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// top returns the current (deepest) menu level.
func (m Model) top() menu.MenuLevel {
	return m.stack[len(m.stack)-1]
}
