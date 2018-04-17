package evm

import (
	"fmt"

	"github.com/benchlab/bvmCreate"
)

// why are gas costs needed?
// to enable the compiler to estimate the gas usage of different execution paths
// can display to user at compile time

type EVMGenerator struct {
}

const (
	gasZero    = 0
	gasBase    = 2
	gasVeryLow = 3
	gasLow     = 5
	gasMid     = 8
	gasHigh    = 10

	gasJumpDest = 1

	gasQuickStep   = 2
	gasFastestStep = 3
	gasFastStep    = 5
	gasMidStep     = 8
	gasSlowStep    = 10
	gasExtStep     = 20

	gasContractByte = 200

	gasBlockhash = 20
)

type Specification struct {
	opcode []byte
	cost   func(interface{}) int
}

func constantGas(gas int) func(interface{}) int {
	return func(interface{}) int {
		return gas
	}
}

func generateDups() (im bvmCreate.SpecificationMap) {
	number := 16
	for i := 1; i <= number; i++ {
		im[fmt.Sprintf("DUP%d", i)] = bvmCreate.Specification{Opcode: uint(0x80 + i), Cost: constantGas(gasZero)}
	}
	return im
}

func generateSwaps() (im bvmCreate.SpecificationMap) {
	number := 16
	for i := 1; i <= number; i++ {
		im[fmt.Sprintf("SWAP%d", i)] = bvmCreate.Specification{Opcode: uint(0x90 + i), Cost: constantGas(gasZero)}
	}
	return im
}

func generatePushes() (im bvmCreate.SpecificationMap) {
	number := 32
	for i := 1; i <= number; i++ {
		im[fmt.Sprintf("SWAP%d", i)] = bvmCreate.Specification{Opcode: uint(0x60 + i), Cost: constantGas(gasZero)}
	}
	return im
}

func generateLogs() (im bvmCreate.SpecificationMap) {
	number := 5
	for i := 0; i < number; i++ {
		im[fmt.Sprintf("SWAP%d", i)] = bvmCreate.Specification{Opcode: uint(0xA0 + i), Cost: constantGas(gasZero)}
	}
	return im
}

func gasExp(interface{}) int {
	return 0
}

func gasSLoad(interface{}) int {
	return 0
}

func gasSStore(interface{}) int {
	return 0
}

func gasCall(interface{}) int {
	return 0
}

func gasBalance(interface{}) int {
	return 0
}

func gasSha3(interface{}) int {
	return 0
}

func gasCodeCopy(interface{}) int {
	return 0
}

func gasCallDataCopy(interface{}) int {
	return 0
}

func gasExtCodeSize(interface{}) int {
	return 0
}

func gasExtCodeCopy(interface{}) int {
	return 0
}

