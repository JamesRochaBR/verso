# Verso - Plano de Implementação

Arquivo de tracking para acompanhamento do progresso da implementação do projeto Verso.

## 📊 Status Atual

- **Milestone Atual:** v0.3 — Router ✅
- **Próximo Milestone:** v0.4 — Skills
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

## Fase 3: v0.4 - Skills

> Skills são a menor unidade reutilizável de conhecimento dentro do framework.

- [ ] Implementar Skill validation completa (expandir `internal/project/validate.go`)
- [ ] Adicionar metadata parsing (frontmatter YAML) em `internal/project/markdown.go`
  - Suportar campos: name, version, author, tags, description
- [ ] Implementar Skill lifecycle tracking
- [ ] Criar testes para Skills (`internal/project/skill_test.go`)

---

## Fase 4: v0.5 - Memory

> Memory armazena conhecimento específico do projeto através das interações.

- [ ] Completar Memory frontmatter parsing
- [ ] Implementar Memory versioning
- [ ] Adicionar tags support no discovery (`internal/project/discovery.go`)
- [ ] Criar testes para Memory (`internal/project/memory_test.go`)

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

## Fase 8: v0.9 - Examples & Validation

> Exemplos que validam a arquitetura e nunca a definem.

- [ ] Expandir hello-world example
- [ ] Criar exemplo multi-skill
- [ ] Criar exemplo com workflow completo
- [ ] Adicionar validação de projetos exemplo
- [ ] Documentar examples em `examples/README.md`

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
