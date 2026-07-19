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
	if args[0] == "--help" || args[0] == "-h" {
		printPromptHelp()
		return nil
	}

	p, err := project.Load(args[0])
	if err != nil {
		return err
	}

	for _, arg := range args[1:] {
		if arg == "--help" || arg == "-h" {
			printPromptHelp()
			return nil
		}
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

func printPromptHelp() {
	fmt.Println(`Usage:
  verso prompt <project> [options]

Options:
  --name <list>      Include only components by name
  --exclude <list>   Exclude component types
  --output <file>    Write prompt to file
  --help             Show this help`)
}
