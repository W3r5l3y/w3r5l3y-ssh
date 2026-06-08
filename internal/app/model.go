package app

type Page int

const (
	PageHome Page = iota
	PageAbout
	PageProjects
	PageProjectDetail
	PageSource
	PageContact
	PageHelp
)

type SessionInfo struct {
	User   string
	Term   string
	Width  int
	Height int
}

type Model struct {
	session SessionInfo

	page     Page
	previous Page

	cursor        int
	projectCursor int
	selectedIndex int

	scroll int
	frame  int
}

func NewModel(session SessionInfo) Model {
	return Model{
		session:       session,
		page:          PageHome,
		previous:      PageHome,
		cursor:        0,
		projectCursor: 0,
		selectedIndex: 0,
		scroll:        0,
		frame:         0,
	}
}

func (m Model) menuItems() []string {
	return []string{
		"About",
		"Projects",
		"Source",
		"Contact",
	}
}

func (m Model) currentPageName() string {
	switch m.page {
	case PageHome:
		return "home"
	case PageAbout:
		return "about"
	case PageProjects:
		return "projects"
	case PageProjectDetail:
		return "project"
	case PageSource:
		return "source"
	case PageContact:
		return "contact"
	case PageHelp:
		return "help"
	default:
		return "home"
	}
}

func (m *Model) openPage(page Page) {
	m.previous = m.page
	m.page = page
	m.scroll = 0
}

func (m *Model) goBack() {
	switch m.page {
	case PageHome:
		return
	case PageProjectDetail:
		m.page = PageProjects
	case PageHelp:
		m.page = m.previous
	default:
		m.page = PageHome
	}

	m.scroll = 0
}
