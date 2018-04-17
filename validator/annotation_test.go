package validator

import (
	"testing"

	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/asteroid/ast"

	"github.com/benchlab/bvmUtils"
	"github.com/benchlab/asteroid/parser"
)

func TestBinaryExpressionAnnotation(t *testing.T) {
	e := parser.ParseExpression("5 == 5")
	v := NewValidator(NewTestBVM())
	v.resolveExpression(e)
	bvmUtils.AssertNow(t, e.Type() == ast.BinaryExpression, "Asteroid Errors: Node Error: Wrong node type. ")
	b := e.(*ast.BinaryExpressionNode)
	bvmUtils.AssertNow(t, b.Resolved != nil, "resolved nil")
	bvmUtils.AssertNow(t, b.Resolved.Compare(typing.Boolean()), "wrong resolved type")
}
