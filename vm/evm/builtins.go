package evm

import (
	"github.com/benchlab/asteroid/ast"

	"github.com/benchlab/asteroid/validator"
	"github.com/benchlab/bvmCreate"
)

var builtins = map[string]validator.BytecodeGenerator{
	// arithmetic
	"addmod":  validator.SimpleSpecification("ADDMOD"),
	"mulmod":  validator.SimpleSpecification("MULMOD"),
	"balance": validator.SimpleSpecification("BALANCE"),
	// transactional
	"transfer":     nil,
	"delegateCall": delegateCall,
	"call":         call,
	//"callcode": callCode,
	// error-checking
	"revert":  validator.SimpleSpecification("REVERT"),
	"throw":   validator.SimpleSpecification("REVERT"),
	"require": require,
	"assert":  assert,
	// cryptographic
	"sha3":      validator.SimpleSpecification("SHA3"),
	"keccak256": nil,
	"sha256":    nil,
	"ecrecover": nil,
	"ripemd160": nil,
	// ending
	"selfDestruct": validator.SimpleSpecification("SELFDESTRUCT"),

	// message
	"calldata":  calldata,
	"gas":       validator.SimpleSpecification("GAS"),
	"sender":    validator.SimpleSpecification("CALLER"),
	"signature": signature,

	// block
	"timestamp": validator.SimpleSpecification("TIMESTAMP"),
	"number":    validator.SimpleSpecification("NUMBER"),
	"blockhash": blockhash,
	"coinbase":  validator.SimpleSpecification("COINBASE"),
	"gasLimit":  validator.SimpleSpecification("GASLIMIT"),
	// tx
	"gasPrice": validator.SimpleSpecification("GASPRICE"),
	"origin":   validator.SimpleSpecification("ORIGIN"),
}

func transfer(bvm validator.BVM) (code bvmCreate.Bytecode) {
	e := bvm.(AsteroidEVM)
	call := e.expression.(*ast.CallExpressionNode)
	// gas
	code.Concat(push(uintAsBytes(uint(2300))))
	// to
	code.Concat(e.traverse(call.Arguments[0]))
	// value
	code.Concat(e.traverse(call.Arguments[1]))
	// in offset
	code.Concat(push(uintAsBytes(uint(0))))
	// in size
	code.Concat(push(uintAsBytes(uint(0))))
	// out offset
	e.allocateMachineMemory("transfer", 1)
	mem := e.lookupMachineMemory("transfer")
	code.Concat(push(uintAsBytes(uint(mem.offset))))
	// out size
	code.Concat(push(uintAsBytes(uint(1))))
	code.Add("CALL")
	return code
}

func call(bvm validator.BVM) (code bvmCreate.Bytecode) {
	e := bvm.(AsteroidEVM)
	call := e.expression.(*ast.CallExpressionNode)
	// gas
	code.Concat(e.traverse(call.Arguments[1]))
	// recipient --> should be on the stack already
	code.Concat(e.traverse(call.Arguments[0]))
	// ether value
	code.Concat(e.traverse(call.Arguments[2]))
	// machineMemory location of start of input data
	code.Add("PUSH")
	// length of input data
	code.Add("PUSH")
	// out offset
	e.allocateMachineMemory("transfer", 1)
	mem := e.lookupMachineMemory("transfer")
	code.Concat(push(uintAsBytes(uint(mem.offset))))
	// out size
	code.Concat(push(uintAsBytes(uint(1))))
	code.Add("CALL")
	return code
}

func calldata(bvm validator.BVM) (code bvmCreate.Bytecode) {
	code.Add("CALLDATA")
	return code
}

func blockhash(bvm validator.BVM) (code bvmCreate.Bytecode) {
	code.Add("BLOCKHASH")
	return code
}

func require(bvm validator.BVM) (code bvmCreate.Bytecode) {
	code.Concat(pushMarker(2))
	code.Add("BOUNCEI")
	code.Add("REVERT")
	return code
}

func assert(bvm validator.BVM) (code bvmCreate.Bytecode) {
	// TODO: invalid opcodes
	code.Concat(pushMarker(2))
	code.Add("BOUNCEI")
	code.Add("INVALID")
	return code
}

func delegateCall(bvm validator.BVM) (code bvmCreate.Bytecode) {
	code.Add("DELEGATECALL")
	return code
}

func callCode(bvm validator.BVM) (code bvmCreate.Bytecode) {
	code.Add("CALLCODE")
	return code
}

func signature(bvm validator.BVM) (code bvmCreate.Bytecode) {
	// get first four bytes of calldata
	return code
}

func length(bvm validator.BVM) (code bvmCreate.Bytecode) {
	// must be an array
	// array size is always at the first index
	evm := bvm.(*AsteroidEVM)
	if evm.inStorage {
		code.Add("SLOAD")
	} else {
		code.Add("MLOAD")
	}
	return code
}

func appendBuiltin(bvm validator.BVM) (code bvmCreate.Bytecode) {
	// must be an array
	// array size is always at the first index
	evm := bvm.(*AsteroidEVM)
	if evm.inStorage {
		//code.Add(push())
		code.Add("SLOAD")
		code.Concat(push(uintAsBytes(uint(1))))
		code.Add("ADD")
		code.Add("SSTORE")

	} else {
		code.Add("MLOAD")
	}
	return code
}
