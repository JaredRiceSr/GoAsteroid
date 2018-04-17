package parser

import (
	"testing"

	"github.com/benchlab/asteroid/ast"

	"github.com/benchlab/bvmUtils"
)

func TestClassSimpleModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static class OrderBook {

		}
	`)
	bvmUtils.AssertNow(t, a != nil, "ast is nil")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertLength(t, a.Declarations.Length(), 1)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertLength(t, len(c.Modifiers.Modifiers), 2)
}

func TestGroupedClassSimpleModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static (
			class OrderBook {

			}

			class assetType {

			}
		)
	`)
	bvmUtils.AssertNow(t, a != nil, "ast is nil")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
	bvmUtils.AssertLength(t, a.Declarations.Length(), 2)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertLength(t, len(c.Modifiers.Modifiers), 2)
}

func TestGroupedClassMultiLevelModifiers(t *testing.T) {
	a, errs := ParseString(`
		public static (
			class OrderBook {
				var name string
			}

			class assetType {
				var name string
			}
		)
	`)
	bvmUtils.AssertNow(t, a != nil, "ast is nil")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
	bvmUtils.AssertLength(t, a.Declarations.Length(), 2)
	n := a.Declarations.Next()
	c := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertLength(t, len(c.Modifiers.Modifiers), 2)
}
