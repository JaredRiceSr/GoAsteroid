package evm

import (
	"fmt"

	"github.com/benchlab/bvmCreate"
)

func invalidBytecodeMessage(actual, expected bvmCreate.Bytecode) string {
	return fmt.Sprintf("Expected: %s\nActual: %s", expected.Format(), actual.Format())
}
