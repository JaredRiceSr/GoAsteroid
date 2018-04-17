package parser

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestParseNonExistentFile(t *testing.T) {
	ParseFile("tests/fake_contract.grd")
}

func TestParseString(t *testing.T) {
	ast, _ := ParseString("event OrderBook()")
	bvmUtils.AssertNow(t, ast != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
}

func TestParseStringDeclaration(t *testing.T) {
	as, _ := ParseString("func hello() {}")
	bvmUtils.AssertNow(t, as != nil, "Asteroid Errors: Configured Scope Should Not Be Nil. In This Case, Nil Returned Represents An Error.")
	bvmUtils.AssertNow(t, as.Declarations != nil, "decls should not be nil")
}

func TestParseFile(t *testing.T) {
	ParseFile("../samples/tests/empty.grd")
}

func TestParseType(t *testing.T) {
	typ := ParseType("map[string]string")
	bvmUtils.AssertNow(t, typ != nil, "typ is nil")
}
