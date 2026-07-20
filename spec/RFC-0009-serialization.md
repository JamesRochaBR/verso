# RFC-0009 — Serialization

## Status

Accepted

## Authors

Verso Core Team

## Summary

This RFC defines how a Verso Project Model is serialized to disk and configuration files.

The reference serialization format is TOML for project configuration, and Markdown with YAML frontmatter for component files.

Alternative serialization formats may exist as long as they represent the same logical Project Model.

## Goals

- Human readable
- Git friendly
- Deterministic
- Extensible
- Backward compatible
- IDE-friendly (easy to edit without tools)

## Non-Goals

This RFC does not define the Project Model itself. That responsibility belongs to RFC-0008.

This RFC does not mandate a single format — it defines the contract that all formats must satisfy.

---

## Serialization Formats

### 1. Project Configuration: TOML

The project root configuration file uses TOML format.

**File name:** `verso.toml` (required)

**Purpose:** Define project metadata, component roots, and global settings.

#### Schema

```toml
# Required: Project identity
name = "my-project"
version = "0.1.0"
description = "A sample Verso project"

# Optional: Component discovery roots
[roots]
skills = "skills/"
memory = "memory/"
workflows = "workflows/"
templates = "templates/"

# Optional: Agent configuration
[agent]
default = "cursor"

[agent.adapters]
enabled = ["github-copilot", "cursor"]

# Optional: Router settings
[router]
default_strategy = "keyword"

# Optional: Metadata
author = "your-name"
license = "MIT"
```

#### Validation Rules

- `name` is required and must be a valid identifier (lowercase, hyphens allowed)
- `version` must follow semantic versioning if present
- `roots` paths are relative to the project root
- Duplicate roots across formats cause an error during discovery

### 2. Component Files: Markdown with YAML Frontmatter

All component types (skill, memory, workflow, template) use Markdown files with YAML frontmatter.

**File naming convention:** `{component-name}.md`

#### Frontmatter Schema

Every component file MUST include a YAML frontmatter block at the top:

```markdown
---
name: my-component
type: skill
version: "1.0.0"
author: "author-name"
tags: [tag1, tag2]
description: "Human-readable description of this component"
---

# Component Title

Markdown content goes here...
```

#### Required Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `name` | string | Yes | Unique identifier within the project |
| `type` | string | Yes | Must match one of: skill, memory, workflow, template |
| `description` | string | No | Human-readable description |

#### Optional Fields

| Field | Type | Description |
|-------|------|-------------|
| `version` | string | Semantic version of the component |
| `author` | string | Component author |
| `tags` | array | Searchable tags for filtering |
| `depends` | array | Names of other components this depends on |

#### Content Format

The Markdown body contains the component's actual content, which varies by type:

- **Skill**: Instructions, prompts, or procedures the agent should follow
- **Memory**: Project knowledge, decisions, and context
- **Workflow**: YAML `steps` array in frontmatter + optional markdown instructions
- **Template**: Template structure with variables for context rendering

### 3. Workflow Steps Format

Workflows use a special `steps` field in frontmatter:

```markdown
---
name: code-review
type: workflow
version: "1.0.0"
steps:
  - component: skill/linter
    action: run
  - component: skill/reviewer
    action: run
    depends-on: linter
  - component: memory/project
    action: read
---

# Code Review Workflow

Review the code and provide feedback.
```

#### Step Schema

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `component` | string | Yes | Reference to a skill, memory, or workflow component (format: `{type}/{name}`) |
| `action` | string | Yes | Action to perform: `run`, `read`, `update` |
| `depends-on` | string | No | Name of a previous step this depends on |

---

## Serialization Contract

### Determinism

- Same Project Model → same serialized output (byte-for-byte deterministic)
- Field ordering in TOML must be consistent (use sorted keys)
- Frontmatter fields must appear in the order defined by this RFC

### Round-Trip Safety

- Parsing a serialized file and re-serializing it must produce equivalent output
- No data loss during parse → serialize → parse cycles

### Error Handling

Serialization errors MUST include:
- The file path that caused the error
- A human-readable description of what went wrong
- Line number if applicable

---

## Implementation

### Go Package Structure

```
internal/serialization/
  serializer.go      — Interface and base types
  toml.go            — TOML serialization/deserialization
  markdown.go        — Markdown frontmatter parsing
  validator.go       — Schema validation
```

### Interface

```go
type Serializer interface {
    Serialize(v any) ([]byte, error)
    Deserialize(data []byte, v any) error
}
```

### Registration

Serializers are registered by file extension:

```go
registry.Register(".toml", TOMLSerializer{})
registry.Register(".md", MarkdownSerializer{})
```

---

## Versioning

- Changes to the serialization format require a version bump in `verso.toml`
- Backward-compatible extensions (new optional fields) do not require version changes
- Breaking changes MUST include a migration path in release notes

---

## Examples

### Example: Skill File

```markdown
---
name: architect
type: skill
version: "1.0.0"
author: "versoteam"
tags: [architecture, design]
description: "Architectural review and system design guidance"
---

# Architect Skill

You are an expert software architect. Analyze the provided code structure and suggest improvements...
```

### Example: Project Configuration

```toml
name = "verso"
version = "0.3.0"
description = "AI agent framework for software development"

[roots]
skills = "skills/"
memory = "memory/"
workflows = "workflows/"
templates = "templates/"

[router]
default_strategy = "keyword"
```

---

## Open Questions

1. Should we support JSON as an alternative configuration format?
2. Should component files support multi-section frontmatter (separate sections for different concerns)?
3. Is binary serialization needed for performance-critical scenarios?

---

*This RFC is accepted. Implementation began in v0.2 and continues through v0.4.*