// SPDX-License-Identifier: Elastic-2.0
// Copyright 2026 Zero Root AI

package main

import (
	"os"
	"testing"
)

// TestBuilds is a smoke test — the binary compiles.
func TestBuilds(t *testing.T) {
	if _, err := os.Stat("main.go"); err != nil {
		t.Fatalf("main.go missing: %v", err)
	}
}
