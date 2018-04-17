package benchvm

import (
	"github.com/benchlab/benchvm"
	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/bvmCreate"
)

// An Commander burns down trees
type Commander struct {
	VM *bvmCreate.VM
}

// TraverseExpression ...
func (a *Commander) TraverseExpression(node ast.ExpressionNode) {
	switch node.Type() {
	case ast.IndexExpression:
		a.traverseIndex(node.(ast.IndexExpressionNode))
		break
	case ast.CallExpression:
		a.traverseCallExpr(node.(ast.CallExpressionNode))
		break
	case ast.BinaryExpression:
		a.traverseBinaryExpr(node.(ast.BinaryExpressionNode))
		break
	case ast.UnaryExpression:
		a.traverseUnaryExpr(node.(ast.UnaryExpressionNode))
		break
	case ast.Literal:
		a.traverseLiteral(node.(ast.LiteralNode))
		break
	case ast.CompositeLiteral:
		a.traverseCompositeLiteral(node.(ast.CompositeLiteralNode))
		break
	case ast.SliceExpression:
		a.traverseSliceExpression(node.(ast.SliceExpressionNode))
		break
	case ast.ArrayLiteral:
		a.traverseArrayLiteral(node.(ast.ArrayLiteralNode))
		break
	case ast.MapLiteral:
		a.traverseMapLiteral(node.(ast.MapLiteralNode))
		break
	}
}

func (a Commander) AddBytecode(op string, params ...byte) {
	t.VM.AddBytecode(op, params...)
}

func NewTraverser() Commander {
	return Commander{}
}

// Traverse ...
func (a Commander) Traverse(node ast.Node) {
	// initialise the vm
	if a.BVM == nil {
		a.BVM = benchvm.NewVM()
	}
	switch node.Type() {
	case ast.ClassDeclaration:
		a.traverseClass(node.(ast.ClassDeclarationNode))
		break
	case ast.InterfaceDeclaration:
		a.traverseInterface(node.(ast.InterfaceDeclarationNode))
		break
	case ast.EnumDeclaration:
		a.traverseEnum(node.(ast.EnumDeclarationNode))
		break
	case ast.EventDeclaration:
		a.traverseEvent(node.(ast.EventDeclarationNode))
		break
	case ast.TypeDeclaration:
		a.traverseType(node.(ast.TypeDeclarationNode))
		break
	case ast.ContractDeclaration:
		a.traverseContract(node.(ast.ContractDeclarationNode))
		break
	case ast.FuncDeclaration:
		a.traverseFunc(node.(ast.FuncDeclarationNode))
		break
	case ast.LifecycleDeclaration:
		a.traverseLifecycle(node.(ast.LifecycleDeclarationNode))
		break

	case ast.ForStatement:
		a.traverseForStatement(node.(ast.ForStatementNode))
		break
	case ast.AssignmentStatement:
		a.traverseAssignmentStatement(node.(ast.AssignmentStatementNode))
		break
	case ast.CaseStatement:
		a.traverseCaseStatement(node.(ast.CaseStatementNode))
		break
	case ast.ReturnStatement:
		a.traverseReturnStatement(node.(ast.ReturnStatementNode))
		break
	case ast.IfStatement:
		a.traverseIfStatement(node.(ast.IfStatementNode))
		break
	case ast.SwitchStatement:
		a.traverseSwitchStatement(node.(ast.SwitchStatementNode))
		break
	}
}
