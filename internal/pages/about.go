package pages

import "github.com/W3r5l3y/w3r5l3y-ssh/internal/style"

func About() string {
	return style.Heading.Render("About") + `

James Worley

Computer Science enthusiast interested in cybersecurity,
networking, Linux systems, and practical infrastructure.

I like building useful tools that connect software,
systems, and security.

` + style.Help.Render("esc back    q quit")
}
