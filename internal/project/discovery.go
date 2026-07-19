package project

import (
	"os"
	"path/filepath"
	"strings"
)

func Discover(path string) ([]Component, error) {
	var components []Component

	types := map[string]ComponentType{
		"skills":    ComponentSkill,
		"memory":    ComponentMemory,
		"workflows": ComponentWorkflow,
		"templates": ComponentTemplate,
	}

	for folder, componentType := range types {
		dir := filepath.Join(path, folder)

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

			components = append(components, Component{
				Name: name,
				Type: componentType,
				Path: filepath.Join(folder, entry.Name()),
			})
		}
	}

	return components, nil
}