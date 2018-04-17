package evm

import (
	"strconv"

	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/asteroid/util"
	"github.com/benchlab/asteroid/validator"
	"github.com/benchlab/bvmCreate"
)

type hookMap map[string]hook

// AsteroidEVM
type AsteroidEVM struct {
	expression         ast.ExpressionNode
	hooks              []hook
	lastSlot           uint
	lastOffset         uint
	storage            map[string]*storageBlock
	freedMemory        []*memoryBlock
	memoryCursor       uint
	memory             map[string]*memoryBlock
	currentlyAssigning string
	internalHooks      hookMap
	externalHooks      hookMap
	globalHooks        hookMap
	eventHooks         hookMap
	lifecycleHooks     hookMap
	inStorage          bool
	mapLiteralCount    int
	arrayLiteralCount  int
}

func push(data []byte) (code bvmCreate.Bytecode) {
	if len(data) > 32 {
		// TODO: error
	}
	m := "PUSH" + strconv.Itoa(len(data))
	code.Add(m, data...)
	return code
}

func bytesRequired(offset int) int {
	count := 0
	for offset > 0 {
		count++
		offset >>= 1
	}
	if count%8 == 0 {
		return count / 8
	}
	return (count / 8) + 1
}

// support all offsets which can be stored in a 64 bit integer
func pushMarker(offset int) (code bvmCreate.Bytecode) {
	//TODO: fix
	code.AddMarker("PUSH"+strconv.Itoa(bytesRequired(offset)), offset)
	return code
}

var (
	builtinScope *ast.ScopeNode
	litMap       validator.LiteralMap
	opMap        validator.OperatorMap
)

func (evm AsteroidEVM) Traverse(node ast.Node) (bvmCreate.Bytecode, util.Errors) {
	// do pre-processing/hooks etc
	code := evm.traverse(node)
	// generate the bytecode
	// finalise the bytecode
	//evm.finalise()
	return code, nil
}

// NewAsteroidEVM ...
func NewBVM() AsteroidEVM {
	return AsteroidEVM{}
}

// A hook conditionally jumps the code to a particular point
//

func (e *AsteroidEVM) finalise() {

	// add external functions
	// add internal functions
	// add events

	// number of specifications =
	/*
		for _, hook := range e.hooks {
			e.BVM.AddBytecode("POP")

			e.BVM.AddBytecode("EQL")
			e.BVM.AddBytecode("JMPI")
		}
		// if the data matches none of the function hooks
		e.BVM.AddBytecode("STOP")
		for _, callable := range e.callables {
			// add function bytecode
		}*/
}

//

// can be called from outside or inside the contract
func (e *AsteroidEVM) hookPublicFunc(h *hook) {

}

// can be
func (e *AsteroidEVM) hookPrivateFunc(h *hook) {

}

func (e AsteroidEVM) traverse(n ast.Node) (code bvmCreate.Bytecode) {
	/* initialise the bvm
	if e.BVM == nil {
		e.BVM = benchvm.NewBVM()
	}*/
	switch node := n.(type) {
	case *ast.ScopeNode:
		return e.traverseScope(node)
	case *ast.ClassDeclarationNode:
		return e.traverseClass(node)
	case *ast.InterfaceDeclarationNode:
		return e.traverseInterface(node)
	case *ast.EnumDeclarationNode:
		return e.traverseEnum(node)
	case *ast.EventDeclarationNode:
		return e.traverseEvent(node)
	case *ast.ExplicitVarDeclarationNode:
		return e.traverseExplicitVarDecl(node)
	case *ast.TypeDeclarationNode:
		return e.traverseType(node)
	case *ast.ContractDeclarationNode:
		return e.traverseContract(node)
	case *ast.FuncDeclarationNode:
		return e.traverseFunc(node)
	case *ast.ForStatementNode:
		return e.traverseForStatement(node)
	case *ast.AssignmentStatementNode:
		return e.traverseAssignmentStatement(node)
	case *ast.CaseStatementNode:
		return e.traverseCaseStatement(node)
	case *ast.ReturnStatementNode:
		return e.traverseReturnStatement(node)
	case *ast.IfStatementNode:
		return e.traverseIfStatement(node)
	case *ast.SwitchStatementNode:
		return e.traverseSwitchStatement(node)
	}
	return code
}

func (e *AsteroidEVM) traverseScope(s *ast.ScopeNode) (code bvmCreate.Bytecode) {
	if s == nil {
		return code
	}
	if s.Declarations != nil {
		for _, d := range s.Declarations.Array() {
			code.Concat(e.traverse(d.(ast.Node)))
		}
	}
	if s.Sequence != nil {
		for _, s := range s.Sequence {
			code.Concat(e.traverse(s))
		}
	}

	return code
}
