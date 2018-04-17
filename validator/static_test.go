package validator

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestStaticVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
        class OrderBook {
            static var name string
        }
        x = OrderBook.name
    `)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestNotStaticVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
        class OrderBook {
            var name string
        }
        x = OrderBook.name
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestEnumVariableAccess(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
        enum Days { Mon }
        x = Days.Mon
    `)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}
