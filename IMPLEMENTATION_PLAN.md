# Verso - Plano de Implementação

Arquivo de tracking para acompanhamento do progresso da implementação do projeto Verso.

## 📊 Status Atual

- **Milestone Atual:** v0.5 — Memory ✅
- **Próximo Milestone:** v0.6 — Workflows
- **Última Atualização:** 2026-07-19

---

## Fase 1: v0.2 - Specifications ✅ CONCLUÍDA

> Todas as especificações da Fase 01 foram criadas/completadas.

- [x] Analisar RFCs existentes e identificar lacunas
- [x] Criar RFC-0007 (Agent Specification)
  - [x] Definir AgentAdapter interface
  - [x] Especificar agentes suportados (GitHub Copilot, Cursor, Claude Code, Cline/Roo)
  - [x] Definir configuração em verso.toml
  - [x] Documentar workflow de seleção de agente
- [x] Criar RFC-0009 (Serialization Specification)
  - [x] Especificar formato TOML para verso.toml
  - [x] Especificar Markdown + YAML frontmatter para componentes
  - [x] Definir schema de frontmatter (campos obrigatórios e opcionais)
  - [x] Definir formato de workflow steps
  - [x] Contrato de serialização (determinismo, round-trip safety)
  - [x] Estrutura do pacote internal/serialization/
- [x] Completar RFC-0010 (Discovery) com detalhes de implementação
  - [x] Definir discovery roots (skills, memory, workflows, templates)
  - [x] Algoritmo de discovery em 3 passos (scan → parse → index)
  - [x] File discovery pattern e exclusões
  - [x] Component parsing e type validation
  - [x] Error handling (fatal errors vs warnings)
  - [x] ComponentIndex structure para lookup rápido
- [x] Adicionar ComponentType "router" e "agent" em `internal/project/component.go`

---

## Fase 2: v0.3 - Router ✅ CONCLUÍDA

> O Router é o cérebro do Verso. Decide quais componentes devem ser carregados com base no contexto.

- [x] Criar pacote `internal/router/`
  - [x] Criar `router.go` — Interface principal e estrutura Router
  - [x] Criar `strategy.go` — Definição de estratégias de routing + utilitários
  - [x] Criar `keyword.go` — Keyword-Based Routing Strategy
  - [x] Criar `workflow.go` — Workflow-Based Routing Strategy
  - [x] Criar `default.go` — Default Routing Strategy
  - [x] Criar `errors.go` — Erros personalizados
  - [x] Criar `router_test.go` — Testes unitários (23 testes passando)

- [x] Implementar interface Router
  ```go
  type Router interface {
      Route(project *Project, options RoutingOptions) (*Project, error)
  }
  
  type RoutingOptions struct {
      Filter   Filter
      Strategy string
      Keywords []string
  }
  ```

- [x] Implementar Keyword-Based Routing Strategy ✅
  - Analisar input do usuário por keywords
  - Match contra metadata de componentes (tags, names, descriptions)
  - Retornar projeto filtrado com componentes correspondentes
  - Sistema de scoring para ranking de relevância

- [x] Implementar Workflow-Based Routing Strategy ✅
  - Quando workflow é solicitado, carregar todos os componentes referenciados
  - Parse steps do workflow para identificar skills e memory referenciadas
  - Carregar em sequência conforme definido no workflow

- [x] Implementar Explicit Filtering (integrado com `internal/project/filter.go`)
  - Integrar com Router
  - Suportar `--name`, `--exclude` via CLI

- [x] Implementar Default Routing ✅
  - Quando nenhuma filter é especificada, carregar todos os componentes
  - Garantir comportamento padrão sensível

- [x] Integrar Router com CLI ✅
  - [x] Atualizar `internal/cli/prompt.go` para usar Router (router.New() + r.Route())
  - [x] Adicionar flags: `--keywords/-k`, `--workflow/-w`, `--strategy/-s` em `internal/cli/flags.go`
  - [x] Help text atualizado com novas opções

