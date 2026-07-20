package router

import (
	"github.com/james-rocha/verso/internal/project"
)

// defaultStrategy implements default routing.
// When no filters are specified and no workflow is requested, it loads all components by default.
type defaultStrategy struct{}

// Route implements the Strategy interface for default routing.
func (s *defaultStrategy) Route(p *project.Project, opts RoutingOptions) (*project.Project, error) {
	if p == nil {
		return nil, ErrNilProject
	}

	// Return all components as-is
	// The project was already loaded with all discovered components
	return p, nil
}