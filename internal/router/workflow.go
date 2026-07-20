package router

import (
	"github.com/james-rocha/verso/internal/project"
)

// workflowStrategy implements workflow-based routing.
// When a specific workflow is requested, it loads all referenced components automatically.
type workflowStrategy struct{}

// Route implements the Strategy interface for workflow-based routing.
func (s *workflowStrategy) Route(p *project.Project, opts RoutingOptions) (*project.Project, error) {
	if p == nil {
		return nil, ErrNilProject
	}

	workflowName := opts.WorkflowName
	if workflowName == "" {
		// No workflow specified, return all components
		return p, nil
	}

	// Extract referenced components from the workflow
	referenced, err := filterComponentsByWorkflow(p, workflowName)
	if err != nil {
		return nil, err
	}

	// Build filtered project with only referenced components + the workflow itself
	workflowComp := findComponentByName(p, workflowName)
	result := &project.Project{
		Metadata:   p.Metadata,
		Components: make([]project.Component, 0, len(referenced)+1),
	}

	// Add the workflow component itself
	if workflowComp != nil {
		result.Components = append(result.Components, *workflowComp)
	}

	// Add referenced components
	result.Components = append(result.Components, referenced...)

	return result, nil
}