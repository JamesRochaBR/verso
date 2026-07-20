package router

import (
	"github.com/james-rocha/verso/internal/project"
)

// Router is the main interface for component routing.
// The Router decides which components (Skills, Memory, Workflows) should be
// loaded based on context and user intent.
type Router interface {
	// Route analyzes the project and options to return a filtered project
	// containing only the components that should be loaded.
	Route(p *project.Project, opts RoutingOptions) (*project.Project, error)
}

// RoutingOptions contains all parameters that influence routing decisions.
type RoutingOptions struct {
	// Filter contains explicit name and type filters from CLI flags.
	Filter project.Filter

	// Strategy selects the routing strategy to use.
	// Valid values: "keyword", "workflow", "default"
	Strategy string

	// Keywords are used for keyword-based routing matching.
	Keywords []string

	// WorkflowName is the name of a workflow to route through.
	WorkflowName string
}

// New creates a Router with default configuration.
// The default strategy is "default" which loads all components.
func New() Router {
	return &router{
		strategies: map[string]Strategy{
			"keyword":   &keywordStrategy{},
			"workflow":  &workflowStrategy{},
			"default":   &defaultStrategy{},
		},
	}
}

// router is the default implementation of the Router interface.
type router struct {
	strategies map[string]Strategy
}

// Route implements the Router interface.
func (r *router) Route(p *project.Project, opts RoutingOptions) (*project.Project, error) {
	if p == nil {
		return nil, ErrNilProject
	}

	// Apply explicit filters first if provided
	if len(opts.Filter.Names) > 0 || len(opts.Filter.Exclude) > 0 {
		return project.ApplyFilter(p, opts.Filter), nil
	}

	// Select strategy based on options
	var strategy Strategy
	switch opts.Strategy {
	case "keyword":
		strategy = r.strategies["keyword"]
	case "workflow":
		strategy = r.strategies["workflow"]
	default:
		// Auto-detect strategy if not specified
		if len(opts.Keywords) > 0 {
			strategy = r.strategies["keyword"]
		} else if opts.WorkflowName != "" {
			strategy = r.strategies["workflow"]
		} else {
			strategy = r.strategies["default"]
		}
	}

	return strategy.Route(p, opts)
}