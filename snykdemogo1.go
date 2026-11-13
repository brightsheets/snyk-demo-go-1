// Package snykdemogo1 is a fixture library with cross-repo Go module dependencies.
package snykdemogo1

import (
	deps "github.com/brightsheets/deps-01"
)

// ID is this module's name.
const ID = "snyk-demo-go-1"

// Name reports this library and its direct dependencies.
func Name() string { return deps.Greet("snyk-demo-go-1") }
