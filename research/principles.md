# Architectural Principles

This document consolidates the lessons learned from the architectural research.

## Principle 1

Separate specification from implementation.

The project model must be independent from any runtime.

---

## Principle 2

Keep the core model immutable whenever possible.

Changes should happen through evolution, not mutation.

---

## Principle 3

Build small composable primitives.

Complexity should emerge from composition.

---

## Principle 4

The specification is the product.

Implementations are consumers of the specification.

---

## Principle 5

A project should be deterministic.

The same project must always resolve to the same context graph.

---

## Principle 6

Extensibility must never compromise compatibility.