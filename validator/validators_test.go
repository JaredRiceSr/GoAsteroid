package validator

import (
	"fmt"
	"testing"

	"github.com/benchlab/asteroid/parser"

	"github.com/benchlab/bvmUtils"
)

func TestTypeValidateValid(t *testing.T) {
	scope, errs := parser.ParseString(`
            var a OrderBook
            type OrderBook int
        `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	le := scope.Declarations.Length()
	bvmUtils.AssertNow(t, len(errs) == 0, "Parser: "+errs.Format())
	bvmUtils.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
	errs = Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, "Validator: "+errs.Format())
}

func TestTypeValidateInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
            var b assetType
            type OrderBook int
    `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	le := scope.Declarations.Length()
	bvmUtils.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}
