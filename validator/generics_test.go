package validator

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
        class List<T> {

        }

        a = new List<string>()
    `)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInheritanceConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class OrderBook {}
		class DEX inherits OrderBook {}
        class ConstrainedList<T inherits OrderBook> {

        }

        a = new ConstrainedList<string>()
		b = new ConstrainedList<DEX>()
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestImplementationConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {}
		class DEX is OrderBook {}
        class ConstrainedList<T is OrderBook> {

        }

        a = new ConstrainedList<string>()
		b = new ConstrainedList<DEX>()
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestMultipleInheritanceConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class OrderBook {}
		class DEX inherits OrderBook {}
        class ConstrainedList<T inherits OrderBook|S inherits OrderBook> {

        }

        a = new ConstrainedList<string|int>()
		b = new ConstrainedList<DEX|DEX>()
    `)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestMultipleImplementationConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {}
		class DEX is OrderBook {}
        class ConstrainedList<T is OrderBook|S is OrderBook> {

        }

        a = new ConstrainedList<string|int>()
		b = new ConstrainedList<DEX|DEX>()
    `)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestMultipleMixedConstraintGenericAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {}
		class Animal {}
		class DEX inherits Animal is OrderBook {}
		class FakeOrderBook is OrderBook {}
		class assetType inherits Animal {}
        class ConstrainedList<T is OrderBook inherits Animal|S inherits Animal is OrderBook> {

        }

        a = new ConstrainedList<string|int>()
		b = new ConstrainedList<DEX|DEX>()
		c = new ConstrainedList<FakeOrderBook|assetType>()
    `)
	bvmUtils.AssertNow(t, len(errs) == 4, errs.Format())
}

func TestSingleGenericInSignature(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class List<T> {

		}
		class OrderBook<K> inherits List<K> {

		}
    `)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMultipleGenericInSignature(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class List<T> {

		}
		class OrderBook<K|V> inherits List<K> {

		}
    `)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTypesWithoutParameters(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		enum OrderBook {}
		x = OrderBook<string>
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFunctionGenericParameter(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func main(){
			var a int
			var b []int
			a = len(b)
			var c []string
			a = len(c)
		}

		func <T> lengthOf(items []T) int {
			return 0
		}
    `)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFunctionTwoGenericParameters(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func main(){
			var a []int
			var b []int
			merge(a, b)
			var c []string
			// should error
			merge(a, c)
		}

		func <T> merge(items []T, others []T) {

		}
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFunctionTwoGenericParametersReturnType(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func main(){
			var a []int
			var b []int
			a = thing(a, b)
			var c []string
			// should error
			c = thing(a, b)
		}

		func <T> thing(items []T) []T {
			return items
		}
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestGenericFunctionLiteralAssignment(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func main(){
			var y func(int)int
			x = func <T> (a T) {

			}
			x = y
		}

		func <T> thing(items []T) []T {
			return items
		}
    `)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}
