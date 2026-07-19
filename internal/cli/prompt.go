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

	var filter project.Filter

	for i := 1; i < len(args); i++ {

		if args[i] == "--name" && i+1 < len(args) {
			filter.Names = splitCSV(args[i+1])
			i++
		}
	}

	p = project.ApplyFilter(p, filter)

	fmt.Print(render.Prompt(p))

	return nil
}
