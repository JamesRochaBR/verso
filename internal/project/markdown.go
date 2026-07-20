package project

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// FrontmatterMetadata represents the YAML frontmatter fields for a Verso component.
// Fields follow the RFC-0009 Serialization specification.
type FrontmatterMetadata struct {
	Name        string   `yaml:"name"`
	Type        string   `yaml:"type"`
	Version     string   `yaml:"version"`
	Author      string   `yaml:"author"`
	Tags        []string `yaml:"tags"`
	Description string   `yaml:"description"`
	Depends     []string `yaml:"depends"`
	Status      string   `yaml:"status"`
}

// IsEmpty returns true if the metadata has no meaningful fields set.
func (m *FrontmatterMetadata) IsEmpty() bool {
	return m == nil || (m.Name == "" && m.Type == "" && m.Version == "" &&
		m.Author == "" && len(m.Tags) == 0 && m.Description == "" &&
		len(m.Depends) == 0 && m.Status == "")
}

// ParseFrontmatter extracts and parses YAML frontmatter from a markdown content string.
// It expects the content to start with "---" delimiters wrapping the YAML block.
// Returns nil if no valid frontmatter is found.
func ParseFrontmatter(content string) (*FrontmatterMetadata, error) {
	trimmed := strings.TrimSpace(content)
	if !strings.HasPrefix(trimmed, "---") {
		return nil, nil
	}

	// Find the closing delimiter
	rest := trimmed[3:]
	closeIdx := strings.Index(rest, "\n---")
	if closeIdx == -1 {
		// Try without newline (same line closing)
		closeIdx = strings.Index(rest, "---")
		if closeIdx == -1 {
			return nil, nil
		}
		rest = rest[:closeIdx]
	} else {
		rest = rest[:closeIdx]
	}

	yamlContent := strings.TrimSpace(rest)
	if len(yamlContent) == 0 {
		return nil, nil
	}

	var metadata FrontmatterMetadata
	if err := yaml.Unmarshal([]byte(yamlContent), &metadata); err != nil {
		return nil, fmt.Errorf("invalid frontmatter YAML in content: %w", err)
	}

	if metadata.IsEmpty() {
		return nil, nil
	}

	return &metadata, nil
}

// ExtractBody returns the markdown body content without the YAML frontmatter.
// If no frontmatter is detected, the entire content is returned unchanged.
func ExtractBody(content string) string {
	metadata, err := ParseFrontmatter(content)
	if err != nil || metadata == nil {
		return content
	}

	// Find content after closing "---" delimiter
	trimmed := strings.TrimSpace(content)
	rest := trimmed[3:]
	closeIdx := strings.Index(rest, "\n---")
	if closeIdx == -1 {
		return strings.TrimSpace(rest)
	}

	body := strings.TrimSpace(rest[closeIdx+4:])
	return body
}

// ExtractTitle extracts the first H1 heading from markdown content.
// It automatically skips YAML frontmatter if present.
func ExtractTitle(content string) string {
	body := ExtractBody(content)
	lines := strings.Split(body, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "# ") {
			return strings.TrimPrefix(line, "# ")
		}
	}

	return ""
}

// lowercaseString converts a string to lowercase.
func lowercaseString(s string) string {
	return strings.ToLower(s)
}

// containsString checks if haystack contains substring (case-insensitive).
func containsString(haystack, substr string) bool {
	return strings.Contains(haystack, substr)
}

// containsTag checks if any tag in tags matches the given keyword (case-insensitive).
func containsTag(tags []string, keyword string) bool {
	for _, tag := range tags {
		if lowercaseString(tag) == keyword || strings.Contains(lowercaseString(tag), keyword) {
			return true
		}
	}
	return false
}
