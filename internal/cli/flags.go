package cli

import (
	"fmt"
	"github.com/james-rocha/verso/internal/project"
	"strings"
)

type PromptOptions struct {
	Filter   project.Filter
	Output   string
	Keywords []string
	Strategy string
	Workflow string
}

func ParsePromptOptions(args []string) (PromptOptions, error) {
	var opts PromptOptions

	for i := 0; i < len(args); i++ {

		switch args[i] {
		case "--name":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("missing value for --name")
			}

			opts.Filter.Names = splitCSV(args[i+1])
			i++

		case "--exclude":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("missing value for --exclude")
			}

			opts.Filter.Exclude = parseComponentTypes(args[i+1])
			i++

		case "--output":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("missing value for --output")
			}

			opts.Output = args[i+1]
			i++

		case "--keywords", "-k":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("missing value for --keywords")
			}

			opts.Keywords = splitCSV(args[i+1])
			opts.Strategy = "keyword"
			i++

		case "--workflow", "-w":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("missing value for --workflow")
			}

			opts.Workflow = args[i+1]
			opts.Strategy = "workflow"
			i++

		case "--strategy", "-s":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("missing value for --strategy")
			}

			opts.Strategy = args[i+1]
			i++

		default:
			return opts, fmt.Errorf("unknown flag: %s", args[i])

		}
	}

	return opts, nil
}

func splitCSV(value string) []string {
	if value == "" {
		return nil
	}

	parts := strings.Split(value, ",")

	result := make([]string, 0, len(parts))

	for _, p := range parts {

		p = strings.TrimSpace(p)

		if p == "" {
			continue
		}

		result = append(result, p)
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func parseComponentTypes(value string) []project.ComponentType {
	names := splitCSV(value)

	result := make([]project.ComponentType, 0, len(names))

	for _, name := range names {
		switch name {

		case "skill":
			result = append(result, project.ComponentSkill)

		case "memory":
			result = append(result, project.ComponentMemory)

		case "workflow":
			result = append(result, project.ComponentWorkflow)

		case "template":
			result = append(result, project.ComponentTemplate)
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
