package project

import (
	"os"
	"path/filepath"
	"testing"
)

// --- ParseSemanticVersion Tests ---

func TestParseSemanticVersion_Valid(t *testing.T) {
	tests := []struct {
		input    string
		expected *SemVer
	}{
		{"1.0.0", &SemVer{Major: 1, Minor: 0, Patch: 0}},
		{"0.0.1", &SemVer{Major: 0, Minor: 0, Patch: 1}},
		{"2.3.4", &SemVer{Major: 2, Minor: 3, Patch: 4}},
		{"v1.0.0", &SemVer{Major: 1, Minor: 0, Patch: 0}},
		{"1.0.0-alpha", &SemVer{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha"}},
		{"1.0.0-beta.1", &SemVer{Major: 1, Minor: 0, Patch: 0, Prerelease: "beta.1"}},
		{"1.0.0-rc.1+build.123", &SemVer{Major: 1, Minor: 0, Patch: 0, Prerelease: "rc.1+build.123"}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ParseSemanticVersion(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Major != tt.expected.Major || result.Minor != tt.expected.Minor || result.Patch != tt.expected.Patch {
				t.Errorf("expected %+v, got %+v", *tt.expected, *result)
			}
			if result.Prerelease != tt.expected.Prerelease {
				t.Errorf("expected prerelease %q, got %q", tt.expected.Prerelease, result.Prerelease)
			}
		})
	}
}

func TestParseSemanticVersion_Empty(t *testing.T) {
	_, err := ParseSemanticVersion("")
	if err == nil {
		t.Fatal("expected error for empty version")
	}
}

func TestParseSemanticVersion_InvalidFormat(t *testing.T) {
	tests := []string{"1.0", "abc", "1.0.x", "one.two.three", "1..0"}
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			_, err := ParseSemanticVersion(tt)
			if err == nil {
				t.Fatalf("expected error for invalid version %q", tt)
			}
		})
	}
}

// --- SemVer.String() Tests ---

