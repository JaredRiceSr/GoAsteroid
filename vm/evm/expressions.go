package evm

import (
	"fmt"

	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/asteroid/token"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/asteroid/ast"
)

func (e *AsteroidEVM) traverseValue(n ast.ExpressionNode) (code bvmCreate.Bytecode) {
	return code
}

func (e *AsteroidEVM) traverseLocation(n ast.ExpressionNode) (code bvmCreate.Bytecode) {
	return code
}

func (e *AsteroidEVM) traverseExpression(n ast.ExpressionNode) (code bvmCreate.Bytecode) {
	switch node := n.(type) {
	case *ast.ArrayLiteralNode:
		return e.traverseArrayLiteral(node)
	case *ast.FuncLiteralNode:
		return e.traverseFuncLiteral(node)
	case *ast.MapLiteralNode:
		return e.traverseMapLiteral(node)
	case *ast.CompositeLiteralNode:
		return e.traverseCompositeLiteral(node)
	case *ast.UnaryExpressionNode:
		return e.traverseUnaryExpr(node)
	case *ast.BinaryExpressionNode:
		return e.traverseBinaryExpr(node)
	case *ast.CallExpressionNode:
		return e.traverseCallExpr(node)
	case *ast.IndexExpressionNode:
		return e.traverseIndex(node)
	case *ast.SliceExpressionNode:
		return e.traverseSliceExpression(node)
	case *ast.IdentifierNode:
		return e.traverseIdentifier(node)
	case *ast.ReferenceNode:
		return e.traverseReference(node)
	case *ast.LiteralNode:
		return e.traverseLiteral(node)
	}
	return code
}

func (e *AsteroidEVM) traverseArrayLiteral(n *ast.ArrayLiteralNode) (code bvmCreate.Bytecode) {
	/*
		// encode the size first
		code.Add(uintAsBytes(len(n.Data))...)

		for _, expr := range n.Data {
			code.Concat(e.traverseExpression(expr))
		}*/

	return code
}

func (e *AsteroidEVM) traverseSliceExpression(n *ast.SliceExpressionNode) (code bvmCreate.Bytecode) {
	// evaluate the original expression first

	// get the data
	code.Concat(e.traverseExpression(n.Expression))

	// ignore the first (item size) * lower

	// ignore the ones after

	return code
}

func (e *AsteroidEVM) traverseCompositeLiteral(n *ast.CompositeLiteralNode) (code bvmCreate.Bytecode) {
	/*
		var ty validator.Class
		for _, f := range ty.Fields {
			if n.Fields[f] != nil {

			} else {
				n.Fields[f].Size()
			}
		}

		for _, field := range n.Fields {
			// evaluate each field
			code.Concat(e.traverseExpression(field))
		}*/

	return code
}

var binaryOps = map[token.Type]BinaryOperator{
	token.Add:        additionOrConcatenation,
	token.Sub:        simpleOperator("SUB"),
	token.Mul:        simpleOperator("MUL"),
	token.Div:        signedOperator("DIV", "SDIV"),
	token.Mod:        signedOperator("MOD", "SMOD"),
	token.Shl:        simpleOperator("SHL"),
	token.Shr:        simpleOperator("SHR"),
	token.And:        simpleOperator("AND"),
	token.Or:         simpleOperator("OR"),
	token.Xor:        simpleOperator("XOR"),
	token.As:         ignoredOperator(),
	token.Gtr:        signedOperator("GT", "SGT"),
	token.Lss:        signedOperator("LT", "SLT"),
	token.Eql:        simpleOperator("EQL"),
	token.Neq:        reversedOperator("EQL"),
	token.Geq:        reversedSignedOperator("LT", "SLT"),
	token.Leq:        reversedSignedOperator("GT", "SGT"),
	token.LogicalAnd: ignoredOperator(),
	token.LogicalOr:  ignoredOperator(),
}

type BinaryOperator func(n *ast.BinaryExpressionNode) bvmCreate.Bytecode

func reversedSignedOperator(unsigned, signed string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
		left, lok := typing.ResolveUnderlying(n.Left.ResolvedType()).(*typing.NumericType)
		right, rok := typing.ResolveUnderlying(n.Left.ResolvedType()).(*typing.NumericType)
		if lok && rok {
			if left.Signed || right.Signed {
				code.Add(signed)
			} else {
				code.Add(unsigned)
			}
			code.Add("NOT")
		}
		return code
	}
}

