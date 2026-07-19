package render

import (
	"strings"
	"testing"

	"github.com/james-rocha/verso/internal/project"
)

func TestMarkdownRendererRegistered(t *testing.T) {

	r, ok := Get("markdown")

	if !ok {
		t.Fatal("renderer not registered")
	}

	if r.Name() != "markdown" {
		t.Fatal("unexpected renderer")
	}
}

func TestPromptReturnsError(t *testing.T) {

	p := &project.Project{
		Metadata: project.Metadata{
			Name:        "test",
			Version:     "0.1.0",
			Description: "Test project",
		},
		Components: []project.Component{},
	}

	out, err := Prompt(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "# Project") {
		t.Fatal("expected prompt to contain project header")
	}

	if !strings.Contains(out, "Name: test") {
		t.Fatal("expected prompt to contain project name")
	}
}

func TestPromptWithComponents(t *testing.T) {

	p := &project.Project{
		Metadata: project.Metadata{
			Name:    "demo",
			Version: "1.0.0",
		},
		Components: []project.Component{
			{
				Name:    "reviewer",
				Title:   "Code Reviewer",
				Type:    project.ComponentSkill,
				Content: "Review code carefully.",
			},
			{
				Name:    "architecture",
				Title:   "Architecture Memory",
				Type:    project.ComponentMemory,
				Content: "Use clean architecture.",
			},
		},
	}

	out, err := Prompt(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(out, "# Skills") {
		t.Fatal("expected prompt to contain skills section")
	}

	if !strings.Contains(out, "## Code Reviewer") {
		t.Fatal("expected prompt to contain skill title")
	}

	if !strings.Contains(out, "Review code carefully.") {
		t.Fatal("expected prompt to contain skill content")
	}

	if !strings.Contains(out, "# Memory") {
		t.Fatal("expected prompt to contain memory section")
	}

	if !strings.Contains(out, "Use clean architecture.") {
		t.Fatal("expected prompt to contain memory content")
	}
}
