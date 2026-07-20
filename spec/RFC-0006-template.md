# RFC-0006

## Title

Template Specification

---

## Status

Draft

---

## Purpose

Define the official contract for a Verso Template component.

Templates define how context is assembled and presented to the end user (LLM or coding agent).

While Skills, Memory and Workflows provide knowledge, Templates define the structure of the final output.

---

## Responsibilities

A Template may:

- Define the structure and format of generated prompts.
- Include conditional sections based on available components.
- Use variables to inject dynamic content.
- Support multiple output formats (Markdown, plain text, JSON).

A Template must not:

- Store reusable knowledge (that is a Skill's responsibility).
- Make routing decisions (that is a Router's responsibility).
- Orchestrate processes (that is a Workflow's responsibility).
- Depend implicitly on another Template component.

---

## Principles

### Structured

Templates define clear structure for generated output.

They separate content from presentation, allowing the same knowledge to be rendered in different formats.

### Variable-Based

Templates use variables to inject dynamic content from components.

Variables are resolved at render time using data from the Project model.

### Conditional

Templates can include conditional sections that only appear when certain conditions are met.

This allows flexible output based on available components.

### Extensible

New template formats and engines can be added without breaking existing templates.

The Template system supports pluggable rendering engines.

---

## Physical Structure

Template components live in the `templates/` directory of a Verso project.

Each template is stored as a separate file:

```
project/
  templates/
    feature.md
    bugfix.md
    review.md
```

### File Format

Templates use Go text/template syntax with optional frontmatter:

```markdown
---
name: feature
version: "1.0.0"
author: "team"
tags: ["development", "feature"]
format: markdown
---

# Project Context

Name: {{ .Project.Metadata.Name }}
Version: {{ .Project.Metadata.Version }}

{{ if .Skills }}
## Skills

{{ range .Skills }}
### {{ .Title }}

{{ .Content }}
{{ end }}
{{ end }}

{{ if .Memory }}
## Memory

{{ range .Memory }}
### {{ .Title }}

{{ .Content }}
{{ end }}
{{ end }}
```

Frontmatter is optional. When absent, metadata is derived from the file name and first heading.

### Metadata Fields

| Field   | Required | Type   | Description                    |
|---------|----------|--------|--------------------------------|
| name    | No       | string | Component identifier           |
| version | No       | string | Semantic version               |
| author  | No       | string | Creator or owner               |
| tags    | No       | array  | Classification keywords        |
| format  | Yes      | string | Output format (markdown, text) |

---

## Template Variables

Templates have access to the following variables through the Go template engine:

### Project Metadata

```go
{{ .Project.Metadata.Name }}
{{ .Project.Metadata.Version }}
{{ .Project.Metadata.Description }}
```

### Components

```go
{{ range .Components }}
  {{ .Name }}
  {{ .Title }}
  {{ .Type }}
  {{ .Content }}
{{ end }}
```

### Filtered Collections

```go
{{ range .Skills }}    // All skill components
{{ range .Memory }}   // All memory components
{{ range .Workflows }} // All workflow components
{{ range .Templates }} // All template components
```

---

## Template Functions

The Go template engine provides built-in functions:

| Function | Description | Example |
|----------|-------------|---------|
| `len` | Get length of slice/string | `{{ len .Skills }}` |
| `printf` | Format string | `{{ printf "%s v%s" .Name .Version }}` |
| `eq` | Equality check | `{{ if eq .Type "skill" }}` |
| `ne` | Not equal check | `{{ if ne .Type "template" }}` |
| `and` | Logical AND | `{{ if and .Skills .Memory }}` |
| `or` | Logical OR | `{{ if or .Skills .Workflows }}` |
| `not` | Logical NOT | `{{ if not .Templates }}` |

---

## Discovery

Template components are discovered automatically from the `templates/` directory.

The discovery process:

1. Reads all files in the `templates/` directory
2. Extracts metadata (from frontmatter or file name)
3. Loads template content for rendering
4. Registers each as a Component with type `ComponentTemplate`

---

## Usage

Templates are used by the Renderer to generate final output:

### CLI Usage

```bash
# Use default template (markdown renderer)
verso prompt ./project

# Use specific template
verso prompt ./project --template feature

# Generate in different format
verso prompt ./project --format text
```

### Programmatic Usage

```go
p, _ := project.Load("./my-project")

template, _ := template.Get("feature")
output, _ := template.Render(p)
fmt.Print(output)
```

---

## Built-in Templates

The Verso framework includes built-in templates:

### Default Markdown Template

Generates a structured Markdown prompt with all sections:

```markdown
# Project Context

Name: {{ .Project.Metadata.Name }}
Version: {{ .Project.Metadata.Version }}

## Skills
...

## Memory
...

## Workflows
...

## Templates
...
```

### Plain Text Template

Generates a plain text version without Markdown formatting.

---

## Examples

### Simple Template

```markdown
# Project: {{ .Project.Metadata.Name }}

{{ range .Skills }}
## {{ .Title }}

{{ .Content }}
{{ end }}
```

### Conditional Template

```markdown
# Feature Implementation Guide

{{ if .Memory }}
## Project Context

{{ range .Memory }}
### {{ .Title }}

{{ .Content }}
{{ end }}
{{ end }}

## Skills

{{ range .Skills }}
### {{ .Title }}

{{ .Content }}
{{ end }}
```

---

## Future Enhancements

### JSON Output

Support for JSON-formatted output for programmatic consumption.

### Custom Functions

Allow projects to define custom template functions.

### Template Inheritance

Support for template inheritance and composition.

---

## Philosophy

Templates are the final layer of Verso.

They take all the organized knowledge and present it in a clear, consistent format that any LLM can understand and act upon.

A good template makes complex context simple to consume.
