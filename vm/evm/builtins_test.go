package evm

import (
	"testing"

	"github.com/benchlab/asteroid/validator"

	"github.com/benchlab/bvmUtils"
)

func TestBuiltinRequire(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "require(5 > 3)")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())

	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"LT",
		"PUSH1",
		"BOUNCEI",
		"REVERT",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinAssert(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "assert(5 > 3)")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"PUSH1",
		"LT",
		"PUSH",
		"BOUNCEI",
		"INVALID",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinArrayLength(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, `len("hello")`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH7",
		"MLOAD",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinAddmod(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "addmod(uint(1), uint(2), uint(3))")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH32",
		"PUSH32",
		"PUSH32",
		"ADDMOD",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinMulmod(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "mulmod(uint(1), uint(2), uint(3))")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH32",
		"PUSH32",
		"PUSH32",
		"ADDMOD",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinCall(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "call(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinDelegateCall(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "delegate(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinBalance(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, "balance(0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF)")
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH1",
		"BALANCE",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinSha3(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, `sha3("hello")`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"PUSH",
		"SHA3",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestBuiltinRevert(t *testing.T) {
	e := NewBVM()
	a, errs := validator.ValidateExpression(e, `revert()`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	code := e.traverseExpression(a)
	expected := []string{
		"REVERT",
	}
	bvmUtils.Assert(t, code.CompareMnemonics(expected), code.Format())
}

func TestCalldata(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "msg.data")
	bytecode := e.traverseExpression(expr)
	expected := []string{"CALLDATA"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestGas(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "msg.gas")
	bytecode := e.traverseExpression(expr)
	expected := []string{"GAS"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestCaller(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "msg.sender")
	bytecode := e.traverseExpression(expr)
	expected := []string{"CALLER"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestSignature(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "msg.sig")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"CALLDATA"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTimestamp(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "block.timestamp")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"TIMESTAMP"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestNumber(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "block.timestamp")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"NUMBER"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestCoinbase(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "block.timestamp")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"COINBASE"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestGasLimit(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "block.gasLimit")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"GASLIMIT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestBlockhash(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "block.gasLimit")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"GASLIMIT"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestGasPrice(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "tx.gasPrice")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"GASPRICE"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestOrigin(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, "tx.origin")
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{"ORIGIN"}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestTransfer(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, `
		transfer(0x123f681646d4a755815f9cb19e1acc8565a0c2ac, 1000)
	`)
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}

func TestCall(t *testing.T) {
	e := new(AsteroidEVM)
	expr, _ := validator.ValidateExpression(e, `
		Call(0x123f681646d4a755815f9cb19e1acc8565a0c2ac).gas(1000).value(1000).sig("aaaaa").call()
	`)
	bytecode := e.traverseExpression(expr)
	// should get first 4 bytes of calldata
	expected := []string{
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"PUSH",
		"CALL",
	}
	bvmUtils.Assert(t, bytecode.CompareMnemonics(expected), bytecode.Format())
}
