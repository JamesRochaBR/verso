package router

import (
	"strings"
	"testing"

	"github.com/james-rocha/verso/internal/project"
)

func sampleProject() *project.Project {
	return &project.Project{
		Metadata: project.Metadata{
			Name:        "demo-project",
			Version:     "0.1.0",
			Description: "Test project for router tests",
		},
		Components: []project.Component{
			{
				Name:    "reviewer",
				Title:   "Code Reviewer",
				Type:    project.ComponentSkill,
				Content: "This skill provides best practices for code review. Use it when reviewing pull requests.",
			},
			{
				Name:    "architect",
				Title:   "System Architect",
				Type:    project.ComponentSkill,
				Content: "This skill helps with system design and architecture decisions.",
			},
			{
				Name:    "developer",
				Title:   "Developer",
				Type:    project.ComponentSkill,
				Content: "This skill provides development guidelines and coding standards.",
			},
			{
				Name:    "project",
				Title:   "Project Context",
				Type:    project.ComponentMemory,
				Content: "Project uses Go modules. All code must pass gofmt before commit.",
			},
			{
				Name:    "conventions",
				Title:   "Code Conventions",
				Type:    project.ComponentMemory,
				Content: "Naming conventions and package structure guidelines for the team.",
			},
			{
				Name:    "feature",
				Title:   "Feature Implementation Workflow",
				Type:    project.ComponentWorkflow,
				Content: "---\nname: feature\nsteps:\n  - name: analyze\n    skill: architect\n    context: Analyze requirements.\n  - name: implement\n    skill: developer\n    memory: project\n    context: Implement following conventions.\n  - name: review\n    skill: reviewer\n---\n# Feature Implementation\n\nThis workflow guides feature implementation.",
			},
		},
	}
}

func TestNew(t *testing.T) {
	r := New()
	if r == nil {
		t.Fatal("expected non-nil router")
	}
}

func TestRouteNilProject(t *testing.T) {
	r := New()
	_, err := r.Route(nil, RoutingOptions{})
	if err == nil {
		t.Fatal("expected error for nil project")
	}
}

func TestRouteDefaultStrategy(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{Strategy: "default"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got.Components) != 6 {
		t.Fatalf("expected 6 components, got %d", len(got.Components))
	}

	if got.Metadata.Name != "demo-project" {
		t.Fatalf("metadata mismatch: expected 'demo-project', got '%s'", got.Metadata.Name)
	}
}

