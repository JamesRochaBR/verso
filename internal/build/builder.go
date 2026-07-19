package build

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/james-rocha/verso/internal/project"
)

func Build(path string) error {

	p, err := project.Load(path)
	if err != nil {
		return err
	}

	outDir := filepath.Join(path, ".verso")

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(
		filepath.Join(outDir, "context.json"),
		data,
		0644,
	)
}