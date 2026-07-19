package cli

import (
	"fmt"

	"github.com/james-rocha/verso/internal/project"
)

func printComponents(title string, components []project.Component, componentType project.ComponentType) {
	count := 0

	for _, c := range components {
		if c.Type == componentType {
			count++
		}
	}

	fmt.Printf("%s (%d)\n", title, count)

	for _, c := range components {
		if c.Type == componentType {
			fmt.Printf("- %s (%d bytes)\n", c.Name, len(c.Content))
		}
	}

	fmt.Println()
}