package main

import "github.com/charmbracelet/lipgloss"

var (
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("82")).Bold(true)
	hintStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	warnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("214"))
)
