// SPDX-License-Identifier: Elastic-2.0
// Copyright 2026 Zero Root AI

// Command zda-ast — AST-grounded code retrieval CLI for zeroroot-ai agents.
// Slice 3.9 of the production-readiness epic (gibson#173 → board #16).
package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `zda-ast — AST-grounded code retrieval CLI for zeroroot-ai agents

USAGE:
	zda-ast <subcommand> [args...]

SUBCOMMANDS:
	where-called <symbol>     find call sites of a fully-qualified Go symbol
	methods-of <type>         list methods on a Go type
	import-graph              show Go import edges (--from / --to / --check)
	proto-authz <rpc>         authz contract for an RPC from registry
	imports-of <path>         typed import list
	package-deps <pkg>        direct + transitive package dependencies
	walker-coverage           report what fraction of internal/ a walker scans
	catalog                   emit a markdown catalog of all gates / walkers / queries

OUTPUT: JSON by default; pass -h for human-readable.

REFS:
	Parent PRD: https://github.com/zeroroot-ai/gibson/issues/173
	Slice:      https://github.com/zeroroot-ai/.github/issues/52
	Board:      https://github.com/orgs/zeroroot-ai/projects/16
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}
	switch os.Args[1] {
	case "where-called":
		runWhereCalled(os.Args[2:])
	case "methods-of":
		stub("methods-of")
	case "import-graph":
		stub("import-graph")
	case "proto-authz":
		stub("proto-authz")
	case "imports-of":
		stub("imports-of")
	case "package-deps":
		stub("package-deps")
	case "walker-coverage":
		stub("walker-coverage")
	case "catalog":
		stub("catalog")
	case "-h", "--help", "help":
		fmt.Print(usage)
	default:
		fmt.Fprintf(os.Stderr, "unknown subcommand: %q\n%s", os.Args[1], usage)
		os.Exit(2)
	}
}

// stub emits a JSON status:not-yet-implemented marker. v0.1.0 ships
// scaffolding; per-subcommand implementations follow in v0.2+.
func stub(name string) {
	fmt.Printf("{\"subcommand\":%q,\"status\":\"not-yet-implemented\"}\n", name)
}

func runWhereCalled(args []string) {
	fs := flag.NewFlagSet("where-called", flag.ExitOnError)
	human := fs.Bool("h", false, "human-readable output")
	_ = fs.Parse(args)
	if len(fs.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "usage: zda-ast where-called <symbol>")
		os.Exit(2)
	}
	symbol := fs.Arg(0)
	if *human {
		fmt.Printf("where-called %s: (v0.1.0 — implementation pending; see slice 3.9)\n", symbol)
		return
	}
	fmt.Printf("{\"subcommand\":\"where-called\",\"symbol\":%q,\"results\":[],\"status\":\"not-yet-implemented\"}\n", symbol)
}
