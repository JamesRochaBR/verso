package project

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// MemoryVersion represents a versioned memory component snapshot.
type MemoryVersion struct {
	Version   string
	Timestamp string
	Content   string
	Path      string
}

// SemVer represents a semantic version (major.minor.patch).
type SemVer struct {
	Major      int
	Minor      int
	Patch      int
	Prerelease string
}

// ParseSemanticVersion parses a semantic version string.
// Accepts formats like "1.0.0", "2.1.3", "0.5.0-pre".
func ParseSemanticVersion(version string) (*SemVer, error) {
	version = strings.TrimSpace(version)
	if version == "" {
		return nil, fmt.Errorf("empty version string")
	}

	// Remove leading 'v' if present
	version = strings.TrimPrefix(version, "v")

	// Regex for semantic versioning with optional prerelease and build metadata
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)(?:-([a-zA-Z0-9.+.-]+))?(?:\+([a-zA-Z0-9.+.-]+))?$`)
	matches := re.FindStringSubmatch(version)
	if matches == nil {
		return nil, fmt.Errorf("invalid semantic version format: %q", version)
	}

	major, _ := strconv.Atoi(matches[1])
	minor, _ := strconv.Atoi(matches[2])
	patch, _ := strconv.Atoi(matches[3])

	return &SemVer{
		Major:      major,
		Minor:      minor,
		Patch:      patch,
		Prerelease: matches[4],
	}, nil
}

// String returns the string representation of a SemVer.
func (s *SemVer) String() string {
	version := fmt.Sprintf("%d.%d.%d", s.Major, s.Minor, s.Patch)
	if s.Prerelease != "" {
		version = fmt.Sprintf("%s-%s", version, s.Prerelease)
	}
	return version
}

// BumpMajor returns a new SemVer with the major version incremented.
func (s *SemVer) BumpMajor() *SemVer {
	return &SemVer{
		Major:      s.Major + 1,
		Minor:      0,
		Patch:      0,
		Prerelease: "",
	}
}

// BumpMinor returns a new SemVer with the minor version incremented.
func (s *SemVer) BumpMinor() *SemVer {
	return &SemVer{
		Major:      s.Major,
		Minor:      s.Minor + 1,
		Patch:      0,
		Prerelease: "",
	}
}

// BumpPatch returns a new SemVer with the patch version incremented.
func (s *SemVer) BumpPatch() *SemVer {
	return &SemVer{
		Major:      s.Major,
		Minor:      s.Minor,
		Patch:      s.Patch + 1,
		Prerelease: "",
	}
}

// CompareVersions compares two semantic version strings.
// Returns: -1 if a < b, 0 if a == b, 1 if a > b.
func CompareVersions(a, b string) (int, error) {
	semA, err := ParseSemanticVersion(a)
	if err != nil {
		return 0, fmt.Errorf("invalid version %q: %w", a, err)
	}

	semB, err := ParseSemanticVersion(b)
	if err != nil {
		return 0, fmt.Errorf("invalid version %q: %w", b, err)
	}

	if semA.Major != semB.Major {
		if semA.Major > semB.Major {
			return 1, nil
		}
		return -1, nil
	}

	if semA.Minor != semB.Minor {
		if semA.Minor > semB.Minor {
			return 1, nil
		}
		return -1, nil
	}

	if semA.Patch != semB.Patch {
		if semA.Patch > semB.Patch {
			return 1, nil
		}
		return -1, nil
	}

	// Handle prerelease versions (no prerelease > with prerelease)
	if semA.Prerelease == "" && semB.Prerelease != "" {
		return 1, nil
	}
	if semA.Prerelease != "" && semB.Prerelease == "" {
		return -1, nil
	}
	if semA.Prerelease == "" && semB.Prerelease == "" {
		return 0, nil
	}

	// Both have prerelease — lexicographic comparison
	if semA.Prerelease < semB.Prerelease {
		return -1, nil
	}
	if semA.Prerelease > semB.Prerelease {
		return 1, nil
	}
	return 0, nil
}

// IsVersioned returns true if the component has a valid semantic version.
func IsVersioned(component Component) bool {
	if component.Version == "" {
		return false
	}
	_, err := ParseSemanticVersion(component.Version)
	return err == nil
}

// GetLatestVersion scans the memory directory for all versions of a component
// and returns the latest one based on semantic versioning.
func GetLatestVersion(component Component, projectPath string) (*MemoryVersion, error) {
	memoryDir := filepath.Join(projectPath, "memory")

	entries, err := os.ReadDir(memoryDir)
	if err != nil {
		return nil, fmt.Errorf("cannot read memory directory: %w", err)
	}

	prefix := component.Name + "."
	var versions []MemoryVersion

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasPrefix(name, prefix) || !strings.HasSuffix(name, ".md") {
			continue
		}

		filePath := filepath.Join(memoryDir, name)
		contentBytes, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		content := string(contentBytes)
		meta, _ := ParseFrontmatter(content)

		mv := MemoryVersion{
			Path:   filePath,
			Content: content,
		}

		if meta != nil && meta.Version != "" {
			mv.Version = meta.Version
		} else {
			// Default version for unversioned files
			mv.Version = "0.0.0"
		}

		versions = append(versions, mv)
	}

	if len(versions) == 0 {
		return nil, fmt.Errorf("no versions found for memory %q", component.Name)
	}

	// Sort by semantic version (latest first)
	sort.SliceStable(versions, func(i, j int) bool {
		cmp, _ := CompareVersions(versions[i].Version, versions[j].Version)
		return cmp > 0
	})

	return &versions[0], nil
}

// ListMemoryVersions returns all versioned memory files for a component.
func ListMemoryVersions(component Component, projectPath string) ([]MemoryVersion, error) {
	memoryDir := filepath.Join(projectPath, "memory")

	entries, err := os.ReadDir(memoryDir)
	if err != nil {
		return nil, fmt.Errorf("cannot read memory directory: %w", err)
	}

	prefix := component.Name + "."
	var versions []MemoryVersion

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasPrefix(name, prefix) || !strings.HasSuffix(name, ".md") {
			continue
		}

		filePath := filepath.Join(memoryDir, name)
		contentBytes, err := os.ReadFile(filePath)
		if err != nil {
			continue
		}

		content := string(contentBytes)
		meta, _ := ParseFrontmatter(content)

		mv := MemoryVersion{
			Path:    filePath,
			Content: content,
		}

		if meta != nil && meta.Version != "" {
			mv.Version = meta.Version
		} else {
			mv.Version = "0.0.0"
		}

		versions = append(versions, mv)
	}

	return versions, nil
}