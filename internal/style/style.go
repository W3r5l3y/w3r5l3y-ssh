package style

import (
	"os"

	"charm.land/lipgloss/v2"
)

type Theme struct {
	Name      string
	Accent    string
	AccentAlt string
	Text      string
	Muted     string
	VeryMuted string
	Success   string
	Warning   string
	Error     string
}

var themes = map[string]Theme{
	"violet": {
		Name:      "violet",
		Accent:    "63",
		AccentAlt: "86",
		Text:      "252",
		Muted:     "245",
		VeryMuted: "240",
		Success:   "82",
		Warning:   "220",
		Error:     "196",
	},
	"cyan": {
		Name:      "cyan",
		Accent:    "86",
		AccentAlt: "63",
		Text:      "252",
		Muted:     "245",
		VeryMuted: "240",
		Success:   "82",
		Warning:   "220",
		Error:     "196",
	},
	"amber": {
		Name:      "amber",
		Accent:    "214",
		AccentAlt: "220",
		Text:      "252",
		Muted:     "245",
		VeryMuted: "240",
		Success:   "82",
		Warning:   "220",
		Error:     "196",
	},
	"rose": {
		Name:      "rose",
		Accent:    "211",
		AccentAlt: "219",
		Text:      "252",
		Muted:     "245",
		VeryMuted: "240",
		Success:   "82",
		Warning:   "220",
		Error:     "196",
	},
	"green": {
		Name:      "green",
		Accent:    "82",
		AccentAlt: "86",
		Text:      "252",
		Muted:     "245",
		VeryMuted: "240",
		Success:   "82",
		Warning:   "220",
		Error:     "196",
	},
	"mono": {
		Name:      "mono",
		Accent:    "252",
		AccentAlt: "245",
		Text:      "252",
		Muted:     "245",
		VeryMuted: "240",
		Success:   "252",
		Warning:   "245",
		Error:     "245",
	},
}

var Current = selectTheme()

func selectTheme() Theme {
	name := os.Getenv("SSH_THEME")

	if name == "" {
		name = "violet"
	}

	theme, ok := themes[name]
	if !ok {
		return themes["violet"]
	}

	return theme
}

var (
	Frame = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(Current.Accent)).
		Padding(1, 3)

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(Current.Accent))

	Subtitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Text))

	Muted = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Current.Muted))

	VeryMutedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.VeryMuted))

	AccentText = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(Current.AccentAlt))

	MenuItem = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Text))

	MenuSelected = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(Current.AccentAlt))

	Heading = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(Current.Accent))

	Label = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Current.Muted))

	Value = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Current.Text))

	Help = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Current.VeryMuted))

	Divider = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Current.VeryMuted))

	SuccessText = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Success))

	WarningText = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Warning))

	ErrorText = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Error))

	Link = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(Current.AccentAlt)).
		Underline(true)
)

func Hyperlink(label string, url string) string {
	return "\x1b]8;;" + url + "\x1b\\" + Link.Render(label) + "\x1b]8;;\x1b\\"
}