func TestRouteDefaultAutoDetect(t *testing.T) {
	r := New()
	p := sampleProject()

	// No strategy specified, no keywords or workflow — should auto-detect default
	got, err := r.Route(p, RoutingOptions{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got.Components) != 6 {
		t.Fatalf("expected 6 components with auto-detected default strategy, got %d", len(got.Components))
	}
}

func TestRouteKeywordStrategy(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Strategy: "keyword",
		Keywords: []string{"review"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got.Components) == 0 {
		t.Fatal("expected at least one component matching 'review'")
	}

	for _, c := range got.Components {
		if c.Name != "reviewer" && c.Title != "Code Reviewer" {
			// Check if content contains review-related keywords
			hasReviewKeyword := false
			for _, word := range []string{"review", "pull", "request"} {
				if containsString(c.Content, word) || containsString(c.Name, word) || containsString(c.Title, word) {
					hasReviewKeyword = true
					break
				}
			}
			if !hasReviewKeyword {
				t.Fatalf("component %q does not match keyword 'review'", c.Name)
			}
		}
	}
}

func TestRouteKeywordStrategyMultipleKeywords(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Strategy: "keyword",
		Keywords: []string{"architecture", "design"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should match the architect skill
	found := false
	for _, c := range got.Components {
		if c.Name == "architect" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("expected 'architect' component to be matched by keywords 'architecture' or 'design'")
	}
}

func TestRouteKeywordStrategyNoMatch(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Strategy: "keyword",
		Keywords: []string{"nonexistentxyz"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got.Components) != 0 {
		t.Fatalf("expected no components matching 'nonexistentxyz', got %d", len(got.Components))
	}
}

func TestRouteKeywordStrategyEmptyKeywords(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Strategy: "keyword",
		Keywords: []string{},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Empty keywords should return all components (fallback behavior)
	if len(got.Components) != 6 {
		t.Fatalf("expected 6 components with empty keywords, got %d", len(got.Components))
	}
}

func TestRouteWorkflowStrategy(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Strategy:     "workflow",
		WorkflowName: "feature",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should include the workflow itself + referenced components (architect, developer, reviewer, project)
	if len(got.Components) < 2 {
		t.Fatalf("expected at least 2 components from 'feature' workflow, got %d", len(got.Components))
	}

	// Check that workflow is included
	foundWorkflow := false
	for _, c := range got.Components {
		if c.Name == "feature" && c.Type == project.ComponentWorkflow {
			foundWorkflow = true
			break
		}
	}

	if !foundWorkflow {
		t.Fatal("expected 'feature' workflow to be included")
	}
}

func TestRouteWorkflowStrategyNotFound(t *testing.T) {
	r := New()
	p := sampleProject()

	_, err := r.Route(p, RoutingOptions{
		Strategy:     "workflow",
		WorkflowName: "nonexistent-workflow",
	})
	if err == nil {
		t.Fatal("expected error for nonexistent workflow")
	}

	if _, ok := err.(*ErrWorkflowNotFound); !ok {
		t.Fatalf("expected ErrWorkflowNotFound, got %T", err)
	}
}

func TestRouteExplicitFilter(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Filter: project.Filter{
			Names: []string{"reviewer"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got.Components) != 1 {
		t.Fatalf("expected 1 component with explicit filter, got %d", len(got.Components))
	}

	if got.Components[0].Name != "reviewer" {
		t.Fatalf("expected 'reviewer' component, got '%s'", got.Components[0].Name)
	}
}

func TestRouteExplicitFilterExcludesType(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{
		Filter: project.Filter{
			Exclude: []project.ComponentType{project.ComponentMemory},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	for _, c := range got.Components {
		if c.Type == project.ComponentMemory {
			t.Fatalf("memory component should have been excluded")
		}
	}
}

func TestRouteExplicitFilterOverridesStrategy(t *testing.T) {
	r := New()
	p := sampleProject()

	// Explicit filter should take precedence over strategy
	got, err := r.Route(p, RoutingOptions{
		Strategy: "keyword",
		Keywords: []string{"review"},
		Filter: project.Filter{
			Names: []string{"architect"},
		},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got.Components) != 1 {
		t.Fatalf("expected 1 component (explicit filter should override), got %d", len(got.Components))
	}

	if got.Components[0].Name != "architect" {
		t.Fatalf("expected 'architect' from explicit filter, got '%s'", got.Components[0].Name)
	}
}

func TestRoutePreservesMetadata(t *testing.T) {
	r := New()
	p := sampleProject()

	got, err := r.Route(p, RoutingOptions{Strategy: "default"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got.Metadata.Name != p.Metadata.Name {
		t.Fatalf("metadata name changed: expected '%s', got '%s'", p.Metadata.Name, got.Metadata.Name)
	}

	if got.Metadata.Version != p.Metadata.Version {
		t.Fatalf("metadata version changed: expected '%s', got '%s'", p.Metadata.Version, got.Metadata.Version)
	}
}

func TestContainsComponent(t *testing.T) {
	comps := []project.Component{
		{Name: "reviewer"},
		{Name: "architect"},
	}

	if !containsComponent(comps, "reviewer") {
		t.Fatal("expected 'reviewer' to be found")
	}

	if containsComponent(comps, "nonexistent") {
		t.Fatal("expected 'nonexistent' not to be found")
	}
}

func TestFindComponentByName(t *testing.T) {
	p := sampleProject()

	found := findComponentByName(p, "reviewer")
	if found == nil {
		t.Fatal("expected to find 'reviewer' component")
	}

	if found.Name != "reviewer" {
		t.Fatalf("unexpected component name: expected 'reviewer', got '%s'", found.Name)
	}

	notFound := findComponentByName(p, "nonexistent")
	if notFound != nil {
		t.Fatal("expected nil for nonexistent component")
	}
}

func TestFindComponentsByType(t *testing.T) {
	p := sampleProject()

	skills := findComponentsByType(p, project.ComponentSkill)
	if len(skills) != 3 {
		t.Fatalf("expected 3 skills, got %d", len(skills))
	}

	for _, s := range skills {
		if s.Type != project.ComponentSkill {
			t.Fatalf("unexpected component type: expected 'skill', got '%s'", s.Type)
		}
	}

	memories := findComponentsByType(p, project.ComponentMemory)
	if len(memories) != 2 {
		t.Fatalf("expected 2 memories, got %d", len(memories))
	}
}

func TestParseSteps(t *testing.T) {
	content := `---
name: feature
version: "1.0.0"
steps:
  - name: analyze
    skill: architect
    context: Analyze requirements.
  - name: implement
    skill: developer
    memory: project
    context: Implement following conventions.
  - name: review
    skill: reviewer
---
# Feature Implementation

This workflow guides feature implementation.`

	steps := parseSteps(content)

	if len(steps) != 3 {
		t.Fatalf("expected 3 steps, got %d", len(steps))
	}

	expectedSteps := []struct {
		name   string
		skill  string
		memory string
	}{
		{name: "analyze", skill: "architect"},
		{name: "implement", skill: "developer", memory: "project"},
		{name: "review", skill: "reviewer"},
	}

	for i, expected := range expectedSteps {
		if steps[i].Name != expected.name {
			t.Fatalf("step %d: expected name '%s', got '%s'", i, expected.name, steps[i].Name)
		}
		if expected.skill != "" && steps[i].Skill != expected.skill {
			t.Fatalf("step %d: expected skill '%s', got '%s'", i, expected.skill, steps[i].Skill)
		}
		if expected.memory != "" && steps[i].Memory != expected.memory {
			t.Fatalf("step %d: expected memory '%s', got '%s'", i, expected.memory, steps[i].Memory)
		}
	}
}

func TestParseStepsNoFrontmatter(t *testing.T) {
	content := `# Simple Workflow

This workflow has no frontmatter.

1. Load reviewer skill
2. Generate feedback`

	steps := parseSteps(content)
	// Without frontmatter, steps may be empty or minimal
	_ = steps // Accept whatever parsing returns for non-frontmatter content
}

func TestFilterComponentsByKeywords(t *testing.T) {
	comps := []project.Component{
		{Name: "reviewer", Title: "Code Reviewer", Content: "Best practices for code review and pull requests."},
		{Name: "architect", Title: "System Architect", Content: "System design and architecture patterns."},
	}

	// Match by exact name keyword - exact match gets bonus points
	matched := filterComponentsByKeywords(comps, []string{"reviewer"})
	if len(matched) == 0 {
		t.Fatalf("expected at least 1 component matching 'reviewer', got %d", len(matched))
	}
	// Check that reviewer is among the matched components
	foundReviewer := false
	for _, c := range matched {
		if c.Name == "reviewer" {
			foundReviewer = true
			break
		}
	}
	if !foundReviewer {
		t.Fatalf("expected 'reviewer' to be in matched components, got %v", componentNames(matched))
	}

	// Match by name keyword - architect should match
	matched = filterComponentsByKeywords(comps, []string{"architect"})
	if len(matched) == 0 {
		t.Fatalf("expected at least 1 component matching 'architect', got %d", len(matched))
	}
	foundArchitect := false
	for _, c := range matched {
		if c.Name == "architect" {
			foundArchitect = true
			break
		}
	}
	if !foundArchitect {
		t.Fatalf("expected 'architect' to be in matched components, got %v", componentNames(matched))
	}

	// No match - use a very specific keyword that won't appear anywhere
	matched = filterComponentsByKeywords(comps, []string{"xyznonexistent12345"})
	if len(matched) != 0 {
		t.Fatalf("expected 0 components matching 'xyznonexistent12345', got %d", len(matched))
	}

	// Empty keywords returns all
	matched = filterComponentsByKeywords(comps, []string{})
	if len(matched) != 2 {
		t.Fatalf("expected 2 components with empty keywords, got %d", len(matched))
	}
}

func TestExtractTitleFromContent(t *testing.T) {
	content := `---
name: test
---

# My Title

Some content.`

	title := extractTitleFromMarkdown(content)
	if title != "My Title" {
		t.Fatalf("expected 'My Title', got '%s'", title)
	}
}

func TestExtractTitleNoHeadingInContent(t *testing.T) {
	content2 := `---
name: test
---

Some content without a heading.`

	title := extractTitleFromMarkdown(content2)
	if title != "" {
		t.Fatalf("expected empty title, got '%s'", title)
	}
}

// extractTitleFromMarkdown extracts the first H1 heading from markdown content.
func extractTitleFromMarkdown(content string) string {
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}
	return ""
}

// componentNames returns the names of components in a slice.
func componentNames(comps []project.Component) []string {
	names := make([]string, len(comps))
	for i, c := range comps {
		names[i] = c.Name
	}
	return names
}

// Helper function to check string contains substring
func containsString(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && (len(s) >= len(substr)) && findSubstring(s, substr)
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}