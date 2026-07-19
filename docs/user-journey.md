# User Journey

## Overview

Verso is a specification and toolchain for building AI context projects.

A project is composed of structured components such as skills, memories,
templates and workflows.

The goal is to transform those components into a deterministic Context Graph
that can be consumed by AI agents.

---

# 1. Create a project

```bash
verso init my-assistant
```

Creates:

```
my-assistant/
├── verso.toml
├── skills/
├── memory/
├── workflows/
└── templates/
```

---

# 2. Add components

Example:

```
skills/reviewer.md
```

---

# 3. Validate

```bash
verso validate
```

Checks:

- Project structure
- Invalid references
- Missing components
- Cyclic dependencies

---

# 4. Build

```bash
verso build
```

Produces:

```
.context/

manifest.json

graph.json
```

---

# 5. Consume

An AI Agent loads the generated Context Graph instead of scanning the entire project.

Examples:

- Cursor
- GitHub Copilot
- Continue
- Cline
- Claude Code

---

# Philosophy

The Context Graph is the product.

The filesystem is only the source.