package validator

import (
	"testing"

	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/asteroid/parser"

	"github.com/benchlab/bvmUtils"
)

func TestValidateClassDecl(t *testing.T) {

}

func TestValidateInterfaceDecl(t *testing.T) {

}

func TestValidateEnumDecl(t *testing.T) {

}

func TestValidateEventDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("event Exchange()")
	bvmUtils.AssertNow(t, scope != nil, "Asteriod Errors: Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateEventDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("event OrderBook(a int)")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())

}

func TestValidateEventDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("event OrderBook(a int, b string)")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateEventDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("event OrderBook(c assetType)")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateEventDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("event OrderBook(c assetType, a Animal)")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateEventDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("event OrderBook(a int, b assetType)")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateFuncDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("func OrderBook() {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("func OrderBook(a int) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("func OrderBook(a int, b string) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateFuncDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("func exchange(a assetType) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateFuncDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("func OrderBook(a assetType, b Animal) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateFuncDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("func OrderBook(a int, b assetType) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateConstructorDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("constructor() {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int, b string) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateConstructorDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a assetType) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateConstructorDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a assetType, b Animal) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateConstructorDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("constructor(a int, b assetType) {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateContractDeclEmpty(t *testing.T) {
	scope, _ := parser.ParseString("contract OrderBook {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclValidSingle(t *testing.T) {
	scope, _ := parser.ParseString("contract OrderBook{} contract OrderBook inherits assetType {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclValidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("contract assetType {} contract Animal {} contract OrderBook inherits assetType, Animal {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateContractDeclInvalidSingle(t *testing.T) {
	scope, _ := parser.ParseString("contract OrderBook inherits assetType {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateContractDeclInvalidMultiple(t *testing.T) {
	scope, _ := parser.ParseString("contract OrderBook inherits assetType, Animal {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())

}

func TestValidateContractDeclMixed(t *testing.T) {
	scope, _ := parser.ParseString("contract OrderBook{} contract OrderBook inherits assetType, Animal {}")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateExplicitVarDecl(t *testing.T) {
	scope, _ := parser.ParseString("var hi uint8")
	bvmUtils.AssertNow(t, scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, scope.Declarations != nil, "Asteroid Errors: Your declarations shouldn't be nil!")
	errs := Validate(NewTestBVM(), scope, nil)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateModifiersValidAccess(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "public var name string")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidateModifiersDuplicateAccess(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "public public var name string")
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateModifiersMEAccess(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "public private var name string")
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateModifiersInvalidNodeType(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "test var name string")
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidateModifiersInvalidUnrecognised(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), "elephant var name string")
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestInterfaceParents(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface Switchable{}
		interface Deletable{}
		interface Light inherits Switchable, Deletable {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInterfaceParentsImplemented(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface Switchable{
			on()
			off()
		}
		interface Deletable{}
		interface Light inherits Switchable, Deletable {

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMapTypeValid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var a map[int]int
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestMapTypeInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		var a map[zoop]doop
	`)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestEnumValid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		enum Weekday {
			Mon, Tue, Wed, Thurs, Fri
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestEnumInheritanceValid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		enum Weekday {
			Mon, Tue, Wed, Thurs, Fri
		}

		enum Weekend {
			Sat, Sun
		}

		enum Day inherits Weekday, Weekend {

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestEnumInheritanceInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		enum Weekday {
			Mon, Tue, Wed, Thurs, Fri
		}

		interface Weekend {

		}

		enum Day inherits Weekday, Weekend {

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassInheritanceValid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class OrderBook {}
		class DEXOB inherits OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestClassInheritanceInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {}
		class DEXOB inherits OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationValid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {}
		class DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestClassImplementationInvalid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class OrderBook {}
		class DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationValid(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {}
		contract DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestContractImplementationInvalidWrongType(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		contract OrderBook {}
		contract DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidWrongType(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class OrderBook {}
		class DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationInvalidMissingMethods(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}
		contract DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidMissingMethods(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}
		class DEXOB is OrderBook {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationInvalidWrongMethodParameters(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}
		contract DEXOB is OrderBook {
			func trade(m string){

			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidWrongMethodParameters(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}
		class DEXOB is OrderBook {
			func trade(m string){

			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestContractImplementationValidThroughParent(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}

		contract X {
			func trade(){

			}
		}

		contract DEXOB is OrderBook inherits X {

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestClassImplementationValidThroughParent(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}

		class X {
			func trade(){

			}
		}

		class DEXOB is OrderBook inherits X {

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestContractImplementationInvalidMissingParentMethods(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}

		interface assetType inherits OrderBook {

		}

		contract DEXOB is assetType {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestClassImplementationInvalidMissingParentMethods(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		interface OrderBook {
			trade()
		}

		interface assetType inherits OrderBook {

		}

		class DEXOB is assetType {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestInterfaceMethods(t *testing.T) {
	a, errs := ValidateString(NewTestBVM(), `
		interface Calculator {
			add(a, b int) int
			sub(a, b int) int
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
	n := a.Declarations.Next().(*ast.InterfaceDeclarationNode)
	i := n.Resolved.(*typing.Interface)
	bvmUtils.AssertLength(t, len(i.Funcs), 2)
	add := i.Funcs["add"]
	bvmUtils.AssertNow(t, add != nil, "add is nil")
	sub := i.Funcs["sub"]
	bvmUtils.AssertNow(t, sub != nil, "sub is nil")
}

func TestCancellationEnums(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		enum Weekday { Mon, Tue, Wed, Thu, Fri }
		enum WeekdayTwo { Mon, Tue, Wed, Thu, Fri }
		enum Cancelled inherits Weekday, WeekdayTwo {}
		x = Cancelled.Mon
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCancellationClass(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class UserA { var name string }
		class UserB { var name string }
		class AtomicSwapMembers inherits UserA, UserB {
			func executeSwap()() {
				x = name
			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestCancellationContract(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		contract UserA { var name string }
		contract UserB { var name string }
		contract AtomicSwapMembers inherits UserA, UserB {
			func executeSwap()() {
				x = name
			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFuncDeclarationSingleReturn(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func hi() string {
			return "hi"
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFuncDeclarationMultipleReturn(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func hi() (int, int) {
			return 6, 6
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFuncDeclarationInvalidMultipleReturn(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func hi() (int, int) {
			return 6, "hi"
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestFuncDeclarationVoidEmptyReturn(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func hi() {
			return
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestFuncDeclarationVoidSingleInvalidReturn(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		func hi() {
			return a
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestAccessBaseContractProperty(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		contract A {
			var b uint256
			constructor(){
				this.balance = this.b
			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}
