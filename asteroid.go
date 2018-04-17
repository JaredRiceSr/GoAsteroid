package asteroid

import (
	"fmt"

	"github.com/benchlab/bvmCreate"

	"github.com/benchlab/asteroid/lexer"
	"github.com/benchlab/asteroid/parser"
	"github.com/benchlab/asteroid/util"
	"github.com/benchlab/asteroid/validator"
)

func reportErrors(category string, errs util.Errors) {
	msg := fmt.Sprintf("%s Errors\n", category)
	msg += errs.Format()
	fmt.Println(msg)
}

// CompileBytes ...
func CompileBytes(bvm validator.BVM, bytes []byte) bvmCreate.Bytecode {
	tokens, errs := lexer.Lex(bytes)

	if errs != nil {
		reportErrors("Lexing", errs)
	}

	ast, errs := parser.Parse(tokens, errs)

	if errs != nil {
		reportErrors("Parsing", errs)
	}

	errs = validator.Validate(ast, bvm)

	if errs != nil {
		reportErrors("Type Validation", errs)
	}

	bytecode, errs := bvm.Traverse(ast)

	if errs != nil {
		reportErrors("Bytecode Generation", errs)
	}
	return bytecode
}

// CompileString ...
func CompileString(bvm validator.BVM, data string) bvmCreate.Bytecode {
	return CompileBytes(bvm, []byte(data))
}

func CompileFilesData(bvm validator.BVM, data [][]byte) (bvmCreate.Bytecode, util.Errors) {

}

/* EVM ...
func EVM() Traverser {
	return evm.NewTraverser()
}

// BenchVM ...
func BenchVM() Traverser {
	return benchvm.NewTraverser()
}*/
