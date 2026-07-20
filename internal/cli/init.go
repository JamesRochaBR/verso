package cli

import (
	"fmt"

	"github.com/james-rocha/verso/internal/project"
)

type InitCommand struct{}

func (InitCommand) Name() string {
	return "init"
}

func (InitCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing project name")
	}

	if err := project.Create(args[0]); err != nil {
		return err
	}

	fmt.Printf("Project '%s' created successfully.\n", args[0])

	return nil
}
