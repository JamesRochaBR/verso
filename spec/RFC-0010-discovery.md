# RFC-0010 — Component Discovery

## Status

Accepted

## Authors

Verso Core Team

## Summary

This RFC defines how Verso discovers components inside a project.

Discovery must be deterministic. A project should always resolve to the same component graph regardless of execution environment.

Discovery scans predefined roots, reads component metadata from files, and builds an in-memory index of all available components.

---

## Design Principles

- Discovery is read-only: it never modifies files
- Discovery is deterministic: same files → same component list
- Discovery is extensible: new component types can be added without changing discovery logic
- Discovery validates during scan: malformed components are reported as errors, not silently ignored

---

## Discovery Roots

### Initial Discovery Roots

| Root | Component Type | Default Path | Description |
|------|---------------|--------------|-------------|
| `skills` | ComponentSkill | `skills/` | Skill component files |
| `memory` | ComponentMemory | `memory/` | Memory/knowledge files |
| `workflows` | ComponentWorkflow | `workflows/` | Workflow definition files |
| `templates` | ComponentTemplate | `templates/` | Template definition files |

Root paths are defined in `verso.toml` under `[roots]`. If not specified, defaults above are used.

### Root Registration

Roots are registered when the project is loaded:

```go
type DiscoveryConfig struct {
    Roots map[ComponentType]string  // type → path
}

func DefaultDiscoveryConfig() DiscoveryConfig {
    return DiscoveryConfig{
        Roots: map[ComponentType]string{
            ComponentSkill:    "skills/",
            ComponentMemory:   "memory/",
            ComponentWorkflow: "workflows/",
            ComponentTemplate: "templates/",
        },
    }
}
```

---

## Discovery Algorithm

### Step 1: Scan Roots

For each registered root path:

1. If the path does not exist, skip it (no error)
2. Recursively scan all `.md` files in the directory
3. For each file, attempt to parse YAML frontmatter

### Step 2: Parse Component Metadata

For each discovered file:

1. Extract YAML frontmatter (content between `---` delimiters at top of file)
2. Validate required fields (`name`, `type`)
3. Create a `Component` struct with parsed metadata and file path
4. Store the raw markdown content for later use

### Step 3: Build Component Index

After scanning all roots:

1. Group components by type
2. Index components by name (for fast lookup)
3. Detect and report duplicate names (error if same name in same type)
4. Return the final component list

### Pseudocode

```go
func Discover(projectPath string, config DiscoveryConfig) ([]Component, error) {
    var components []Component
    
    for compType, rootPath := range config.Roots {
        fullPath := filepath.Join(projectPath, rootPath)
        
        // Skip non-existent roots
        if !exists(fullPath) {
            continue
        }
        
        // Scan all .md files recursively
        files, err := findMarkdownFiles(fullPath)
        if err != nil {
            return nil, fmt.Errorf("discovery scan failed for %s: %w", rootPath, err)
        }
        
        for _, file := range files {
            comp, err := parseComponent(file, compType)
            if err != nil {
                // Report error but continue scanning
                logWarning(fmt.Sprintf("skipping %s: %v", file, err))
                continue
            }
            
            components = append(components, *comp)
        }
    }
    
    // Validate uniqueness
    if err := validateUniqueness(components); err != nil {
        return nil, err
    }
    
    return components, nil
}
```

---

## File Discovery

### Pattern

Files are discovered using the pattern: `**/*.md` (all Markdown files recursively)

### Exclusions

The following are excluded from discovery:

- Files starting with `_` (e.g., `_template.md`) — considered internal/temporary
- Files in directories starting with `.` (e.g., `.git/`, `.cache/`)
- Files without valid YAML frontmatter (logged as warning, not error)

### Symlinks

Symbolic links are followed during discovery. This allows shared component libraries.

---

## Component Parsing

### Frontmatter Requirements

Each discovered file MUST have:

```yaml
name: <string>    # Required: unique within type
type: <string>    # Required: must match expected component type
```

Optional fields:

```yaml
version: <string>   # Semantic version
author: <string>    # Author name
tags: [<string>]    # Searchable tags
description: <string>  # Human-readable description
depends: [<string>]  # Dependencies on other components
```

### Type Validation

The `type` field in frontmatter is validated against the expected type for that root. If a file is found in `skills/`, its `type` must be `"skill"`. Mismatches produce an error.

### Content Extraction

Everything after the closing `---` delimiter is treated as component content (Markdown body). Empty bodies are allowed.

---

## Error Handling

### Fatal Errors (stop discovery)

- Duplicate component names within the same type
- Invalid YAML frontmatter that cannot be parsed
- Root path is a file instead of a directory

### Warnings (continue discovery)

- File has no frontmatter
- File `type` doesn't match root type
- Missing required `name` field
- Malformed optional fields (tags, depends)

### Error Format

All errors include:

```
discovery: <file-path>: <description>
  line <N>: <detail>
```

Example:

```
discovery: skills/architect.md: duplicate component name "architect"
  line 3: name conflicts with skills/designer.md
```

---

## Component Index

After discovery, components are indexed for fast access:

```go
type ComponentIndex struct {
    byType map[ComponentType][]*Component
    byName map[string]*Component       // global unique key
    all    []*Component                 // flat list
}

func (idx *ComponentIndex) GetByType(t ComponentType) []*Component
func (idx *ComponentIndex) GetByName(name string) (*Component, bool)
func (idx *ComponentIndex) All() []*Component
```

---

## Implementation

### Go Package Structure

Discovery is implemented in `internal/project/discovery.go`:

```go
package project

type Discovery struct {
    config DiscoveryConfig
    index  *ComponentIndex
}

func NewDiscovery(config DiscoveryConfig) *Discovery
func (d *Discovery) Discover(projectPath string) ([]*Component, error)
func (d *Discovery) Index() *ComponentIndex
```

### Public API

```go
// LoadProject combines loading and discovery
func Load(path string) (*Project, error)

// DiscoverComponents runs discovery without loading full project
func DiscoverComponents(path string) ([]*Component, error)
```

---

## Examples

### Example: Project with Skills Only

```
my-project/
  verso.toml
  skills/
    architect.md
    reviewer.md
```

Discovery result:

```go
[]Component{
    {Name: "architect", Type: "skill", Path: "skills/architect.md"},
    {Name: "reviewer", Type: "skill", Path: "skills/reviewer.md"},
}
```

### Example: Empty Project

```
my-project/
  verso.toml
```

Discovery result: empty list (no error, no warnings)

### Example: Missing Root Directory

```
my-project/
  verso.toml
  skills/       ← exists
  memory/       ← does NOT exist (skipped silently)
```

Discovery result: only components from `skills/` are discovered

---

## Extensibility

### Adding New Component Types

To add a new component type:

1. Add the type constant in `component.go`
2. Add the root mapping in `DefaultDiscoveryConfig()`
3. Update this RFC if the discovery semantics change

No changes to the core discovery algorithm are needed.

---

## Open Questions

1. Should discovery support external sources (e.g., git repos, registries)?
2. Should we cache discovery results for performance on large projects?
3. Should discovery report component dependency graphs?

---

*This RFC is accepted. Implementation exists in `internal/project/discovery.go` and is used by the Router.*