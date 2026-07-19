package cli

import (
	"fmt"

	"github.com/james-rocha/verso/internal/project"
	"github.com/james-rocha/verso/internal/render"
)

type PromptCommand struct{}

func (PromptCommand) Name() string {
	return "prompt"
}

func (PromptCommand) Run(args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("missing project path")
	}

	p, err := project.Load(args[0])
	if err != nil {
		return err
	}

	filter := ParseFilter(args[1:])

	p = project.ApplyFilter(p, filter)

	fmt.Print(render.Prompt(p))

	return nil
}
