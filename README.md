# Verso

> Less Context. Better Reasoning.

Verso is an open-source framework that helps AI coding agents load only the context they actually need.

Instead of relying on monolithic prompts, Verso organizes knowledge into modular skills, workflows, memory and intelligent context routing.

The goal is simple:

> Load less. Reason better.

---

## Why Verso?

Modern AI coding agents are incredibly capable.

However, they all suffer from the same problem:

- Too much context.
- Too many instructions.
- Too many responsibilities.

Large prompts become difficult to maintain, expensive to process and easy to break.

Verso solves this problem by decomposing knowledge into small reusable modules that can be loaded only when necessary.

---

## Core Principles

Verso is built around a few fundamental ideas.

- Context is expensive.
- Load only what is necessary.
- Small modules are better than monolithic prompts.
- Every Skill has a single responsibility.
- Workflows orchestrate Skills.
- Memory stores knowledge.
- Templates generate artifacts.
- Routing decides what should be loaded.

---

## Architecture

```
User
    │
    ▼
Coding Agent
    │
    ▼
Verso Router
    │
    ├── Skills
    ├── Memory
    ├── Templates
    └── Workflows
    │
    ▼
Response
```

---

## Compatibility

Verso is designed to work with any modern coding agent.

Current targets include:

- GitHub Copilot Agent
- Cursor
- Continue
- Cline
- Roo Code
- Claude Code

The framework itself remains agent-agnostic.

---

## Dogfooding

Verso follows one simple rule:

Every new version of Verso must be built using the capabilities provided by the previous version.

The framework evolves by using itself.

---

## Roadmap

Current milestone:

**v0.1 — Foundation**

Next milestone:

**v0.2 — Specifications**

---

## License

MIT License.