package comet

import "github.com/benchlab/asteroid/compiler/ast"

// VerifyInvariant ...
func VerifyInvariant() {

}

// Verifier ...
type Verifier struct {
	Invariants []Invariant
}

// Invariant ...
type Invariant struct {
	Condition func(ast.Station) bool
}
