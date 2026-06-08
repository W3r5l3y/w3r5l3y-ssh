package app

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/W3r5l3y/w3r5l3y-ssh/internal/content"
)

type animationTickMsg time.Time

func animationTick() tea.Cmd {
	return tea.Tick(650*time.Millisecond, func(t time.Time) tea.Msg {
		return animationTickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return animationTick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.session.Width = msg.Width
		m.session.Height = msg.Height

	case animationTickMsg:
		m.frame++
		return m, animationTick()

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "?":
			if m.page != PageHelp {
				m.previous = m.page
				m.page = PageHelp
				m.scroll = 0
			}

		case "esc", "backspace", "b":
			m.goBack()

		case "enter":
			m.handleEnter()

		case "up", "k":
			m.handleUp()

		case "down", "j":
			m.handleDown()

		case "pgup":
			m.scroll -= 8
			if m.scroll < 0 {
				m.scroll = 0
			}

		case "pgdown":
			m.scroll += 8

		case "home":
			m.scroll = 0

		case "end":
			m.scroll = 999
		}
	}

	return m, nil
}

func (m *Model) handleEnter() {
	switch m.page {
	case PageHome:
		switch m.cursor {
		case 0:
			m.openPage(PageAbout)
		case 1:
			m.openPage(PageProjects)
		case 2:
			m.openPage(PageSource)
		case 3:
			m.openPage(PageContact)
		}

	case PageProjects:
		if m.projectCursor >= 0 && m.projectCursor < len(content.Projects) {
			m.selectedIndex = m.projectCursor
			m.openPage(PageProjectDetail)
		}
	}
}

func (m *Model) handleUp() {
	switch m.page {
	case PageHome:
		if m.cursor > 0 {
			m.cursor--
		}

	case PageProjects:
		if m.projectCursor > 0 {
			m.projectCursor--
		}

	default:
		if m.scroll > 0 {
			m.scroll--
		}
	}
}

func (m *Model) handleDown() {
	switch m.page {
	case PageHome:
		if m.cursor < len(m.menuItems())-1 {
			m.cursor++
		}

	case PageProjects:
		if m.projectCursor < len(content.Projects)-1 {
			m.projectCursor++
		}

	default:
		m.scroll++
	}
}
