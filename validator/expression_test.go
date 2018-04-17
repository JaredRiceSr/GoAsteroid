package validator

import (
	"testing"

	"github.com/benchlab/asteroid/parser"

	"github.com/benchlab/bvmUtils"
)

func TestCallExpressionValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        func f(a, b int8) int8 {
			if a == 0 or b == 0 {
                return 0
            }
            return f(a - 1, b - 1)
        }
    `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        interface Open {

        }

        Open(5, 5)
    `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCallExpressionEmptyConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class OrderBook {

        }

        d = new OrderBook()
        `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionSingleArgumentConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class OrderBook {

            var buyOrder int8

            constructor(age int8){
                buyOrder = age
            }
        }

        d = new OrderBook(10)
        `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionMultipleArgumentConstructorValid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class OrderBook {

            var buyOrder int8
            var fullName string

            constructor(name string, age int8){

            }
        }

        d = new OrderBook("alan", int8(10))
        `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestCallExpressionConstructorInvalid(t *testing.T) {
	scope, _ := parser.ParseString(`
        class OrderBook {

        }

        d = new OrderBook(6, 6)
        `)
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestOnlyExpressionResolved(t *testing.T) {
	_, errs := ValidateExpression(NewTestBVM(), "5 < 4")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMapLiteralInvalidKeysAndValues(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var m map[string]string
		m = map[string]string {
			1: 2,
			3: 5,
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 4, errs.Format())
}

func TestMapLiteralInvalidKeys(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var m map[string]string
		m = map[string]string {
			1: "hi",
			3: "hi",
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestMapLiteralInvalidValues(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var m map[string]string
		m = map[string]string {
			"hi": 1,
			"hi": 3,
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestMapLiteralValidValues(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var m map[string]string
		m = map[string]string {
			"hi": "bye",
			"hi": "bye",
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestArrayLiteralValidValues(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var m []string
		m = []string { "hi", "bye" }
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}
