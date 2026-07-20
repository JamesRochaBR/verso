package router

import (
	"strings"

	"github.com/james-rocha/verso/internal/project"
)

// Strategy defines the interface for routing decisions.
// Each strategy implements different logic to select components from a project.
type Strategy interface {
	Route(p *project.Project, opts RoutingOptions) (*project.Project, error)
}

// findComponentByType returns all components of a given type.
func findComponentsByType(p *project.Project, typ project.ComponentType) []project.Component {
	var result []project.Component
	for _, c := range p.Components {
		if c.Type == typ {
			result = append(result, c)
		}
	}
	return result
}

// findComponentByName returns a component by name, or nil if not found.
func findComponentByName(p *project.Project, name string) *project.Component {
	for _, c := range p.Components {
		if strings.EqualFold(c.Name, name) {
			return &c
		}
	}
	return nil
}

// containsComponent checks if a component list contains a component with the given name.
func containsComponent(components []project.Component, name string) bool {
	for _, c := range components {
		if strings.EqualFold(c.Name, name) {
			return true
		}
	}
	return false
}

// extractKeywords analyzes component content and metadata to build a keyword map.
// Each component is indexed by its name, title, type, and content words.
func extractKeywords(components []project.Component) map[string][]string {
	keywords := make(map[string][]string)

	for _, c := range components {
		var keys []string

		// Index by name and title
		keys = append(keys, strings.ToLower(c.Name))
		if c.Title != "" {
			keys = append(keys, strings.ToLower(c.Title))
		}

		// Index by type
		keys = append(keys, string(c.Type))

		// Index content words (first 500 chars to avoid noise)
		content := c.Content
		if len(content) > 500 {
			content = content[:500]
		}
		for _, word := range strings.Fields(strings.ToLower(content)) {
			cleaned := strings.TrimFunc(word, func(r rune) bool {
				return !(r >= 'a' && r <= 'z') && !(r >= 'A' && r <= 'Z') && !(r >= '0' && r <= '9')
			})
			if len(cleaned) > 2 {
				keys = append(keys, cleaned)
			}
		}

		keywords[c.Name] = keys
	}

	return keywords
}

// filterComponentsByKeywords returns components whose keywords match any of the search terms.
func filterComponentsByKeywords(components []project.Component, keywords []string) []project.Component {
	if len(keywords) == 0 {
		return components
	}

	index := extractKeywords(components)
	scoreMap := make(map[string]int)

	// Score each component based on keyword matches
	for _, kw := range keywords {
		searchTerm := strings.ToLower(kw)
		for compName, compKeys := range index {
			for _, key := range compKeys {
				// Only match if the search term is at least 4 characters
				// to avoid false positives from short substrings
				if len(searchTerm) < 4 {
					continue
				}

				// Exact match gets highest score
				if strings.EqualFold(key, searchTerm) {
					scoreMap[compName] += 3
					continue
				}

				// Key is contained in search term (e.g., "architect" in "architecture")
				if len(key) >= 4 && strings.Contains(searchTerm, key) {
					scoreMap[compName] += 1
				}
			}
		}
	}

	// Collect components with score > 0
	var result []project.Component
	for _, c := range components {
		if scoreMap[c.Name] > 0 {
			result = append(result, c)
		}
	}

	return result
}

// filterComponentsByWorkflow extracts referenced components from workflow steps.
func filterComponentsByWorkflow(p *project.Project, workflowName string) ([]project.Component, error) {
	workflow := findComponentByName(p, workflowName)
	if workflow == nil {
		return nil, &ErrWorkflowNotFound{Name: workflowName}
	}

	var referenced []project.Component
	seen := make(map[string]bool)

	// Parse steps from frontmatter or content
	steps := parseSteps(workflow.Content)

	for _, step := range steps {
		// Add referenced skill
		if step.Skill != "" && !seen[step.Skill] {
			seen[step.Skill] = true
			if comp := findComponentByName(p, step.Skill); comp != nil {
				referenced = append(referenced, *comp)
			}
		}

		// Add referenced memory
		if step.Memory != "" && !seen[step.Memory] {
			seen[step.Memory] = true
			if comp := findComponentByName(p, step.Memory); comp != nil {
				referenced = append(referenced, *comp)
			}
		}

		// Add referenced context skill from content
		for _, ctx := range step.ContextSkills {
			if !seen[ctx] {
				seen[ctx] = true
				if comp := findComponentByName(p, ctx); comp != nil {
					referenced = append(referenced, *comp)
				}
			}
		}
	}

	return referenced, nil
}

