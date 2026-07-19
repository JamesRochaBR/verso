package cli

import (
	"strings"

	"github.com/james-rocha/verso/internal/project"
)

func ParseFilter(args []string) project.Filter {
	var filter project.Filter

	for i := 0; i < len(args); i++ {

		switch args[i] {

		case "--name":
			if i+1 < len(args) {
				filter.Names = splitCSV(args[i+1])
				i++
			}

		}
	}

	return filter
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
