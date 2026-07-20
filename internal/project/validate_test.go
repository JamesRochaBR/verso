package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidate(t *testing.T) {

	t.Run("missing verso.toml", func(t *testing.T) {
		dir := t.TempDir()

		if err := Validate(dir); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("valid project", func(t *testing.T) {

		dir := t.TempDir()

		os.WriteFile(filepath.Join(dir, "verso.toml"), []byte(`
name="demo"
version="0.1.0"
description=""
`), 0644)

		os.Mkdir(filepath.Join(dir, "skills"), 0755)
		os.Mkdir(filepath.Join(dir, "memory"), 0755)
		os.Mkdir(filepath.Join(dir, "workflows"), 0755)
		os.Mkdir(filepath.Join(dir, "templates"), 0755)

		if err := Validate(dir); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
