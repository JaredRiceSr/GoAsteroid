package token

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestIsUnaryOperator(t *testing.T) {
	x := Not
	bvmUtils.Assert(t, x.IsUnaryOperator(), "Asteroid Token Error: should not be unary")
	x = Add
	bvmUtils.Assert(t, !x.IsUnaryOperator(), "Asteroid Token Error: add should not be unary")
}

func TestIsBinaryOperator(t *testing.T) {
	x := Add
	bvmUtils.Assert(t, x.IsBinaryOperator(), "Asteroid Token Message: add should be binary")
	x = Not
	bvmUtils.Assert(t, !x.IsBinaryOperator(), "Asteroid Token Message: not should not be binary")
}

func TestIsIdentifierByte(t *testing.T) {
	bvmUtils.Assert(t, isIdentifierByte('A'), "Asteroid Token Message: upper letter not id byte")
	bvmUtils.Assert(t, isIdentifierByte('t'), "Asteroid Token Message: lower letter not id byte")
	bvmUtils.Assert(t, isIdentifierByte('2'), "Asteroid Token Message: number not id byte")
	bvmUtils.Assert(t, isIdentifierByte('_'), "Asteroid Token Message: underscore not id byte")
	bvmUtils.Assert(t, !isIdentifierByte(' '), "Asteroid Token Message: space should not be id byte")
}

func TestIsNumber(t *testing.T) {
	byt := []byte(`9`)
	b := &bytecode{bytes: byt}
	bvmUtils.Assert(t, isInteger(b), "Asteroid Token Message: positive integer")
	byt = []byte(`-9`)
	b = &bytecode{bytes: byt}
	bvmUtils.Assert(t, isInteger(b), "Asteroid Token Message: negative integer")
	byt = []byte(`9.0`)
	b = &bytecode{bytes: byt}
	bvmUtils.Assert(t, isFloat(b), "Asteroid Token Message: positive float")
	byt = []byte(`-9.0`)
	b = &bytecode{bytes: byt}
	bvmUtils.Assert(t, isFloat(b), "Asteroid Token Message: negative float")
	byt = []byte(`-.9`)
	b = &bytecode{bytes: byt}
	bvmUtils.Assert(t, isFloat(b), "Asteroid Token Message: negative float")
}

func TestIsWhitespace(t *testing.T) {
	byt := []byte(` `)
	b := &bytecode{bytes: byt}
	bvmUtils.Assert(t, isWhitespace(b), "space")

	byt = []byte(`	`)
	b = &bytecode{bytes: byt}
	bvmUtils.Assert(t, isWhitespace(b), "tab")
}

func TestIsAssignment(t *testing.T) {
	a := Assign
	bvmUtils.AssertNow(t, a.IsAssignment(), "Asteroid Token Message: assign not assignment")
}

func TestIsDeclaration(t *testing.T) {
	a := Event
	bvmUtils.AssertNow(t, a.IsDeclaration(), "Asteroid Token Message: event not decl")
}
