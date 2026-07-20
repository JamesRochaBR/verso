package project

type Metadata struct {
	Name        string `toml:"name" json:"name"`
	Version     string `toml:"version" json:"version"`
	Description string `toml:"description" json:"description"`
}
