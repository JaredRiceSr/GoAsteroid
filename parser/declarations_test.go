package parser

import (
	"fmt"
	"testing"

	"github.com/benchlab/asteroid/ast"

	"github.com/benchlab/bvmUtils"
)

func TestParseInterfaceDeclarationEmpty(t *testing.T) {
	p := createParser(`interface Wagable {}`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	bvmUtils.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	bvmUtils.AssertNow(t, p.scope != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, i.Supers == nil, "wrong supers")
	bvmUtils.AssertLength(t, len(i.Signatures), 0)
}

func TestParseInterfaceDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible {}`)
	bvmUtils.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseInterfaceDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`interface Wagable inherits Visible, Movable {}`)
	bvmUtils.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseInterfaceDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract interface Wagable {}`)
	bvmUtils.Assert(t, isModifier(p), "should detect modifier")
	parseModifiers(p)
	bvmUtils.Assert(t, isInterfaceDeclaration(p), "should detect interface decl")
	parseInterfaceDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.InterfaceDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseInterfaceInheriting(t *testing.T) {
	_, errs := ParseString(`
		interface Switchable{}
		interface Deletable{}
		interface Light inherits Switchable, Deletable {}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseContractDeclarationEmpty(t *testing.T) {
	p := createParser(`contract Wagable {}`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 4, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseContractDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible inherits Object {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Interfaces) == 1, "wrong interfaces length")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits A,B is Visible, Movable  {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible, Movable {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Visible is A, B {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`contract Wagable inherits Object is Visible {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect interface decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseContractDeclarationMultipleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
	bvmUtils.AssertNow(t, len(i.Interfaces) == 2, "wrong interfaces length")
}

func TestParseContractDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`contract Wagable is Visible, Movable inherits A,B {}`)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseContractDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract contract Wagable {}`)
	bvmUtils.Assert(t, isModifier(p), "should detect modifier")
	parseModifiers(p)
	bvmUtils.Assert(t, isContractDeclaration(p), "should detect contract decl")
	parseContractDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ContractDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ContractDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationMultipleInterfaces(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 0, "wrong supers length")
}

func TestParseClassDeclarationSingleInterfaceSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable is Visible inherits Object {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInterfaceMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits A,B is Visible, Movable  {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritance(t *testing.T) {
	p := createParser(`class Wagable inherits Visible, Movable {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Visible {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationSingleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable inherits Object is Visible {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect interface decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritanceSingleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 1, "wrong supers length")
}

func TestParseClassDeclarationMultipleInheritanceMultipleInterface(t *testing.T) {
	p := createParser(`class Wagable is Visible, Movable inherits A,B {}`)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
	bvmUtils.AssertNow(t, len(i.Supers) == 2, "wrong supers length")
}

func TestParseClassDeclarationAbstract(t *testing.T) {
	p := createParser(`abstract class Wagable {}`)
	bvmUtils.Assert(t, isModifier(p), "should detect modifier")
	parseModifiers(p)
	bvmUtils.Assert(t, isClassDeclaration(p), "should detect class decl")
	parseClassDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ClassDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	i := n.(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, i.Identifier == "Wagable", "wrong identifier")
}

func TestParseTypeDeclaration(t *testing.T) {
	p := createParser(`type Wagable int`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 3, fmt.Sprintf("wrong token length: %d", len(p.lexer.Tokens)))
	bvmUtils.Assert(t, isTypeDeclaration(p), "should detect type decl")
	parseTypeDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.TypeDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.TypeDeclarationNode)
	bvmUtils.AssertNow(t, e.Identifier == "Wagable", "wrong type name")

}

func TestParseExplicitVarDeclaration(t *testing.T) {
	p := createParser(`var a string`)
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "should detect expvar decl")
	parseExplicitVarDeclaration(p)
}

func TestParseEventDeclarationEmpty(t *testing.T) {
	p := createParser(`event Notification()`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 0, "wrong param length")
}

