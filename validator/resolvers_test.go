package validator

import (
	"fmt"
	"testing"

	"github.com/benchlab/asteroid/util"

	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/asteroid/parser"

	"github.com/benchlab/bvmUtils"
)

func TestResolveLiteralExpressionBoolean(t *testing.T) {
	v := NewValidator(NewTestBVM())
	p := parser.ParseExpression("true")
	bvmUtils.AssertNow(t, p != nil, "expr should not be nil")
	bvmUtils.AssertNow(t, p.Type() == ast.Literal, "wrong expression type")
	a := p.(*ast.LiteralNode)
	bvmUtils.Assert(t, v.resolveExpression(a).Compare(typing.Boolean()), "wrong true expression type")
}

func TestResolveCallExpression(t *testing.T) {
	v := NewValidator(NewTestBVM())
	fn := new(typing.Func)
	fn.Params = typing.NewTuple(typing.Boolean(), typing.Boolean())
	fn.Results = typing.NewTuple(typing.Boolean())
	v.declareVar(util.Location{}, "hello", &typing.Func{
		Params:  typing.NewTuple(),
		Results: typing.NewTuple(typing.Boolean()),
	})
	p := parser.ParseExpression("hello(5, 5)")
	bvmUtils.AssertNow(t, p.Type() == ast.CallExpression, "wrong expression type")
	a := p.(*ast.CallExpressionNode)
	resolved, ok := v.resolveExpression(a).(*typing.Tuple)
	bvmUtils.Assert(t, ok, "wrong base type")
	bvmUtils.Assert(t, len(resolved.Types) == 1, "wrong type length")
	bvmUtils.Assert(t, fn.Results.Compare(resolved), "should be equal")
}

