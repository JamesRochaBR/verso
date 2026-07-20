package project

import "encoding/json"

// ComponentType represents the type of a Verso component.
type ComponentType string

const (
	ComponentSkill    ComponentType = "skill"
	ComponentMemory   ComponentType = "memory"
	ComponentWorkflow ComponentType = "workflow"
	ComponentTemplate ComponentType = "template"
	ComponentRouter   ComponentType = "router"
	ComponentAgent    ComponentType = "agent"
)

// LifecycleState represents the lifecycle state of a component.
type LifecycleState string

const (
	StateCreated    LifecycleState = "created"
	StateReviewed   LifecycleState = "reviewed"
	StateApproved   LifecycleState = "approved"
	StateDeprecated LifecycleState = "deprecated"
)

// IsValid checks if the lifecycle state is a recognized value.
func (s LifecycleState) IsValid() bool {
	switch s {
	case StateCreated, StateReviewed, StateApproved, StateDeprecated:
		return true
	default:
		return false
	}
}

// Component represents a Verso project component (skill, memory, workflow, template, router, agent).
// It includes both core fields and optional metadata parsed from YAML frontmatter.
type Component struct {
	Name        string                `json:"name"`
	Title       string                `json:"title"`
	Type        ComponentType         `json:"type"`
	Path        string                `json:"path"`
	Content     string                `json:"content,omitempty"`
	Metadata    *FrontmatterMetadata  `json:"metadata,omitempty"`
	Version     string                `json:"version,omitempty"`
	Author      string                `json:"author,omitempty"`
	Tags        []string              `json:"tags,omitempty"`
	Description string                `json:"description,omitempty"`
	Status      LifecycleState        `json:"status,omitempty"`
}

// MarshalJSON provides custom JSON marshaling to include derived fields.
func (c Component) MarshalJSON() ([]byte, error) {
	type Alias Component
	return json.Marshal(&struct {
		*Alias
	}{
		Alias: (*Alias)(&c),
	})
}

// HasTags returns true if any of the given tags match this component's tags.
func (c Component) HasTags(tags []string) bool {
	if len(c.Tags) == 0 || len(tags) == 0 {
		return false
	}

	tagSet := make(map[string]struct{}, len(c.Tags))
	for _, t := range c.Tags {
		tagSet[t] = struct{}{}
	}

	for _, tag := range tags {
		if _, ok := tagSet[tag]; ok {
			return true
		}
	}

	return false
}

// ContainsKeyword returns true if the component's name, title, description, or tags contain the keyword.
func (c Component) ContainsKeyword(keyword string) bool {
	keyword = lowercaseString(keyword)
	if lowercaseString(c.Name) == keyword {
		return true
	}
	if lowercaseString(c.Title) != "" && containsString(lowercaseString(c.Title), keyword) {
		return true
	}
	if lowercaseString(c.Description) != "" && containsString(lowercaseString(c.Description), keyword) {
		return true
	}
	if len(c.Tags) > 0 && containsTag(c.Tags, keyword) {
		return true
	}
	return false
}

// IsDeprecated returns true if the component has been deprecated.
func (c Component) IsDeprecated() bool {
	return c.Status == StateDeprecated
}
