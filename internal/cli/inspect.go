package cli

import (
	"fmt"

	"github.com/james-rocha/verso/internal/project"
)

type InspectCommand struct{}

func (InspectCommand) Name() string {
	return "inspect"
}

func (InspectCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing project path")
	}

	p, err := project.Load(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Project: %s\n", p.Metadata.Name)
	fmt.Printf("Version: %s\n\n", p.Metadata.Version)

	printComponents("Skills", p.Components, project.ComponentSkill)
	printComponents("Memory", p.Components, project.ComponentMemory)
	printComponents("Workflows", p.Components, project.ComponentWorkflow)
	printComponents("Templates", p.Components, project.ComponentTemplate)

	return nil
}