func TestParseEventDeclarationSingle(t *testing.T) {
	p := createParser(`event Notification(a string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
}

func TestParseEventDeclarationMultiple(t *testing.T) {
	p := createParser(`event Notification(a string, b string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
}

func TestParseEnum(t *testing.T) {
	p := createParser(`enum Weekday {}`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 4, "wrong token length")
	bvmUtils.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EnumDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EnumDeclarationNode)
	bvmUtils.AssertNow(t, e.Identifier == "Weekday", "wrong identifier")
}

func TestParseEnumInheritsSingle(t *testing.T) {
	p := createParser(`enum Day inherits Weekday {}`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 6, "wrong token length")
	bvmUtils.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EnumDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EnumDeclarationNode)
	bvmUtils.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseEnumInheritsMultiple(t *testing.T) {
	p := createParser(`enum Day inherits Weekday, Weekend {}`)
	bvmUtils.AssertNow(t, len(p.lexer.Tokens) == 8, "wrong token length")
	bvmUtils.Assert(t, isEnumDeclaration(p), "should detect enum decl")
	parseEnumDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EnumDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EnumDeclarationNode)
	bvmUtils.AssertNow(t, e.Identifier == "Day", "wrong identifier")
}

func TestParseVarDeclarationSimple(t *testing.T) {
	p := createParser("a int")
	d := p.parseVarDeclaration()
	bvmUtils.AssertNow(t, len(d.Identifiers) == 1, "wrong id length")
	bvmUtils.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	dt := d.DeclaredType
	bvmUtils.AssertNow(t, dt.Type() == ast.PlainType, "Asteroid Errors: Node Error: Wrong node type. ")
}

func TestParseVarDeclarationMultiple(t *testing.T) {
	p := createParser("a, b int")
	d := p.parseVarDeclaration()
	bvmUtils.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	bvmUtils.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	bvmUtils.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
}

func TestParseVarDeclarationMultipleExternal(t *testing.T) {
	p := createParser("a, b pkg.Type")
	d := p.parseVarDeclaration()
	bvmUtils.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	bvmUtils.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	bvmUtils.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	bvmUtils.AssertNow(t, dt.Type() == ast.PlainType, "Asteroid Errors: Node Error: Wrong node type. ")
}

func TestParseVarDeclarationMap(t *testing.T) {
	p := createParser("a, b map[string]string")
	d := p.parseVarDeclaration()
	bvmUtils.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	bvmUtils.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	bvmUtils.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	bvmUtils.AssertNow(t, dt.Type() == ast.MapType, "Asteroid Errors: Node Error: Wrong node type. ")
	m := dt.(*ast.MapTypeNode)
	bvmUtils.AssertNow(t, m.Key.Type() == ast.PlainType, "wrong key type")
	bvmUtils.AssertNow(t, m.Value.Type() == ast.PlainType, "wrong value type")
}

func TestParseVarDeclarationArray(t *testing.T) {
	p := createParser("a, b []string")
	d := p.parseVarDeclaration()
	bvmUtils.AssertNow(t, len(d.Identifiers) == 2, "wrong id length")
	bvmUtils.AssertNow(t, d.Identifiers[0] == "a", "wrong id 0 value")
	bvmUtils.AssertNow(t, d.Identifiers[1] == "b", "wrong id 1 value")
	dt := d.DeclaredType
	bvmUtils.AssertNow(t, dt.Type() == ast.ArrayType, "Asteroid Errors: Node Error: Wrong node type. ")
	m := dt.(*ast.ArrayTypeNode)
	bvmUtils.AssertNow(t, m.Value.Type() == ast.PlainType, "wrong key type")
}

func TestParseFuncNoParameters(t *testing.T) {
	p := createParser(`func foo(){}`)
	bvmUtils.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n != nil, "node is nil")
	bvmUtils.AssertNow(t, n.Type() == ast.FuncDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	f := n.(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, f.Signature != nil, "signature isn't nil")
	bvmUtils.AssertLength(t, len(f.Signature.Parameters), 0)
}

