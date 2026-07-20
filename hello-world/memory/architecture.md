---
name: architecture
type: memory
version: "1.0.0"
author: "versoteam"
tags: [architecture, design, system]
description: "System architecture overview and design decisions"
status: approved
---

# System Architecture

## Overview

Verso is an AI agent orchestration framework that manages project context through reusable components.

## Core Components

### Router
The brain of Verso. Decides which components should be loaded based on context using strategies like keyword-based, workflow-based, or default routing.

### Skills
The smallest reusable unit of knowledge. Each skill defines exactly one responsibility and remains independent from others.

### Memory
Project-specific knowledge that persists across interactions. Stores architectural decisions, business rules, conventions, and technical debt.

### Workflows
Orchestrate skills and memory components to execute defined processes.

### Templates
Define how context is assembled and presented to the AI agent.

## Directory Structure

```
project/
  verso.toml          # Project configuration
  skills/             # Reusable knowledge units
  memory/             # Project-specific knowledge
  workflows/          # Process orchestration
  templates/          # Context presentation templates
```

## Data Flow

1. User initiates interaction via CLI
2. Router analyzes context and selects components
3. Selected components are loaded and rendered
4. Context is presented to the AI agent
5. Agent executes based on provided context