package ui

import "github.com/charmbracelet/lipgloss"

var (
	SuccessStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
	HintStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	WarnStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
	TitleStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("255"))
	SelectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true)
	DimStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("238"))
	BoxStyle      = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240")).
			Padding(1, 2)
)
