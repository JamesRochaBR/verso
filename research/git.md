# Git

## The problem it solves

Git solves distributed version control by allowing developers to track changes, collaborate safely, work offline, and maintain a complete history of a project's evolution.

## Primary artifact

- `.git/` directory
- Commit graph
- Git repository

The primary artifact is the repository itself, which represents the complete history and state of a project.

## What we can learn

- Distributed architecture without central dependency.
- Immutable history built from simple primitives.
- Small and composable commands.
- Everything is based on well-defined objects.
- Clear separation between repository format and user tooling.

## What we should NOT copy

- Complex command set with a steep learning curve.
- Inconsistent naming accumulated over decades.
- Internal implementation details exposed to end users.
- Features optimized for historical compatibility instead of simplicity.