func TestParseFuncOneParameter(t *testing.T) {
	p := createParser(`func foo(a int){}`)
	bvmUtils.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n != nil, "node is nil")
	bvmUtils.AssertNow(t, n.Type() == ast.FuncDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	f := n.(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, f.Signature != nil, "signature isn't nil")
	bvmUtils.AssertNow(t, len(f.Signature.Parameters) == 1, "wrong param length")
}

func TestParseFuncParameters(t *testing.T) {
	p := createParser(`func foo(a int, b string){}`)
	bvmUtils.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n != nil, "node is nil")
	bvmUtils.AssertNow(t, n.Type() == ast.FuncDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	f := n.(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, f.Signature != nil, "signature isn't nil")
	bvmUtils.AssertLength(t, len(f.Signature.Parameters), 2)
}

func TestParseFuncMultiplePerType(t *testing.T) {
	p := createParser(`func foo(a, b int){}`)
	bvmUtils.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n != nil, "node is nil")
	bvmUtils.AssertNow(t, n.Type() == ast.FuncDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	f := n.(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, f.Signature != nil, "signature isn't nil")
	bvmUtils.AssertNow(t, len(f.Signature.Parameters) == 1, "wrong param length")
}

func TestParseFuncMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`func foo(a, b int, c string){}`)
	bvmUtils.Assert(t, isFuncDeclaration(p), "should detect func decl")
	parseFuncDeclaration(p)

	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n != nil, "node is nil")
	bvmUtils.AssertNow(t, n.Type() == ast.FuncDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	f := n.(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, f.Signature != nil, "signature isn't nil")
	bvmUtils.AssertNow(t, len(f.Signature.Parameters) == 2, "wrong param length")
}

func TestParseConstructorNoParameters(t *testing.T) {
	p := createParser(`constructor(){}`)
	bvmUtils.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n != nil, "node is nil")
	bvmUtils.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	c := n.(*ast.LifecycleDeclarationNode)
	bvmUtils.AssertNow(t, len(c.Parameters) == 0, "wrong param length")
}

