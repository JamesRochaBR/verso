package cli

import (
	"fmt"
	"github.com/james-rocha/verso/internal/project"
	"github.com/james-rocha/verso/internal/render"
	"os"
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

	opts, err := ParsePromptOptions(args[1:])
	if err != nil {
		return err
	}

	p = project.ApplyFilter(p, opts.Filter)

	out := render.Prompt(p)

	if opts.Output != "" {
		return os.WriteFile(opts.Output, []byte(out), 0644)
	}

	fmt.Print(out)
	return nil
}
