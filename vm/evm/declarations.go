package evm

import (
	"fmt"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/asteroid/ast"
)

func (e *AsteroidEVM) traverseType(n *ast.TypeDeclarationNode) (code bvmCreate.Bytecode) {
	// do nothing
	return code
}

func (e *AsteroidEVM) traverseClass(n *ast.ClassDeclarationNode) (code bvmCreate.Bytecode) {
	// create constructor hooks
	// create function hooks
	if n.Body.Declarations != nil {
		for _, d := range n.Body.Declarations.Map() {
			switch a := d.(type) {
			case *ast.ExplicitVarDeclarationNode:
				e.traverseExplicitVarDecl(a)
				break
			case *ast.LifecycleDeclarationNode:
				//e.addLifecycleHook(n.Identifier, a)
				break
			case *ast.FuncDeclarationNode:
				e.traverseFunc(a)
				break
			default:
				e.traverse(a.(ast.Node))
			}
		}
	}
	return code
}

func (e *AsteroidEVM) traverseInterface(n *ast.InterfaceDeclarationNode) (code bvmCreate.Bytecode) {
	// don't need to be interacted with
	// all interfaces are dealt with by the type system
	return code
}

func (e *AsteroidEVM) traverseEnum(n *ast.EnumDeclarationNode) (code bvmCreate.Bytecode) {
	// don't create anything
	return code
}

func (e *AsteroidEVM) traverseContract(n *ast.ContractDeclarationNode) (code bvmCreate.Bytecode) {

	e.inMachineStorage = false

	// create hooks for functions
	// create hooks for constructors
	// create hooks for events
	// traverse everything else?
	if n.Body.Declarations != nil {
		for _, d := range n.Body.Declarations.Map() {
			switch a := d.(type) {
			case *ast.LifecycleDeclarationNode:
				//	e.addLifecycleHook(n.Identifier, a)
				break
			case *ast.FuncDeclarationNode:
				e.addFunctionHook(n.Identifier, a)
				break
			case *ast.EventDeclarationNode:
				e.addEventHook(n.Identifier, a)
				break
			default:
				e.traverse(a.(ast.Node))
			}
		}
	}

	return code
}

func (e *AsteroidEVM) addFunctionHook(parent string, node *ast.FuncDeclarationNode) {
	/*e.hooks[parent][node.Identifier] = hook {
		name: node.Identifier,
	}*/

}

func (e *AsteroidEVM) addEventHook(parent string, node *ast.EventDeclarationNode) {
	/*e.hooks[parent][node.Identifier] = hook{
		name: node.Identifier,
	}*/
}

func (e *AsteroidEVM) addHook(name string) {
	h := hook{
		name: name,
	}
	if e.hooks == nil {
		e.hooks = make([]hook, 0)
	}
	e.hooks = append(e.hooks, h)
}

func (e *AsteroidEVM) traverseEvent(n *ast.EventDeclarationNode) (code bvmCreate.Bytecode) {
	indexed := 0

	if hasModifier(n.Modifiers.Modifiers, "indexed") {
		// all parameters will be indexed
		for _ = range n.Parameters {
			for _ = range n.Identifier {
				indexed++
			}
		}
	} else {
		for _, param := range n.Parameters {
			for _ = range n.Identifier {
				if hasModifier(param.Modifiers.Modifiers, "indexed") {
					indexed++

				}
			}
		}
	}

	topicLimit := 4
	if indexed > topicLimit {
		// TODO: add error
	}

	code.Add("BOUNCEDEST")
	code.Add("CALLER")
	// other parameters should be on the stack already
	code.Add(fmt.Sprintf("LOG%d", indexed))

	//e.addEventHook(n.Identifier, )

	return code
}

func (e *AsteroidEVM) traverseParameters(params []*ast.ExplicitVarDeclarationNode) (code bvmCreate.Bytecode) {
	machineStorage := false
	for _, p := range params {
		// function parameters are passed in on the stack and then assigned
		// default: machineMemory, can be overriden to be machineStorage
		// check if it's in machineStorage
		for _, i := range p.Identifiers {
			if machineStorage {
				e.allocateStorageToMachine(i, p.Resolved.Size())
				//code.Push(e.lookupMachineStorage(i))
				code.Add("SSTORE")
			} else {
				e.allocateMachineMemory(i, p.Resolved.Size())
				//code.Push(uintAsBytes(location)...)
				code.Add("MSTORE")
			}
		}
		machineStorage = false
	}
	return code
}

func (e *AsteroidEVM) traverseFunc(node *ast.FuncDeclarationNode) (code bvmCreate.Bytecode) {

	e.inMachineStorage = true

	// don't worry about hooking
	if hasModifier(node.Modifiers.Modifiers, "external") {
		code.Concat(e.traverseExternalFunction(node))
	} else if hasModifier(node.Modifiers.Modifiers, "internal") {
		code.Concat(e.traverseInternalFunction(node))
	} else if hasModifier(node.Modifiers.Modifiers, "global") {
		code.Concat(e.traverseGlobalFunction(node))
	}

	return code
}

func (e *AsteroidEVM) traverseExplicitVarDecl(n *ast.ExplicitVarDeclarationNode) (code bvmCreate.Bytecode) {
	// variable declarations don't require machineStorage (yet), just have to designate a slot
	for _, id := range n.Identifiers {
		if e.inMachineStorage {
			e.allocateStorageToMachine(id, n.Resolved.Size())
		} else {
			e.allocateMachineMemory(id, n.Resolved.Size())
		}
	}
	return code
}
