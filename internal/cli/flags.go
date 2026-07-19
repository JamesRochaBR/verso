package cli

import (
	"strings"

	"github.com/james-rocha/verso/internal/project"
)

type PromptOptions struct {
	Filter project.Filter
	Output string
}

func ParsePromptOptions(args []string) PromptOptions {
	var opts PromptOptions

	for i := 0; i < len(args); i++ {

		switch args[i] {

		case "--name":
			if i+1 < len(args) {
				opts.Filter.Names = splitCSV(args[i+1])
				i++
			}

		case "--exclude":
			if i+1 < len(args) {
				opts.Filter.Exclude = parseComponentTypes(args[i+1])
				i++
			}

		case "--output":
			if i+1 < len(args) {
				opts.Output = args[i+1]
				i++
			}

		}
	}

	return opts
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
