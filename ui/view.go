package ui

import (
	"strings"

	styled "git-tui/styles"
)

const sep = "──────────────────────────────────────────────"

func (m Model) View() string {

	// Menu view
	var titleText string
	if m.repoWarning != "" {
		titleText = styled.Title("git-tui") + "  " + styled.Warn("⚠ not a git repository")
	} else {
		titleText = styled.Title("git-tui") + "  " + styled.Success("✓ git detected")
	}

	header := titleText + "\n" + styled.Hint(sep) + "\n\n"

	switch {
	case m.loading:
		frame := spinnerFrames[m.spinnerFrame]
		return styled.Box(frame+" working...") + "\n"
	case m.input != nil:
		return m.renderInput() + "\n"
	case m.confirm != nil:
		return m.renderConfirm() + "\n"
	case m.result != "" || m.info != "":
		text := ""
		if m.result != "" {
			text = m.result
		} else {
			text = m.info
		}
		content := header + text + "\n\n" + styled.Hint("esc: back  q: quit")
		return styled.Box(content) + "\n"
	}

	top := m.top()
	var sb strings.Builder
	for i, item := range top.Items {
		label := item.DisplayLabel()
		switch {
		case m.repo == nil:
			sb.WriteString("  " + styled.Dim(label) + "\n")
		case i == top.Cursor:
			sb.WriteString(styled.Selected("> "+label) + "\n")
		default:
			sb.WriteString("  " + label + "\n")
		}
	}
	menuLines := sb.String()

	var footerHint string
	if len(m.stack) > 1 {
		footerHint = "↑↓ / jk: navigate  enter: select  esc: back  q: quit"
	} else {
		footerHint = "↑↓ / jk: navigate  enter: select  q: quit"
	}

	content := header + menuLines + "\n" + styled.Hint(footerHint)
	return styled.Box(content) + "\n"
}

func (m Model) renderInput() string {
	flow := m.input
	var sb strings.Builder
	for i, step := range flow.Steps {
		switch {
		case i < flow.Current:
			// Completed: show label + value (dimmed)
			sb.WriteString(styled.Hint(step.Label+": "+step.Value) + "\n")
		case i == flow.Current:
			// Active: show label + accumulated value + cursor
			hint := ""
			if step.Hint != "" {
				hint = "\n  " + styled.Dim(step.Hint)
			}
			sb.WriteString(step.Label + ": " + step.Value + "_" + hint + "\n")
		default:
			// Upcoming: label only, dimmed
			sb.WriteString(styled.Dim(step.Label+":") + "\n")
		}
	}
	lines := sb.String()

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
