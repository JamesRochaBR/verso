# RFC-0005

## Title

Router Specification

---

## Status

Draft

---

## Purpose

Define the official contract for a Verso Router component.

The Router decides which components (Skills, Memory, Workflows) should be loaded based on context and user intent.

It is the intelligence layer that determines what knowledge is relevant for a given situation.

---

## Responsibilities

A Router may:

- Analyze user input or context to determine intent.
- Select appropriate Skills based on keywords and patterns.
- Load relevant Memory components for project-specific guidance.
- Choose Workflows that match the described process.
- Apply filtering rules to include or exclude components.

A Router must not:

- Generate final output (that is a Template's responsibility).
- Store reusable knowledge (that is a Skill's responsibility).
- Orchestrate execution flow (that is a Workflow's responsibility).
- Make decisions based on undocumented behavior.

---

## Principles

### Context-Aware

The Router analyzes the context provided by the user to make intelligent decisions about what components are relevant.

It uses heuristics, keywords and patterns to determine intent.

### Deterministic

Given the same input and project configuration, the Router should always produce the same component selection.

This ensures predictable and reproducible results.

### Configurable

The Router respects explicit user filters (`--name`, `--exclude`) while providing sensible defaults when no filters are specified.

Users can override automatic decisions at any time.

### Extensible

New routing strategies can be added without breaking existing behavior.

The Router architecture supports pluggable routing logic.

---

## Routing Strategies

The Router supports multiple routing strategies:

### 1. Keyword-Based Routing

Analyzes user input for keywords and matches them against component metadata (tags, names, descriptions).

Example:
- Input: "review this code"
- Matched Skill: `reviewer` (contains keyword "review")

### 2. Workflow-Based Routing

When a specific workflow is requested, loads all referenced components automatically.

Example:
- Request: "implement feature using feature workflow"
- Loads: architect skill + developer skill + reviewer skill + project memory

### 3. Explicit Filtering

Users can explicitly specify which components to include or exclude.

Example:
- `--name reviewer,architect` — only load these skills
- `--exclude memory` — skip all memory components

### 4. Default Routing

When no filters are specified and no workflow is requested, loads all components by default.

This ensures maximum context is available when the user has not provided specific guidance.

---

## Physical Structure

The Router is implemented as a core library component, not as a file-based component.

It lives in the `internal/router/` package:

```
project/
  internal/
    router/
      router.go
      strategy.go
      keyword.go
```

### Public API

```go
type Router interface {
    Route(project *Project, options RoutingOptions) (*Project, error)
}

type RoutingOptions struct {
    Filter   Filter
    Strategy string
    Keywords []string
}
```

---

## Discovery

The Router discovers available components through the Project model.

It does not perform file system operations directly — it works with already-discovered component data.

---

## Usage

The Router is invoked by CLI commands that need to select specific components:

### CLI Usage

```bash
# Default routing (all components)
verso prompt ./project

# Keyword-based routing
verso prompt ./project --keywords "review,architecture"

# Explicit filtering
verso prompt ./project --name reviewer,architect
verso prompt ./project --exclude memory

# Workflow-based routing
verso prompt ./project --workflow feature
```

### Programmatic Usage

```go
p, _ := project.Load("./my-project")

router := router.New()
filtered, _ := router.Route(p, router.RoutingOptions{
    Keywords: []string{"review", "architecture"},
})

prompt, _ := render.Prompt(filtered)
fmt.Print(prompt)
```

---

## Routing Decision Flow

```
User Input
    ↓
Parse Options (--name, --exclude, --keywords, --workflow)
    ↓
Options Specified?
    ├─ Yes → Apply Explicit Filters
    └─ No  → Default Strategy
                ↓
        Workflow Requested?
            ├─ Yes → Load Workflow Components
            └─ No  → Load All Components
```

---

## Examples

### Example 1: Keyword-Based Routing

Input: "I need help reviewing this code"

Router Analysis:
- Keywords detected: "reviewing", "code"
- Matched Skills: `reviewer` (tags include "review")
- Matched Memory: `conventions` (tags include "code-style")

Result: Only reviewer skill and conventions memory are loaded.

### Example 2: Workflow-Based Routing

Input: "implement a new feature using the feature workflow"

Router Analysis:
- Workflow detected: "feature"
- Workflow steps reference: architect, developer, reviewer skills + project memory

Result: All referenced components are loaded in sequence.

### Example 3: Explicit Filtering

Command: `verso prompt ./project --name reviewer --exclude templates`

Router Analysis:
- Filter applied: only component named "reviewer"
- Exclusion applied: skip all template components

Result: Only the reviewer skill is included, no templates.

---

## Future Enhancements

### Machine Learning Routing

Future versions may use ML models to improve routing accuracy based on historical usage patterns.

### Dependency-Aware Routing

The Router could analyze component dependencies and automatically load required dependencies when a component is selected.

### Priority-Based Routing

Components could have priority levels, allowing the Router to select the most relevant components when context limits exist.

---

## Philosophy

The Router is the brain of Verso.

It transforms user intent into actionable knowledge by selecting exactly what is needed and nothing more.

A good router loads less so the agent can reason better.
