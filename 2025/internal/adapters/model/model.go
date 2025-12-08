package model

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	Days   []string
	cursor int
}

func New() *Model {
	return &Model{
		Days: []string{
			"01",
		},
		cursor: 0,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
