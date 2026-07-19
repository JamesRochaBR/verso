# Use Cases

This document describes how developers are expected to interact with Verso.

These use cases are implementation-independent.

They exist to validate the architecture before code is written.

---

## UC-001

A developer installs Verso in an existing project.

Expected outcome:

Verso discovers the project.

Verso identifies the available Components.

Verso prepares the context for the selected AI coding agent.

---

## UC-002

A developer asks an AI agent to implement a new feature.

Expected outcome:

Only the relevant Components are loaded.

The agent receives focused context.

No unnecessary knowledge is included.

---

## UC-003

A team shares reusable Components across multiple projects.

Expected outcome:

Knowledge becomes portable.

Projects remain independent.

Components evolve without duplication.