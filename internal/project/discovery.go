package project

import (
	"os"
	"path/filepath"
	"strings"
)

// Discover scans the project directory for components (skills, memory, workflows, templates).
// It parses YAML frontmatter to extract metadata and populates Component fields accordingly.
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

			filePath := filepath.Join(dir, entry.Name())
			contentBytes, err := os.ReadFile(filePath)
			if err != nil {
				continue
			}

			content := string(contentBytes)
			name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))

			// Parse frontmatter metadata if present
			metadata, _ := ParseFrontmatter(content)

			component := Component{
				Name:    name,
				Title:   ExtractTitle(content),
				Type:    componentType,
				Path:    filePath,
				Content: content,
			}

			// Populate metadata fields if available
			if metadata != nil {
				if metadata.Name != "" {
					component.Name = metadata.Name
				}
				component.Type = ComponentType(metadata.Type)
				if component.Type == "" {
					component.Type = componentType
				}
				component.Version = metadata.Version
				component.Author = metadata.Author
				component.Tags = metadata.Tags
				component.Description = metadata.Description
				component.Metadata = metadata

				// Map lifecycle status if present
				if metadata.Status != "" {
					component.Status = LifecycleState(metadata.Status)
				}
			}

			components = append(components, component)
		}
	}

	return components, nil
}