func (e EVMGenerator) Opcodes() bvmCreate.SpecificationMap {
	m := bvmCreate.SpecificationMap{
		"STOP":       bvmCreate.Specification{Opcode: 0x00, Cost: constantGas(gasZero)},
		"ADD":        bvmCreate.Specification{Opcode: 0x01, Cost: constantGas(gasVeryLow)},
		"MUL":        bvmCreate.Specification{Opcode: 0x02, Cost: constantGas(gasLow)},
		"SUB":        bvmCreate.Specification{Opcode: 0x03, Cost: constantGas(gasVeryLow)},
		"DIV":        bvmCreate.Specification{Opcode: 0x04, Cost: constantGas(gasLow)},
		"SDIV":       bvmCreate.Specification{Opcode: 0x05, Cost: constantGas(gasLow)},
		"MOD":        bvmCreate.Specification{Opcode: 0x06, Cost: constantGas(gasLow)},
		"SMOD":       bvmCreate.Specification{Opcode: 0x07, Cost: constantGas(gasLow)},
		"ADDMOD":     bvmCreate.Specification{Opcode: 0x08, Cost: constantGas(gasMid)},
		"MULMOD":     bvmCreate.Specification{Opcode: 0x09, Cost: constantGas(gasMid)},
		"EXP":        bvmCreate.Specification{Opcode: 0x0A, Cost: gasExp},
		"SIGNEXTEND": bvmCreate.Specification{Opcode: 0x0B, Cost: constantGas(gasLow)},

		"LT":     bvmCreate.Specification{Opcode: 0x10, Cost: constantGas(gasVeryLow)},
		"GT":     bvmCreate.Specification{Opcode: 0x11, Cost: constantGas(gasVeryLow)},
		"SLT":    bvmCreate.Specification{Opcode: 0x12, Cost: constantGas(gasVeryLow)},
		"SGT":    bvmCreate.Specification{Opcode: 0x13, Cost: constantGas(gasVeryLow)},
		"EQ":     bvmCreate.Specification{Opcode: 0x14, Cost: constantGas(gasVeryLow)},
		"ISZERO": bvmCreate.Specification{Opcode: 0x15, Cost: constantGas(gasVeryLow)},
		"AND":    bvmCreate.Specification{Opcode: 0x16, Cost: constantGas(gasVeryLow)},
		"OR":     bvmCreate.Specification{Opcode: 0x17, Cost: constantGas(gasVeryLow)},
		"XOR":    bvmCreate.Specification{Opcode: 0x18, Cost: constantGas(gasVeryLow)},
		"NOT":    bvmCreate.Specification{Opcode: 0x19, Cost: constantGas(gasVeryLow)},
		"BYTE":   bvmCreate.Specification{Opcode: 0x1A, Cost: constantGas(gasVeryLow)},

		"SHA3": bvmCreate.Specification{Opcode: 0x20, Cost: gasSha3},

		"ADDRESS":      bvmCreate.Specification{Opcode: 0x30, Cost: constantGas(gasBase)},
		"BALANCE":      bvmCreate.Specification{Opcode: 0x31, Cost: gasBalance},
		"ORIGIN":       bvmCreate.Specification{Opcode: 0x32, Cost: constantGas(gasBase)},
		"CALLER":       bvmCreate.Specification{Opcode: 0x33, Cost: constantGas(gasBase)},
		"CALLVALUE":    bvmCreate.Specification{Opcode: 0x34, Cost: constantGas(gasBase)},
		"CALLDATALOAD": bvmCreate.Specification{Opcode: 0x35, Cost: constantGas(gasVeryLow)},
		"CALLDATASIZE": bvmCreate.Specification{Opcode: 0x36, Cost: constantGas(gasBase)},
		"CALLDATACOPY": bvmCreate.Specification{Opcode: 0x37, Cost: gasCallDataCopy},
		"CODESIZE":     bvmCreate.Specification{Opcode: 0x38, Cost: constantGas(gasBase)},
		"CODECOPY":     bvmCreate.Specification{Opcode: 0x39, Cost: gasCodeCopy},
		"GASPRICE":     bvmCreate.Specification{Opcode: 0x3A, Cost: constantGas(gasBase)},
		"EXTCODESIZE":  bvmCreate.Specification{Opcode: 0x3B, Cost: gasExtCodeSize},
		"EXTCODECOPY":  bvmCreate.Specification{Opcode: 0x3C, Cost: gasExtCodeCopy},

		"BLOCKHASH":  bvmCreate.Specification{Opcode: 0x40, Cost: constantGas(gasExtStep)},
		"COINBASE":   bvmCreate.Specification{Opcode: 0x41, Cost: constantGas(gasBase)},
		"TIMESTAMP":  bvmCreate.Specification{Opcode: 0x42, Cost: constantGas(gasBase)},
		"NUMBER":     bvmCreate.Specification{Opcode: 0x43, Cost: constantGas(gasBase)},
		"DIFFICULTY": bvmCreate.Specification{Opcode: 0x44, Cost: constantGas(gasBase)},
		"GASLIMIT":   bvmCreate.Specification{Opcode: 0x45, Cost: constantGas(gasBase)},

		"POP":      bvmCreate.Specification{Opcode: 0x50, Cost: constantGas(gasBase)},
		"MLOAD":    bvmCreate.Specification{Opcode: 0x51, Cost: constantGas(gasVeryLow)},
		"MSTORE":   bvmCreate.Specification{Opcode: 0x52, Cost: constantGas(gasVeryLow)},
		"MSTORE8":  bvmCreate.Specification{Opcode: 0x53, Cost: constantGas(gasVeryLow)},
		"SLOAD":    bvmCreate.Specification{Opcode: 0x54, Cost: gasSLoad},
		"SSTORE":   bvmCreate.Specification{Opcode: 0x55, Cost: gasSStore},
		"BOUNCE":     bvmCreate.Specification{Opcode: 0x56, Cost: constantGas(gasMid)},
		"BOUNCEI":    bvmCreate.Specification{Opcode: 0x57, Cost: constantGas(gasHigh)},
		"PC":       bvmCreate.Specification{Opcode: 0x58, Cost: constantGas(gasBase)},
		"MSIZE":    bvmCreate.Specification{Opcode: 0x59, Cost: constantGas(gasBase)},
		"GAS":      bvmCreate.Specification{Opcode: 0x5A, Cost: constantGas(gasBase)},
		"BOUNCEDEST": bvmCreate.Specification{Opcode: 0x5B, Cost: constantGas(gasJumpDest)},

		"CREATE":       bvmCreate.Specification{Opcode: 0xF0, Cost: constantGas(gasJumpDest)},
		"CALL":         bvmCreate.Specification{Opcode: 0xF1, Cost: constantGas(gasJumpDest)},
		"CALLCODE":     bvmCreate.Specification{Opcode: 0xF2, Cost: constantGas(gasJumpDest)},
		"RETURN":       bvmCreate.Specification{Opcode: 0xF3, Cost: constantGas(gasJumpDest)},
		"DELEGATECALL": bvmCreate.Specification{Opcode: 0xF4, Cost: constantGas(gasJumpDest)},

		"SELFDESTRUCT": bvmCreate.Specification{Opcode: 0xFF, Cost: constantGas(gasJumpDest)},
	}

	m.AddAll(generatePushes())
	m.AddAll(generateDups())
	m.AddAll(generateLogs())
	m.AddAll(generateSwaps())

	return m
}
