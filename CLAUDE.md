# zda-ast — CLAUDE.md

> **Workflow rules:** see [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) — canonical for branching / commits / PRs / releases / merging. Conventional Commits MANDATORY. Never push to main. Never force-push.

This file is the per-repo addendum. Workspace-wide concerns live in the workspace `CLAUDE.md`; architectural decisions in ``docs/adr/`` (local docs → `adr`).

## TL;DR

AST-grounded retrieval CLI (`github.com/zeroroot-ai/zda-ast`) that agents use instead of grep for typed-symbol questions. Entry point: `zda-ast <subcommand>`. Subcommands: `where-called`, `methods-of`, `import-graph`, `proto-authz`, `imports-of`, `package-deps`, `walker-coverage`, `catalog`.

## Architecture

A single Go module under `cmd/`. Each subcommand is a typed query over a parsed Go AST / import graph, returning structured answers (callers of a symbol, methods on a type, import edges, proto-authz lookups). It is the typed-search arm of the workspace search strategy — see ``docs/agents/search-strategy.md`` (local docs → `agents/search-strategy.md`).

## Regen commands

No proto/codegen in this repo. Standard Go build:

```bash
make build        # compile the CLI
make check        # fmt vet test-race
```

## Gotchas

- Internal Go module — not part of the open (Apache) tier; do not add it as a dependency of any Apache-licensed repo.
- Results are only as fresh as the parsed tree; run against the current checkout, not a stale GOPATH copy.

## Links

- Org-level workflow: [`AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md)
- Workspace map: workspace `CLAUDE.md`
- Search strategy: ``docs/agents/search-strategy.md`` (local docs → `agents/search-strategy.md`)
- PR checklist: ``docs/agents/pr-checklist.md`` (local docs → `agents/pr-checklist.md`)
