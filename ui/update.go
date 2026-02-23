package ui

import (
	menu "git-tui/menu"
	"git-tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Bracketed paste arrives as KeyMsg with Paste==true; handle it before the key switch.
	if key, ok := msg.(tea.KeyMsg); ok && key.Paste && m.input != nil {
		step := &m.input.Steps[m.input.Current]
		if step.Options == nil { // only paste into text fields, not selects
			for _, r := range key.Runes {
				if r >= ' ' { // filter control characters
					step.Value += string(r)
				}
			}
		}
		return m, nil
	}

	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	switch {
	case m.input != nil:
		return m.updateInput(key)
	case m.confirm != nil:
		return m.updateConfirm(key)
	case m.result != "":
		return m.updateResult(key)
	default:
		return m.updateMenu(key)
	}
}

func (m Model) updateResult(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		m.result = ""
	}
	return m, nil
}

func (m Model) updateConfirm(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "y", "Y":
		m.result = m.confirm.OnYes(m.repo)
		m.confirm = nil
	case "n", "N", "esc":
		m.confirm = nil
	}
	return m, nil
}

func (m Model) updateInput(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	flow := m.input
	step := &flow.Steps[flow.Current]

	switch key.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.input = nil
	case "up", "k":
		if step.Options != nil && step.Cursor > 0 {
			step.Cursor--
		}
	case "down", "j":
		if step.Options != nil && step.Cursor < len(step.Options)-1 {
			step.Cursor++
		}
	case "enter":
		if step.Options != nil {
			step.Value = step.Options[step.Cursor].Value
		}
		if flow.Current < len(flow.Steps)-1 {
			flow.Current++
		} else {
			Values := make([]string, len(flow.Steps))
			for i, s := range flow.Steps {
				Values[i] = s.Value
			}
			m.result = flow.OnSubmit(m.repo, Values)
			m.input = nil
		}
	case "backspace":
		if step.Options == nil && len(step.Value) > 0 {
			step.Value = step.Value[:len(step.Value)-1]
		}
	default:
		if step.Options == nil && isPrintable(key.String()) {
			step.Value += key.String()
		}
	}

	return m, nil
}

func (m Model) updateMenu(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	last := len(m.stack) - 1

	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "esc":
		if last > 0 {
			m.stack = m.stack[:last]
		}
	case "up", "k":
		if m.stack[last].Cursor > 0 {
			m.stack[last].Cursor--
		}
	case "down", "j":
		if m.stack[last].Cursor < len(m.stack[last].Items)-1 {
			m.stack[last].Cursor++
		}
	case "enter":
		item := m.stack[last].Items[m.stack[last].Cursor]
		m = m.activateItem(item)
	}

	return m, nil
}

func (m Model) activateItem(item menu.MenuItem) Model {
	switch {
	case item.Children != nil:
		m.stack = append(m.stack, menu.MenuLevel{Items: item.Children})
	case item.Submenu != nil:
		if m.repo == nil {
			m.result = styles.WarnStyle.Render("⚠ not in a git repository")
		} else {
			m.stack = append(m.stack, menu.MenuLevel{Items: item.Submenu(m.repo)})
		}
	case item.Operation != nil:
		if m.repo == nil {
			m.result = styles.WarnStyle.Render("⚠ not in a git repository")
		} else {
			op := item.Operation(m.repo)
			m.result = op.Title
			m.confirm = op.ConfirmPrompt
			m.input = op.Flow
		}
	}
	return m
}

func isPrintable(s string) bool {
	return len(s) == 1 && s[0] >= ' ' && s[0] <= '~'
}