---

## Fase 3: v0.4 — Skills ✅ CONCLUÍDA

> Skills são a menor unidade reutilizável de conhecimento dentro do framework.

- [x] Adicionar metadata parsing (frontmatter YAML) em `internal/project/markdown.go`
  - [x] Criar struct `FrontmatterMetadata` com campos: name, type, version, author, tags, description, depends, status
  - [x] Implementar `ParseFrontmatter()` — extrai e parseia YAML frontmatter
  - [x] Implementar `ExtractBody()` — retorna markdown sem frontmatter
  - [x] Atualizar `ExtractTitle()` para pular frontmatter automaticamente
  - [x] Adicionar utilitários: `lowercaseString()`, `containsString()`, `containsTag()`
- [x] Expandir Component struct com metadata em `internal/project/component.go`
  - [x] Adicionar campos: Metadata, Version, Author, Tags, Description, Status
  - [x] Definir `LifecycleState` e constantes (created, reviewed, approved, deprecated)
  - [x] Implementar `HasTags()`, `ContainsKeyword()`, `IsDeprecated()`
- [x] Atualizar Discovery para usar metadata em `internal/project/discovery.go`
  - [x] Parse frontmatter ao descobrir componentes
  - [x] Preencher campos do Component com valores do frontmatter
  - [x] Mapear lifecycle status quando presente
- [x] Implementar Skill validation completa em `internal/project/validate.go`
  - [x] `ValidateSkill()` — valida regras do RFC-0002 (name, type, content)
  - [x] `ValidateLifecycle()` — verifica estado válido de lifecycle
  - [x] `ValidateComponent()` — validação genérica por tipo
  - [x] `ValidateComponents()` — validação em lote
- [x] Criar testes para Skills (`internal/project/skill_test.go`) — **49 testes passando**
  - [x] 7 testes de ParseFrontmatter (valid, empty, invalid YAML, minimal, tags as list)
  - [x] 5 testes de ExtractBody/ExtractTitle com frontmatter
  - [x] 6 testes de ValidateSkill
  - [x] 4 testes de ValidateLifecycle/LifecycleState
  - [x] 8 testes de Component helper methods (HasTags, ContainsKeyword, IsDeprecated)
  - [x] 2 testes de Discover integration (with/without frontmatter)
  - [x] 5 testes de ValidateComponent/ValidateComponents
- [x] Atualizar exemplos hello-world com frontmatter
  - [x] `hello-world/skills/architect.md` — metadata completo com tags e status
  - [x] `hello-world/skills/reviewer.md` — metadata completo com tags e status
  - [x] `hello-world/memory/project.md` — metadata para memory component

**Novas dependências:** `gopkg.in/yaml.v3`

**Total de testes no pacote `internal/project`:** 49 testes passando

---

## Fase 4: v0.5 - Memory ✅ CONCLUÍDA

> Memory armazena conhecimento específico do projeto através das interações.

- [x] Implementar `ValidateMemory()` em `internal/project/validate.go`
  - [x] Validar name é obrigatório e não vazio
  - [x] Validar type deve ser "memory"
  - [x] Validar content não pode ser vazio
- [x] Expandir `ValidateComponent()` para suportar `ComponentMemory`
- [x] Expandir `Validate()` para validar memory components no projeto
- [x] Criar pacote `internal/project/memory.go` com versioning utilities
  - [x] Struct `SemVer` com parsing de versão semântica
  - [x] `ParseSemanticVersion()` — suporta formato MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]
  - [x] `CompareVersions()` — comparação de versões semânticas
  - [x] `SemVer.String()`, `BumpMajor()`, `BumpMinor()`, `BumpPatch()`
  - [x] `IsVersioned()` — verifica se component tem versão válida
  - [x] `GetLatestVersion()` — encontra última versão baseada em semver
  - [x] `ListMemoryVersions()` — lista todas as versões de um memory component
