package ui

import "github.com/charmbracelet/lipgloss"

var (
	inputStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFAA")).
		Border(lipgloss.NormalBorder()).
		Width(70).
		Padding(0, 1)
)
