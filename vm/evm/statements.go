package evm

import (
	"github.com/benchlab/asteroid/typing"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/asteroid/ast"
)

func (e *AsteroidEVM) traverseSwitchStatement(n *ast.SwitchStatementNode) (code bvmCreate.Bytecode) {
	// always traverse the target
	var generatedExpressions []bvmCreate.Bytecode
	if n.Target != nil {
		e.traverse(n.Target)
		for _, c := range n.Cases.Sequence {
			switch cas := c.(type) {
			case *ast.CaseStatementNode:
				for _, exp := range cas.Expressions {
					generatedExpressions = append(generatedExpressions, e.traverseExpression(exp))
				}
				break
			}
		}
		sum := 0
		for _, exp := range generatedExpressions {
			sum += 1 // for duplicating the top of the stack
			sum += exp.Length()
			sum += 3 // for EQ, PUSH, BOUNCE
		}
		for _, exp := range generatedExpressions {
			code.Concat(exp)
			code.Concat(pushMarker(sum))
			code.Add("EQ")
			code.Add("BOUNCEI")
		}
	} else {
		for _, c := range n.Cases.Sequence {
			switch cas := c.(type) {
			case *ast.CaseStatementNode:
				for _, exp := range cas.Expressions {
					generatedExpressions = append(generatedExpressions, e.traverseExpression(exp))
				}
				break
			}
		}
		sum := 0
		for _, exp := range generatedExpressions {
			sum += 1 // for duplicating the top of the stack
			sum += exp.Length()
			sum += 3 // for EQ, PUSH, BOUNCE
		}
		for _, exp := range generatedExpressions {
			code.Concat(exp)
			code.Concat(pushMarker(sum))
			code.Add("BOUNCEI")
		}
	}

	// switch statements are implicitly converted to if statements
	// may be a better way to do this
	// Solidity doesn't have a switch so shrug
	return code
}

func (e *AsteroidEVM) traverseCaseStatement(n *ast.CaseStatementNode) (code bvmCreate.Bytecode) {
	return code
}

func (e *AsteroidEVM) traverseForStatement(n *ast.ForStatementNode) (code bvmCreate.Bytecode) {

	// init statement
	// jumpdest
	// condition
	// jump to end
	// regular loop processes would occur here
	// post statement
	// jump back to the top of the loop
	// jumpdest
	// continue after the loop

	init := e.traverse(n.Init)

	cond := e.traverseExpression(n.Cond)

	block := e.traverseScope(n.Block)

	post := e.traverse(n.Post)

	code.Concat(init)
	code.Add("BOUNCEDEST")
	top := code.Length()
	code.Concat(cond)
	code.Concat(pushMarker(block.Length() + post.Length() + 1))
	code.Add("BOUNCEI")
	code.Concat(block)
	code.Concat(post)
	bottom := code.Length()
	code.Concat(pushMarker(bottom - top))
	code.Add("BOUNCE")
	code.Add("BOUNCEDEST")

	return code
}

func (evm *AsteroidEVM) increment(varName string) (code bvmCreate.Bytecode) {
	code.Concat(push([]byte(varName)))
	code.Add("DUP1")
	code.Add("MLOAD")
	code.Concat(push(uintAsBytes(1)))
	code.Add("ADD")
	code.Add("MSTORE")
	return code
}

func (e *AsteroidEVM) traverseForEachStatement(n *ast.ForEachStatementNode) (code bvmCreate.Bytecode) {
	// starting from index 0
	// same for
	// allocate machineMemory for the index
	// NOTE:
	// can potentially support dmaps by encoding a backing array as well
	// would be more expensive - add a keyword?

	switch n.ResolvedType.(type) {
	case *typing.Array:
		// TODO: index size must be large enough for any vars
		name := n.Variables[0]
		e.allocateMachineMemory(name, 10)
		memloc := e.lookupMachineMemory(name).offset
		increment := e.increment(name)
		block := e.traverseScope(n.Block)
		code.Concat(push(encodeUint(0)))

		code.Concat(push(encodeUint(memloc)))
		code.Add("MSTORE")
		code.Add("BOUNCEDEST")
		code.Concat(push(encodeUint(memloc)))
		code.Add("MLOAD")
		code.Add("LT")
		code.Concat(pushMarker(1 + increment.Length() + block.Length()))
		code.Add("BOUNCEI")

		code.Concat(block)
		code.Concat(increment)

		code.Add("BOUNCEDEST")
		break
	case *typing.Map:
		break
	}
	return code
}

func (e *AsteroidEVM) traverseReturnStatement(n *ast.ReturnStatementNode) (code bvmCreate.Bytecode) {
	for _, r := range n.Results {
		// leave each of them on the stack in turn
		code.Concat(e.traverse(r))
	}
	// jump back to somewhere
	// top of stack should now be return address
	code.Add("BOUNCE")
	return code
}

func (e *AsteroidEVM) traverseControlFlowStatement(n *ast.FlowStatementNode) (code bvmCreate.Bytecode) {
	return code
}

func (e *AsteroidEVM) traverseIfStatement(n *ast.IfStatementNode) (code bvmCreate.Bytecode) {
	conds := make([]bvmCreate.Bytecode, len(n.Conditions))
	blocks := make([]bvmCreate.Bytecode, len(n.Conditions))
	end := 0

	for _, c := range n.Conditions {
		cond := e.traverse(c.Condition)
		conds = append(conds, cond)
		body := e.traverseScope(c.Body)
		blocks = append(blocks, body)
		end += cond.Length() + body.Length() + 3 + 1
	}

	for i := range n.Conditions {
		code.Concat(conds[i])
		code.Add("ISZERO")
		code.Concat(pushMarker(blocks[i].Length() + 1))
		code.Add("BOUNCEI")
		code.Concat(blocks[i])
		code.Concat(pushMarker(end))
	}

	code.Concat(e.traverseScope(n.Else))

	return code
}

func (e *AsteroidEVM) traverseAssignmentStatement(n *ast.AssignmentStatementNode) (code bvmCreate.Bytecode) {
	for i, l := range n.Left {
		r := n.Right[i]
		code.Concat(e.assign(l, r, e.inMachineStorage))
	}
	return code
}

func (e *AsteroidEVM) assign(l, r ast.ExpressionNode, inStorage bool) (code bvmCreate.Bytecode) {
	// get the location
	//code.Concat(e.traverseExpression(l))
	// do the calculation
	code.Concat(e.traverseExpression(r))
	if inStorage {
		code.Add("SSTORE")
	} else {
		code.Add("MSTORE")
	}
	return code
}