func TestSemVer_String(t *testing.T) {
	tests := []struct {
		input    *SemVer
		expected string
	}{
		{&SemVer{Major: 1, Minor: 0, Patch: 0}, "1.0.0"},
		{&SemVer{Major: 2, Minor: 3, Patch: 4}, "2.3.4"},
		{&SemVer{Major: 1, Minor: 0, Patch: 0, Prerelease: "alpha"}, "1.0.0-alpha"},
		{&SemVer{Major: 0, Minor: 1, Patch: 0, Prerelease: "beta.1"}, "0.1.0-beta.1"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := tt.input.String()
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// --- SemVer Bump Methods Tests ---

func TestSemVer_BumpMajor(t *testing.T) {
	sem := &SemVer{Major: 1, Minor: 2, Patch: 3}
	bumped := sem.BumpMajor()
	if bumped.Major != 2 || bumped.Minor != 0 || bumped.Patch != 0 {
		t.Errorf("expected 2.0.0, got %s", bumped.String())
	}
}

func TestSemVer_BumpMinor(t *testing.T) {
	sem := &SemVer{Major: 1, Minor: 2, Patch: 3}
	bumped := sem.BumpMinor()
	if bumped.Major != 1 || bumped.Minor != 3 || bumped.Patch != 0 {
		t.Errorf("expected 1.3.0, got %s", bumped.String())
	}
}

func TestSemVer_BumpPatch(t *testing.T) {
	sem := &SemVer{Major: 1, Minor: 2, Patch: 3}
	bumped := sem.BumpPatch()
	if bumped.Major != 1 || bumped.Minor != 2 || bumped.Patch != 4 {
		t.Errorf("expected 1.2.4, got %s", bumped.String())
	}
}

// --- CompareVersions Tests ---

func TestCompareVersions_Equal(t *testing.T) {
	cmp, err := CompareVersions("1.0.0", "1.0.0")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmp != 0 {
		t.Errorf("expected 0, got %d", cmp)
	}
}

func TestCompareVersions_AGreaterB(t *testing.T) {
	tests := []struct {
		a, b string
	}{
		{"2.0.0", "1.0.0"},
		{"1.1.0", "1.0.0"},
		{"1.0.1", "1.0.0"},
	}
	for _, tt := range tests {
		t.Run(tt.a+">"+tt.b, func(t *testing.T) {
			cmp, err := CompareVersions(tt.a, tt.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cmp != 1 {
				t.Errorf("expected 1, got %d", cmp)
			}
		})
	}
}

func TestCompareVersions_ALessThanB(t *testing.T) {
	tests := []struct {
		a, b string
	}{
		{"1.0.0", "2.0.0"},
		{"1.0.0", "1.1.0"},
		{"1.0.0", "1.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.a+"<"+tt.b, func(t *testing.T) {
			cmp, err := CompareVersions(tt.a, tt.b)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if cmp != -1 {
				t.Errorf("expected -1, got %d", cmp)
			}
		})
	}
}

func TestCompareVersions_Prerelease(t *testing.T) {
	cmp, err := CompareVersions("1.0.0", "1.0.0-alpha")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmp != 1 {
		t.Errorf("expected 1 (release > prerelease), got %d", cmp)
	}

	cmp, err = CompareVersions("1.0.0-alpha", "1.0.0-beta")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cmp != -1 {
		t.Errorf("expected -1 (alpha < beta), got %d", cmp)
	}
}

func TestCompareVersions_Invalid(t *testing.T) {
	_, err := CompareVersions("invalid", "1.0.0")
	if err == nil {
		t.Fatal("expected error for invalid version")
	}
}

// --- IsVersioned Tests ---

func TestIsVersioned_WithValidVersion(t *testing.T) {
	component := Component{Version: "1.0.0"}
	if !IsVersioned(component) {
		t.Error("expected true for valid version")
	}
}

func TestIsVersioned_WithoutVersion(t *testing.T) {
	component := Component{Version: ""}
	if IsVersioned(component) {
		t.Error("expected false for empty version")
	}
}

func TestIsVersioned_WithInvalidVersion(t *testing.T) {
	component := Component{Version: "invalid"}
	if IsVersioned(component) {
		t.Error("expected false for invalid version")
	}
}

// --- ValidateMemory Tests ---

func TestValidateMemory_Valid(t *testing.T) {
	component := Component{
		Name:    "project",
		Type:    ComponentMemory,
		Content: "# Project Memory\nSome content here.",
	}
	err := ValidateMemory(component)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateMemory_ValidWithFrontmatter(t *testing.T) {
	component := Component{
		Name:        "architecture",
		Type:        ComponentMemory,
		Version:     "1.0.0",
		Author:      "team",
		Tags:        []string{"architecture"},
		Description: "System architecture overview",
		Content:     "# Architecture\nDetailed content.",
	}
	err := ValidateMemory(component)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateMemory_EmptyName(t *testing.T) {
	component := Component{
		Name:    "",
		Type:    ComponentMemory,
		Content: "# Some Memory",
	}
	err := ValidateMemory(component)
	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestValidateMemory_WrongType(t *testing.T) {
	component := Component{
		Name:    "test",
		Type:    ComponentSkill,
		Content: "# Some Skill",
	}
	err := ValidateMemory(component)
	if err == nil {
		t.Fatal("expected error for wrong type")
	}
}

func TestValidateMemory_EmptyContent(t *testing.T) {
	component := Component{
		Name:    "test",
		Type:    ComponentMemory,
		Content: "",
	}
	err := ValidateMemory(component)
	if err == nil {
		t.Fatal("expected error for empty content")
	}
}

// --- ValidateComponent with Memory Tests ---

func TestValidateComponent_Memory(t *testing.T) {
	component := Component{
		Name:    "test",
		Type:    ComponentMemory,
		Content: "# Test Memory\nContent here.",
	}
	err := ValidateComponent(component)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateComponent_MemoryInvalid(t *testing.T) {
	component := Component{
		Name:    "",
		Type:    ComponentMemory,
		Content: "",
	}
	err := ValidateComponent(component)
	if err == nil {
		t.Fatal("expected error for invalid memory component")
	}
}

// --- GetLatestVersion Tests ---

func TestGetLatestVersion_SingleFile(t *testing.T) {
	dir := t.TempDir()
	memoryDir := filepath.Join(dir, "memory")
	os.MkdirAll(memoryDir, 0755)

	content := `---
name: project
version: "1.0.0"
---

# Project Memory
Content here.`

	err := os.WriteFile(filepath.Join(memoryDir, "project.md"), []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	component := Component{Name: "project", Type: ComponentMemory}
	mv, err := GetLatestVersion(component, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mv.Version != "1.0.0" {
		t.Errorf("expected version 1.0.0, got %s", mv.Version)
	}
}

func TestGetLatestVersion_MultipleVersions(t *testing.T) {
	dir := t.TempDir()
	memoryDir := filepath.Join(dir, "memory")
	os.MkdirAll(memoryDir, 0755)

	content1 := `---
name: project
version: "1.0.0"
---

# Project Memory v1`

	content2 := `---
name: project
version: "2.0.0"
---

# Project Memory v2`

	os.WriteFile(filepath.Join(memoryDir, "project.md"), []byte(content1), 0644)
	os.WriteFile(filepath.Join(memoryDir, "project.v2.md"), []byte(content2), 0644)

	component := Component{Name: "project", Type: ComponentMemory}
	mv, err := GetLatestVersion(component, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if mv.Version != "2.0.0" {
		t.Errorf("expected version 2.0.0 (latest), got %s", mv.Version)
	}
}

func TestGetLatestVersion_NoVersions(t *testing.T) {
	dir := t.TempDir()
	component := Component{Name: "nonexistent", Type: ComponentMemory}
	_, err := GetLatestVersion(component, dir)
	if err == nil {
		t.Fatal("expected error for nonexistent component")
	}
}

// --- ListMemoryVersions Tests ---

func TestListMemoryVersions_MultipleFiles(t *testing.T) {
	dir := t.TempDir()
	memoryDir := filepath.Join(dir, "memory")
	os.MkdirAll(memoryDir, 0755)

	content1 := `---
name: project
version: "1.0.0"
---

# Project Memory v1`

	content2 := `---
name: project
version: "2.0.0"
---

# Project Memory v2`

	os.WriteFile(filepath.Join(memoryDir, "project.md"), []byte(content1), 0644)
	os.WriteFile(filepath.Join(memoryDir, "project.v2.md"), []byte(content2), 0644)

	component := Component{Name: "project", Type: ComponentMemory}
	versions, err := ListMemoryVersions(component, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(versions) != 2 {
		t.Errorf("expected 2 versions, got %d", len(versions))
	}
}

func TestListMemoryVersions_ExcludesNonMatching(t *testing.T) {
	dir := t.TempDir()
	memoryDir := filepath.Join(dir, "memory")
	os.MkdirAll(memoryDir, 0755)

	content1 := `---
name: project
version: "1.0.0"
---

# Project Memory`

	content2 := `---
name: conventions
version: "1.0.0"
---

# Conventions`

	os.WriteFile(filepath.Join(memoryDir, "project.md"), []byte(content1), 0644)
	os.WriteFile(filepath.Join(memoryDir, "conventions.md"), []byte(content2), 0644)

	component := Component{Name: "project", Type: ComponentMemory}
	versions, err := ListMemoryVersions(component, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(versions) != 1 {
		t.Errorf("expected 1 version for 'project', got %d", len(versions))
	}
}

func TestListMemoryVersions_ExcludesDirectories(t *testing.T) {
	dir := t.TempDir()
	memoryDir := filepath.Join(dir, "memory")
	os.MkdirAll(memoryDir, 0755)

	// Create a subdirectory (should be ignored)
	os.MkdirAll(filepath.Join(memoryDir, "subdir"), 0755)

	content := `---
name: project
version: "1.0.0"
---

# Project Memory`

	os.WriteFile(filepath.Join(memoryDir, "project.md"), []byte(content), 0644)

	component := Component{Name: "project", Type: ComponentMemory}
	versions, err := ListMemoryVersions(component, dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(versions) != 1 {
		t.Errorf("expected 1 version (directory excluded), got %d", len(versions))
	}
}

// --- ValidateLifecycle for Memory Tests ---

func TestValidateLifecycle_Memory_Default(t *testing.T) {
	component := Component{
		Name: "test",
		Type: ComponentMemory,
	}
	err := ValidateLifecycle(component)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Note: ValidateLifecycle receives component by value, so it modifies a copy.
	// We verify no error is returned and the function handles empty status gracefully.
}

func TestValidateLifecycle_Memory_Approved(t *testing.T) {
	component := Component{
		Name:   "test",
		Type:   ComponentMemory,
		Status: StateApproved,
	}
	err := ValidateLifecycle(component)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateLifecycle_Memory_Invalid(t *testing.T) {
	component := Component{
		Name:   "test",
		Type:   ComponentMemory,
		Status: LifecycleState("invalid-state"),
	}
	err := ValidateLifecycle(component)
	if err == nil {
		t.Fatal("expected error for invalid lifecycle state")
	}
}

// --- Discover Memory Integration Tests ---

func TestDiscover_MemoryComponent(t *testing.T) {
	dir := t.TempDir()

	// Create required directories and files
	os.MkdirAll(filepath.Join(dir, "memory"), 0755)
	os.WriteFile(filepath.Join(dir, "verso.toml"), []byte("name = \"test\"\nversion = \"1.0.0\""), 0644)
	os.MkdirAll(filepath.Join(dir, "skills"), 0755)
	os.MkdirAll(filepath.Join(dir, "workflows"), 0755)
	os.MkdirAll(filepath.Join(dir, "templates"), 0755)

	content := `---
name: project
type: memory
version: "1.0.0"
tags: [project]
description: "Project overview"
status: approved
---

# Project Memory
Important information.`

	os.WriteFile(filepath.Join(dir, "memory", "project.md"), []byte(content), 0644)

	components, err := Discover(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var found bool
	for _, comp := range components {
		if comp.Name == "project" && comp.Type == ComponentMemory {
			found = true
			if comp.Version != "1.0.0" {
				t.Errorf("expected version 1.0.0, got %s", comp.Version)
			}
			if comp.Description != "Project overview" {
				t.Errorf("expected description 'Project overview', got %q", comp.Description)
			}
			if comp.Status != StateApproved {
				t.Errorf("expected status 'approved', got %q", comp.Status)
			}
		}
	}

	if !found {
		t.Fatal("expected to discover memory component 'project'")
	}
}

func TestValidateComponents_MemoryBatch(t *testing.T) {
	components := []Component{
		{Name: "project", Type: ComponentMemory, Content: "# Project"},
		{Name: "conventions", Type: ComponentMemory, Content: "# Conventions"},
		{Name: "architecture", Type: ComponentSkill}, // No content — this should fail ValidateSkill
	}
	err := ValidateComponents(components)
	if err == nil {
		t.Fatal("expected error for invalid skill component in batch")
	}
}

func TestValidateComponents_MemoryBatchValid(t *testing.T) {
	components := []Component{
		{Name: "project", Type: ComponentMemory, Content: "# Project"},
		{Name: "conventions", Type: ComponentMemory, Content: "# Conventions"},
	}
	err := ValidateComponents(components)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

// --- FSTest-based tests for edge cases ---

func TestGetLatestVersion_EmptyMemoryDir(t *testing.T) {
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "memory"), 0755)

	component := Component{Name: "test", Type: ComponentMemory}
	_, err := GetLatestVersion(component, dir)
	if err == nil {
		t.Fatal("expected error for empty memory directory")
	}
}