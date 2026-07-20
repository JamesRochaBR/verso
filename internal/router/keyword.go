package router

import (
	"github.com/james-rocha/verso/internal/project"
)

// keywordStrategy implements keyword-based routing.
// It analyzes user input for keywords and matches them against component metadata
// (names, titles, types, and content words).
type keywordStrategy struct{}

// Route implements the Strategy interface for keyword-based routing.
func (s *keywordStrategy) Route(p *project.Project, opts RoutingOptions) (*project.Project, error) {
	if p == nil {
		return nil, ErrNilProject
	}

	keywords := opts.Keywords
	if len(keywords) == 0 {
		// If no keywords provided, return all components (fallback to default behavior)
		return p, nil
	}

	// Filter all components by keywords
	matched := filterComponentsByKeywords(p.Components, keywords)

	// Build filtered project
	result := &project.Project{
		Metadata:   p.Metadata,
		Components: matched,
	}

	return result, nil
}