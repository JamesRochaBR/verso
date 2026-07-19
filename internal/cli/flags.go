package cli

import "strings"

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
