package pages

import "github.com/W3r5l3y/w3r5l3y-ssh/internal/style"

func Contact() string {
	return style.Heading.Render("Contact") + `

Website    w3r5l3y.com
GitHub     github.com/W3r5l3y
LinkedIn   ###
Email      placeholder@example.com

` + style.Help.Render("esc back    q quit")
}
