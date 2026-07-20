---
name: conventions
type: memory
version: "1.0.0"
author: "versoteam"
tags: [conventions, coding-standards, project]
description: "Project coding conventions and standards"
status: approved
---

# Project Conventions

## Coding Standards

- All Go code must pass `gofmt` before commit
- Use Go modules for dependency management
- Follow effective-go guidelines
- Keep functions small and focused
- Prefer explicit over implicit

## File Organization

- Place new skills in `skills/` directory
- Store project memory in `memory/` directory
- Define workflows in `workflows/` directory
- Create templates in `templates/` directory

## Commit Conventions

- Use conventional commits format
- Prefix with type: feat, fix, docs, chore, test, refactor
- Keep subject line under 72 characters
- Reference issue numbers when applicable

## Review Process

- All changes require at least one review
- Tests must pass before merging
- Documentation updates required for API changes