func reversedOperator(mnemonic string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
		code.Add(mnemonic)
		code.Add("NOT")
		return code
	}
}

func ignoredOperator() BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
		return code
	}
}

func simpleOperator(mnemonic string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
		code.Add(mnemonic)
		return code
	}
}

func signedOperator(unsigned, signed string) BinaryOperator {
	return func(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
		//fmt.Println(typing.WriteType(n.Left.ResolvedType()))
		left, lok := typing.ResolveUnderlying(n.Left.ResolvedType()).(*typing.NumericType)
		right, rok := typing.ResolveUnderlying(n.Right.ResolvedType()).(*typing.NumericType)
		if lok && rok {
			if left.Signed || right.Signed {
				code.Add(signed)
			} else {
				code.Add(unsigned)
			}
		}
		return code
	}
}

func additionOrConcatenation(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
	switch n.Resolved.(type) {
	case *typing.NumericType:
		code.Add("ADD")
		return code
	default:
		// otherwise must be a string
		return code

	}
}

func (e *AsteroidEVM) traverseBinaryExpr(n *ast.BinaryExpressionNode) (code bvmCreate.Bytecode) {
	/* alter stack:

	| Operand 1 |
	| Operand 2 |
	| Operator  |

	Note that these operands may contain further expressions of arbitrary depth.
	*/
	code.Concat(e.traverseExpression(n.Left))
	code.Concat(e.traverseExpression(n.Right))

	code.Concat(binaryOps[n.Operator](n))

	return code
}

var unaryOps = map[token.Type]string{
	token.Not: "NOT",
}

func (e *AsteroidEVM) traverseUnaryExpr(n *ast.UnaryExpressionNode) (code bvmCreate.Bytecode) {
	/* alter stack:

	| Expression 1 |
	| Operand      |

	Note that these expressions may contain further expressions of arbitrary depth.
	*/
	code.Concat(e.traverseExpression(n.Operand))
	code.Add(unaryOps[n.Operator])
	return code
}

func (e *AsteroidEVM) traverseContractCall(n *ast.CallExpressionNode) (code bvmCreate.Bytecode) {
	// calls another contract, The CREATE opcode takes three values:
	// value (ie. initial amount of ether),
	code.Add("GAS")
	// machineMemory start
	code.Add("PUSH")
	// machineMemory length
	code.Add("PUSH")
	code.Add("CREATE")
	return code
}

func (e *AsteroidEVM) traverseClassCall(n *ast.CallExpressionNode) (code bvmCreate.Bytecode) {
	return code
}

func (e *AsteroidEVM) traverseFunctionCall(n *ast.CallExpressionNode) (code bvmCreate.Bytecode) {
	for _, arg := range n.Arguments {
		code.Concat(e.traverseExpression(arg))
	}

	// traverse the call expression
	// should leave the function address on top of the stack

	// need to get annotations through this process

	if n.Call.Type() == ast.Identifier {
		i := n.Call.(*ast.IdentifierNode)
		if b, ok := builtins[i.Name]; ok {
			code.Concat(b(e))
			return code
		}
	}

	call := e.traverse(n.Call)

	code.Concat(call)

	// parameters are at the top of the stack
	// jump to the top of the function
	return code
}

func (e *AsteroidEVM) traverseCallExpr(n *ast.CallExpressionNode) (code bvmCreate.Bytecode) {
	e.expression = n

	switch typing.ResolveUnderlying(n.Call.ResolvedType()).(type) {
	case *typing.Func:
		return e.traverseFunctionCall(n)
	case *typing.Contract:
		return e.traverseContractCall(n)
	case *typing.Class:
		return e.traverseClassCall(n)
	}
	return code
}

func (e *AsteroidEVM) traverseLiteral(n *ast.LiteralNode) (code bvmCreate.Bytecode) {

	// Literal Nodes are directly converted to push specifications
	// these nodes must be divided into blocks of 16 bytes
	// in order to maintain

	// maximum number size is 256 bits (32 bytes)
	switch n.LiteralType {
	case token.Integer, token.Float:
		if len(n.Data) > 32 {
			// error
		} else {
			fmt.Println("HERE")
			// TODO: type size or data size
			code.Add(fmt.Sprintf("PUSH%d", bytesRequired(int(n.Resolved.Size()))), []byte(n.Data)...)
		}
		break
	case token.String:
		bytes := []byte(n.Data)
		max := 32
		size := 0
		for size = len(bytes); size > max; size -= max {
			code.Add("PUSH32", bytes[len(bytes)-size:len(bytes)-size+max]...)
		}
		op := fmt.Sprintf("PUSH%d", size)
		code.Add(op, bytes[size:len(bytes)]...)
		break
	}
	return code
}

