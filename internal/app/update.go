package app

import tea "charm.land/bubbletea/v2"

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.session.Width = msg.Width
		m.session.Height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "esc", "backspace":
			m.page = PageHome

		case "up", "k":
			if m.page == PageHome && m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.page == PageHome && m.cursor < len(m.menuItems())-1 {
				m.cursor++
			}

		case "enter":
			if m.page == PageHome {
				switch m.cursor {
				case 0:
					m.page = PageAbout
				case 1:
					m.page = PageProjects
				case 2:
					m.page = PageHomelab
				case 3:
					m.page = PageContact
				}
			}
		}
	}

	return m, nil
}
