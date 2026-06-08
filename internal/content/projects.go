package content

type Project struct {
	Name    string
	Kind    string
	Stack   string
	Focus   string
	Summary string
	Details []string
}

var Projects = []Project{
	{
		Name:    "AI Reference Verifier",
		Kind:    "Research tool",
		Stack:   "Python, SQLite, GROBID, Crossref, OpenAlex, Scopus",
		Focus:   "Detecting fabricated academic references",
		Summary: "A tool for checking whether academic references are real, valid, and correctly represented.",
		Details: []string{
			"Extracts references from academic papers and validates them against scholarly metadata sources.",
			"Compares titles, authors, years, DOI values, and source URLs to produce a confidence score.",
			"Designed around a practical problem: AI-generated writing can include references that look convincing but do not exist.",
		},
	},
	{
		Name:    "Self-hosted Infrastructure",
		Kind:    "Personal systems project",
		Stack:   "Linux, TrueNAS SCALE, ZFS, Docker, DNS, VPN, monitoring",
		Focus:   "Storage, networking, service management, and reliability",
		Summary: "A personal infrastructure environment for learning Linux administration, networking, and self-hosting.",
		Details: []string{
			"Built around reliable storage, service isolation, local networking, DNS filtering, and system monitoring.",
			"Used as a practical environment for learning how services are deployed, maintained, secured, and observed over time.",
			"Focuses on infrastructure fundamentals rather than exposing specific private services or internal network details.",
		},
	},
	{
		Name:    "exEco",
		Kind:    "Web application",
		Stack:   "Django, Python, APIs, QR scanning",
		Focus:   "Sustainability-focused software",
		Summary: "A Django web app exploring sustainability features through QR scanning and API integration.",
		Details: []string{
			"Built as a practical web application with user-facing features, backend logic, and external API integration.",
			"Explored how sustainability-related data can be connected to simple interactions such as scanning and lookup flows.",
			"Helped develop experience with Django, structured backend development, and integrating external services.",
		},
	},
}
