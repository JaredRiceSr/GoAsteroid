package validator

import (
	"fmt"
	"testing"

	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/asteroid/token"

	"github.com/benchlab/bvmUtils"
)

func TestAdd(t *testing.T) {
	m := OperatorMap{}

	bvmUtils.Assert(t, len(m) == 0, "Asteroid Errors: Configured With The Wrong Length")
	// numericalOperator with floats/ints

	m.Add(BinaryNumericOperator, token.Sub, token.Mul, token.Div)

	bvmUtils.Assert(t, len(m) == 3, fmt.Sprintf("Asteroid= Errors: Configured With The Wrong Length %d", len(m)))

	// integers only
	m.Add(BinaryIntegerOperator, token.Shl, token.Shr)

	bvmUtils.Assert(t, len(m) == 5, fmt.Sprintf("Asteroid Errors: Final Length Configured Incorrectly: %d", len(m)))
}

func TestImportVM(t *testing.T) {
	v := new(Validator)
	tvm := NewTestBVM()
	v.importBVM(tvm)
	bvmUtils.AssertNow(t, len(v.errs) == 0, v.errs.Format())
	bvmUtils.AssertNow(t, v.primitives != nil, "Asteroid Errors: Asteroid Primitives Returned a Nil Error. Nil Error Represents An Error In This Case.")
	typ, _ := v.isTypeVisible("address")
	bvmUtils.AssertNow(t, typ != typing.Unknown(), "Asteroid Errors: Address Unrecognized")
}

func TestCastingValidToUnsigned(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "x = uint(5)")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCastingInvalidToUnsigned(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "x = uint(-5)")
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}
