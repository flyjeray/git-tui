package ui

import "git-tui/styles"

func (m Model) View() string {
	const sep = "──────────────────────────────────────────────"

	switch {
	case m.input != nil:
		return m.renderInput(sep) + "\n"
	case m.confirm != nil:
		return m.renderConfirm() + "\n"
	case m.result != "":
		content := m.result + "\n\n" + styles.HintStyle.Render("esc: back  q: quit")
		return styles.BoxStyle.Render(content) + "\n"
	}

	// Menu view
	var titleText string
	if m.repoWarning != "" {
		titleText = styles.TitleStyle.Render("git-tui") + "  " + styles.WarnStyle.Render("⚠ not a git repository")
	} else {
		titleText = styles.TitleStyle.Render("git-tui") + "  " + styles.SuccessStyle.Render("✓ git detected")
	}

	top := m.top()
	var menuLines string
	for i, item := range top.Items {
		label := item.Label
		if item.Children != nil || item.Submenu != nil {
			label += " ›"
		}
		var line string
		switch {
		case m.repo == nil && item.Operation != nil:
			line = "  " + styles.DimStyle.Render(label)
		case i == top.Cursor:
			line = styles.SelectedStyle.Render("> " + label)
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

	content := titleText + "\n" + styles.HintStyle.Render(sep) + "\n\n" + menuLines + "\n" + styles.HintStyle.Render(footerHint)
	return styles.BoxStyle.Render(content) + "\n"
}

func (m Model) renderInput(sep string) string {
	flow := m.input
	var lines string
	for i, step := range flow.Steps {
		switch {
		case i < flow.Current:
			// Completed: show label + value (dimmed)
			display := step.Value
			for _, opt := range step.Options {
				if opt.Value == step.Value {
					display = step.Value + " — " + opt.Label
					break
				}
			}
			lines += styles.HintStyle.Render(step.Label+": "+display) + "\n"

		case i == flow.Current:
			if step.Options != nil {
				// Active select: show label then navigable options list
				lines += step.Label + ":\n"
				for j, opt := range step.Options {
					if j == step.Cursor {
						lines += styles.SelectedStyle.Render("> "+opt.Value+" — "+opt.Label) + "\n"
					} else {
						lines += "  " + styles.DimStyle.Render(opt.Value+" — "+opt.Label) + "\n"
					}
				}
			} else {
				// Active text field: show label + accumulated value + cursor
				hint := ""
				if step.Hint != "" {
					hint = "\n  " + styles.DimStyle.Render(step.Hint)
				}
				lines += step.Label + ": " + step.Value + "_" + hint + "\n"
			}

		default:
			// Upcoming: label only, dimmed
			lines += styles.DimStyle.Render(step.Label+":") + "\n"
		}
	}

	footerVerb := "next"
	if flow.Current == len(flow.Steps)-1 {
		footerVerb = "submit"
	}
	footer := styles.HintStyle.Render("esc: cancel  enter: " + footerVerb)

	content := styles.TitleStyle.Render(flow.Title) + "\n" + styles.HintStyle.Render(sep) + "\n\n" + lines + "\n" + footer
	return styles.BoxStyle.Render(content)
}

func (m Model) renderConfirm() string {
	content := styles.WarnStyle.Render("⚠ "+m.confirm.Prompt) + "\n\n" +
		styles.HintStyle.Render("y: confirm  n / esc: cancel")
	return styles.BoxStyle.Render(content)
}
