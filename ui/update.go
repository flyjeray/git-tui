package ui

import (
	"time"

	menu "git-tui/menu"
	styled "git-tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

type resultMsg string
type tickMsg struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle async operation result.
	if r, ok := msg.(resultMsg); ok {
		m.result = string(r)
		m.loading = false
		return m, nil
	}

	// Advance spinner while loading.
	if _, ok := msg.(tickMsg); ok {
		if m.loading {
			m.spinnerFrame = (m.spinnerFrame + 1) % len(spinnerFrames)
			return m, tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg{} })
		}
		return m, nil
	}

	// Bracketed paste arrives as KeyMsg with Paste==true; handle it before the key switch.
	if key, ok := msg.(tea.KeyMsg); ok && key.Paste && m.input != nil {
		step := m.input.CurrentStep()
		for _, r := range key.Runes {
			if r >= ' ' { // filter control characters
				step.AppendText(string(r))
			}
		}
		return m, nil
	}

	key, ok := msg.(tea.KeyMsg)
	if !ok {
		return m, nil
	}

	// Only allow quit while loading; all other input is ignored.
	if m.loading {
		if key.String() == "ctrl+c" {
			return m, tea.Quit
		}
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
		m.stack = menu.GetStartMenu(m.repo)
		m.result = ""
	}
	return m, nil
}

func (m Model) updateConfirm(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch key.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "y", "Y":
		onYes := m.confirm.OnYes
		repo := m.repo
		m.confirm = nil
		m.loading = true
		m.stack = menu.GetStartMenu(m.repo)
		return m, tea.Batch(
			func() tea.Msg { return resultMsg(onYes(repo)) },
			tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg { return tickMsg{} }),
		)
	case "n", "N", "esc":
		m.confirm = nil
	}
	return m, nil
}

func (m Model) updateInput(key tea.KeyMsg) (tea.Model, tea.Cmd) {
	flow := m.input
	step := flow.CurrentStep()

	switch key.String() {
	case "ctrl+c":
		return m, tea.Quit
	case "esc":
		m.input = nil
	case "enter":
		if flow.IsLast() {
			m.result = flow.OnSubmit(m.repo, flow.CollectValues())
			m.input = nil
		} else {
			flow.Advance()
		}
	case "backspace":
		step.Backspace()
	default:
		if isPrintable(key.String()) {
			step.AppendText(key.String())
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
		m.stack[last].MoveUp()
	case "down", "j":
		m.stack[last].MoveDown()
	case "enter":
		m = m.activateItem(m.stack[last].Selected())
	}

	return m, nil
}

func (m Model) activateItem(item menu.MenuItem) Model {
	switch {
	case item.Submenu != nil:
		m.stack = append(m.stack, menu.MenuLevel{Items: item.Submenu(m.repo)})
	case m.repo == nil:
		m.result = styled.Warn("⚠ not in a git repository")
	case item.Result != nil:
		m.result = item.Result(m.repo)
	case item.Confirm != nil:
		m.confirm = item.Confirm
	case item.Flow != nil:
		m.input = item.Flow(m.repo)
	}
	return m
}

func isPrintable(s string) bool {
	return len(s) == 1 && s[0] >= ' ' && s[0] <= '~'
}
