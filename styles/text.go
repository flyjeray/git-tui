package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	successStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
	hintStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	warnStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	titleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	dimStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
	boxStyle      = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1, 2)
)

func Title(val string) string    { return titleStyle.Render(val) }
func Warn(val string) string     { return warnStyle.Render(val) }
func Success(val string) string  { return successStyle.Render(val) }
func Box(val string) string      { return boxStyle.Render(val) }
func Hint(val string) string     { return hintStyle.Render(val) }
func Dim(val string) string      { return dimStyle.Render(val) }
func Selected(val string) string { return selectedStyle.Render(val) }
