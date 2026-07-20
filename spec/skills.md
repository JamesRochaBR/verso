# Skills

> **Este documento define a filosofia das Skills do Verso.**
>
> Este NÃO é um documento de implementação.
>
> Este documento define o propósito de cada Skill e como elas devem se comportar.
>
> Antes de implementar qualquer Skill, uma IA deve ler este documento juntamente com os RFCs.

---

# O que é uma Skill?

Uma Skill representa uma especialização.

Ela responde à pergunta:

> **"Se este fosse um especialista, quais conhecimentos ele possuiria?"**

Uma Skill nunca representa um projeto.

Uma Skill nunca representa uma tarefa.

Uma Skill representa apenas conhecimento.

---

# Objetivos

Uma Skill deve ser:

- reutilizável
- independente
- especializada
- pequena
- previsível

Ela nunca deve depender de outra Skill para existir.

---

# O que uma Skill NÃO deve fazer

Uma Skill NÃO deve:

- conhecer o projeto
- conhecer regras de negócio
- conhecer arquivos específicos
- conhecer diretórios
- conhecer uma IA específica
- executar código
- modificar arquivos

Ela apenas fornece conhecimento.

---

# Estrutura

Uma Skill normalmente será composta por um único arquivo Markdown.

Exemplo:

```
skills/
    architect.md
```

ou

```
skills/
    golang.md
```

O conteúdo deve ser totalmente textual.

---

# Filosofia

Uma Skill deve ser escrita como se fosse um especialista ensinando outro especialista.

Ela não é documentação.

Ela não é tutorial.

Ela representa experiência.

---

# Composição

Várias Skills podem ser carregadas simultaneamente.

Exemplo:

```
architect
golang
backend
postgres
security
```

Cada uma adiciona conhecimento ao contexto final.

Nenhuma delas deve sobrescrever outra.

O objetivo é composição.

Nunca substituição.

---

# Categorias

As Skills do Verso são divididas em categorias.

---

# Architecture

Responsável pelas decisões estruturais.

Exemplos:

- architect
- software-architect
- system-design
- clean-architecture
- ddd

Responsabilidades:

- organização do projeto
- divisão em módulos
- responsabilidades
- baixo acoplamento
- alta coesão
- escalabilidade

Nunca:

- escrever regras de negócio

---

# Code Review

Responsável pela qualidade.

Exemplos:

- reviewer
- code-review
- quality

Responsabilidades:

- identificar problemas
- sugerir melhorias
- verificar padrões
- reduzir complexidade
- encontrar duplicações

Nunca:

- alterar arquitetura

---

# Backend

Especialização em desenvolvimento backend.

Exemplos:

- backend
- api
- rest
- grpc

Responsabilidades:

- APIs
- serviços
- arquitetura backend
- validação
- autenticação
- autorização

Nunca:

- definir interface visual

---

# Frontend

Especialização em UI.

Responsabilidades:

- UX
- componentes
- acessibilidade
- performance
- organização visual

Nunca:

- definir arquitetura backend

---

# Golang

Especialização na linguagem Go.

Responsabilidades:

- organização em packages
- interfaces
- concorrência
- tratamento de erros
- testes
- performance

Nunca:

- definir regras de negócio

---

# TypeScript

Especialização em TypeScript.

Responsabilidades:

- tipos
- generics
- organização
- boas práticas

---

# React

Especialização em React.

Responsabilidades:

- componentes
- hooks
- composição
- performance

---

# Database

Especialização em persistência.

Responsabilidades:

- modelagem
- índices
- consultas
- normalização
- performance

---

# Security

Especialização em segurança.

Responsabilidades:

- autenticação
- autorização
- criptografia
- validação
- OWASP

Nunca:

- alterar regras de negócio

---

# DevOps

Especialização em infraestrutura.

Responsabilidades:

- Docker
- CI/CD
- Deploy
- Observabilidade
- Logs
- Monitoramento

---

# Testing

Especialização em testes.

Responsabilidades:

- testes unitários
- integração
- cobertura
- qualidade

---

# Documentation

Especialização em documentação.

Responsabilidades:

- README
- RFCs
- exemplos
- documentação técnica

---

# Product

Especialização em produto.

Responsabilidades:

- experiência do usuário
- priorização
- funcionalidades
- visão do produto

---

# Como as Skills trabalham juntas

Exemplo:

```
Feature

↓

Architect

↓

Backend

↓

Golang

↓

Testing

↓

Reviewer
```

Cada Skill contribui apenas com seu conhecimento.

Nenhuma Skill assume a responsabilidade da outra.

---

# Conflitos

Caso duas Skills entrem em conflito:

1. Arquitetura vence implementação.
2. Segurança vence conveniência.
3. Simplicidade vence complexidade.
4. Clareza vence abstração.

---

# Dependências

Uma Skill nunca deve depender da existência de outra.

Ela deve funcionar isoladamente.

A composição é responsabilidade do Verso.

Nunca da Skill.

---

# Independência

Uma Skill deve poder ser utilizada em qualquer projeto.

Se ela depende de um projeto específico, ela provavelmente deveria ser uma Memory.

---

# Skills x Memory

Esta é uma das regras mais importantes do Verso.

## Skill

Conhecimento reutilizável.

Exemplos:

- Clean Architecture
- Go
- React
- PostgreSQL

---

## Memory

Conhecimento específico.

Exemplos:

- Arquitetura do Projeto X
- Convenções da Empresa Y
- Regra de Negócio Z

Nunca misturar os dois conceitos.

---

# Skills x Templates

Skill responde:

> O que deve ser dito?

Template responde:

> Como deve ser apresentado?

---

# Skills x Workflow

Skill responde:

> O que eu sei?

Workflow responde:

> Em qual ordem executar?

---

# Evolução

Novas Skills podem ser adicionadas sem alterar o núcleo do Verso.

Esse é um dos principais objetivos da arquitetura.

O framework deve crescer horizontalmente.

Nunca verticalmente.

---

# Para IAs

Antes de criar uma nova Skill, responda:

- Ela representa conhecimento?
- Ela é reutilizável?
- Ela funciona isoladamente?
- Ela pode ser usada em dezenas de projetos?
- Ela não depende de um projeto específico?

Se qualquer resposta for "não", provavelmente ela não é uma Skill.

---

# Objetivo Final

O objetivo das Skills não é gerar prompts.

O objetivo das Skills é representar especialistas.

O Verso apenas reúne esses especialistas, combina seus conhecimentos com as Memories do projeto, aplica um Template e produz um contexto consistente para qualquer LLM.

As Skills são o conhecimento.

As Memories são o contexto.

Os Templates são a apresentação.

Os Workflows são a orquestração.

O Verso é o compositor que une tudo isso.