// Step represents a single step in a workflow.
type Step struct {
	Name          string
	Skill         string
	Memory        string
	Context       string
	ContextSkills []string
}

// parseSteps extracts steps from workflow markdown content.
// It looks for YAML frontmatter "steps" field or numbered lists.
func parseSteps(content string) []Step {
	var steps []Step

	// Try to parse YAML-like frontmatter steps
	inFrontmatter := false
	inSteps := false
	currentStep := Step{}

	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)

		// Detect frontmatter start
		if trimmed == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				inFrontmatter = false
				continue
			}
		}

		if inFrontmatter && strings.HasPrefix(trimmed, "steps:") {
			inSteps = true
			continue
		}

		if inSteps {
			// Check for list item (- name: ...)
			if strings.HasPrefix(trimmed, "- name:") {
				if currentStep.Name != "" {
					steps = append(steps, currentStep)
				}
				currentStep = Step{
					Name: strings.TrimSpace(strings.TrimPrefix(trimmed, "- name:")),
				}
				continue
			}

			if strings.HasPrefix(trimmed, "skill:") {
				currentStep.Skill = strings.TrimSpace(strings.TrimPrefix(trimmed, "skill:"))
				continue
			}

			if strings.HasPrefix(trimmed, "memory:") {
				currentStep.Memory = strings.TrimSpace(strings.TrimPrefix(trimmed, "memory:"))
				continue
			}

			if strings.HasPrefix(trimmed, "context:") {
				currentStep.Context = strings.TrimSpace(strings.TrimPrefix(trimmed, "context:"))
				continue
			}

			// If we hit a non-list line that's not empty, stop steps parsing
			if trimmed != "" && !strings.HasPrefix(trimmed, "-") &&
				!strings.Contains(trimmed, ":") || (strings.Contains(trimmed, ":") &&
				!strings.HasPrefix(trimmed, "skill:") && !strings.HasPrefix(trimmed, "memory:") &&
				!strings.HasPrefix(trimmed, "context:") && !strings.HasPrefix(trimmed, "name:")) {
				if currentStep.Name != "" {
					steps = append(steps, currentStep)
					currentStep = Step{}
				}
				inSteps = false
			}
		}
	}

	// Don't forget the last step
	if currentStep.Name != "" {
		steps = append(steps, currentStep)
	}

	// If no steps found from frontmatter, try to parse numbered list from content
	if len(steps) == 0 {
		steps = parseNumberedListSteps(content)
	}

	return steps
}

// parseNumberedListSteps extracts steps from a numbered list in markdown.
func parseNumberedListSteps(content string) []Step {
	var steps []Step
	currentStep := Step{}

	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)

		// Match "1. Skill name" or "1. Load skill_name"
		if len(steps) > 0 || (len(currentStep.Name) == 0 && len(steps) == 0) {
			// Try to extract step info from numbered list
		}

		// Look for skill/memory references in the content
		lower := strings.ToLower(trimmed)
		if strings.Contains(lower, "skill:") || strings.Contains(lower, "load") && strings.Contains(lower, "skill") {
			if currentStep.Name == "" {
				currentStep.Name = extractStepName(trimmed)
			}
			if skill := extractComponentRef(trimmed, "skill"); skill != "" {
				currentStep.Skill = skill
			}
		}

		if strings.Contains(lower, "memory:") || strings.Contains(lower, "load") && strings.Contains(lower, "memory") {
			if memory := extractComponentRef(trimmed, "memory"); memory != "" {
				currentStep.Memory = memory
			}
		}
	}

	return steps
}

// extractStepName extracts a human-readable step name from a line.
func extractStepName(line string) string {
	// Remove numbering
	for i := 0; i < 100; i++ {
		if len(line) > 1 && line[0] == byte('0'+i) && line[1] == '.' {
			line = strings.TrimSpace(line[2:])
			break
		}
	}
	return line
}

// extractComponentRef extracts a component reference from a line like "skill: architect".
func extractComponentRef(line string, refType string) string {
	lower := strings.ToLower(line)
	idx := strings.Index(lower, refType+":")
	if idx == -1 {
		return ""
	}
	val := strings.TrimSpace(line[idx+len(refType)+1:])
	// Remove quotes if present
	val = strings.Trim(val, "\"'")
	return val
}