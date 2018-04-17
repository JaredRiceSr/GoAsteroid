package evm

import (
	"testing"

	"github.com/benchlab/asteroid/validator"

	"github.com/benchlab/bvmUtils"
)

func TestTraverseIdentifierExpression(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		hello = 5
		x = hello
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteralsBinaryExpression(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		x = 1 + 2
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIdentifierBinaryExpression(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		a = 1
		b = 2
		x = a + b
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseCallBinaryExpression(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `

		func a() int {
			return 1
		}

		func b() int {
			return 2
		}

		x = a() + b()
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierLiteral(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		var b [5]int
		x = b[1]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push a
		"PUSH",
		// push index
		"PUSH",
		// push size of int
		"PUSH",
		// calculate offset
		"MUL",
		// create final position
		"ADD",
		"SLOAD",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseTwoDimensionalArray(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		b [5][5]int
		x = b[2][3]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierIdentifier(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		b [5]int
		x = b[a]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierCall(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		b [5]int
		x = b[a()]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionIdentifierIndex(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		b [5]int
		a [4]int
		c = 3
		x = b[a[c]]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionCallIdentifier(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		x = a()[b]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionCallLiteral(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		x = a()[1]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseIndexExpressionCallCall(t *testing.T) {
	e := NewBVM()
	a, _ := validator.ValidateString(e, `
		x = a()[b()]
	`)
	bytecode, _ := e.Traverse(a)
	expected := []string{
		// push x
		"PUSH",
		// push b
		"PUSH",
		// push index (2)
		"PUSH",
		// push size of type
		"PUSH",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteral(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "0")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteralTwoBytes(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "256")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH2"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTraverseLiteralThirtyTwoBytes(t *testing.T) {
	e := new(AsteroidEVM)
	// 2^256
	expr, _ := validator.ValidateExpression(e, "115792089237316195423570985008687907853269984665640564039457584007913129639936")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH32"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySignedLess(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 < 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SLT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySignedLessEqual(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 <= 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SGT", "NOT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySignedGreater(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 < 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SLT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySignedGreaterEqual(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 >= 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SLT", "NOT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryEqual(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 == 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "EQ"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryNotEqual(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 == 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "EQ", "NOT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryAnd(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 & 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "AND"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryOr(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 | 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "OR"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryXor(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 ^ 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "XOR"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryLogicalAnd(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "true and false")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "OR"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryLogicalOr(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "true or false")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "OR"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryAddition(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 + 5")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "ADD"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySubtraction(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 - 5")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SUB"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryMultiplication(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "3 * 5")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "MUL"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySignedDivision(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "4 / 2")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "DIV"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryDivision(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "4 as uint / 2 as uint")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "DIV"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinarySignedMod(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "4 % 2")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "SMOD"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryMod(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "4 as uint % 2 as uint")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "MOD"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBinaryExp(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "2 ** 4")
	bytecode := e.traverseExpression(expr)
	expected := []string{"PUSH1", "PUSH1", "EXP"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestEmptyConstructorCall(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateString(e, `
		contract OrderBook {

		}

		d = OrderBook()
	`)
	bytecode, _ := e.Traverse(expr)
	expected := []string{"PUSH1", "PUSH1", "PUSH1", "CREATE"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestLiteralParameterConstructorCall(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateString(e, `
		contract OrderBook {

			var name string

			constructor(n string){
				this.name = n
			}

		}

		d = OrderBook("Alex")
	`)
	bytecode, _ := e.Traverse(expr)
	expected := []string{"PUSH1", "PUSH1", "PUSH1", "CREATE"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
