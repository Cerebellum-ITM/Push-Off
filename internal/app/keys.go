package app

import "github.com/charmbracelet/bubbles/v2/key"

type KeyMap struct {
	Quit key.Binding
}

func mainMenuKeys() KeyMap {
	return KeyMap{
		Quit: key.NewBinding(key.WithKeys("ctrl+x"), key.WithHelp("ctrl+x", "quit")),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	b := []key.Binding{}
	if k.Quit.Enabled() {
		b = append(b, k.Quit)
	}
	return b
}

func (k KeyMap) FullHelp() [][]key.Binding {
	b := []key.Binding{}
	if k.Quit.Enabled() {
		b = append(b, k.Quit)
	}
	return [][]key.Binding{b}
}