- [x] Criar testes para Memory (`internal/project/memory_test.go`) — **42 testes passando**
  - [x] 8 testes de ParseSemanticVersion (valid, empty, invalid)
  - [x] 4 testes de SemVer.String()
  - [x] 3 testes de bump methods (BumpMajor, BumpMinor, BumpPatch)
  - [x] 6 testes de CompareVersions (equal, greater, less, prerelease, invalid)
  - [x] 3 testes de IsVersioned
  - [x] 5 testes de ValidateMemory
  - [x] 2 testes de ValidateComponent com Memory
  - [x] 3 testes de GetLatestVersion (single, multiple, none)
  - [x] 3 testes de ListMemoryVersions (multiple, excludes non-matching, excludes dirs)
  - [x] 3 testes de ValidateLifecycle para Memory
  - [x] 1 teste de Discover integration
  - [x] 2 testes de ValidateComponents batch
  - [x] 1 teste de edge case

- [x] Expandir exemplos em `hello-world/memory/`
  - [x] `project.md` — existente (mantido)
  - [x] `conventions.md` — coding conventions e standards
  - [x] `architecture.md` — system architecture overview
  - [x] `decisions.md` — architectural decisions log

**Novas funções implementadas:** 10 novas funções públicas em `memory.go`
**Novos testes:** 42 testes unitários
**Novos exemplos:** 3 novos memory components

---

## Fase 5: v0.6 - Workflows

> Workflows orquestram Skills e Memory components para executar processos definidos.

- [ ] Implementar Workflow step parsing (YAML steps no frontmatter)
- [ ] Implementar Workflow validation
- [ ] Implementar Workflow execution engine básico
- [ ] Criar testes para Workflows (`internal/workflow/engine_test.go`)

---

## Fase 6: v0.7 - Templates

> Templates definem como o contexto é montado e apresentado ao agente de IA.

- [ ] Criar pacote `internal/render/engines/`
  - [ ] Implementar Go template rendering engine
  - [ ] Criar `markdown.go` — Markdown output template
  - [ ] Criar `text.go` — Plain text output template
- [ ] Implementar variáveis de template (.Project, .Components, etc.)
- [ ] Implementar conditional sections
- [ ] Adicionar built-in templates (markdown, text)
- [ ] Criar testes para Templates (`internal/render/engines/engine_test.go`)

---

## Fase 7: v0.8 - Agent Compatibility

> Compatibilidade com agentes de codificação modernos.

- [ ] Definir formato de output por agente
- [ ] Implementar adapter para GitHub Copilot
- [ ] Implementar adapter para Cursor
- [ ] Implementar adapter para Claude Code
- [ ] Implementar adapter para Cline/Roo Code
- [ ] Criar pacote `internal/agent/` com interfaces comuns

---

## Fase 8: v0.9 — Examples & Validation

> Exemplos que validam a arquitetura e nunca a definem.

### Objetivo

Criar um conjunto abrangente de skills reutilizáveis que definem comportamentos padrão do agente Verso, organizadas por domínio de competência.

### Skills Planejadas (15 total)

#### 🧪 Testing (2 skills)
| # | Arquivo | Nome | Descrição | Tags |
|---|---------|------|-----------|------|
| 1 | `testing/test-driven-development.md` | TDD Cycle | RED-GREEN-REFACTOR cycle completo, incluindo anti-patterns de testing | `testing`, `tdd`, `red-green-refactor`, `anti-patterns` |
| 2 | `testing/verification-before-completion.md` | Verification Gate | Garantir que tudo está realmente fixo antes de declarar conclusão | `testing`, `verification`, `quality-gate` |

