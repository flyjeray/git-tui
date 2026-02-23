package ui

import styled "git-tui/styles"

func (m Model) View() string {
	const sep = "──────────────────────────────────────────────"

	switch {
	case m.input != nil:
		return m.renderInput(sep) + "\n"
	case m.confirm != nil:
		return m.renderConfirm() + "\n"
	case m.result != "":
		content := m.result + "\n\n" + styled.Hint("esc: back  q: quit")
		return styled.Box(content) + "\n"
	}

	// Menu view
	var titleText string
	if m.repoWarning != "" {
		titleText = styled.Title("git-tui") + "  " + styled.Warn("⚠ not a git repository")
	} else {
		titleText = styled.Title("git-tui") + "  " + styled.Success("✓ git detected")
	}

	top := m.top()
	var menuLines string
	for i, item := range top.Items {
		label := item.Label
		if item.Submenu != nil {
			label += " ›"
		}
		var line string
		switch {
		case m.repo == nil && item.Submenu == nil:
			line = "  " + styled.Dim(label)
		case i == top.Cursor:
			line = styled.Selected("> " + label)
		default:
			line = "  " + label
		}
		menuLines += line + "\n"
	}

	var footerHint string
	if len(m.stack) > 1 {
		footerHint = "↑↓ / jk: navigate  enter: select  esc: back  q: quit"
	} else {
		footerHint = "↑↓ / jk: navigate  enter: select  q: quit"
	}

	content := titleText + "\n" + styled.Hint(sep) + "\n\n" + menuLines + "\n" + styled.Hint(footerHint)
	return styled.Box(content) + "\n"
}

func (m Model) renderInput(sep string) string {
	flow := m.input
	var lines string
	for i, step := range flow.Steps {
		switch {
		case i < flow.Current:
			// Completed: show label + value (dimmed)
			lines += styled.Hint(step.Label+": "+step.Value) + "\n"
		case i == flow.Current:
			// Active: show label + accumulated value + cursor
			hint := ""
			if step.Hint != "" {
				hint = "\n  " + styled.Dim(step.Hint)
			}
			lines += step.Label + ": " + step.Value + "_" + hint + "\n"
		default:
			// Upcoming: label only, dimmed
			lines += styled.Dim(step.Label+":") + "\n"
		}
	}

	footerVerb := "next"
	if flow.IsLast() {
		footerVerb = "submit"
	}
	footer := styled.Hint("esc: cancel  enter: " + footerVerb)

	content := styled.Title(flow.Title) + "\n" + styled.Hint(sep) + "\n\n" + lines + "\n" + footer
	return styled.Box(content)
}

func (m Model) renderConfirm() string {
	content := styled.Warn("⚠ "+m.confirm.Prompt) + "\n\n" + styled.Hint("y: confirm  n / esc: cancel")
	return styled.Box(content)
}
