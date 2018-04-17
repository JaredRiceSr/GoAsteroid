package benchvm

import (
	"testing"

	"github.com/benchlab/bvmUtils"
	"github.com/benchlab/asteroid"
)

func TestBinaryExpressionBytecodeLiterals(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, "x = 1 + 5")
	bvmUtils.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD",
	})
	checkStack(t, a.BVM.Stack, [][]byte{
		[]byte{byte(6)},
	})
}

func TestBinaryExpressionBytecodeReferences(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, "a + b")
	bvmUtils.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		"PUSH", // push hash(x)
		"ADD",
	})
}

func TestBinaryExpressionBytecodeStringLiteralConcat(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `"my name is" + " who knows tbh"`)
	bvmUtils.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push hash(x)
		"CONCAT", // concatenate bytes
	})
}

func TestBinaryExpressionBytecodeStringReferenceConcat(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `
		var (
			a = "hello"
			b = "world"
			c = a + b
		)
		`)
	bvmUtils.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH",   // push string data
		"PUSH",   // push a reference
		"STORE",  // store at a
		"PUSH",   // push string data
		"PUSH",   // push b reference
		"PUSH",   // push a reference
		"LOAD",   // load a data
		"PUSH",   // push b reference
		"LOAD",   // load b data
		"CONCAT", //
	})
}

func TesExtendedBinaryExpressionBytecodeStringReferenceConcat(t *testing.T) {
	a := new(Arsonist)
	bvmUtils.AssertNow(t, a.VM != nil, "vm shouldn't be nil")
	asteroid.CompileString(a, `
		var (
			a = "hello"
			b = "world"
			c = "www"
			d = "comma"
		)
		a + b + c + d
		`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH",   // push hello
		"PUSH",   // push world
		"CONCAT", //
		"PUSH",   // push www
		"CONCAT", //
		"PUSH",   // push comma
		"CONCAT", //
	})
}

func TestUnaryExpressionBytecodeLiteral(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, "!1")
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push data
		"NOT",
	})
}

func TestCallExpressionBytecodeLiteral(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `
		doSomething("data")
		`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		"NOT",
	})
}

func TestCallExpressionBytecodeUseResult(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `
		s := doSomething("data")
		`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		"NOT",
	})
}

func TestCallExpressionBytecodeUseMultipleResults(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `
		s, a, p := doSomething("data")
		`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		"PUSH", // push p
		"SET",  // set p
		"PUSH", // push a
		"SET",  // set a
		"PUSH", // push s
		"SET",  // set s
	})
}

func TestCallExpressionBytecodeIgnoredResult(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `
		s, _, p := doSomething("data")
		`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		// stack ends: {R3, R2, R1}
		"PUSH", // push p
		"SET",  // set p
		"POP",  // ignore second return value
		"PUSH", // push s
		"SET",  // set s
	})
}

func TestCallExpressionBytecodeNestedCall(t *testing.T) {
	a := new(Arsonist)
	asteroid.CompileString(a, `
		err := saySomething(doSomething("data"))
		`)
	checkMnemonics(t, a.BVM.Instructions, []string{
		"PUSH", // push string data
		// this is where the function code would go
		// stack ends: {R1}
		// second function code goes here
		"PUSH", // push err
		"SET",  // set err
	})
}
