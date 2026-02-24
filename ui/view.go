package ui

import (
	"strings"

	styled "git-tui/styles"
)

const sep = "──────────────────────────────────────────────"

func (m Model) View() string {
	return styled.Box(m.renderHeader()+m.renderBody()) + "\n"
}

func (m Model) renderHeader() string {
	var title string
	var subtitle string
	length := len(m.stack)
	if length > 1 {
		for i := range m.stack[:length-1] {
			title += m.stack[i].Selected().Label
			if i < length-2 {
				title += " > "
			}
		}
	} else {
		title = "git-tui"
		if m.repoWarning != "" {
			subtitle = "  " + styled.Warn("⚠ not a git repository")
		} else {
			subtitle = "  " + styled.Success("✓ git detected")
		}
	}

	return styled.Title(title) + subtitle + "\n" + styled.Hint(sep) + "\n\n"
}

func (m Model) renderBody() string {
	switch {
	case m.loading:
		frame := spinnerFrames[m.spinnerFrame]
		return frame + " working..."
	case m.input != nil:
		return m.renderInput()
	case m.confirm != nil:
		return m.renderConfirm()
	case m.result != "" || m.info != "":
		text := m.result
		if text == "" {
			text = m.info
		}
		return text + "\n\n" + styled.Hint("esc: back  q: quit")
	default:
		return m.renderMenu()
	}
}

func (m Model) renderMenu() string {
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

	var footerHint string
	if len(m.stack) > 1 {
		footerHint = "↑↓ / jk: navigate  enter: select  esc: back  q: quit"
	} else {
		footerHint = "↑↓ / jk: navigate  enter: select  q: quit"
	}

	return sb.String() + "\n" + styled.Hint(footerHint)
}

func (m Model) renderInput() string {
	flow := m.input
	var sb strings.Builder
	for i, step := range flow.Steps {
		switch {
		case i < flow.Current:
			sb.WriteString(styled.Hint(step.Label+": "+step.Value) + "\n")
		case i == flow.Current:
			hint := ""
			if step.Hint != "" {
				hint = "\n  " + styled.Dim(step.Hint)
			}
			sb.WriteString(step.Label + ": " + step.Value + "_" + hint + "\n")
		default:
			sb.WriteString(styled.Dim(step.Label+":") + "\n")
		}
	}

	footerVerb := "next"
	if flow.IsLast() {
		footerVerb = "submit"
	}

	return sb.String() + "\n" + styled.Hint("esc: cancel  enter:"+footerVerb)
}

func (m Model) renderConfirm() string {
	return styled.Warn("⚠ "+m.confirm.Prompt) + "\n\n" + styled.Hint("y: confirm  n / esc: cancel")
}
