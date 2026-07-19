package project

type ComponentType string

const (
	ComponentSkill    ComponentType = "skill"
	ComponentMemory   ComponentType = "memory"
	ComponentWorkflow ComponentType = "workflow"
	ComponentTemplate ComponentType = "template"
)

type Component struct {
	Name    string        `json:"name"`
	Title   string        `json:"title"`
	Type    ComponentType `json:"type"`
	Path    string        `json:"path"`
	Content string        `json:"content,omitempty"`
}
