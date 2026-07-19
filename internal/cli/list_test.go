package cli

import (
	"os"
	"path/filepath"
	"testing"
)

func TestListCommandRequiresPath(t *testing.T) {
	cmd := ListCommand{}

	if err := cmd.Run(nil); err == nil {
		t.Fatal("expected error when project path is missing")
	}
}

func TestListCommandFailsOnInvalidProject(t *testing.T) {
	cmd := ListCommand{}

	err := cmd.Run([]string{"/nonexistent/path"})
	if err == nil {
		t.Fatal("expected error for invalid project path")
	}
}

func TestListCommandListsComponents(t *testing.T) {
	dir := t.TempDir()

	os.Mkdir(filepath.Join(dir, "skills"), 0755)
	os.Mkdir(filepath.Join(dir, "memory"), 0755)
	os.WriteFile(
		filepath.Join(dir, "verso.toml"),
		[]byte("name = \"test\""),
		0644,
	)

	os.WriteFile(
		filepath.Join(dir, "skills", "reviewer.md"),
		[]byte("# Reviewer\n\nReview code."),
		0644,
	)

	cmd := ListCommand{}
	err := cmd.Run([]string{dir})
	if err != nil {
		t.Fatal(err)
	}
}

func TestListCommandEmptyProject(t *testing.T) {
	dir := t.TempDir()

	os.Mkdir(filepath.Join(dir, "skills"), 0755)
	os.WriteFile(
		filepath.Join(dir, "verso.toml"),
		[]byte("name = \"empty\""),
		0644,
	)

	cmd := ListCommand{}
	err := cmd.Run([]string{dir})
	if err != nil {
		t.Fatal(err)
	}
}
