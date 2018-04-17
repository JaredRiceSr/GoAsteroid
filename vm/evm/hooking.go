package evm

import (
	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/bvmCreate"
)

type hook struct {
	name     string
	position int
	bytecode bvmCreate.Bytecode
}

func (e *AsteroidEVM) createFunctionBody(node *ast.FuncDeclarationNode) (body bvmCreate.Bytecode) {
	// all function bodies look the same
	// traverse the scope
	body.Concat(e.traverseScope(node.Body))

	// return address should be on top of the stack
	body.Add("BOUNCE")
	return body
}

func (e *AsteroidEVM) createExternalFunctionComponents(node *ast.FuncDeclarationNode) (params, body bvmCreate.Bytecode) {
	// move params from calldata to machineMemory

	return params, e.createFunctionBody(node)
}

func (e *AsteroidEVM) createExternalParameters(node *ast.FuncDeclarationNode) (code bvmCreate.Bytecode) {
	offset := uint(0)
	for _, param := range node.Signature.Parameters {
		exp := param.(*ast.ExplicitVarDeclarationNode)
		for _ = range exp.Identifiers {
			offset += exp.Resolved.Size()

		}
	}
	return code
}

func (e *AsteroidEVM) traverseExternalFunction(node *ast.FuncDeclarationNode) (code bvmCreate.Bytecode) {

	params := e.createExternalParameters(node)

	body := e.createFunctionBody(node)

	code.Concat(params)
	code.Concat(body)

	e.addExternalHook(node.Signature.Identifier, code)

	return code
}

/*
| return location |
| param 1 |
| param 2 |
*/
func (e *AsteroidEVM) traverseInternalFunction(node *ast.FuncDeclarationNode) (code bvmCreate.Bytecode) {

	params := e.createInternalParameters(node)

	body := e.createFunctionBody(node)

	code.Concat(params)
	code.Concat(body)

	e.addInternalHook(node.Signature.Identifier, code)

	return code
}

func (e *AsteroidEVM) createInternalParameters(node *ast.FuncDeclarationNode) (params bvmCreate.Bytecode) {
	// as internal functions can only be called from inside the contract
	// no need to have a hook
	// can just jump to the location
	params.Add("BOUNCEDEST")
	// load all parameters into machineMemory blocks
	for _, param := range node.Signature.Parameters {
		exp := param.(*ast.ExplicitVarDeclarationNode)
		for _, i := range exp.Identifiers {
			e.allocateMachineMemory(i, exp.Resolved.Size())
			// all parameters must be on the stack
			loc := e.lookupMachineMemory(i)
			params.Concat(loc.retrieve())
		}
	}

	return params
}

func (e *AsteroidEVM) traverseGlobalFunction(node *ast.FuncDeclarationNode) (code bvmCreate.Bytecode) {
	// hook here
	// get all parameters out of calldata and into machineMemory
	// then jump over the parameter init in the internal declaration
	// leave the hook for later
	external := e.createExternalParameters(node)

	internal := e.createInternalParameters(node)

	body := e.createFunctionBody(node)

	code.Concat(pushMarker(external.Length()))

	code.Concat(external)

	code.Add("BOUNCEDEST")

	code.Concat(internal)

	code.Concat(body)

	e.addGlobalHook(node.Signature.Identifier, code)
	return code
}

func (e *AsteroidEVM) addGlobalHook(id string, code bvmCreate.Bytecode) {
	if e.globalHooks == nil {
		e.globalHooks = make(map[string]hook)
	}
	e.globalHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}

func (e *AsteroidEVM) addInternalHook(id string, code bvmCreate.Bytecode) {
	if e.internalHooks == nil {
		e.internalHooks = make(map[string]hook)
	}
	e.internalHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}

func (e *AsteroidEVM) addExternalHook(id string, code bvmCreate.Bytecode) {
	if e.externalHooks == nil {
		e.externalHooks = make(map[string]hook)
	}
	e.externalHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}

func (e *AsteroidEVM) addLifecycleHook(id string, code bvmCreate.Bytecode) {
	if e.lifecycleHooks == nil {
		e.lifecycleHooks = make(map[string]hook)
	}
	e.lifecycleHooks[id] = hook{
		name:     id,
		bytecode: code,
	}
}
