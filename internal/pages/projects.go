package pages

import "github.com/W3r5l3y/w3r5l3y-ssh/internal/style"

func Projects() string {
	return style.Heading.Render("Projects") + `

AI Reference Verifier
  Detects fabricated academic references using GROBID,
  Crossref, OpenAlex, Scopus, SQLite, and confidence scoring.

Homelab
  Self-hosted infrastructure for XXX

exEco
  Django sustainability web app with QR scanning and API
  integration.

` + style.Help.Render("esc back    q quit")
}
