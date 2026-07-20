# RFC-0004

## Title

Workflow Specification

---

## Status

Draft

---

## Purpose

Define the official contract for a Verso Workflow component.

A Workflow orchestrates Skills and Memory components to execute a defined process.

While Skills provide knowledge, Workflows define how that knowledge is applied in sequence.

---

## Responsibilities

A Workflow may:

- Define a sequence of steps to execute.
- Reference Skills and Memory components by name.
- Specify conditional logic for step execution.
- Provide context for each step.

A Workflow must not:

- Store reusable knowledge (that is a Skill's responsibility).
- Make routing decisions (that is a Router's responsibility).
- Generate final output directly (that is a Template's responsibility).
- Depend implicitly on another Workflow component.

---

## Principles

### Sequential

Workflows execute steps in a defined order.

Each step builds upon the previous one, creating a logical progression.

### Declarative

Workflows declare what should happen, not how to do it.

The Router and Renderer handle execution details.

### Composable

Multiple Workflows can reference the same Skills and Memory components.

This avoids duplication and promotes reuse.

### Self-Documenting

A Workflow is a process description that anyone can read and understand.

It serves as both specification and documentation.

---

## Physical Structure

Workflow components live in the `workflows/` directory of a Verso project.

Each workflow is stored as a separate file:

```
project/
  workflows/
    feature.md
    bugfix.md
    review.md
```

### File Format

Each Workflow file uses Markdown with optional frontmatter:

```markdown
---
name: feature
version: "1.0.0"
author: "team"
tags: ["development", "feature"]
steps:
  - name: analyze
    skill: architect
    context: "Analyze requirements and design approach."
  - name: implement
    skill: developer
    memory: project
    context: "Implement following project conventions."
  - name: review
    skill: reviewer
    context: "Review implementation for quality."
---

# Implement Feature

This workflow guides the implementation of a new feature...
```

Frontmatter is optional. When absent, metadata is derived from the file name and first heading.

### Metadata Fields

| Field   | Required | Type   | Description                    |
|---------|----------|--------|--------------------------------|
| name    | No       | string | Component identifier           |
| version | No       | string | Semantic version               |
| author  | No       | string | Creator or owner               |
| tags    | No       | array  | Classification keywords        |
| steps   | Yes      | object | List of workflow steps         |

### Step Structure

Each step in a Workflow contains:

| Field   | Required | Type   | Description                    |
|---------|----------|--------|--------------------------------|
| name    | Yes      | string | Step identifier                |
| skill   | No       | string | Skill component to load        |
| memory  | No       | string | Memory component to load       |
| context | No       | string | Additional context for the step|

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

Workflows evolve as processes change.

They should be reviewed periodically to ensure relevance.

---

## Discovery

Workflow components are discovered automatically from the `workflows/` directory.

The discovery process:

1. Reads all files in the `workflows/` directory
2. Extracts metadata (from frontmatter or file name)
3. Parses step definitions
4. Loads content for rendering
5. Registers each as a Component with type `ComponentWorkflow`

---

## Usage

Workflows are loaded by the Router when a specific process is needed.

The Router:

1. Identifies which Workflow to use based on context
2. Extracts steps from the Workflow definition
3. Loads referenced Skills and Memory components
4. Assembles the final prompt using a Template

---

## Examples

### Simple Workflow

```markdown
# Code Review

This workflow guides a code review process...

Steps:
1. Load reviewer skill
2. Load project memory for conventions
3. Generate review feedback
```

### Workflow with Frontmatter and Steps

```markdown
---
name: feature
version: "1.0.0"
steps:
  - name: analyze
    skill: architect
    context: "Analyze requirements."
  - name: implement
    skill: developer
    memory: project
    context: "Implement following conventions."
  - name: review
    skill: reviewer
---

# Implement Feature

This workflow guides feature implementation...
```

---

## Philosophy

Workflows turn knowledge into action.

They transform isolated Skills and Memory into a coherent process that produces consistent results.

A well-designed Workflow makes complex processes simple and repeatable.
