package parser

import (
	"fmt"
	"testing"

	"github.com/benchlab/asteroid/ast"

	"github.com/benchlab/bvmUtils"
)

func TestParseSimpleClassGeneric(t *testing.T) {
	p := createParser("class List<T> {}")
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseExtendingClassGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementingClassGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseExtendsImplementsClassGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item is Comparable> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementsExtendsClassGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable inherits Item> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseMultipleExtendsandImplementsClassGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable, Real inherits Item, Dog> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseSimpleClassMultipleGeneric(t *testing.T) {
	p := createParser("class List<T|S|R> {}")
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 3, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseExtendingClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item | S inherits Dog> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementingClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Item | S is Dog> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))
}

func TestParseExtendsImplementsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T inherits Item is Comparable | S is Comparable> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseImplementsExtendsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable inherits Item | S inherits Item> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 2, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseMultipleExtendsandImplementsClassMultipleGeneric(t *testing.T) {
	p := createParser(`class List<T is Comparable, Real inherits Item, Dog | S | R is Comparable inherits Item, Dog> {}`)
	parseClassDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.ClassDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 3, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))

}

func TestParseSingleGenericFunction(t *testing.T) {
	p := createParser(`func <T> hello(){}`)
	parseFuncDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations != nil, "declarations is nil")
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertNow(t, len(c.Generics) == 1, fmt.Sprintf("wrong gen len: %d", len(c.Generics)))
}

func TestParseMultipleGenericFunction(t *testing.T) {
	p := createParser(`func <T|S|R> hello(){}`)
	parseFuncDeclaration(p)
	bvmUtils.AssertNow(t, p.scope.Declarations != nil, "declarations is nil")
	bvmUtils.AssertNow(t, p.scope.Declarations.Length() == 1, "wrong length")
	c := p.scope.Declarations.Next().(*ast.FuncDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 3)
}

func TestParseGenerics(t *testing.T) {
	p := createParser("<T>")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 1)
}

func TestParseGenericsDouble(t *testing.T) {
	p := createParser("<T|S>")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 2)
}

func TestParseGenericsTriple(t *testing.T) {
	p := createParser("<T|S|R>")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 3)
}

func TestParseGenericsInheritance(t *testing.T) {
	p := createParser("<T inherits A>")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 1)
	bvmUtils.AssertLength(t, len(gens[0].Inherits), 1)
}

func TestParseGenericsImplementation(t *testing.T) {
	p := createParser("<T is A>")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 1)
	bvmUtils.AssertLength(t, len(gens[0].Implements), 1)
}

func TestParseGenericsImplementationDouble(t *testing.T) {
	p := createParser("<T is A|S is B>")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 2)
	bvmUtils.AssertLength(t, len(gens[0].Implements), 1)
	bvmUtils.AssertLength(t, len(gens[1].Implements), 1)
}

func TestParseGenericComplex(t *testing.T) {
	text := "<T is Comparable, Real inherits Item, Dog | S | R is Comparable inherits Item, Dog>"
	p := createParser(text)
	bvmUtils.AssertNow(t, p.index == 0, "wrong starting index")
	gens := p.parseGenerics()
	bvmUtils.AssertNow(t, gens != nil, "nil generics")
	bvmUtils.AssertLength(t, len(gens), 3)
	bvmUtils.AssertNow(t, p.index == len(p.lexer.Tokens), "wrong ending index")
}

func TestParseEventGenericSimpleMultiple(t *testing.T) {
	a, errs := ParseString(`event <T|S|R> hello(a T, b S, c R)`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations.Length() == 1, "wrong length")
	c := a.Declarations.Next().(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 3)
}

func TestParseEventGenericSimpleSingle(t *testing.T) {
	a, errs := ParseString(`event <T> hello(a T)`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations.Length() == 1, "wrong length")
	c := a.Declarations.Next().(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 1)
}

func TestParseEventGenericComplexSingle(t *testing.T) {
	a, errs := ParseString(`event <T> hello(a T)`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations.Length() == 1, "wrong length")
	c := a.Declarations.Next().(*ast.EventDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 1)
}

func TestParseInterfaceGenericsSimpleSingle(t *testing.T) {
	a, errs := ParseString(`interface Dog<T>{}`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations.Length() == 1, "wrong length")
	c := a.Declarations.Next().(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 1)
}

func TestParseInterfaceGenericsSimpleMultiple(t *testing.T) {
	a, errs := ParseString(`interface Dog<T|S|R>{}`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations.Length() == 1, "wrong length")
	c := a.Declarations.Next().(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 3)
}

func TestParseInterfaceGenericsComplexSingle(t *testing.T) {
	a, errs := ParseString(`interface Dog<T inherits Cat>{}`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
	bvmUtils.AssertNow(t, a.Declarations.Length() == 1, "wrong length")
	c := a.Declarations.Next().(*ast.InterfaceDeclarationNode)
	bvmUtils.AssertNow(t, c.Generics != nil, "nil generics")
	bvmUtils.AssertLength(t, len(c.Generics), 1)
}

func TestParseSimpleGenericAssignment(t *testing.T) {
	_, errs := ParseString(`x = new List<string>()`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
}

func TestParseMultipleGenericAssignment(t *testing.T) {
	_, errs := ParseString(`x = new List<string|int>()`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
}

func TestParseFullGenerics(t *testing.T) {
	_, errs := ParseString(`
		class List<T> {

		}

		class Dog<T> inherits List<T> {

		}

		a = new Dog<string>()
		b = new Dog<List<string>>()
		c = new Dog<Dog<Dog<string>>>()
		d = new Dog<List<List<List<string>>>>()
	`)
	bvmUtils.AssertNow(t, errs == nil, errs.Format())
}
