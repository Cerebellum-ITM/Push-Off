package main

import (
	"os"

	"main/internal/app"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/log"
)

func main() {
	initialModel, err := app.NewPushOffModel()
	if err != nil {
		log.Fatal("An error occurred while trying to initialize the app model")
	}
	p := tea.NewProgram(initialModel, tea.WithOutput(os.Stderr))

	finalModel, err := p.Run()
	if err != nil {
		log.Fatal("Oh no! There was an error", "error", err)
	}

	if _, ok := finalModel.(*app.Model); ok {
		log.Info("The program has ended")
	}
}
