package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func Validate(path string) error {
	required := []string{
		"verso.toml",
		"skills",
		"memory",
		"workflows",
		"templates",
	}

	for _, item := range required {
		if _, err := os.Stat(filepath.Join(path, item)); err != nil {
			return fmt.Errorf("missing %s", item)
		}
	}

	return nil
}