func TestParseConstructorOneParameter(t *testing.T) {
	p := createParser(`constructor(a int){}`)
	bvmUtils.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	c := n.(*ast.LifecycleDeclarationNode)
	bvmUtils.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorParameters(t *testing.T) {
	p := createParser(`constructor(a int, b string){}`)
	bvmUtils.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	c := n.(*ast.LifecycleDeclarationNode)
	bvmUtils.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
}

func TestParseConstructorMultiplePerType(t *testing.T) {
	p := createParser(`constructor(a, b int){}`)
	bvmUtils.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	c := n.(*ast.LifecycleDeclarationNode)
	bvmUtils.AssertNow(t, len(c.Parameters) == 1, "wrong param length")
}

func TestParseConstructorMultiplePerTypeExtra(t *testing.T) {
	p := createParser(`constructor(a, b int, c []string){}`)
	bvmUtils.Assert(t, isLifecycleDeclaration(p), "should detect Constructor decl")
	parseLifecycleDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.LifecycleDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	c := n.(*ast.LifecycleDeclarationNode)
	bvmUtils.AssertNow(t, len(c.Parameters) == 2, "wrong param length")
	first := c.Parameters[0]
	second := c.Parameters[1]
	bvmUtils.AssertNow(t, first.DeclaredType != nil, "first dt shouldn't be nil")
	bvmUtils.AssertNow(t, first.DeclaredType.Type() == ast.PlainType, "first dt should be plain")
	bvmUtils.AssertNow(t, second.DeclaredType != nil, "second dt shouldn't be nil")
	bvmUtils.AssertNow(t, second.DeclaredType.Type() == ast.ArrayType, "second dt should be array")
}

func TestParseParametersSingleVarSingleType(t *testing.T) {
	p := createParser(`(a string)`)
	exps := p.parseParameters()
	bvmUtils.AssertNow(t, exps != nil, "params not nil")
	bvmUtils.AssertNow(t, len(exps) == 1, "params of length 1")
	bvmUtils.AssertNow(t, exps[0].DeclaredType != nil, "declared type shouldn't be nil")
	bvmUtils.AssertNow(t, exps[0].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	bvmUtils.AssertNow(t, len(exps[0].Identifiers) == 1, "wrong parameter length")
}

func TestParseParametersMultipleVarSingleType(t *testing.T) {
	p := createParser(`(a, b string)`)
	exps := p.parseParameters()
	bvmUtils.AssertNow(t, exps != nil, "params not nil")
	bvmUtils.AssertNow(t, len(exps) == 1, "params of length 1")
	bvmUtils.AssertNow(t, exps[0].DeclaredType != nil, "declared type shouldn't be nil")
	bvmUtils.AssertNow(t, exps[0].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	bvmUtils.AssertNow(t, len(exps[0].Identifiers) == 2, "wrong parameter length")
}

func TestParseParametersSingleVarMultipleType(t *testing.T) {
	p := createParser(`(a string, b int)`)
	exps := p.parseParameters()
	bvmUtils.AssertNow(t, exps != nil, "params not nil")
	bvmUtils.AssertNow(t, len(exps) == 2, "params of length 1")
	bvmUtils.AssertNow(t, exps[0].DeclaredType != nil, "declared type shouldn't be nil")
	bvmUtils.AssertNow(t, exps[0].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	bvmUtils.AssertNow(t, exps[1].DeclaredType != nil, "declared type shouldn't be nil")
	bvmUtils.AssertNow(t, exps[1].DeclaredType.Type() == ast.PlainType, "wrong declared type")
	bvmUtils.AssertNow(t, len(exps[0].Identifiers) == 1, "wrong parameter length")
	bvmUtils.AssertNow(t, len(exps[1].Identifiers) == 1, "wrong parameter length")
}

func TestParseFuncSignatureNoParamOneResult(t *testing.T) {
	p := createParser("hi() string")
	sig := p.parseFuncSignature()
	bvmUtils.AssertNow(t, sig != nil, "nil signature")
	bvmUtils.AssertLength(t, len(sig.Parameters), 0)
	bvmUtils.AssertLength(t, len(sig.Results), 1)
}

func TestParseComponentFuncType(t *testing.T) {
	p := createParser("func(string, string) (int, int)")
	sig := p.parseFuncType()
	bvmUtils.AssertNow(t, sig != nil, "nil signature")
	bvmUtils.AssertLength(t, len(sig.Parameters), 2)
	bvmUtils.AssertLength(t, len(sig.Results), 2)
}

func TestParseComplexFuncParams(t *testing.T) {
	p := createParser("c int, b func(string, string) (int, int), d int")
	params := p.parseFuncTypeParameters()
	bvmUtils.AssertLength(t, len(params), 3)
	names := []string{"c", "b", "d"}
	for i, p := range params {
		bvmUtils.Assert(t, p.Type() == ast.ExplicitVarDeclaration, fmt.Sprintf("param %d of wrong type"))
		e := p.(*ast.ExplicitVarDeclarationNode)
		bvmUtils.AssertLength(t, len(e.Identifiers), 1)
		bvmUtils.AssertNow(t, e.Identifiers[0] == names[i], "wrong name")
	}
}

func TestParseComplexFuncType(t *testing.T) {
	p := createParser("func(c int, b func(string, string) (int, int), d int) (int, float, int)")
	sig := p.parseFuncType()
	bvmUtils.AssertNow(t, sig != nil, "nil signature")
	bvmUtils.AssertLength(t, len(sig.Parameters), 3)
	bvmUtils.AssertLength(t, len(sig.Results), 3)
	p = createParser("func(int, func(string, string) (int, int), int) (a int, b float, c int)")
	sig = p.parseFuncType()
	bvmUtils.AssertNow(t, sig != nil, "nil signature")
	bvmUtils.AssertLength(t, len(sig.Parameters), 3)
	bvmUtils.AssertLength(t, len(sig.Results), 3)
}

func TestParseSimplePlainType(t *testing.T) {
	p := createParser("int")
	pt := p.parsePlainType()
	bvmUtils.Assert(t, !pt.Variable, "should be variable")
	bvmUtils.AssertLength(t, len(pt.Names), 1)
	bvmUtils.AssertLength(t, len(pt.Parameters), 0)
}

func TestParseReferencePlainType(t *testing.T) {
	p := createParser("int.a")
	pt := p.parsePlainType()
	bvmUtils.Assert(t, !pt.Variable, "should be variable")
	bvmUtils.AssertLength(t, len(pt.Names), 2)
	bvmUtils.AssertLength(t, len(pt.Parameters), 0)
}

func TestParseSimpleGenericPlainType(t *testing.T) {
	p := createParser("int<string>")
	pt := p.parsePlainType()
	bvmUtils.Assert(t, !pt.Variable, "should be variable")
	bvmUtils.AssertLength(t, len(pt.Names), 1)
	bvmUtils.AssertLength(t, len(pt.Parameters), 1)
}

func TestParseReferenceGenericPlainType(t *testing.T) {
	p := createParser("int.a<string>")
	pt := p.parsePlainType()
	bvmUtils.Assert(t, !pt.Variable, "should be variable")
	bvmUtils.AssertLength(t, len(pt.Names), 2)
	bvmUtils.AssertLength(t, len(pt.Parameters), 1)
}

func TestParseReferenceMultipleGenericPlainType(t *testing.T) {
	p := createParser("int.a<string|int>")
	pt := p.parsePlainType()
	bvmUtils.Assert(t, !pt.Variable, "should be variable")
	bvmUtils.AssertLength(t, len(pt.Names), 2)
	bvmUtils.AssertLength(t, len(pt.Parameters), 2)
}

func TestIsNamedParameter(t *testing.T) {
	p := createParser("a string")
	bvmUtils.Assert(t, p.isNamedParameter(), "plain should be np")
	p = createParser("a, b string")
	bvmUtils.Assert(t, p.isNamedParameter(), "multiple should be np")
}

func TestExplicitVarAssignment(t *testing.T) {
	p := createParser("const a = 5")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "expvar not recognised")
	v := p.parseOptionallyTypedVarDeclaration()
	bvmUtils.Assert(t, v != nil, "shouldn't be nil")
	bvmUtils.AssertLength(t, len(v.Identifiers), 1)
	p = createParser("const a = 5")
	parseExplicitVarDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ExplicitVarDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	v = n.(*ast.ExplicitVarDeclarationNode)
	bvmUtils.AssertLength(t, len(v.Identifiers), 1)
	bvmUtils.AssertNow(t, v.Value != nil, "nil value")
	bvmUtils.AssertNow(t, v.IsConstant, "should be constant")
}

func TestExplicitVarAssignmentGrouped(t *testing.T) {
	p := createParser(`const (
		a = uint(5)
	)`)
	bvmUtils.AssertNow(t, isExplicitVarDeclaration(p), "not recognised")
	parseExplicitVarDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.ExplicitVarDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	v := n.(*ast.ExplicitVarDeclarationNode)
	bvmUtils.AssertLength(t, len(v.Identifiers), 1)
	bvmUtils.AssertNow(t, v.Value != nil, "nil value")
	bvmUtils.AssertNow(t, v.IsConstant, "should be constant")
}

func TestParseVariableDeclarations(t *testing.T) {
	ast, errs := ParseString(`
		var balance func(a address) uint256
		var transfer func(a address, amount uint256) uint
		var send func(a address, amount uint256) bool
		var call func(a address) bool
		var delegateCall func(a address)
	`)
	bvmUtils.AssertNow(t, ast != nil, "nil ast")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestInterfaceMethods(t *testing.T) {
	a, errs := ParseString(`
		interface Calculator {
			add(a, b int) int
			sub(a, b int) int
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
	f := a.Declarations.Next()
	ifc := f.(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertLength(t, len(ifc.Signatures), 2)
}

func TestFuncDeclarationEnclosedParams(t *testing.T) {
	a, errs := ParseString(`
		func hi() (int, int, string){

		}
		var x func() (int, int, string)
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
}

func TestInvalidArrayTypeSize(t *testing.T) {
	a, errs := ParseString(`
		var x ["hi"]string
	`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
}

func TestFuncDeclarationResults(t *testing.T) {
	a, errs := ParseString(`
		func hi(){

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations != nil, "nil declarations")
	f := a.Declarations.Next().(*ast.FuncDeclarationNode)
	bvmUtils.AssertLength(t, len(f.Signature.Parameters), 0)
	bvmUtils.AssertLength(t, len(f.Signature.Results), 0)
}

func TestParseFuncTypeNoResults(t *testing.T) {
	p := createParser("trade()")
	sig := p.parseFuncSignature()
	bvmUtils.AssertNow(t, sig != nil, "signature is nil")
	bvmUtils.AssertNow(t, sig.Identifier == "trade", "wrong identifier: "+sig.Identifier)
	bvmUtils.AssertNow(t, len(sig.Parameters) == 0, "wrong parameter length")
	bvmUtils.AssertNow(t, len(sig.Results) == 0, "wrong result length")
}

func TestParseFuncTypeEmptyBracketResults(t *testing.T) {
	p := createParser("trade() ()")
	sig := p.parseFuncSignature()
	bvmUtils.AssertNow(t, sig != nil, "signature is nil")
	bvmUtils.AssertNow(t, sig.Identifier == "trade", "wrong identifier: "+sig.Identifier)
	bvmUtils.AssertNow(t, len(sig.Parameters) == 0, "wrong parameter length")
	bvmUtils.AssertNow(t, len(sig.Results) == 0, "wrong result length")
}

func TestParseEventParameterSimple(t *testing.T) {
	p := createParser(`event Notification(a, b string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
	one := e.Parameters[0]
	bvmUtils.AssertLength(t, len(one.Modifiers.Modifiers), 0)
	bvmUtils.AssertLength(t, len(one.Identifiers), 2)
}

func TestParseEventParametersSimple(t *testing.T) {
	p := createParser(`event Notification(a, b string, c int)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
	one := e.Parameters[0]
	bvmUtils.AssertLength(t, len(one.Modifiers.Modifiers), 0)
	bvmUtils.AssertLength(t, len(one.Identifiers), 2)
	two := e.Parameters[1]
	bvmUtils.AssertLength(t, len(two.Modifiers.Modifiers), 0)
	bvmUtils.AssertLength(t, len(two.Identifiers), 1)
}

func TestParseEventParameterSingleModifier(t *testing.T) {
	p := createParser(`event Notification(indexed a string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
	one := e.Parameters[0]
	bvmUtils.AssertLength(t, len(one.Modifiers.Modifiers), 1)
}

func TestParseEventParameterDoubleModifier(t *testing.T) {
	p := createParser(`event Notification(indexed indexed a string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 1, "wrong param length")
	one := e.Parameters[0]
	bvmUtils.AssertLength(t, len(one.Modifiers.Modifiers), 2)
}

func TestParseEventParametersSingleModifier(t *testing.T) {
	p := createParser(`event Notification(indexed a string, b, c string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
	one := e.Parameters[0]
	bvmUtils.AssertLength(t, len(one.Modifiers.Modifiers), 1)
	two := e.Parameters[1]
	bvmUtils.AssertLength(t, len(two.Modifiers.Modifiers), 0)
	bvmUtils.AssertLength(t, len(two.Identifiers), 2)
}

func TestParseEventParametersDoubleModifier(t *testing.T) {
	p := createParser(`event Notification(indexed indexed a string, indexed b string)`)
	bvmUtils.Assert(t, isEventDeclaration(p), "should detect event decl")
	parseEventDeclaration(p)
	n := p.scope.NextDeclaration()
	bvmUtils.AssertNow(t, n.Type() == ast.EventDeclaration, "Asteroid Errors: Node Error: Wrong node type. ")
	e := n.(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 2, "wrong param length")
	one := e.Parameters[0]
	bvmUtils.AssertLength(t, len(one.Modifiers.Modifiers), 2)
	two := e.Parameters[1]
	bvmUtils.AssertLength(t, len(two.Modifiers.Modifiers), 1)
}

func TestParseFullProblematicFunc(t *testing.T) {
	_, errs := ParseString(`
		external func burnFrom(from address, value uint256) bool {
			this.allowance[from][msg.sender] -= value;     // Subtract from the sender's allowance
		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseLifecycleParameters(t *testing.T) {
	a, errs := ParseString(`
		constructor(age int8){

		}
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
	bvmUtils.AssertNow(t, a != nil, "a is nil")
	n := a.Declarations.Next().(*ast.LifecycleDeclarationNode)
	bvmUtils.AssertNow(t, len(n.Parameters) == 1, "wrong param length")
	bvmUtils.AssertNow(t, len(n.Parameters[0].Identifiers) == 1, "wrong id length")
	bvmUtils.AssertNow(t, n.Parameters[0].Identifiers[0] == "age", fmt.Sprintf("wrong param name: %s", n.Parameters[0].Identifiers[0]))
}

func TestParseParameterIdentifiers(t *testing.T) {
	p := createParser(`(age int8)`)
	params := p.parseParameters()
	bvmUtils.AssertNow(t, len(params) == 1, "wrong param length")
	bvmUtils.AssertNow(t, len(params[0].Identifiers) == 1, "wrong id length")
	bvmUtils.AssertNow(t, params[0].Identifiers[0] == "age", fmt.Sprintf("wrong param name: %s", params[0].Identifiers[0]))
}

func TestParseUnimplementedFuncDeclaration(t *testing.T) {
	_, errs := ParseString(`@Builtin("a") func add(a, b int) int`)
	bvmUtils.AssertNow(t, len(errs) == 0, "wrong error length")
}

func TestParseUnimplementedFuncDeclarationNoBuiltin(t *testing.T) {
	_, errs := ParseString(`func add(a, b int) int`)
	// one for each missing bracket
	bvmUtils.AssertNow(t, len(errs) == 2, "wrong error length")
}

func TestParseVarDeclarationWithValue(t *testing.T) {
	_, errs := ParseString(`var (
        WEIPERETH uint256  = 1000000000000000000
    )`)
	// one for each missing bracket
	bvmUtils.AssertNow(t, len(errs) == 0, "wrong error length")
}

func TestParseInterfaceDeclarationWithModifiers(t *testing.T) {
	_, errs := ParseString(`interface Switchable {
			global on()
			internal off()
			internal external ok(a, b string) (a int)
		}`)
	// one for each missing bracket
	bvmUtils.AssertNow(t, len(errs) == 0, "wrong error length")
}

func TestParseEventDeclarationWithModifiers(t *testing.T) {
	a, errs := ParseString(`
		event Transfer(indexed from address, indexed to address, value uint)
	`)
	bvmUtils.AssertNow(t, len(errs) == 0, "wrong error length")
	e := a.NextDeclaration().(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, len(e.Parameters) == 3, "wrong parameter length")
	bvmUtils.AssertNow(t, len(e.Parameters[0].Identifiers) == 1, "wrong parameter 0 length")
	bvmUtils.AssertNow(t, len(e.Parameters[1].Identifiers) == 1, "wrong parameter 1 length")
	bvmUtils.AssertNow(t, len(e.Parameters[2].Identifiers) == 1, "wrong parameter 2 length")
}
