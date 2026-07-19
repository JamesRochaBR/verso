# RFC-0001

## Title

Skill Specification

---

## Status

Draft

---

## Purpose

Define the official contract for a Verso Skill.

A Skill is the smallest reusable unit of knowledge inside the framework.

Skills are designed to perform exactly one responsibility and must remain independent from one another.

---

## Responsibilities

A Skill may:

- Provide domain knowledge.
- Describe a process.
- Explain a concept.
- Execute a reusable procedure.
- Define best practices.

A Skill must not:

- Store persistent information.
- Decide when it should be loaded.
- Execute workflows.
- Depend implicitly on another Skill.

---

## Principles

### Single Responsibility

Each Skill must solve one problem.

---

### Independent

A Skill must be understandable without requiring another Skill.

---

### Reusable

Skills are designed to be reused by multiple workflows.

---

### Deterministic

Given the same input and context, a Skill should always produce consistent guidance.

---

## Lifecycle

Created

↓

Reviewed

↓

Approved

↓

Published

↓

Deprecated

---

## Metadata

Every Skill should define:

- Name
- Description
- Version
- Author
- Tags

---

## Future Structure

The physical representation of a Skill is intentionally left undefined by this RFC.

This specification defines the contract only.

Directory structure and file format will be specified by future RFCs.

---

## Philosophy

A Skill defines knowledge.

It never defines orchestration.