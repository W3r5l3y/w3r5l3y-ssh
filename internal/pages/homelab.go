package pages

import "github.com/W3r5l3y/w3r5l3y-ssh/internal/style"

func Homelab() string {
	return style.Heading.Render("Homelab") + `

A small self-hosted infrastructure setup used for XXX

Summary:
  Storage       ###
  Services      ###
  Focus         ###


` + style.Help.Render("esc back    q quit")
}
