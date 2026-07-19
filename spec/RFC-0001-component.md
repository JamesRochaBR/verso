# RFC-0001

## Title

Component Model

---

## Status

Draft

---

## Purpose

Define the fundamental architectural unit of the Verso framework.

Every element inside Verso is considered a Component.

Components expose a well-defined responsibility and collaborate through explicit contracts.

---

## Component Types

The initial Component types are:

- Skill
- Memory
- Workflow
- Template
- Router
- Agent

New Component types may be introduced by future RFCs.

---

## Properties

Every Component:

- has one responsibility;
- has a specification;
- may evolve independently;
- is versionable;
- is discoverable.

---

## Rules

A Component must never rely on undocumented behavior.

A Component must never assume the existence of another Component.

Interactions must always occur through published contracts.

---

## Philosophy

Everything in Verso is a Component.

Specifications define Components.

Implementations realize Components.