package benchvm

import (
	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/asteroid/lexer"
)

func (a *Arsonist) traverseArrayLiteral(n ast.ArrayLiteralNode) {

}

func (a *Arsonist) traverseSliceExpression(n ast.SliceExpressionNode) {

}

func (a *Arsonist) traverseCompositeLiteral(n ast.CompositeLiteralNode) {

}

var binaryOps = map[token.Type]string{
	token.Add: "ADD",
	token.Sub: "SUB",
	token.Mul: "MUL",
	token.Div: "DIV",
	token.Mod: "MOD",
	token.Shl: "SHL",
	token.Shr: "SHR",
	token.And: "AND",
	token.Or:  "OR",
	token.Xor: "XOR",
}

func (a *Arsonist) traverseBinaryExpr(n ast.BinaryExpressionNode) {
	a.TraverseExpression(n.Left)
	a.TraverseExpression(n.Right)
	// operation
	a.BVM.AddBytecode(binaryOps[n.Operator])
}

var unaryOps = map[token.Type]string{
	token.Not: "NOT",
}

func (a *Arsonist) traverseUnaryExpr(n ast.UnaryExpressionNode) {
	a.BVM.AddBytecode(unaryOps[n.Operator])
	a.Traverse(n.Operand)
}

func (a *Arsonist) traverseCallExpr(n ast.CallExpressionNode) {
	for _, arg := range n.Arguments {
		a.Traverse(arg)
	}
	// parameters are at the top of the stack
	// jump to the top of the function
}

func (a *Arsonist) traverseLiteral(n ast.LiteralNode) {
	// Literal Nodes are directly converted to push instructions
	var parameters []byte
	bytes := n.GetBytes()
	parameters = append(parameters, byte(len(bytes)))
	parameters = append(parameters, bytes...)
	a.BVM.AddBytecode("PUSH", parameters...)
}

func (a *Arsonist) traverseIndex(n ast.IndexExpressionNode) {
	// evaluate the index
	a.Traverse(n.Index)
	// then MLOAD it at the index offset
	a.Traverse(n.Expression)
	a.BVM.AddBytecode("GET")
}

func (a *Arsonist) traverseMapLiteral(n ast.MapLiteralNode) {
	for k, v := range n.Data {
		a.Traverse(k)
		a.Traverse(v)
	}
	// push the size of the map
	a.BVM.AddBytecode("PUSH", byte(1), byte(len(n.Data)))
	a.BVM.AddBytecode("MAP")
}

func (a *Arsonist) traverseReference(n ast.ReferenceNode) {

	// reference e.g. dog.tail.wag()
	// get the object
	/*if n.InStorage {
		// if in storage
		// only the top level name is accessible in storage
		// everything else is accessed
		//a.BVM.AddBytecode("PUSH", len(n.Names[0]), n.Names[0])
		a.BVM.AddBytecode("LOAD")

		// now get the sub-references
		// a.BVM.AddBytecode("", params)
	} else {
		a.BVM.AddBytecode("GET")
	}*/
}
