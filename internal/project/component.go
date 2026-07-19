package project

type ComponentType string

const (
	ComponentSkill    ComponentType = "skill"
	ComponentMemory   ComponentType = "memory"
	ComponentWorkflow ComponentType = "workflow"
	ComponentTemplate ComponentType = "template"
)

type Component struct {
	Name    string
	Type    ComponentType
	Path    string
	Content string
}
