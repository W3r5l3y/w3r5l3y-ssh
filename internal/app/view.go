package app

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/W3r5l3y/w3r5l3y-ssh/internal/content"
	"github.com/W3r5l3y/w3r5l3y-ssh/internal/style"
	"github.com/muesli/reflow/wordwrap"
)

func (m Model) View() tea.View {
	width := m.session.Width
	height := m.session.Height

	if width <= 0 {
		width = 80
	}

	if height <= 0 {
		height = 24
	}

	var output string

	if width < 70 || height < 20 {
		output = m.renderTooSmall(width, height)
	} else {
		output = m.renderApp(width, height)
	}

	view := tea.NewView(output)
	view.AltScreen = true

	return view
}

func (m Model) liveIndicator() string {
	frames := []string{
		style.SuccessText.Render("● live"),
		style.AccentText.Render("● live"),
		style.Muted.Render("● live"),
		style.AccentText.Render("● live"),
	}

	return frames[m.frame%len(frames)]
}

func (m Model) renderTooSmall(width int, height int) string {
	body := style.Title.Render("terminal too small") +
		"\n\n" +
		style.Muted.Render(fmt.Sprintf("current size: %dx%d", width, height)) +
		"\n" +
		style.Muted.Render("minimum size: 70x20") +
		"\n\n" +
		style.Help.Render("resize terminal or press q to quit")

	panel := style.Frame.Width(46).Render(body)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, panel)
}

func (m Model) renderApp(width int, height int) string {
	panelWidth := clamp(width-6, 70, 92)
	contentWidth := panelWidth - 8
	bodyHeight := clamp(height-12, 6, 18)

	header := m.renderHeader(contentWidth)
	body := m.renderBody(contentWidth, bodyHeight)
	footer := m.renderFooter()

	divider := style.Divider.Render(strings.Repeat("─", contentWidth))

	panelContent := strings.Join([]string{
		header,
		divider,
		body,
		divider,
		footer,
	}, "\n")

	panel := style.Frame.Width(panelWidth).Render(panelContent)

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, panel)
}

func (m Model) renderHeader(width int) string {
	left := style.Title.Render("w3r5l3y.com")
	rightText := "guest"

	if m.session.User != "" {
		rightText = m.session.User
	}

	right := style.Muted.Render(rightText + " / " + m.currentPageName())

	rawLeft := "w3r5l3y.com"
	rawRight := rightText + " / " + m.currentPageName()

	gap := width - lipgloss.Width(rawLeft) - lipgloss.Width(rawRight)
	if gap < 1 {
		gap = 1
	}

	return left + strings.Repeat(" ", gap) + right
}

func (m Model) renderFooter() string {
	switch m.page {
	case PageHome:
		return style.Help.Render("↑/↓ or j/k move    enter select    ? help    q quit")
	case PageProjects:
		return style.Help.Render("↑/↓ or j/k move    enter open    esc back    ? help    q quit")
	case PageProjectDetail:
		return style.Help.Render("↑/↓ scroll    pgup/pgdn    esc projects    q quit")
	default:
		return style.Help.Render("↑/↓ scroll    pgup/pgdn    esc back    ? help    q quit")
	}
}

func (m Model) renderBody(width int, bodyHeight int) string {
	switch m.page {
	case PageHome:
		return m.renderHome(width, bodyHeight)
	case PageAbout:
		return m.renderScrollable(m.aboutLines(), width, bodyHeight)
	case PageProjects:
		return m.renderProjects(width, bodyHeight)
	case PageProjectDetail:
		return m.renderScrollable(m.projectDetailLines(), width, bodyHeight)
	case PageSource:
		return m.renderScrollable(m.sourceLines(), width, bodyHeight)
	case PageContact:
		return m.renderScrollable(m.contactLines(), width, bodyHeight)
	case PageHelp:
		return m.renderScrollable(m.helpLines(), width, bodyHeight)
	default:
		return m.renderHome(width, bodyHeight)
	}
}

func (m Model) renderHome(width int, bodyHeight int) string {
	var lines []string

	lines = append(lines, "")
	lines = append(lines, style.Subtitle.Render("systems, security, software"))
	lines = append(lines, style.Muted.Render("a personal site served over ssh"))
	lines = append(lines, "")

	for i, item := range m.menuItems() {
		prefix := "  "
		itemStyle := style.MenuItem

		if i == m.cursor {
			prefix = "› "
			itemStyle = style.MenuSelected
		}

		lines = append(lines, itemStyle.Render(prefix+item))
	}

	lines = append(lines, "")
	lines = append(lines, style.VeryMutedStyle.Render("minimal terminal interface"))
	lines = append(lines, style.VeryMutedStyle.Render("keyboard-first, public-facing, and self-hosted"))

	return strings.Join(padLines(lines, bodyHeight), "\n")
}

func (m Model) renderProjects(width int, bodyHeight int) string {
	var lines []string

	lines = append(lines, style.Heading.Render("Projects"))
	lines = append(lines, "")

	for i, project := range content.Projects {
		prefix := "  "
		nameStyle := style.Value

		if i == m.projectCursor {
			prefix = "› "
			nameStyle = style.MenuSelected
		}

		lines = append(lines, nameStyle.Render(prefix+project.Name))
		lines = append(lines, style.Muted.Render("  "+project.Summary))
		lines = append(lines, "")
	}

	lines = append(lines, style.VeryMutedStyle.Render("enter opens project details"))

	return strings.Join(padLines(lines, bodyHeight), "\n")
}

