package cli

import (
	"fmt"

	"github.com/james-rocha/verso/internal/project"
)

type ValidateCommand struct{}

func (ValidateCommand) Name() string {
	return "validate"
}

func (ValidateCommand) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing project path")
	}

	if err := project.Validate(args[0]); err != nil {
		return err
	}

	fmt.Println("✓ Project is valid")

	return nil
}
