# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- RFC-0007: Agent Specification (AgentAdapter interface, supported agents)
- RFC-0009: Serialization Specification (TOML + Markdown frontmatter)
- RFC-0010: Component Discovery (3-step algorithm)
- Router implementation (`internal/router/`)
  - Keyword-Based Routing Strategy
  - Workflow-Based Routing Strategy
  - Default Routing Strategy
  - Explicit Filtering integration
- CLI flags: `--keywords/-k`, `--workflow/-w`, `--strategy/-s`
- ComponentTypes: `router`, `agent` added to `internal/project/component.go`
- 23 unit tests for Router (all passing)
- IMPLEMENTATION_PLAN.md with phase tracking
- AI_ENGINEERING_GUIDE.md with commit/tag workflow for open-source (MIT)

## [v0.1.0] — 2026-07-19

### Added

- Initial project structure
- CLI framework (`internal/cli/`)
- Project model (`internal/project/`)
- Render engine (`internal/render/`)
- Discovery implementation
- Filter implementation
- Markdown parsing with frontmatter
- Validation system
- Hello-world example project
- Documentation: ARCHITECTURE.md, VISION.md, CONTRIBUTING.md, README.md
- MIT License

---

## Versioning

- **MAJOR**: Incompatible API changes (never during development)
- **MINOR**: New backward-compatible functionality (per phase completion)
- **PATCH**: Backward-compatible bug fixes

## Tags

| Tag | Meaning | Status |
|-----|---------|--------|
| `v0.1.0` | Foundation — Initial setup | ✅ Published |
| `v0.2.0` | Specifications — RFCs, Serialization, Discovery, Router | 🔄 Pending release |
| `v0.3.0` | Skills — Reusable knowledge unit | Planned |
| `v0.4.0` | Memory — Project-specific knowledge | Planned |
| `v1.0.0` | Stable — Production-ready | Planned |