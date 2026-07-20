# RFC-0003

## Title

Memory Specification

---

## Status

Draft

---

## Purpose

Define the official contract for a Verso Memory component.

Memory stores project-specific knowledge that must be remembered across interactions.

Unlike Skills, which are reusable and general-purpose, Memory is tied to a single project.

---

## Responsibilities

A Memory may:

- Store architectural decisions.
- Record business rules.
- Maintain code conventions.
- Track technical debt.
- Preserve historical context.

A Memory must not:

- Define reusable knowledge (that is a Skill's responsibility).
- Orchestrate processes (that is a Workflow's responsibility).
- Make routing decisions (that is a Router's responsibility).
- Depend implicitly on another Memory component.

---

## Principles

### Project-Specific

Memory belongs to the project.

It should not be reused across different projects.

Each project has its own unique context that must be preserved independently.

### Persistent

Memory represents knowledge that persists over time.

It captures decisions, patterns and constraints that outlive individual interactions.

### Structured

Memory components follow a consistent structure:

- Title (first line of the file)
- Description (optional second section)
- Content (structured knowledge)

### Composable

Multiple Memory components can coexist within a project.

They are loaded together to form a complete picture of project context.

---

## Physical Structure

Memory components live in the `memory/` directory of a Verso project.

Each memory is stored as a separate file:

```
project/
  memory/
    architecture.md
    conventions.md
    decisions.md
```

### File Format

Each Memory file uses Markdown with optional frontmatter:

```markdown
---
name: architecture
version: "1.0.0"
author: "team"
tags: ["architecture", "design"]
---

# Architecture

Content of the memory component...
```

Frontmatter is optional. When absent, metadata is derived from the file name and first heading.

### Metadata Fields

| Field     | Required | Type   | Description                    |
|-----------|----------|--------|--------------------------------|
| name      | No       | string | Component identifier           |
| version   | No       | string | Semantic version               |
| author    | No       | string | Creator or owner               |
| tags      | No       | array  | Classification keywords        |
| description | No     | string | Short description              |

When frontmatter is absent:
- `name` is derived from the file name (without extension)
- `title` is extracted from the first `# Heading` in the content
- Other fields default to empty values

---

## Lifecycle

```
Created
  ↓
Reviewed
  ↓
Approved
  ↓
Active
  ↓
Deprecated
```

Memory components evolve as the project evolves.

They should be reviewed periodically to ensure accuracy.

---

## Discovery

Memory components are discovered automatically from the `memory/` directory.

The discovery process:

1. Reads all files in the `memory/` directory
2. Extracts metadata (from frontmatter or file name)
3. Loads content for rendering
4. Registers each as a Component with type `ComponentMemory`

---

## Usage

Memory is loaded by the Router when context about the project is needed.

It can be filtered by:

- Name (`--name architecture,conventions`)
- Type exclusion (`--exclude memory`)
- Tags (future enhancement)

---

## Examples

### Simple Memory

```markdown
# Project Conventions

We use Go modules for dependency management.
All code must pass `gofmt` before commit.
```

### Memory with Frontmatter

```markdown
---
name: architecture
version: "2.0.0"
author: "team-lead"
tags: ["architecture", "backend"]
description: "System architecture overview"
---

# Architecture

The system follows a layered architecture...
```

---

## Philosophy

Memory is the project's voice.

It speaks for the project, preserving what matters and discarding what does not.

A well-maintained Memory makes every interaction smarter.
