---
name: decisions
type: memory
version: "1.0.0"
author: "versoteam"
tags: [decisions, ADR, architecture]
description: "Architectural decisions log"
status: approved
---

# Architectural Decisions

## Decision Log

### D001 - Use TOML for Project Configuration

**Status:** Approved  
**Date:** 2026-01-01

TOML was chosen over JSON and YAML because it is more human-readable and designed for configuration files. It supports comments and has a clean syntax.

### D002 - Markdown with Frontmatter for Components

**Status:** Approved  
**Date:** 2026-01-01

All components (skills, memory, workflows) use Markdown with optional YAML frontmatter. This ensures:
- Human-readable content
- Machine-parseable metadata
- Version control friendly
- Easy to edit in any text editor

### D003 - Semantic Versioning for Memory

**Status:** Approved  
**Date:** 2026-01-01

Memory components use semantic versioning (major.minor.patch) to track changes:
- Major: Breaking changes
- Minor: New features, backward compatible
- Patch: Bug fixes, backward compatible

### D004 - Component Type Separation

**Status:** Approved  
**Date:** 2026-01-01

Components are separated into distinct directories (skills/, memory/, workflows/, templates/) to enforce clear boundaries and prevent implicit coupling.