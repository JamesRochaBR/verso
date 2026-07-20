package project

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Validate checks if a project directory is a valid Verso project.
// It verifies required directories exist and loads the project for further validation.
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

	project, err := Load(path)
	if err != nil {
		return err
	}

	// Validate all components in the project
	for _, comp := range project.Components {
		if comp.Type == ComponentSkill {
			if err := ValidateSkill(comp); err != nil {
				return fmt.Errorf("invalid skill %s: %w", comp.Name, err)
			}
		}
	}

	return nil
}

// ValidateSkill validates a Skill component according to RFC-0002 specifications.
//
// Rules:
//   - Name is required and must be non-empty
//   - Type must be "skill"
//   - Content must not be empty
//   - Description is recommended (warning only, not error)
func ValidateSkill(component Component) error {
	// 1. Name é obrigatório
	if strings.TrimSpace(component.Name) == "" {
		return fmt.Errorf("skill name is required")
	}

	// 2. Type deve ser "skill"
	if component.Type != ComponentSkill {
		return fmt.Errorf("component %q has type %q, expected %q",
			component.Name, component.Type, ComponentSkill)
	}

	// 3. Content não pode ser vazio
	if strings.TrimSpace(component.Content) == "" {
		return fmt.Errorf("skill %q has empty content", component.Name)
	}

	// 4. Description é recomendada (warning-level check)
	//    Não é erro, mas podemos registrar um aviso se necessário
	if component.Description == "" && component.Metadata != nil {
		// Warning: description is recommended but not required
		_ = "description missing — this is a warning, not an error"
	}

	return nil
}

// ValidateLifecycle checks if the component's lifecycle state is valid.
func ValidateLifecycle(component Component) error {
	if component.Status == "" {
		// Default to created if no status set
		component.Status = StateCreated
		return nil
	}

	if !component.Status.IsValid() {
		return fmt.Errorf("invalid lifecycle state %q for component %q",
			component.Status, component.Name)
	}

	return nil
}

// ValidateComponent validates any Verso component based on its type.
func ValidateComponent(component Component) error {
	switch component.Type {
	case ComponentSkill:
		return ValidateSkill(component)
	default:
		// Basic validation for all types
		if strings.TrimSpace(component.Name) == "" {
			return fmt.Errorf("component %q: name is required", component.Path)
		}
		return nil
	}
}

// ValidateComponents validates a slice of components, returning the first error found.
func ValidateComponents(components []Component) error {
	for _, comp := range components {
		if err := ValidateComponent(comp); err != nil {
			return err
		}
	}
	return nil
}