func (m Model) renderScrollable(lines []string, width int, bodyHeight int) string {
	wrapped := wrapLines(lines, width)

	if len(wrapped) <= bodyHeight {
		return strings.Join(padLines(wrapped, bodyHeight), "\n")
	}

	contentHeight := bodyHeight - 1
	if contentHeight < 1 {
		contentHeight = 1
	}

	maxStart := len(wrapped) - contentHeight
	start := clamp(m.scroll, 0, maxStart)
	end := start + contentHeight

	if end > len(wrapped) {
		end = len(wrapped)
	}

	visible := append([]string{}, wrapped[start:end]...)
	visible = padLines(visible, contentHeight)

	position := fmt.Sprintf("showing %d-%d of %d", start+1, end, len(wrapped))
	visible = append(visible, style.VeryMutedStyle.Render(position))

	return strings.Join(visible, "\n")
}

func (m Model) aboutLines() []string {
	return []string{
		style.Heading.Render("About"),
		"",
		style.Value.Render("James Worley"),
		"",
		"Cybersecurity and systems enthusiast interested in",
		"networking, Linux infrastructure, and practical software",
		"projects.",
		"",
		"I like building useful tools that connect software,",
		"systems, and security.",
		"",
		"This SSH site is a small experiment in making a personal",
		"website feel more like a terminal application.",
	}
}

func (m Model) contactLines() []string {
	return []string{
		style.Heading.Render("Contact"),
		"",
		style.Value.Render("Website    ") + style.Hyperlink("w3r5l3y.com", "https://w3r5l3y.com"),
		style.Value.Render("GitHub     ") + style.Hyperlink("github.com/W3r5l3y", "https://github.com/W3r5l3y"),
		style.Value.Render("LinkedIn   ") + style.Hyperlink("add profile URL", "https://linkedin.com/in/add-profile-url"),
		style.Value.Render("Email      ") + style.Hyperlink("you@example.com", "mailto:you@example.com"),
		"",
	}
}

func (m Model) helpLines() []string {
	return []string{
		style.Heading.Render("Help"),
		"",
		style.Value.Render("Navigation"),
		"",
		"↑ / k        move up or scroll up",
		"↓ / j        move down or scroll down",
		"enter        select or open",
		"esc / b      go back",
		"?            open help",
		"q            quit",
		"",
		style.Value.Render("Scrolling"),
		"",
		"pageup       scroll up quickly",
		"pagedown     scroll down quickly",
		"home         jump to top",
		"end          jump to bottom",
		"",
		style.Value.Render("Links"),
		"",
		"Some terminals support clickable links.",
		"If clicking does not work, try Ctrl + click.",
	}
}

func (m Model) sourceLines() []string {
	return []string{
		style.Heading.Render("Source"),
		"",
		"This site is a terminal interface served over SSH.",
		"",
		style.Value.Render("Built with:"),
		"",
		"Go",
		"Wish",
		"Bubble Tea",
		"Lip Gloss",
		"Docker",
		"",
		style.Value.Render("Repository:"),
		"",
		style.Hyperlink("github.com/W3r5l3y/w3r5l3y-ssh", "https://github.com/W3r5l3y/w3r5l3y-ssh"),
		"",
	}
}

func (m Model) projectDetailLines() []string {
	if len(content.Projects) == 0 {
		return []string{"No projects available."}
	}

	index := clamp(m.selectedIndex, 0, len(content.Projects)-1)
	project := content.Projects[index]

	lines := []string{
		style.Heading.Render(project.Name),
		"",
		style.Label.Render("Type    ") + project.Kind,
		style.Label.Render("Stack   ") + project.Stack,
		style.Label.Render("Focus   ") + project.Focus,
		"",
		project.Summary,
		"",
		style.Value.Render("Details"),
		"",
	}

	for _, detail := range project.Details {
		lines = append(lines, "• "+detail)
		lines = append(lines, "")
	}

	return lines
}

func padLines(lines []string, height int) []string {
	out := append([]string{}, lines...)

	for len(out) < height {
		out = append(out, "")
	}

	if len(out) > height {
		return out[:height]
	}

	return out
}

func wrapLines(lines []string, width int) []string {
	if width < 20 {
		width = 20
	}

	var wrapped []string

	for _, line := range lines {
		if line == "" {
			wrapped = append(wrapped, "")
			continue
		}

		if strings.HasPrefix(line, "• ") {
			wrapped = append(wrapped, wrapBullet(line, width)...)
			continue
		}

		text := wordwrap.String(line, width)
		parts := strings.Split(strings.TrimRight(text, "\n"), "\n")

		wrapped = append(wrapped, parts...)
	}

	return wrapped
}

func wrapBullet(line string, width int) []string {
	text := strings.TrimPrefix(line, "• ")

	firstPrefix := "• "
	nextPrefix := "  "

	firstWidth := width - lipgloss.Width(firstPrefix)
	nextWidth := width - lipgloss.Width(nextPrefix)

	if firstWidth < 10 {
		firstWidth = 10
	}

	if nextWidth < 10 {
		nextWidth = 10
	}

	firstWrapped := wordwrap.String(text, firstWidth)
	parts := strings.Split(strings.TrimRight(firstWrapped, "\n"), "\n")

	if len(parts) == 0 {
		return []string{firstPrefix}
	}

	out := []string{firstPrefix + parts[0]}

	if len(parts) > 1 {
		remaining := strings.Join(parts[1:], " ")
		nextWrapped := wordwrap.String(remaining, nextWidth)
		nextParts := strings.Split(strings.TrimRight(nextWrapped, "\n"), "\n")

		for _, part := range nextParts {
			out = append(out, nextPrefix+part)
		}
	}

	return out
}

func clamp(value int, minValue int, maxValue int) int {
	if value < minValue {
		return minValue
	}

	if value > maxValue {
		return maxValue
	}

	return value
}
