package ui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	last := len(m.stack) - 1

	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		if m.result != "" {
			m.result = ""
		} else if last > 0 {
			m.stack = m.stack[:last]
		}
	case "up", "k":
		if m.result == "" && m.stack[last].cursor > 0 {
			m.stack[last].cursor--
		}
	case "down", "j":
		if m.result == "" && m.stack[last].cursor < len(m.stack[last].items)-1 {
			m.stack[last].cursor++
		}
	case "enter":
		if m.result != "" {
			break
		}
		item := m.stack[last].items[m.stack[last].cursor]
		switch {
		case item.Children != nil:
			m.stack = append(m.stack, menuLevel{items: item.Children})
		case item.Action != nil:
			if m.repo == nil {
				m.result = WarnStyle.Render("⚠ not in a git repository")
			} else {
				m.result = item.Action(m.repo)
			}
		}
	}

	return m, nil
}