func (e *AsteroidEVM) traverseIndex(n *ast.IndexExpressionNode) (code bvmCreate.Bytecode) {

	// TODO: bounds checking?

	// load the data
	code.Concat(e.traverseExpression(n.Expression))

	typ := n.Expression.ResolvedType()

	// calculate offset
	// evaluate index
	code.Concat(e.traverseExpression(n.Index))
	// get size of type
	code.Concat(push(encodeUint(typ.Size())))
	// offset = size of type * index
	code.Add("MUL")

	code.Add("ADD")

	return code
}

const (
	mapLiteralReserved   = "gevm_map_literal_%d"
	arrayLiteralReserved = "gevm_array_literal_%d"
)

// reserve name for map literals: gevm_map_literal_{count}
// if you try to name things it won't let you
func (e *AsteroidEVM) traverseMapLiteral(n *ast.MapLiteralNode) (code bvmCreate.Bytecode) {

	fakeKey := fmt.Sprintf(mapLiteralReserved, e.mapLiteralCount)

	// must be deterministic iteration here
	for _, v := range n.Data {
		// each machineStorage slot must be 32 bytes regardless of contents
		slot := EncodeName(fakeKey + e.traverse(n))
		code.Concat(push(slot))
		code.Add("SSTORE")
	}

	e.mapLiteralCount++

	return code
}

func (e *AsteroidEVM) traverseFuncLiteral(n *ast.FuncLiteralNode) (code bvmCreate.Bytecode) {
	// create an internal hook

	// parameters should have been pushed onto the stack by the caller
	// take them off and put them in machineMemory
	for _, p := range n.Parameters {
		for _, i := range p.Identifiers {
			e.allocateMachineMemory(i, p.Resolved.Size())
			code.Add("MSTORE")
		}
	}

	code.Concat(e.traverseScope(n.Scope))

	for _, p := range n.Parameters {
		for _, i := range p.Identifiers {
			e.freeMachineMemory(i)
		}
	}

	return code
}

func (e *AsteroidEVM) traverseIdentifier(n *ast.IdentifierNode) (code bvmCreate.Bytecode) {

	if e.inMachineStorage {
		s := e.lookupMachineStorage(n.Name)
		if s != nil {
			return s.retrieve()
		}
		e.allocateStorageToMachine(n.Name, n.Resolved.Size())
		s = e.lookupMachineStorage(n.Name)
		if s != nil {
			return s.retrieve()
		}
	} else {
		m := e.lookupMachineMemory(n.Name)
		if m != nil {
			return m.retrieve()
		}
		e.allocateMachineMemory(n.Name, n.Resolved.Size())
		m = e.lookupMachineMemory(n.Name)
		if m != nil {
			return m.retrieve()
		}
	}
	return code
}

func (e *AsteroidEVM) traverseContextual(t typing.Type, expr ast.ExpressionNode) (code bvmCreate.Bytecode) {
	switch expr.(type) {
	case *ast.IdentifierNode:
		switch t.(type) {
		case *typing.Class:

			break
		case *typing.Interface:
			break
		}
		break
	case *ast.CallExpressionNode:
		break
	}
	return code
}

func (e *AsteroidEVM) traverseReference(n *ast.ReferenceNode) (code bvmCreate.Bytecode) {
	code.Concat(e.traverse(n.Parent))

	resolved := n.Parent.ResolvedType()

	ctx := e.traverseContextual(resolved, n.Reference)

	code.Concat(ctx)

	if e.inMachineStorage {
		code.Add("SLOAD")
	} else {
		code.Add("MLOAD")

	}

	// reference e.g. exchange.tail.wag()
	// get the object
	/*if n.InStorage {
		// if in machineStorage
		// only the top level name is accessible in machineStorage
		// everything else is accessed
		e.AddBytecode("PUSH", len(n.Names[0]), n.Names[0])
		e.AddBytecode("LOAD")

		// now get the sub-references
		// e.AddBytecode("", params)
	} else {
		e.AddBytecode("GET")
	}*/
	return code
}