#### 🐛 Debugging (1 skill)
| # | Arquivo | Nome | Descrição | Tags |
|---|---------|------|-----------|------|
| 3 | `debugging/systematic-debugging.md` | Systematic Debugging | Processo de 4 fases para root cause analysis, incluindo técnicas de root-cause-tracing, defense-in-depth, condition-based-waiting | `debugging`, `root-cause`, `defense-in-depth`, `condition-based-waiting` |

#### 🤝 Collaboration (9 skills)
| # | Arquivo | Nome | Descrição | Tags |
|---|---------|------|-----------|------|
| 4 | `collaboration/brainstorming.md` | Brainstorming | Refinamento de design via perguntas socráticas | `collaboration`, `design`, `socratic-method` |
| 5 | `collaboration/writing-plans.md` | Writing Plans | Criação de planos de implementação detalhados | `planning`, `documentation`, `architecture` |
| 6 | `collaboration/executing-plans.md` | Executing Plans | Batch execution com checkpoints | `execution`, `checkpoints`, `workflow` |
| 7 | `collaboration/dispatching-parallel-agents.md` | Parallel Agents | Workflows concorrentes com subagents | `parallel`, `subagent`, `concurrent` |
| 8 | `collaboration/requesting-code-review.md` | Request Review | Checklist prévia para solicitar code review | `code-review`, `checklist`, `quality` |
| 9 | `collaboration/receiving-code-review.md` | Receive Review | Como responder a feedback de forma construtiva | `code-review`, `feedback`, `iteration` |
| 10 | `collaboration/using-git-worktrees.md` | Git Worktrees | Branches paralelos para desenvolvimento simultâneo | `git`, `worktrees`, `parallel-development` |
| 11 | `collaboration/finishing-a-development-branch.md` | Finish Branch | Workflow de merge/PR com critério de decisão | `git`, `merge`, `pr`, `workflow` |
| 12 | `collaboration/subagent-driven-development.md` | Subagent Dev | Iteração rápida com review em duas etapas (spec compliance → code quality) | `subagent`, `two-stage-review`, `iteration` |

#### 🧠 Meta (2 skills)
| # | Arquivo | Nome | Descrição | Tags |
|---|---------|------|-----------|------|
| 13 | `meta/writing-skills.md` | Writing Skills | Criar novas skills seguindo best practices (inclui methodology de testing) | `skills`, `authoring`, `best-practices` |
| 14 | `meta/using-superpowers.md` | Using Superpowers | Introdução ao sistema de skills do Verso | `onboarding`, `introduction`, `skills-system` |

#### 💡 Philosophy (1 skill)
| # | Arquivo | Nome | Descrição | Tags |
|---|---------|------|-----------|------|
| 15 | `philosophy/core-principles.md` | Core Principles | Test-Driven Development, Systematic over ad-hoc, Complexity reduction, Evidence over claims | `philosophy`, `principles`, `tdd`, `simplicity` |

### Tarefas de Implementação

- [ ] Criar estrutura de diretórios para skills
  - `testing/` — Skills de testing e verificação
  - `debugging/` — Skills de debugging sistemático
  - `collaboration/` — Skills de colaboração e trabalho em equipe
  - `meta/` — Skills sobre o próprio sistema Verso
  - `philosophy/` — Princípios fundamentais

- [ ] Implementar Testing skills (2 arquivos)
  - [x] `testing/test-driven-development.md` — TDD RED-GREEN-REFACTOR cycle
  - [x] `testing/verification-before-completion.md` — Verification gate

- [ ] Implementar Debugging skills (1 arquivo)
  - [x] `debugging/systematic-debugging.md` — 4-phase root cause process

