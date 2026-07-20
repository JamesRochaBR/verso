package project

import (
	"os"
	"path/filepath"
	"testing"
)

// ========== Frontmatter Parsing Tests ==========

func TestParseFrontmatter_Valid(t *testing.T) {
	content := `---
name: architect
type: skill
version: "1.0.0"
author: "test-author"
tags: [architecture, design]
description: "Architectural review guidance"
depends: [memory, base]
status: approved
---

# Architect Skill

Content goes here.
`

	metadata, err := ParseFrontmatter(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if metadata == nil {
		t.Fatal("expected non-nil metadata")
	}
	if metadata.Name != "architect" {
		t.Errorf("expected name 'architect', got '%s'", metadata.Name)
	}
	if metadata.Type != "skill" {
		t.Errorf("expected type 'skill', got '%s'", metadata.Type)
	}
	if metadata.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got '%s'", metadata.Version)
	}
	if metadata.Author != "test-author" {
		t.Errorf("expected author 'test-author', got '%s'", metadata.Author)
	}
	if len(metadata.Tags) != 2 || metadata.Tags[0] != "architecture" || metadata.Tags[1] != "design" {
		t.Errorf("expected tags [architecture, design], got %v", metadata.Tags)
	}
	if metadata.Description != "Architectural review guidance" {
		t.Errorf("unexpected description: %s", metadata.Description)
	}
	if len(metadata.Depends) != 2 {
		t.Errorf("expected 2 depends, got %d", len(metadata.Depends))
	}
	if metadata.Status != "approved" {
		t.Errorf("expected status 'approved', got '%s'", metadata.Status)
	}
}

