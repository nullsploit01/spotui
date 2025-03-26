package ui

import "github.com/charmbracelet/lipgloss"

var (
	SearchInput = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFAA")).
			Border(lipgloss.ThickBorder()).
			Padding(0, 1)

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#00FFAA")).
			Border(lipgloss.NormalBorder()).
			Padding(0, 1)
)
