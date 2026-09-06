# zda-ast

AST-grounded retrieval CLI for zeroroot-ai agents. Subcommands: where-called, methods-of, import-graph, proto-authz, imports-of, package-deps, walker-coverage, catalog. Agents use this instead of grep for typed-symbol questions.

Internal Go module under the zeroroot-ai workspace. See [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) for workflow conventions (branching, PRs, releases, agent merge autonomy).

## Status

Bootstrap repo. Initial implementation lands via the corresponding production-readiness slice on board #16. Until then, this README + LICENSE + Makefile contract are the only contents.

## Install

```bash
go get github.com/zeroroot-ai/zda-ast@latest
```

## License

[BUSL-1.1](./LICENSE).

## License and history

Elastic License 2.0. See [LICENSE](LICENSE). Zero Root AI is the licensor.

Issue and pull request numbers cited in comments and documents dated before 2026-09-05 refer to the tracker before the history reset, archived offline. They do not resolve on GitHub.
