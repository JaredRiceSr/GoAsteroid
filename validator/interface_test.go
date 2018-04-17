package validator

import (
	"testing"

	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/bvmUtils"
)

func TestMakeName(t *testing.T) {
	// single name
	names := []string{"hi"}
	bvmUtils.Assert(t, makeName(names) == "hi", "wrong single make name")
	names = []string{"hi", "you"}
	bvmUtils.Assert(t, makeName(names) == "hi.you", "wrong multiple make name")
}

func TestValidateString(t *testing.T) {
	scope, errs := ValidateString(NewTestBVM(), `
		if x = 0; x > 5 {

		} else {

		}
	`)

	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
}

func TestValidateExpression(t *testing.T) {
	expr, errs := ValidateExpression(NewTestBVM(), "5")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, expr != nil, "expr should not be nil")
	bvmUtils.AssertNow(t, expr.ResolvedType() != nil, "resolved is nil")
	_, ok := expr.ResolvedType().(*typing.NumericType)
	bvmUtils.AssertNow(t, ok, "wrong type")
}

func TestNewValidator(t *testing.T) {
	te := NewTestBVM()
	v := NewValidator(te)
	bvmUtils.AssertLength(t, len(v.operators), len(operators()))
	bvmUtils.AssertLength(t, len(v.literals), len(te.Literals()))
	bvmUtils.AssertNow(t, len(v.primitives) > 0, "no primitives")
}

func TestDeclarationAndCall(t *testing.T) {
	te := NewTestBVM()
	_, errs := ValidateString(te, `
		func hi(a bool){

		}
		hi(6 > 7)
		hi(false)
	`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
}

func TestValidatePackageFileData(t *testing.T) {
	a := `
	package x asteroid.0.0.1
	class DEX {}
	`
	b := `
	package x asteroid.0.0.1
	class OrderBook inherits x {}
	`
	errs := ValidateFileData(NewTestBVM(), []string{a, b})
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
}
