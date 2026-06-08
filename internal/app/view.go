package app

import (
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/W3r5l3y/w3r5l3y-ssh/internal/pages"
	"github.com/W3r5l3y/w3r5l3y-ssh/internal/style"
)

func (m Model) View() tea.View {
	var body string

	switch m.page {
	case PageHome:
		body = m.homeView()
	case PageAbout:
		body = pages.About()
	case PageProjects:
		body = pages.Projects()
	case PageHomelab:
		body = pages.Homelab()
	case PageContact:
		body = pages.Contact()
	default:
		body = m.homeView()
	}

	content := style.Frame.Render(body)

	view := tea.NewView(content)
	view.AltScreen = true

	return view
}

func (m Model) homeView() string {
	var b strings.Builder

	b.WriteString(style.Title.Render("w3r5l3y.com"))
	b.WriteString("\n\n")
	b.WriteString(style.Subtitle.Render("Cybersecurity | Linux | Networking | Software"))
	b.WriteString("\n")
	b.WriteString(style.Muted.Render("A personal site served over SSH."))
	b.WriteString("\n\n")

	for i, item := range m.menuItems() {
		cursor := "  "
		lineStyle := style.MenuItem

		if m.cursor == i {
			cursor = "› "
			lineStyle = style.MenuSelected
		}

		b.WriteString(lineStyle.Render(cursor + item))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(style.Help.Render("↑/↓ or j/k move    enter select    esc back    q quit"))

	return b.String()
}
