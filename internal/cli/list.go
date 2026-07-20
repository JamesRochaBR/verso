package cli

import (
	"fmt"
	"os"

	"github.com/james-rocha/verso/internal/project"
)

type ListCommand struct{}

func (ListCommand) Name() string {
	return "list"
}

func (ListCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing project path")
	}

	p, err := project.Load(args[0])
	if err != nil {
		return err
	}

	if len(p.Components) == 0 {
		fmt.Fprintf(os.Stderr, "No components found in %s\n", args[0])
		return nil
	}

	printSection("Skills", p.Components, project.ComponentSkill)
	printSection("Memory", p.Components, project.ComponentMemory)
	printSection("Workflows", p.Components, project.ComponentWorkflow)
	printSection("Templates", p.Components, project.ComponentTemplate)

	return nil
}

func printSection(title string, components []project.Component, ctype project.ComponentType) {
	count := 0

	for _, c := range components {
		if c.Type == ctype {
			count++
		}
	}

	if count == 0 {
		return
	}

	fmt.Printf("%s (%d)\n", title, count)

	for _, c := range components {
		if c.Type != ctype {
			continue
		}

		title := c.Title
		if title == "" {
			title = "(no title)"
		}

		fmt.Printf("  - %s | %s\n", c.Name, title)
	}

	fmt.Println()
}
