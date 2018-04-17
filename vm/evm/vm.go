package evm

import (
	"fmt"

	"github.com/benchlab/asteroid/util"

	"github.com/benchlab/asteroid/ast"
	"github.com/benchlab/asteroid/parser"
	"github.com/benchlab/asteroid/token"
	"github.com/benchlab/asteroid/typing"
	"github.com/benchlab/asteroid/validator"
)

func (evm AsteroidEVM) Builtins() *ast.ScopeNode {
	if builtinScope == nil {
		builtinScope, _ = parser.ParseFile("builtins.grd")
	}
	return builtinScope
}

func (evm AsteroidEVM) BooleanName() string {
	return "bool"
}

func (evm AsteroidEVM) Literals() validator.LiteralMap {
	if litMap == nil {
		litMap = validator.LiteralMap{
			token.String:  validator.SimpleLiteral("string"),
			token.True:    validator.BooleanLiteral,
			token.False:   validator.BooleanLiteral,
			token.Integer: resolveIntegerLiteral,
			token.Float:   resolveFloatLiteral,
		}
	}
	return litMap
}

func resolveIntegerLiteral(v *validator.Validator, data string) typing.Type {
	if len(data) == len("0x")+20 {
		// this might be an address
	}
	x := typing.BitsNeeded(len(data))
	return v.SmallestNumericType(x, false)
}

func resolveFloatLiteral(v *validator.Validator, data string) typing.Type {
	// convert to float
	return typing.Unknown()
}

func (evm AsteroidEVM) Primitives() map[string]typing.Type {

	const maxSize = 256
	m := map[string]typing.Type{}

	const increment = 8
	for i := increment; i <= maxSize; i += increment {
		ints := fmt.Sprintf("int%d", i)
		uints := "u" + ints
		m[uints] = &typing.NumericType{Name: uints, BitSize: i, Signed: false, Integer: true}
		m[ints] = &typing.NumericType{Name: ints, BitSize: i, Signed: true, Integer: true}
	}

	return m
}

func (evm AsteroidEVM) ValidDeclarations() []ast.NodeType {
	return ast.AllDeclarations
}

func (evm AsteroidEVM) ValidExpressions() []ast.NodeType {
	return ast.AllExpressions
}

func (evm AsteroidEVM) ValidStatements() []ast.NodeType {
	return ast.AllStatements
}

// TODO: context stacks
// want to validate that a modifier can only be applied to a function in a contract
var mods = []*validator.ModifierGroup{
	&validator.ModifierGroup{
		Name:       "Visibility",
		Modifiers:  []string{"external", "internal", "global"},
		RequiredOn: []ast.NodeType{ast.FuncDeclaration},
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
	&validator.ModifierGroup{
		Name:       "Payable",
		Modifiers:  []string{"payable", "nonpayable"},
		RequiredOn: []ast.NodeType{},
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
	&validator.ModifierGroup{
		Name:       "Payable",
		Modifiers:  []string{"payable", "nonpayable"},
		RequiredOn: []ast.NodeType{},
		AllowedOn:  []ast.NodeType{ast.FuncDeclaration},
		Maximum:    1,
	},
}

func (evm AsteroidEVM) Modifiers() []*validator.ModifierGroup {
	return mods
}

func (evm AsteroidEVM) Annotations() []*typing.Annotation {
	return nil
}

func (evm AsteroidEVM) Assignable(val *validator.Validator, left, right typing.Type, fromExpression ast.ExpressionNode) bool {
	t, _ := val.IsTypeVisible("address")
	if t.Compare(right) {
		switch left.(type) {
		case *typing.Contract:
			return true
		case *typing.Interface:
			return true
		}
	}
	if !typing.AssignableTo(left, right, true) {
		if validator.LiteralAssignable(left, right, fromExpression) {
			return true
		}
		return false
	}
	return true
}

func (evm AsteroidEVM) Castable(val *validator.Validator, to, from typing.Type, fromExpression ast.ExpressionNode) bool {
	// can cast all addresses to all contracts
	t, _ := val.IsTypeVisible("address")
	if t.Compare(from) {
		switch to.(type) {
		case *typing.Contract:
			return true
		case *typing.Interface:
			return true
		}
	}
	// can cast all uints to addresses
	if t.Compare(to) {
		switch a := typing.ResolveUnderlying(from).(type) {
		case *typing.NumericType:
			if !a.Signed {
				// check size
				return true
			}
			if l, ok := fromExpression.(*ast.LiteralNode); ok {
				hasSign := (l.Data[0] == '-')
				if !hasSign {
					return true
				}
			}
		}
	}
	if typing.AssignableTo(to, from, false) {
		return true
	}
	if validator.LiteralAssignable(to, from, fromExpression) {
		return true
	}
	return false
}

func (evm AsteroidEVM) BaseContract() (*ast.ContractDeclarationNode, util.Errors) {
	s, errs := parser.ParseString(`
		contract Base {
		    var balance uint
			var address address
		}
	`)
	c := s.Declarations.Next().(*ast.ContractDeclarationNode)
	return c, errs
}

func (evm AsteroidEVM) BaseClass() (*ast.ContractDeclarationNode, util.Errors) {
	s, errs := parser.ParseString(`
		class Object {

		}
	`)
	c := s.Declarations.Next().(*ast.ContractDeclarationNode)
	return c, errs
}

func (evm AsteroidEVM) BytecodeGenerators() map[string]validator.BytecodeGenerator {
	return builtins
}
