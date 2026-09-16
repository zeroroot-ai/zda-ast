# zda-ast

AST-grounded retrieval CLI for zeroroot-ai agents. The plan is a typed query over a parsed Go AST, so an agent can ask "who calls this symbol" instead of running grep.

Go module under the zeroroot-ai workspace. See [`zeroroot-ai/.github` → `AGENTS.md`](https://github.com/zeroroot-ai/.github/blob/main/AGENTS.md) for workflow conventions (branching, PRs, releases, agent merge autonomy).

## Status

**This repo is a stub. No subcommand returns a real answer yet.**

`cmd/zda-ast` parses the command line and nothing else. Seven of the eight planned subcommands print `{"status":"not-yet-implemented"}`. The eighth, `where-called`, prints the same status with an empty `results` array. An empty `results` array means "not implemented", never "no call sites".

Planned subcommands: `where-called`, `methods-of`, `import-graph`, `proto-authz`, `imports-of`, `package-deps`, `walker-coverage`, `catalog`.

Do not build anything on this CLI until the implementation lands.

## Install

```bash
go install github.com/zeroroot-ai/zda-ast/cmd/zda-ast@latest
```

## License and history

Elastic License 2.0. See [LICENSE](LICENSE). Zero Root AI is the licensor. The Elastic License 2.0 is source-available, not open source.

Issue and pull request numbers cited in comments and documents dated before 2026-09-05 refer to the tracker before the history reset, archived offline. They do not resolve on GitHub.
