package parser

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

// mini-parser tests belong here

func TestHasTokens(t *testing.T) {
	p := createParser("this is data")
	bvmUtils.Assert(t, len(p.lexer.Tokens) == 3, "wrong length of tokens")
	bvmUtils.Assert(t, p.hasTokens(0), "should have 0 tokens")
	bvmUtils.Assert(t, p.hasTokens(1), "should have 1 tokens")
	bvmUtils.Assert(t, p.hasTokens(2), "should have 2 tokens")
	bvmUtils.Assert(t, p.hasTokens(3), "should have 3 tokens")
	bvmUtils.Assert(t, !p.hasTokens(4), "should not have 4 tokens")
}

func TestParseIdentifier(t *testing.T) {
	p := createParser("identifier")
	bvmUtils.Assert(t, p.parseIdentifier() == "identifier", "wrong identifier")
	p = createParser("")
	bvmUtils.Assert(t, p.parseIdentifier() == "", "empty should be nil")
	p = createParser("{")
	bvmUtils.Assert(t, p.parseIdentifier() == "", "wrong token should be nil")
}

func TestParserNumDeclarations(t *testing.T) {
	a, _ := ParseString(`
		var b int
		var a string
	`)
	bvmUtils.AssertNow(t, a != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, a.Declarations != nil, "scope declarations should not be nil")
	le := a.Declarations.Length()
	bvmUtils.AssertNow(t, le == 2, fmt.Sprintf("wrong decl length: %d", le))
}

func TestParseSimpleStringAnnotationValid(t *testing.T) {
	_, errs := ParseString(`@Builtin("hello")`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseMultipleStringAnnotationValid(t *testing.T) {
	_, errs := ParseString(`@Builtin("hello", "world")`)
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestParseSingleIntegerAnnotationInvalid(t *testing.T) {
	_, errs := ParseString(`@Builtin(6)`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestParseMultipleIntegerAnnotationInvalid(t *testing.T) {
	_, errs := ParseString(`@Builtin(6, 6)`)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}
