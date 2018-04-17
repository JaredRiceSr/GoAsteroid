package evm

import (
	"fmt"
	"testing"

	"github.com/benchlab/asteroid/validator"

	"github.com/benchlab/bvmUtils"
)

func TestTraverseExplicitVariableDeclaration(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `var name uint8`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.machineStorage)))
}

func TestTraverseExplicitVariableDeclarationFunc(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `var name func(a, b string) int`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.machineStorage)))
}

func TestTraverseExplicitVariableDeclarationFixedArray(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `var name [3]uint8`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.machineStorage)))
}

func TestTraverseExplicitVariableDeclarationVariableArray(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `var name [3]uint8`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 1, fmt.Sprintf("didn't allocate a block: %d", len(e.machineStorage)))
}

func TestTraverseTypeDeclaration(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `type OrderBook int`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 0, fmt.Sprintf("allocate a block: %d", len(e.machineStorage)))
}

func TestTraverseClassDeclaration(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `class OrderBook {}`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 0, fmt.Sprintf("allocate a block: %d", len(e.machineStorage)))
}

func TestTraverseInterfaceDeclaration(t *testing.T) {
	e := NewBVM()
	ast, errs := validator.ValidateString(e, `interface OrderBook {}`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, ast != nil, "ast shouldn't be nil")
	bvmUtils.AssertNow(t, ast.Declarations != nil, "ast decls shouldn't be nil")
	e.Traverse(ast)
	bvmUtils.Assert(t, len(e.machineStorage) == 0, fmt.Sprintf("allocate a block: %d", len(e.machineStorage)))
}
