package render

import (
	"strings"
	"testing"

	"github.com/james-rocha/verso/internal/project"
)

func TestPrompt(t *testing.T) {

	p := &project.Project{
		Metadata: project.Metadata{
			Name:    "demo",
			Version: "0.1.0",
		},
		Components: []project.Component{
			{
				Name:    "reviewer",
				Title:   "Code Reviewer",
				Type:    project.ComponentSkill,
				Content: "Review code.",
			},
		},
	}

	out := Prompt(p)

	if !strings.Contains(out, "# Project") {
		t.Fatal("missing project section")
	}

	if !strings.Contains(out, "Code Reviewer") {
		t.Fatal("missing component")
	}

	if !strings.Contains(out, "Review code.") {
		t.Fatal("missing content")
	}
}
