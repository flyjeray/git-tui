package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.result = ""
	case "up", "k":
		if m.result == "" && m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.result == "" && m.cursor < len(menuItems)-1 {
			m.cursor++
		}
	case "enter":
		if m.result == "" {
			m.result = m.runSelected()
		}
	}

	return m, nil
}

func (m Model) runSelected() string {
	if m.repo == nil {
		return WarnStyle.Render("⚠ not in a git repository")
	}
	switch m.cursor {
	case 0:
		branch, err := m.repo.Branch()
		if err != nil {
			return WarnStyle.Render("error: " + err.Error())
		}
		return "Current branch: " + HintStyle.Render(branch)
	case 1:
		remotes, err := m.repo.Remotes()
		if err != nil {
			return WarnStyle.Render("error: " + err.Error())
		}
		if len(remotes) == 0 {
			return HintStyle.Render("no remotes configured")
		}
		lines := make([]string, len(remotes))
		for i, r := range remotes {
			lines[i] = HintStyle.Render(r)
		}
		return "Remotes:\n" + strings.Join(lines, "\n")
	}
	return ""
}
