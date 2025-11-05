package app

import tea "github.com/charmbracelet/bubbletea/v2"

type appState int

type Model struct {
	State appState
	Keys  KeyMap
}

const (
	stateChoosingRepo appState = iota
)

func NewPushOffModel() (*Model, error) {
	m := &Model{
		State: stateChoosingRepo,
		Keys: mainMenuKeys(),
	}
	return m, nil
}

func (model *Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}
