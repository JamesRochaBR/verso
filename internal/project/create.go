package project

import (
	"fmt"
	"os"
	"path/filepath"
)

func Create(path string) error {
	// Verifica se o diretório já existe
	if _, err := os.Stat(path); err == nil {
		return fmt.Errorf("directory '%s' already exists", path)
	}

	// Cria a estrutura
	dirs := []string{
		path,
		filepath.Join(path, "skills"),
		filepath.Join(path, "memory"),
		filepath.Join(path, "workflows"),
		filepath.Join(path, "templates"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	// Cria o verso.toml
	content := fmt.Sprintf(`name = "%s"
version = "0.1.0"
description = ""
`, filepath.Base(path))

	return os.WriteFile(
		filepath.Join(path, "verso.toml"),
		[]byte(content),
		0644,
	)
}