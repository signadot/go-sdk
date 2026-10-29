//go:build tools

// Package tools pins the tools CI runs, so that Dependabot can bump them. It is
// its own module to keep them out of go-sdk's dependency graph.
package tools

import _ "golang.org/x/vuln/cmd/govulncheck"
