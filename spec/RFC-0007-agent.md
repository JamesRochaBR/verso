# RFC-0007 — Agent Specification

## Status

Draft

## Authors

Verso Core Team

## Summary

This RFC defines the Agent component type in the Verso framework.

An Agent is a compatible AI coding agent that consumes context produced by Verso and executes tasks using the provided skills, memory, and workflows.

Agents are not part of the core Verso runtime. They are external tools that Verso integrates with through adapters.

## Design Principles

- Agents consume, not produce, Verso context
- Verso does not dictate which agents to use
- Every supported agent must have a dedicated adapter
- Agent compatibility is validated through examples, never enforced by the framework

## Agent Architecture

### Agent Interface

An agent adapter must implement the following interface:

```go
type AgentAdapter interface {
    // Name returns the adapter name (e.g., "github-copilot", "cursor")
    Name() string
    
    // FormatContext formats a Verso Project Model into the native format expected by the agent
    FormatContext(project *project.Project) ([]byte, error)
    
    // Validate checks if the agent environment is properly configured
    Validate() error
}
```

### Agent Adapter Contract

Each adapter must:

1. Accept a `*project.Project` as input
2. Transform the project context into the agent's expected format
3. Return formatted output (bytes) that can be written to a file or piped to the agent

## Supported Agents

### v0.8 Initial Agent List

The following agents are targeted for initial support:

| Agent | Adapter Name | Output Format |
|-------|-------------|---------------|
| GitHub Copilot | `github-copilot` | Markdown prompt file |
| Cursor | `cursor` | JSON context bundle |
| Claude Code | `claude-code` | Text prompt with structured metadata |
| Cline/Roo Code | `cline` | Markdown task file |

### Agent Format Specifications

#### GitHub Copilot

- Output: `.github/prompts/verso-context.md`
- Format: Markdown with frontmatter
- Content: Full project context formatted for Copilot's custom instructions

#### Cursor

- Output: `.cursor/verso-context.json`
- Format: JSON bundle containing component metadata and filtered content
- Content: Structured context optimized for Cursor's agent mode

#### Claude Code

- Output: stdout or file specified by `--output`
- Format: Plain text with YAML frontmatter
- Content: Clean prompt ready to pipe to claude command

#### Cline/Roo Code

- Output: `.cline/verso-task.md`
- Format: Markdown task file with structured sections
- Content: Task-oriented context for Cline/Roo workflow

## Agent Configuration

Agents are configured in `verso.toml`:

```toml
[agent]
default = "cursor"

[agent.adapters]
enabled = ["github-copilot", "cursor", "claude-code"]
```

## Agent Selection

The CLI supports agent selection via flags:

```bash
verso prompt . --agent cursor
verso prompt . --agent claude-code --output prompt.txt
```

If no agent is specified, Verso uses the default adapter or falls back to a generic text format.

## Agent Workflow

1. User runs `verso prompt <project> --agent <name>`
2. Verso loads and filters the project (via Router)
3. Verso selects the requested agent adapter
4. Adapter formats the filtered project into agent-specific output
5. Output is written to file or printed to stdout

## Non-Goals

- This RFC does not define how agents execute tasks
- This RFC does not enforce agent behavior
- This RFC does not create a universal agent format — each adapter is independent
- This RFC does not support custom/unknown agents at runtime

## Relationship to Other Components

- **Router**: Determines which components the agent should see
- **Template**: Defines how context sections are rendered before agent formatting
- **Skill/Memory/Workflow**: Content consumed by the agent, not processed by it

## Implementation Notes

- Agent adapters live in `internal/agent/adapters/`
- Each adapter is a separate Go file
- Adapter registration happens in `internal/agent/registry.go`
- The CLI flag `--agent` selects which adapter to use

## Open Questions

1. Should Verso provide a fallback "generic" agent format for unsupported agents?
2. How should agent adapters handle authentication or environment setup validation?
3. Should we support agent-to-agent communication in future versions?

---

*This RFC is a draft. Implementation details may change based on feedback and testing.*