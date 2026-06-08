package app

type Page int

const (
	PageHome Page = iota
	PageAbout
	PageProjects
	PageHomelab
	PageContact
)

type SessionInfo struct {
	User   string
	Term   string
	Width  int
	Height int
}

type Model struct {
	session SessionInfo
	cursor  int
	page    Page
}

func NewModel(session SessionInfo) Model {
	return Model{
		session: session,
		cursor:  0,
		page:    PageHome,
	}
}

func (m Model) menuItems() []string {
	return []string{
		"About",
		"Projects",
		"Homelab",
		"Contact",
	}
}
