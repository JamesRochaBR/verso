# Verso - Implementation Vision

> **Este documento é a fonte de verdade sobre a visão do projeto.**
>
> Antes de implementar qualquer funcionalidade, uma IA deve ler este documento juntamente com os RFCs da pasta `/spec`.
>
> Os RFCs definem **como** cada parte funciona.
> Este documento explica **por que** ela existe.

---

# O que é o Verso?

Verso é um framework open-source para construção de contexto para Large Language Models (LLMs).

Seu objetivo NÃO é conversar com modelos de IA.

Seu objetivo é organizar conhecimento de forma estruturada, reutilizável e previsível para produzir prompts consistentes.

Em outras palavras:

Verso transforma um conjunto de arquivos organizados em um contexto pronto para ser utilizado por qualquer LLM.

---

# O problema que o Verso resolve

Hoje cada desenvolvedor cria seus próprios prompts.

Esses prompts normalmente possuem problemas:

- conhecimento duplicado
- difícil manutenção
- impossível reutilizar
- mistura contexto do projeto com instruções
- dependência de um modelo específico

O Verso resolve isso separando responsabilidades.

O conhecimento deixa de ficar dentro de prompts gigantes e passa a existir em componentes independentes.

---

# O Verso NÃO é

O Verso NÃO é um agente.

O Verso NÃO é um chatbot.

O Verso NÃO é um editor.

O Verso NÃO executa código.

O Verso NÃO possui memória própria.

O Verso NÃO chama APIs da OpenAI, Anthropic ou Google.

O Verso NÃO depende de um LLM específico.

---

# O propósito do Verso

O propósito do projeto é montar contexto.

Nada além disso.

O resultado final do Verso é um prompt consistente.

Quem executa esse prompt é outro sistema.

---

# Filosofia

O conhecimento deve ser organizado.

O conhecimento deve ser reutilizável.

O conhecimento deve ser desacoplado.

Cada arquivo deve possuir uma única responsabilidade.

Tudo deve ser composto.

Nunca duplicado.

---

# Os pilares do Verso

Todo projeto Verso é composto por quatro pilares.

```

Project

├── Skills
├── Memory
├── Templates
└── Workflows

```

Todo o framework gira em torno desses quatro conceitos.

---

# Skills

Uma Skill representa uma capacidade.

Ela responde:

> "O que este especialista sabe fazer?"

Exemplos:

- architect
- reviewer
- golang
- react
- security
- database
- backend
- frontend
- devops
- testing

Uma Skill nunca descreve um projeto.

Ela descreve conhecimento especializado.

Uma Skill deve ser reutilizável entre dezenas de projetos diferentes.

Uma Skill nunca deve conter regras específicas de um projeto.

---

# Memory

Memory representa conhecimento específico.

Ela responde:

> "O que este projeto precisa lembrar?"

Exemplos:

- arquitetura do projeto
- regras de negócio
- padrões de código
- convenções
- decisões arquiteturais
- histórico técnico

Memory pertence ao projeto.

Ela não deve ser reutilizada entre projetos diferentes.

---

# Templates

Templates definem apresentação.

Eles NÃO armazenam conhecimento.

Eles definem como o contexto será montado.

Exemplos:

- feature.md
- bugfix.md
- review.md
- planning.md

---

# Workflows

Workflow representa um processo.

Ele responde:

> "Qual sequência deve ser executada?"

Exemplo:

```

Implement Feature

↓

Load Skills

↓

Load Memory

↓

Load Template

↓

Compose Context

↓

Render Prompt

```

Workflow não contém conhecimento.

Workflow apenas organiza a execução.

---

# Como tudo funciona

Um comando como:

```

verso prompt \
--skill architect \
--skill golang \
--memory project \
--template feature

```

segue o fluxo:

```

Discovery

↓

Validation

↓

Load Skills

↓

Load Memory

↓

Load Template

↓

Compose Context

↓

Render

↓

Output

```

Todo comando do Verso deve respeitar essa arquitetura.

---

# Responsabilidade de cada package

## internal/config

Responsável pelo verso.toml.

Nunca deve conhecer Skills ou Templates.

---

## internal/project

Responsável por descobrir um projeto Verso.

Nunca deve renderizar prompts.

---

## internal/render

Responsável apenas por renderizar.

Nunca deve descobrir arquivos.

---

## internal/build

Responsável por gerar artefatos.

Nunca deve conter lógica de descoberta.

---

## internal/cli

Responsável apenas pela interface de linha de comando.

Toda regra de negócio deve existir em outros packages.

---

# Descoberta (Discovery)

Todo componente do Verso deve ser descoberto automaticamente.

A CLI nunca deve conhecer caminhos específicos.

Exemplo:

```

skills/
memory/
templates/
workflows/

```

Toda descoberta pertence ao mecanismo de Discovery.

---

# Composição

O núcleo do Verso é o processo de composição.

O objetivo não é ler arquivos.

O objetivo é montar contexto.

Toda implementação deve fortalecer esse mecanismo.

Nunca criar atalhos.

Nunca duplicar comportamento.

---

# Princípios Arquiteturais

Sempre seguir:

- Baixo acoplamento
- Alta coesão
- Componentes pequenos
- Interfaces simples
- Responsabilidade única
- Código previsível
- Código testável

---

# Independência de LLM

O Verso nunca deve conhecer:

- ChatGPT
- Claude
- Gemini
- Copilot
- Qwen
- DeepSeek

Todos devem ser consumidores do resultado produzido pelo Verso.

---

# Implementação

Antes de implementar qualquer funcionalidade, responder:

## 1

Qual RFC esta implementação atende?

Se a resposta for "nenhum", provavelmente a implementação está errada.

---

## 2

Esta funcionalidade fortalece o mecanismo de composição?

Se não fortalecer, deve ser revista.

---

## 3

Existe duplicação de responsabilidade?

Se existir, refatorar.

---

## 4

Esta implementação aumenta o acoplamento?

Se sim, procurar outra solução.

---

## 5

Ela mantém o projeto independente de qualquer LLM?

Se não, rejeitar a implementação.

---

# Os RFCs

Os RFCs presentes em `/spec` são a especificação oficial do projeto.

Eles possuem prioridade sobre qualquer implementação existente.

Em caso de conflito:

RFC > Código

Nunca o contrário.

---

# Objetivo Final

O objetivo do Verso NÃO é possuir dezenas de comandos.

O objetivo do Verso é possuir um núcleo pequeno, sólido e previsível para construção de contexto.

Toda funcionalidade implementada deve fortalecer esse núcleo.

Nunca aumentar complexidade desnecessária.

---

# Para IAs

Se você está implementando o Verso, siga esta ordem:

1. Leia este documento.
2. Leia os RFCs relacionados.
3. Entenda a arquitetura.
4. Implemente apenas o que o RFC define.
5. Nunca invente comportamento novo.
6. Nunca altere a filosofia do projeto.
7. Sempre preserve a simplicidade do núcleo.

A implementação deve seguir a arquitetura.

Nunca a arquitetura seguir a implementação.