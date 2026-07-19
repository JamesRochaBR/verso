package project

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func Load(path string) (*Project, error) {
	data, err := os.ReadFile(filepath.Join(path, "verso.toml"))
	if err != nil {
		return nil, err
	}

	var metadata Metadata

	if err := toml.Unmarshal(data, &metadata); err != nil {
		return nil, err
	}

	return &Project{
		Metadata: metadata,
	}, nil
}