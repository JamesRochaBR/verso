package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscover(t *testing.T) {

	dir := t.TempDir()

	os.Mkdir(filepath.Join(dir, "skills"), 0755)
	os.Mkdir(filepath.Join(dir, "memory"), 0755)
	os.Mkdir(filepath.Join(dir, "workflows"), 0755)
	os.Mkdir(filepath.Join(dir, "templates"), 0755)

	os.WriteFile(
		filepath.Join(dir, "skills", "reviewer.md"),
		[]byte("# Code Reviewer\n\nReview code."),
		0644,
	)

	os.WriteFile(
		filepath.Join(dir, "memory", "project.md"),
		[]byte("# Project Memory\n\nImportant."),
		0644,
	)

	components, err := Discover(dir)

	if err != nil {
		t.Fatal(err)
	}

	if len(components) != 2 {
		t.Fatalf("expected 2 components, got %d", len(components))
	}

	foundSkill := false
	foundMemory := false

	for _, c := range components {

		switch c.Name {

		case "reviewer":

			foundSkill = true

			if c.Title != "Code Reviewer" {
				t.Fatalf("unexpected title %q", c.Title)
			}

			if c.Type != ComponentSkill {
				t.Fatalf("unexpected type %v", c.Type)
			}

		case "project":

			foundMemory = true

			if c.Title != "Project Memory" {
				t.Fatalf("unexpected title %q", c.Title)
			}

			if c.Type != ComponentMemory {
				t.Fatalf("unexpected type %v", c.Type)
			}
		}
	}

	if !foundSkill {
		t.Fatal("skill component not found")
	}

	if !foundMemory {
		t.Fatal("memory component not found")
	}
}
