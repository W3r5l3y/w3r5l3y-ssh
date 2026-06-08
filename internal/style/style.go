package style

import "charm.land/lipgloss/v2"

var (
	Frame = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 3).
		Width(62)

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63"))

	Subtitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	Muted = lipgloss.NewStyle().
		Foreground(lipgloss.Color("245"))

	MenuItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252"))

	MenuSelected = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("86"))

	Heading = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("63"))

	Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color("244"))
)