- [ ] Implementar Collaboration skills (9 arquivos)
  - [x] `collaboration/brainstorming.md` — Socratic design refinement
  - [x] `collaboration/writing-plans.md` — Detailed implementation plans
  - [x] `collaboration/executing-plans.md` — Batch execution with checkpoints
  - [x] `collaboration/dispatching-parallel-agents.md` — Concurrent subagent workflows
  - [x] `collaboration/requesting-code-review.md` — Pre-review checklist
  - [x] `collaboration/receiving-code-review.md` — Responding to feedback
  - [x] `collaboration/using-git-worktrees.md` — Parallel development branches
  - [x] `collaboration/finishing-a-development-branch.md` — Merge/PR decision workflow
  - [x] `collaboration/subagent-driven-development.md` — Two-stage review

- [ ] Implementar Meta skills (2 arquivos)
  - [x] `meta/writing-skills.md` — Create new skills following best practices
  - [x] `meta/using-superpowers.md` — Introduction to the skills system

- [ ] Implementar Philosophy skill (1 arquivo)
  - [x] `philosophy/core-principles.md` — Core principles do Verso

- [ ] Criar workflow de exemplo multi-skill
  - [ ] `hello-world/workflows/tdd-cycle.md` — Orquestra: philosophy → testing → debugging → code-review

- [ ] Atualizar validação de projetos exemplo
  - Verificar se todas as skills têm frontmatter válido
  - Validar unicidade de nomes dentro do projeto
  - Garantir consistência de tags

**Total de novas skills:** 15
**Novos diretórios:** 5 (testing, debugging, collaboration, meta, philosophy)

---

## Fase 9: v1.0 - Stable

> Versão estável e production-ready.

- [ ] Cobertura de testes > 80%
- [ ] Documentation completa
- [ ] CI/CD pipeline
- [ ] Release notes e changelog atualizado
- [ ] Performance benchmarks

---

## 📐 Regras Arquiteturais (para referência)

1. Toda nova capacidade começa com um problema documentado
2. Todo problema deve produzir pelo menos um Use Case
3. Toda decisão arquitetural deve ser documentada antes da implementação
4. Toda implementação deve seguir uma especificação aceita
5. Exemplos validam a arquitetura, nunca a definem
6. O repositório deve reduzir complexidade continuamente

---

## 📝 Component Types Necessários

Atualmente implementados:
- [x] `skill` — Unidade menor de conhecimento reutilizável
- [x] `memory` — Conhecimento específico do projeto
- [x] `workflow` — Orquestração de processos
- [x] `template` — Estrutura de output
- [x] `router` — Camada de decisão de contexto (Fase 2)
- [x] `agent` — Compatibilidade com agentes (RFC definida, implementação Fase 7)

Necessários (não implementados):
- [ ] `agent` — Adapters concretos (GitHub Copilot, Cursor, Claude Code, Cline/Roo) → Fase 7

---

## 🔑 Arquivos-Chave para Modificação

| Arquivo | Alteração Necessária | Fase |
|---------|---------------------|------|
| `internal/project/component.go` | Adicionar router e agent types | 1 |
| `internal/project/markdown.go` | Adicionar frontmatter YAML parsing | 3,4 |
| `internal/project/filter.go` | Expandir com suporte a keywords/tags | 2 |
| `internal/cli/prompt.go` | Integrar com Router | 2 |
| `internal/cli/flags.go` | Adicionar novas flags (keywords, workflow) | 2 |
| `internal/render/renderer.go` | Implementar Go template engine | 6 |

---

## 🧪 Testes Existentes

- [x] `internal/cli/flags_test.go` — Tests para flag parsing
- [x] `internal/cli/list_test.go` — Tests para comando list
- [x] `internal/cli/prompt_test.go` — Tests para prompt command
- [x] `internal/project/discovery_test.go` — Tests para discovery
- [x] `internal/project/filter_test.go` — Tests para filter
- [x] `internal/project/markdown_test.go` — Tests para markdown parsing
- [x] `internal/project/validate_test.go` — Tests para validation
- [x] `internal/render/prompt_test.go` — Tests para prompt renderer
- [x] `internal/router/router_test.go` — **CONCLUÍDO** (23 testes passando, Fase 2)
