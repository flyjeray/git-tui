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

	var menuLines string
	for i, item := range menuItems {
		var line string
		switch {
		case m.repo == nil:
			line = "  " + DimStyle.Render(item)
		case i == m.cursor:
			line = SelectedStyle.Render("> " + item)
		default:
			line = "  " + item
		}
		menuLines += line + "\n"
	}

	footer := HintStyle.Render("↑↓ / jk: navigate  enter: select  q: quit")

	content := titleText + "\n" + HintStyle.Render(sep) + "\n\n" + menuLines + "\n" + footer
	return BoxStyle.Render(content) + "\n"
}
