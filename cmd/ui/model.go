package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
		case tea.KeyBackspace:
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}
		default:
			m.input += msg.String()
		}
	}
	return m, nil
}

func (m model) View() string {
	return "🔍 Search Spotify: " + m.input + "\n(press Enter to search, q to quit)"
}

func StartUI() error {
	m := model{}
	p := tea.NewProgram(m)
	_, err := p.Run()
	return err
}
