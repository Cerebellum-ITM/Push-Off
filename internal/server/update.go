package server

import (
	"runtime"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/x/editor"
)

type cmdFinishedMsg struct{ err error }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "e":
			// Open file.txt in the default editor.
			edit, err := editor.Cmd("wish", "file.txt")
			if err != nil {
				m.err = err
				return m, nil
			}
			// Creates a wish.Cmd from the exec.Cmd
			wishCmd := wish.Command(m.sess, edit.Path, edit.Args...)
			// Runs the cmd through Bubble Tea.
			// Bubble Tea should handle the IO to the program, and get it back
			// once the program quits.
			cmd := tea.Exec(wishCmd, func(err error) tea.Msg {
				if err != nil {
					log.Error("editor finished", "error", err)
				}
				return cmdFinishedMsg{err: err}
			})
			return m, cmd
		case "s":
			// We can also execute a shell and give it over to the user.
			// Note that this session won't have control, so it can't run tasks
			// in background, suspend, etc.
			c := wish.Command(m.sess, "sh")
			if runtime.GOOS == "windows" {
				c = wish.Command(m.sess, "powershell")
			}
			cmd := tea.Exec(c, func(err error) tea.Msg {
				if err != nil {
					log.Error("shell finished", "error", err)
				}
				return cmdFinishedMsg{err: err}
			})
			return m, cmd
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case cmdFinishedMsg:
		m.err = msg.err
		return m, nil
	}

	return m, nil
}
