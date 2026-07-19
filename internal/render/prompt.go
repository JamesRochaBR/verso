package render

import (
	"strings"

	"github.com/james-rocha/verso/internal/project"
)

func Prompt(p *project.Project) string {
	var b strings.Builder

	b.WriteString("# Project\n\n")

	b.WriteString("Name: ")
	b.WriteString(p.Metadata.Name)
	b.WriteString("\n")

	b.WriteString("Version: ")
	b.WriteString(p.Metadata.Version)
	b.WriteString("\n\n")

	writeSection(&b, "Skills", p.Components, project.ComponentSkill)
	writeSection(&b, "Memory", p.Components, project.ComponentMemory)
	writeSection(&b, "Workflows", p.Components, project.ComponentWorkflow)
	writeSection(&b, "Templates", p.Components, project.ComponentTemplate)

	return b.String()
}

func writeSection(
	b *strings.Builder,
	title string,
	components []project.Component,
	componentType project.ComponentType,
) {
	first := true

	for _, c := range components {
		if c.Type != componentType {
			continue
		}

		if first {
			b.WriteString("# ")
			b.WriteString(title)
			b.WriteString("\n\n")
			first = false
		}

		b.WriteString("## ")
		b.WriteString(c.Title)
		b.WriteString("\n\n")
		b.WriteString(c.Content)
		b.WriteString("\n\n")
	}
}
