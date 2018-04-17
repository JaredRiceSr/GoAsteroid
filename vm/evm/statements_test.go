package evm

import (
	"testing"

	"github.com/benchlab/asteroid/validator"

	"github.com/benchlab/asteroid/ast"

	"github.com/benchlab/bvmUtils"
)

func TestIncrement(t *testing.T) {

}

func TestSimpleAssignmentStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        i = 0
    `)
	f := scope.Sequence[0].(*ast.AssignmentStatementNode)
	bytecode := e.traverseAssignmentStatement(f)
	expected := []string{
		// push left
		"PUSH",
		// push right
		"PUSH",
		// store (default is machineMemory)
		"MSTORE",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestIndexAssignmentStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        nums [5]int
        nums[3] = 0
    `)

	bytecode := e.traverse(scope)
	expected := []string{
		// push left
		"PUSH",
		// push right
		"PUSH",
		// store (default is machineMemory)
		"MSTORE",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestIfStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        if x = 0; x > 5 {

        }
    `)
	f := scope.Sequence[0].(*ast.IfStatementNode)
	bytecode := e.traverseIfStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "BOUNCEI",
		// loop body
		"BOUNCEDEST",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

/*
func TestElseIfStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        if x = 0; x > 5 {
            x = 1
        } else if x < 3 {
            x = 2
        }
    `)
	f := scope.Sequence[0].(*ast.IfStatementNode)
	bytecode := e.traverseIfStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "BOUNCEI",
		// if body
		"PUSH", "PUSH", "MSTORE",
		// else if condition
		"PUSH", "PUSH", "LT",
		// jumper
		"PUSH", "BOUNCEI",
		// else if body
		"PUSH", "PUSH", "MSTORE",
		"BOUNCEDEST",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
*/
func TestElseStatement(t *testing.T) {
	e := NewBVM()

	validator.ValidateString(e, `
        if x = 0; x > 5 {

        } else {

        }
    `)
	/*f := scope.Sequence[0].(*ast.IfStatementNode)
	bytecode := e.traverseIfStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "BOUNCEI",
		// if body
		// else
		"BOUNCEDEST",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())*/
}

/*
func TestForStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        for i = 0; i < 5; i++ {

        }
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
	bytecode := e.traverseForStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"BOUNCEDEST",
		// condition
		"PUSH", "MLOAD", "PUSH", "LT", "PUSH", "BOUNCEI",
		// body
		// post
		"PUSH", "MLOAD", "PUSH", "ADD", "PUSH", "MSTORE",
		// jump back to top
		"BOUNCE",
		"BOUNCEDEST",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
*/
func TestReturnStatement(t *testing.T) {

}

func TestBreakStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        for x = 0; x < 5; x++ {
			if x == 3 {
				break
			}
		}
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
	bytecode := e.traverseForStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "BOUNCEI",
		// loop body
		// if statement
		"PUSH", "MLOAD", "PUSH", "EQ", "ISZERO", "BOUNCEI",
		"BOUNCE",
		"BOUNCEDEST",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func Test2DLoopBreakStatement(t *testing.T) {
	e := NewBVM()
	scope, _ := validator.ValidateString(e, `
        for x = 0; x < 5; x++ {
			for y = 0; y < 5; y++ {
				if x + y == 3 {
					break
				}
			}
		}
    `)
	f := scope.Sequence[0].(*ast.ForStatementNode)
	bytecode := e.traverseForStatement(f)
	expected := []string{
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "BOUNCEI",
		// loop body
		// init
		"PUSH", "PUSH", "MSTORE",
		// top of loop
		"PUSH", "PUSH", "GT",
		// jumper
		"PUSH", "BOUNCEI",
		// if statement
		"PUSH", "MLOAD", "PUSH", "MLOAD,", "ADD", "PUSH", "EQ", "ISZERO", "BOUNCEI",
		// break statement
		"BOUNCE",
		"BOUNCEDEST",
		"BOUNCEDEST",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
