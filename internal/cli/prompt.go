package cli

import (
	"fmt"
	"os"

	"github.com/james-rocha/verso/internal/project"
	"github.com/james-rocha/verso/internal/render"
	"github.com/james-rocha/verso/internal/router"
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

	// Use Router for intelligent component selection
	r := router.New()

	routingOpts := router.RoutingOptions{
		Filter:   opts.Filter,
		Strategy: opts.Strategy,
		Keywords: opts.Keywords,
	}

	if opts.Workflow != "" {
		routingOpts.Strategy = "workflow"
		routingOpts.WorkflowName = opts.Workflow
	}

	p, err = r.Route(p, routingOpts)
	if err != nil {
		return err
	}

	out, err := render.Prompt(p)
	if err != nil {
		return err
	}

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
  --name <list>          Include only components by name
  --exclude <list>       Exclude component types
  --keywords, -k <list>  Filter components by keywords (uses keyword routing)
  --workflow, -w <name>  Route through a specific workflow (uses workflow routing)
  --strategy, -s <type>  Routing strategy: "keyword", "workflow", or "default"
  --output <file>        Write prompt to file
  --help                 Show this help`)
}
