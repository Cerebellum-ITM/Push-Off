package server

func (m model) View() string {
	if m.err != nil {
		return m.errStyle.Render(m.err.Error() + "\n")
	}
	return m.style.Render("Press 'e' to edit, 's' to hop into a shell, or 'q' to quit...\n")
}
