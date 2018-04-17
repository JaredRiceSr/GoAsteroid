package evm

import (
	"testing"

	"github.com/benchlab/asteroid/validator"

	"github.com/benchlab/asteroid/parser"

	"github.com/benchlab/bvmUtils"
)

func TestBuiltins(t *testing.T) {
	expr, _ := parser.ParseFile("test/builtins.grd")
	errs := validator.Validate(expr, NewBVM())
	bvmUtils.Assert(t, expr != nil, "expr is nil")
	bvmUtils.Assert(t, len(errs) == 0, "expr is nil")
}

func TestGreeter(t *testing.T) {
	expr, _ := parser.ParseFile("test/greeter.grd")
	errs := validator.Validate(expr, NewBVM())
	bvmUtils.Assert(t, expr != nil, "expr is nil")
	bvmUtils.Assert(t, len(errs) == 0, "expr is nil")
}