func TestParseFrontmatter_Empty(t *testing.T) {
	content := `# My Skill

This is a skill without frontmatter.
`

	metadata, err := ParseFrontmatter(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if metadata != nil {
		t.Fatal("expected nil metadata for content without frontmatter")
	}
}

func TestParseFrontmatter_InvalidYAML(t *testing.T) {
	content := `---
name: architect
  invalid: yaml: syntax: error
type: skill
---

# Content
`

	metadata, err := ParseFrontmatter(content)
	if err == nil {
		t.Fatal("expected error for invalid YAML")
	}
	if metadata != nil {
		t.Fatal("expected nil metadata on error")
	}
}

func TestParseFrontmatter_EmptyBody(t *testing.T) {
	content := `---
---

# Title
`

	metadata, err := ParseFrontmatter(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if metadata != nil {
		t.Fatal("expected nil metadata for empty frontmatter body")
	}
}

func TestParseFrontmatter_Minimal(t *testing.T) {
	content := `---
name: my-skill
type: skill
---

# Content
`

	metadata, err := ParseFrontmatter(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if metadata == nil {
		t.Fatal("expected non-nil metadata for minimal frontmatter")
	}
	if metadata.Name != "my-skill" {
		t.Errorf("expected name 'my-skill', got '%s'", metadata.Name)
	}
	if metadata.Type != "skill" {
		t.Errorf("expected type 'skill', got '%s'", metadata.Type)
	}
}

func TestParseFrontmatter_TagsAsList(t *testing.T) {
	content := `---
name: test-skill
type: skill
tags:
  - tag1
  - tag2
  - tag3
---

# Content
`

	metadata, err := ParseFrontmatter(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(metadata.Tags) != 3 {
		t.Errorf("expected 3 tags, got %d: %v", len(metadata.Tags), metadata.Tags)
	}
}

func TestParseFrontmatter_EmptyFrontmatter(t *testing.T) {
	content := `---
---

# Title
`

	metadata, err := ParseFrontmatter(content)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if metadata != nil {
		t.Fatal("expected nil for empty frontmatter")
	}
}

func TestFrontmatterMetadata_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		metadata *FrontmatterMetadata
		expected bool
	}{
		{
			name:     "nil metadata",
			metadata: nil,
			expected: true,
		},
		{
			name:     "empty metadata",
			metadata: &FrontmatterMetadata{},
			expected: true,
		},
		{
			name: "with name",
			metadata: &FrontmatterMetadata{
				Name: "test",
			},
			expected: false,
		},
		{
			name: "with type",
			metadata: &FrontmatterMetadata{
				Type: "skill",
			},
			expected: false,
		},
		{
			name: "with tags",
			metadata: &FrontmatterMetadata{
				Tags: []string{"test"},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.metadata.IsEmpty()
			if result != tt.expected {
				t.Errorf("IsEmpty() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// ========== ExtractBody Tests ==========

func TestExtractBody_NoFrontmatter(t *testing.T) {
	content := `# Title

Body content here.
`

	body := ExtractBody(content)
	if body != content {
		t.Errorf("expected unchanged content, got: %q", body)
	}
}

func TestExtractBody_WithFrontmatter(t *testing.T) {
	content := `---
name: test
type: skill
---

# Title

Body content here.
`

	expected := "# Title\n\nBody content here."
	body := ExtractBody(content)
	if body != expected {
		t.Errorf("expected %q, got %q", expected, body)
	}
}

func TestExtractBody_EmptyFrontmatter(t *testing.T) {
	// When frontmatter is empty (no fields), ParseFrontmatter returns nil,
	// so ExtractBody should return the original content unchanged.
	content := `---
---

# Title

Body.
`

	body := ExtractBody(content)
	if body != content {
		t.Errorf("expected unchanged content for empty frontmatter, got %q", body)
	}
}

// ========== ExtractTitle Tests (with frontmatter) ==========

func TestExtractTitle_WithFrontmatter(t *testing.T) {
	content := `---
name: test
type: skill
---

# My Title

Body content.
`

	title := ExtractTitle(content)
	if title != "My Title" {
		t.Errorf("expected 'My Title', got '%s'", title)
	}
}

func TestExtractTitle_FrontmatterNoH1(t *testing.T) {
	content := `---
name: test
type: skill
description: "A test skill"
---

Just plain text without heading.
`

	title := ExtractTitle(content)
	if title != "" {
		t.Errorf("expected empty title, got '%s'", title)
	}
}

// ========== ValidateSkill Tests ==========

func TestValidateSkill_Valid(t *testing.T) {
	skill := Component{
		Name:    "architect",
		Type:    ComponentSkill,
		Content: "# Architect\n\nSome content.",
	}

	err := ValidateSkill(skill)
	if err != nil {
		t.Errorf("expected no error for valid skill, got: %v", err)
	}
}

func TestValidateSkill_EmptyName(t *testing.T) {
	skill := Component{
		Name:    "",
		Type:    ComponentSkill,
		Content: "# Content",
	}

	err := ValidateSkill(skill)
	if err == nil {
		t.Fatal("expected error for empty skill name")
	}
}

func TestValidateSkill_WhitespaceName(t *testing.T) {
	skill := Component{
		Name:    "   ",
		Type:    ComponentSkill,
		Content: "# Content",
	}

	err := ValidateSkill(skill)
	if err == nil {
		t.Fatal("expected error for whitespace-only skill name")
	}
}

func TestValidateSkill_WrongType(t *testing.T) {
	skill := Component{
		Name:    "test",
		Type:    ComponentMemory,
		Content: "# Content",
	}

	err := ValidateSkill(skill)
	if err == nil {
		t.Fatal("expected error for wrong component type")
	}
}

func TestValidateSkill_EmptyContent(t *testing.T) {
	skill := Component{
		Name:    "test",
		Type:    ComponentSkill,
		Content: "",
	}

	err := ValidateSkill(skill)
	if err == nil {
		t.Fatal("expected error for empty content")
	}
}

func TestValidateSkill_WhitespaceContent(t *testing.T) {
	skill := Component{
		Name:    "test",
		Type:    ComponentSkill,
		Content: "   \n\n  ",
	}

	err := ValidateSkill(skill)
	if err == nil {
		t.Fatal("expected error for whitespace-only content")
	}
}

// ========== ValidateLifecycle Tests ==========

func TestValidateLifecycle_NoStatus(t *testing.T) {
	component := Component{
		Name:   "test",
		Status: "",
	}

	err := ValidateLifecycle(component)
	if err != nil {
		t.Errorf("expected no error for empty status, got: %v", err)
	}
}

func TestValidateLifecycle_ValidStatus(t *testing.T) {
	tests := []LifecycleState{
		StateCreated,
		StateReviewed,
		StateApproved,
		StateDeprecated,
	}

	for _, state := range tests {
		t.Run(string(state), func(t *testing.T) {
			component := Component{Name: "test", Status: state}
			err := ValidateLifecycle(component)
			if err != nil {
				t.Errorf("expected no error for status %q, got: %v", state, err)
			}
		})
	}
}

func TestValidateLifecycle_InvalidStatus(t *testing.T) {
	component := Component{
		Name:   "test",
		Status: "invalid-state",
	}

	err := ValidateLifecycle(component)
	if err == nil {
		t.Fatal("expected error for invalid lifecycle state")
	}
}

func TestLifecycleState_IsValid(t *testing.T) {
	tests := []struct {
		state    LifecycleState
		expected bool
	}{
		{StateCreated, true},
		{StateReviewed, true},
		{StateApproved, true},
		{StateDeprecated, true},
		{LifecycleState("unknown"), false},
		{LifecycleState(""), false},
	}

	for _, tt := range tests {
		t.Run(string(tt.state), func(t *testing.T) {
			result := tt.state.IsValid()
			if result != tt.expected {
				t.Errorf("IsValid(%q) = %v, want %v", tt.state, result, tt.expected)
			}
		})
	}
}

// ========== Component Helper Method Tests ==========

func TestComponent_HasTags_Match(t *testing.T) {
	component := Component{
		Name: "test",
		Tags: []string{"architecture", "design"},
	}

	if !component.HasTags([]string{"architecture"}) {
		t.Error("expected tag match")
	}
}

func TestComponent_HasTags_NoMatch(t *testing.T) {
	component := Component{
		Name: "test",
		Tags: []string{"architecture", "design"},
	}

	if component.HasTags([]string{"testing"}) {
		t.Error("expected no tag match")
	}
}

func TestComponent_HasTags_EmptyComponentTags(t *testing.T) {
	component := Component{
		Name: "test",
		Tags: []string{},
	}

	if component.HasTags([]string{"architecture"}) {
		t.Error("expected no match for empty tags")
	}
}

func TestComponent_HasTags_EmptyFilterTags(t *testing.T) {
	component := Component{
		Name: "test",
		Tags: []string{"architecture"},
	}

	if component.HasTags([]string{}) {
		t.Error("expected no match for empty filter tags")
	}
}

func TestComponent_ContainsKeyword_Name(t *testing.T) {
	component := Component{
		Name: "architect",
	}

	if !component.ContainsKeyword("architect") {
		t.Error("expected keyword to match name")
	}
}

func TestComponent_ContainsKeyword_Title(t *testing.T) {
	component := Component{
		Title: "Architect Skill",
	}

	if !component.ContainsKeyword("architect") {
		t.Error("expected keyword to match title")
	}
}

func TestComponent_ContainsKeyword_Description(t *testing.T) {
	component := Component{
		Description: "System design guidance",
	}

	if !component.ContainsKeyword("design") {
		t.Error("expected keyword to match description")
	}
}

func TestComponent_ContainsKeyword_Tag(t *testing.T) {
	component := Component{
		Tags: []string{"architecture", "system-design"},
	}

	if !component.ContainsKeyword("architecture") {
		t.Error("expected keyword to match tag")
	}
}

func TestComponent_ContainsKeyword_NoMatch(t *testing.T) {
	component := Component{
		Name:    "reviewer",
		Title:   "Code Reviewer",
		Tags:    []string{"code-review"},
	}

	if component.ContainsKeyword("architect") {
		t.Error("expected no keyword match")
	}
}

func TestComponent_IsDeprecated(t *testing.T) {
	component := Component{
		Status: StateDeprecated,
	}

	if !component.IsDeprecated() {
		t.Error("expected IsDeprecated to return true")
	}
}

func TestComponent_NotDeprecated(t *testing.T) {
	component := Component{
		Status: StateApproved,
	}

	if component.IsDeprecated() {
		t.Error("expected IsDeprecated to return false")
	}
}

// ========== Discover with Frontmatter Integration Tests ==========

func TestDiscover_Skills_WithMetadata(t *testing.T) {
	// Create a temporary project directory
	tmpDir, err := os.MkdirTemp("", "verso-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create required directories and files
	dirs := []string{"skills", "memory", "workflows", "templates"}
	for _, dir := range dirs {
		if err := os.Mkdir(filepath.Join(tmpDir, dir), 0755); err != nil {
			t.Fatal(err)
		}
	}

	// Create verso.toml
	tomlContent := `name = "test-project"
version = "0.1.0"
`
	if err := os.WriteFile(filepath.Join(tmpDir, "verso.toml"), []byte(tomlContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a skill with frontmatter
	skillContent := `---
name: test-architect
type: skill
version: "1.0.0"
author: "tester"
tags: [test, architecture]
description: "Test skill description"
status: approved
---

# Test Architect Skill

This is the content of the test skill.
`
	if err := os.WriteFile(filepath.Join(tmpDir, "skills", "architect.md"), []byte(skillContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Discover components
	components, err := Discover(tmpDir)
	if err != nil {
		t.Fatalf("discover failed: %v", err)
	}

	// Find the skill component
	var found *Component
	for i, c := range components {
		if c.Type == ComponentSkill {
			found = &components[i]
			break
		}
	}

	if found == nil {
		t.Fatal("expected to find a skill component")
	}

	// Verify metadata was populated correctly
	if found.Name != "test-architect" {
		t.Errorf("expected name 'test-architect', got '%s'", found.Name)
	}
	if found.Version != "1.0.0" {
		t.Errorf("expected version '1.0.0', got '%s'", found.Version)
	}
	if found.Author != "tester" {
		t.Errorf("expected author 'tester', got '%s'", found.Author)
	}
	if len(found.Tags) != 2 || found.Tags[0] != "test" || found.Tags[1] != "architecture" {
		t.Errorf("unexpected tags: %v", found.Tags)
	}
	if found.Description != "Test skill description" {
		t.Errorf("unexpected description: %s", found.Description)
	}
	if found.Status != StateApproved {
		t.Errorf("expected status 'approved', got '%s'", found.Status)
	}
	if found.Metadata == nil {
		t.Fatal("expected non-nil Metadata")
	}
}

func TestDiscover_Skills_WithoutFrontmatter(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "verso-test-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	dirs := []string{"skills", "memory", "workflows", "templates"}
	for _, dir := range dirs {
		if err := os.Mkdir(filepath.Join(tmpDir, dir), 0755); err != nil {
			t.Fatal(err)
		}
	}

	tomlContent := `name = "test-project"
version = "0.1.0"
`
	if err := os.WriteFile(filepath.Join(tmpDir, "verso.toml"), []byte(tomlContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Create a skill WITHOUT frontmatter (legacy format)
	skillContent := `# Legacy Skill

This is a legacy skill without YAML frontmatter.
`
	if err := os.WriteFile(filepath.Join(tmpDir, "skills", "legacy.md"), []byte(skillContent), 0644); err != nil {
		t.Fatal(err)
	}

	components, err := Discover(tmpDir)
	if err != nil {
		t.Fatalf("discover failed: %v", err)
	}

	var found *Component
	for i, c := range components {
		if c.Type == ComponentSkill {
			found = &components[i]
			break
		}
	}

	if found == nil {
		t.Fatal("expected to find a skill component")
	}

	// Should use filename as name when no frontmatter
	if found.Name != "legacy" {
		t.Errorf("expected name 'legacy', got '%s'", found.Name)
	}
	if found.Title != "Legacy Skill" {
		t.Errorf("expected title 'Legacy Skill', got '%s'", found.Title)
	}
	if found.Metadata != nil {
		t.Error("expected nil Metadata for skill without frontmatter")
	}
}

// ========== ValidateComponent Tests ==========

func TestValidateComponent_Skill(t *testing.T) {
	skill := Component{
		Name:    "test",
		Type:    ComponentSkill,
		Content: "# Content",
	}

	err := ValidateComponent(skill)
	if err != nil {
		t.Errorf("expected no error for valid skill, got: %v", err)
	}
}

func TestValidateComponent_Generic(t *testing.T) {
	component := Component{
		Name:    "test-component",
		Type:    ComponentMemory,
		Content: "# Content",
	}

	err := ValidateComponent(component)
	if err != nil {
		t.Errorf("expected no error for valid component, got: %v", err)
	}
}

func TestValidateComponent_EmptyName(t *testing.T) {
	component := Component{
		Name:    "",
		Type:    ComponentMemory,
		Content: "# Content",
	}

	err := ValidateComponent(component)
	if err == nil {
		t.Fatal("expected error for empty component name")
	}
}

// ========== ValidateComponents Tests ==========

func TestValidateComponents_AllValid(t *testing.T) {
	components := []Component{
		{Name: "skill1", Type: ComponentSkill, Content: "# Content"},
		{Name: "memory1", Type: ComponentMemory, Content: "# Content"},
	}

	err := ValidateComponents(components)
	if err != nil {
		t.Errorf("expected no error for valid components, got: %v", err)
	}
}

func TestValidateComponents_InvalidSkill(t *testing.T) {
	components := []Component{
		{Name: "", Type: ComponentSkill, Content: "# Content"},
	}

	err := ValidateComponents(components)
	if err == nil {
		t.Fatal("expected error for invalid component")
	}
}