func TestResolveArrayLiteralExpression(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	p := parser.ParseExpression("[]exchange{}")
	bvmUtils.AssertNow(t, p.Type() == ast.ArrayLiteral, "wrong expression type")
	a := p.(*ast.ArrayLiteralNode)
	_, ok := v.resolveExpression(a).(*typing.Array)
	bvmUtils.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionCopy(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	p := parser.ParseExpression("[]exchange{}[:]")
	bvmUtils.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	bvmUtils.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionLower(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	p := parser.ParseExpression("[]exchange{}[6:]")
	bvmUtils.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	bvmUtils.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionUpper(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	p := parser.ParseExpression("[]exchange{}[:10]")
	bvmUtils.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	bvmUtils.Assert(t, ok, "wrong base type")
}

func TestResolveArrayLiteralSliceExpressionBoth(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	p := parser.ParseExpression("[]exchange{}[6:10]")
	bvmUtils.AssertNow(t, p.Type() == ast.SliceExpression, "wrong expression type")
	_, ok := v.resolveExpression(p).(*typing.Array)
	bvmUtils.Assert(t, ok, "wrong base type")
}

func TestResolveMapLiteralExpression(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	v.declareType(util.Location{}, "cat", typing.Unknown())
	p := parser.ParseExpression("map[exchange]cat{}")
	bvmUtils.AssertNow(t, p.Type() == ast.MapLiteral, "wrong expression type")
	m, ok := v.resolveExpression(p).(*typing.Map)
	bvmUtils.AssertNow(t, ok, "wrong base type")
	bvmUtils.Assert(t, m.Key.Compare(typing.Unknown()), fmt.Sprintf("wrong key: %s", typing.WriteType(m.Key)))
	bvmUtils.Assert(t, m.Value.Compare(typing.Unknown()), fmt.Sprintf("wrong val: %s", typing.WriteType(m.Value)))
}

func TestResolveIndexExpressionArrayLiteral(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareVar(util.Location{}, "cat", typing.Boolean())
	p := parser.ParseExpression("[]cat{}[0]")
	bvmUtils.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	b := p.(*ast.IndexExpressionNode)
	resolved := v.resolveExpression(b)
	typ, _ := v.isTypeVisible("cat")
	bvmUtils.AssertNow(t, resolved.Compare(typ), "wrong expression type")
}

func TestResolveIndexExpressionMapLiteral(t *testing.T) {
	v := NewValidator(NewTestBVM())
	v.declareType(util.Location{}, "exchange", typing.Unknown())
	v.declareType(util.Location{}, "cat", typing.Unknown())
	p := parser.ParseExpression(`map[exchange]cat{}["hi"]`)
	bvmUtils.AssertNow(t, p.Type() == ast.IndexExpression, "wrong expression type")
	typ, _ := v.isTypeVisible("cat")
	ok := v.resolveExpression(p).Compare(typ)
	bvmUtils.AssertNow(t, ok, "wrong type returned")

}

func TestResolveBinaryExpressionSimpleNumeric(t *testing.T) {
	p := parser.ParseExpression("5 + 5")
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	bvmUtils.AssertNow(t, resolved.Compare(v.SmallestInteger(256, false)), fmt.Sprintf("wrong expression type: %s", typing.WriteType(resolved)))
}

func TestResolveBinaryExpressionConcatenation(t *testing.T) {
	p := parser.ParseExpression(`"a" + "b"`)
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	typ, _ := v.isTypeVisible("string")
	bvmUtils.AssertNow(t, resolved.Compare(typ), "wrong expression type")
}

func TestResolveBinaryExpressionEql(t *testing.T) {
	p := parser.ParseExpression("5 == 5")
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	bvmUtils.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionGeq(t *testing.T) {
	p := parser.ParseExpression("5 >= 5")
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	bvmUtils.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionLeq(t *testing.T) {
	p := parser.ParseExpression("5 <= 5")
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	bvmUtils.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionLss(t *testing.T) {
	p := parser.ParseExpression("5 < 5")
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	bvmUtils.AssertNow(t, resolved.Compare(typing.Boolean()), "wrong expression type")
}

func TestResolveBinaryExpressionGtr(t *testing.T) {
	p := parser.ParseExpression("5 > 5")
	bvmUtils.AssertNow(t, p.Type() == ast.BinaryExpression, "wrong b expression type")
	b := p.(*ast.BinaryExpressionNode)
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(b)
	bvmUtils.AssertNow(t, resolved.Compare(typing.Boolean()), fmt.Sprintf("wrong resolved expression type: %s", typing.WriteType(resolved)))
}

func TestResolveBinaryExpressionCast(t *testing.T) {
	p := parser.ParseExpression("uint8(5)")
	bvmUtils.AssertNow(t, p.Type() == ast.CallExpression, "wrong expression type")
	v := NewValidator(NewTestBVM())
	resolved := v.resolveExpression(p)
	bvmUtils.AssertNow(t, len(v.errs) == 0, v.errs.Format())
	typ, _ := v.isTypeVisible("uint8")
	bvmUtils.AssertNow(t, resolved.Compare(typ), fmt.Sprintf("wrong resolved expression type: %s", typing.WriteType(resolved)))
}

func TestReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			func b() string {

			}
		}
		var x string
		var a A
		x = a.b()
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			var b string
		}
		var x string
		var a A
		x = a.b
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			var b []int
		}
		var x int
		var a A
		x = a.b[2]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			var b []int
		}
		var x []int
		var a A
		x = a.b[:]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			func c() string {

			}
		}
		type B A
		var x string
		var b B
		x = b.c()
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			var c string
		}
		type B A
		var x string
		var b B
		x = b.c
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			var c []int
		}
		type B A
		var x int
		var b B
		x = b.c[2]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestAliasedReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class A {
			var c []int
		}
		type B A
		var x []int
		var b B
		x = b.c[:]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			func c() string {

			}
		}
		class A {
			var b B
		}
		var x string
		var a A
		x = a.b.c()
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			var c string
		}
		class A {
			var b B
		}
		var x string
		var a A
		x = a.b.c
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		var x int
		var a A
		x = a.b.c[2]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		var x []int
		var a A
		x = a.b.c[:]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceCallResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			func c() string {

			}
		}
		class A {
			var b B
		}
		type Z A
		var x string
		var z Z
		x = z.b.c()
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceIdentifierResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			var c string
		}
		class A {
			var b B
		}
		type Z A
		var x string
		var z Z
		x = z.b.c
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceIndexResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		type Z A
		var x int
		var z Z
		x = z.b.c[2]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestTripleAliasedReferenceSliceResolution(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class B {
			var c []int
		}
		class A {
			var b B
		}
		type Z A
		var x []int
		var z Z
		x = z.b.c[:]
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInvalidThis(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		this
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestValidThisClassConstructor(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		class OrderBook {

			var name string

			constructor(n string){
				this.name = n
			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestValidThisContractConstructor(t *testing.T) {
	_, errs := ValidateString(NewTestBVM(), `
		contract OrderBook {

			var name string

			constructor(n string){
				this.name = n
			}
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestIsValidMapKey(t *testing.T) {
	a := typing.Boolean()
	bvmUtils.AssertNow(t, !isValidMapKey(a), "boolean valid")
	a = typing.Unknown()
	bvmUtils.AssertNow(t, !isValidMapKey(a), "unknown valid")
	a = typing.Invalid()
	bvmUtils.AssertNow(t, !isValidMapKey(a), "invalid valid")
	a = &typing.NumericType{BitSize: 8, Signed: true}
	bvmUtils.AssertNow(t, isValidMapKey(a), "int8 invalid")
	a = &typing.NumericType{BitSize: 8, Signed: false}
	bvmUtils.AssertNow(t, isValidMapKey(a), "uint8 invalid")
	a = &typing.Array{Value: &typing.NumericType{BitSize: 8, Signed: false}}
	bvmUtils.AssertNow(t, isValidMapKey(a), "uint8 array invalid")
	a = &typing.Array{Value: &typing.NumericType{BitSize: 8, Signed: true}}
	bvmUtils.AssertNow(t, isValidMapKey(a), "int8 array invalid")
	b := &typing.Aliased{Alias: "string", Underlying: a}
	bvmUtils.AssertNow(t, isValidMapKey(b), "aliased int8 array invalid")
}

func TestValidateSimpleCast(t *testing.T) {
	exp, errs := ValidateExpression(NewTestBVM(), `uint(0)`)
	bvmUtils.AssertNow(t, exp != nil, "exp isn't nil")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}
