package ui

func (m Model) View() string {
	const sep = "──────────────────────────────────────────────"

	if m.result != "" {
		content := m.result + "\n\n" + HintStyle.Render("esc: back  q: quit")
		return BoxStyle.Render(content) + "\n"
	}

	var titleText string
	if m.repoWarning != "" {
		titleText = TitleStyle.Render("git-tui") + "  " + WarnStyle.Render("⚠ not a git repository")
	} else {
		titleText = TitleStyle.Render("git-tui") + "  " + SuccessStyle.Render("✓ git detected")
	}

	top := m.top()
	var menuLines string
	for i, item := range top.items {
		label := item.Label
		if item.Children != nil {
			label += " ›"
		}
		var line string
		switch {
		case m.repo == nil && item.Action != nil:
			line = "  " + DimStyle.Render(label)
		case i == top.cursor:
			line = SelectedStyle.Render("> " + label)
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

	content := titleText + "\n" + HintStyle.Render(sep) + "\n\n" + menuLines + "\n" + HintStyle.Render(footerHint)
	return BoxStyle.Render(content) + "\n"
}
