package server

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
)


func TeaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// Create a lipgloss.Renderer for the session
	renderer := bubbletea.MakeRenderer(s)
	// Set up the model with the current session and styles.
	// We'll use the session to call wish.Command, which makes it compatible
	// with tea.Command.
	m := model{
		sess:     s,
		style:    renderer.NewStyle().Foreground(lipgloss.Color("8")),
		errStyle: renderer.NewStyle().Foreground(lipgloss.Color("3")),
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

type model struct {
	err      error
	sess     ssh.Session
	style    lipgloss.Style
	errStyle lipgloss.Style
}

func (m model) Init() tea.Cmd {
	return nil
}
