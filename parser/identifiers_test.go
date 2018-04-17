package parser

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestIdentifierSafety(t *testing.T) {
	p := createParser("")
	// none of these should crash
	for _, c := range getPrimaryConstructs() {
		c.is(p)
	}
}

func TestIsClassDeclaration(t *testing.T) {
	p := createParser("class OrderBook {")
	bvmUtils.Assert(t, isClassDeclaration(p), "class declaration not recognised")
}

func TestIsInterfaceDeclaration(t *testing.T) {
	p := createParser("interface Box {")
	bvmUtils.Assert(t, isInterfaceDeclaration(p), "interface declaration not recognised")
}

func TestIsContractDeclaration(t *testing.T) {
	p := createParser("contract Box {")
	bvmUtils.Assert(t, isContractDeclaration(p), "contract declaration not recognised")
}

func TestIsFuncDeclaration(t *testing.T) {
	p := createParser("func main(){")
	bvmUtils.Assert(t, isFuncDeclaration(p), "function declaration not recognised")
	p = createParser("func main() int {")
	bvmUtils.Assert(t, isFuncDeclaration(p), "returning function declaration not recognised")
	p = createParser("func main() (int, int) {")
	bvmUtils.Assert(t, isFuncDeclaration(p), "tuple returning function declaration not recognised")
}

func TestIsTypeDeclaration(t *testing.T) {
	p := createParser("type Large int")
	bvmUtils.Assert(t, isTypeDeclaration(p), "type declaration not recognised")
	p = createParser("type Large []int")
	bvmUtils.Assert(t, isTypeDeclaration(p), "array type declaration not recognised")
	p = createParser("type Large map[int]string")
	bvmUtils.Assert(t, isTypeDeclaration(p), "map type declaration not recognised")
}

func TestIsReturnStatement(t *testing.T) {
	p := createParser("return 0")
	bvmUtils.Assert(t, isReturnStatement(p), "return statement not recognised")
	p = createParser("return (0, 0)")
	bvmUtils.Assert(t, isReturnStatement(p), "tuple return statement not recognised")
}

func TestIsForStatement(t *testing.T) {
	p := createParser("for i := 0; i < 10; i++ {}")
	bvmUtils.Assert(t, isForStatement(p), "for statement not recognised")
	p = createParser("for i < 10 {}")
	bvmUtils.Assert(t, isForStatement(p), "cond only for statement not recognised")
	p = createParser("for i in 0...10{}")
	bvmUtils.Assert(t, isForStatement(p), "range for statement not recognised")
	p = createParser("for i, _ in array")
	bvmUtils.Assert(t, isForStatement(p), "")
}

func TestIsIfStatement(t *testing.T) {
	p := createParser("if x > 5 {}")
	bvmUtils.Assert(t, isIfStatement(p), "if statement not recognised")
	p = createParser("if p := getData(); p < 5 {}")
	bvmUtils.Assert(t, isIfStatement(p), "init if statement not recognised")
}

func TestIsExplicitVarDeclaration(t *testing.T) {
	p := createParser("var x string")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "expvar statement not recognised")
	p = createParser("var x, a string")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "multiple var expvar statement not recognised")
	p = createParser("var x map[string]string")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "map expvar statement not recognised")
	p = createParser("var x []string")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "array expvar statement not recognised")
	p = createParser("var transfer func(a address, amount uint256) uint")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "func type statement not recognised")
	p = createParser("x = 5")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise simple assignment")
	p = createParser("a[b] = 5")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise index assignment")
	p = createParser("a[b].c()")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise reference call")
	p = createParser("")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise empty string")
	p = createParser("}")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise empty string")
	p = createParser("contract OrderBook {}")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise contract opening")
	p = createParser("empty()")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise empty call")
}

func TestIsExpVarFunc(t *testing.T) {
	p := createParser("var blockhash func(blockNumber uint) [32]byte")
	bvmUtils.Assert(t, isExplicitVarDeclaration(p), "second func type statement not recognised")
}

func TestIsExpVarCall(t *testing.T) {

	p := createParser(`full("hi", "bye")`)
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not recognise empty call")
}

func TestIsSwitchStatement(t *testing.T) {
	p := createParser("switch x {}")
	bvmUtils.Assert(t, isSwitchStatement(p), "switch statement not recognised")
	p = createParser("exclusive switch x {}")
	bvmUtils.Assert(t, isSwitchStatement(p), "exclusive switch statement not recognised")
}

func TestIsCaseStatement(t *testing.T) {
	p := createParser("case 1, 2, 3 { break }")
	bvmUtils.Assert(t, isCaseStatement(p), "multi case statement not recognised")
	p = createParser("case 1 { break }")
	bvmUtils.Assert(t, isCaseStatement(p), "single case statement not recognised")
}

func TestIsEventDeclaration(t *testing.T) {
	p := createParser("event Notification()")
	bvmUtils.Assert(t, isEventDeclaration(p), "empty event not recognised")
	p = createParser("event Notification(string)")
	bvmUtils.Assert(t, isEventDeclaration(p), "single event not recognised")
	p = createParser("event Notification(string, exchange.OrderBook)")
	bvmUtils.Assert(t, isEventDeclaration(p), "multiple event not recognised")
}

func TestIsNextAssignmentStatement(t *testing.T) {
	p := createParser("++")
	bvmUtils.Assert(t, p.isNextTokenAssignment(), "simple increment not recognised")
	p = createParser("--")
	bvmUtils.Assert(t, p.isNextTokenAssignment(), "simple decrement not recognised")

}

func TestIsMapType(t *testing.T) {
	p := createParser("map[string]string")
	bvmUtils.Assert(t, p.isMapType(), "map type not recognised")
	p = createParser("[string]")
	bvmUtils.Assert(t, !p.isMapType(), "index array type should not be recognised")
	p = createParser("a[b]")
	bvmUtils.Assert(t, !p.isMapType(), "index array type should not be recognised")
}

func TestIsArrayType(t *testing.T) {
	p := createParser("[]string")
	bvmUtils.Assert(t, p.isArrayType(), "array type not recognised")

	p = createParser("a[b]")
	bvmUtils.Assert(t, !p.isArrayType(), "index array type should not be recognised")
}

func TestIsPlainType(t *testing.T) {
	p := createParser("string")
	bvmUtils.Assert(t, p.isPlainType(), "simple type not recognised")

	p = createParser("string.hi")
	bvmUtils.Assert(t, p.isPlainType(), "reference type not recognised")
	p = createParser("[string]")
	bvmUtils.Assert(t, !p.isPlainType(), "index array type should not be recognised")
	p = createParser("a[b]")
	bvmUtils.Assert(t, !p.isPlainType(), "index expr type should not be recognised")
	p = createParser("call()")
	bvmUtils.Assert(t, !p.isPlainType(), "empty call type should not be recognised")
	p = createParser(`full("hi", "bye")`)
	bvmUtils.Assert(t, !p.isPlainType(), "full call type should not be recognised")
	p = createParser(`full("hi", "bye")
	`)
	bvmUtils.Assert(t, !p.isPlainType(), "multiline full call type should not be recognised")
}

func TestVariableTypes(t *testing.T) {
	p := createParser("...map[string]string")
	bvmUtils.Assert(t, p.isMapType(), "variable map type not recognised")
	p = createParser("...[]string")
	bvmUtils.Assert(t, p.isArrayType(), "variable array type not recognised")
	p = createParser("...string")
	bvmUtils.Assert(t, p.isPlainType(), "variable type not recognised")
}

func TestIsFuncType(t *testing.T) {
	p := createParser("func(a address, amount uint256) uint")
	bvmUtils.Assert(t, p.isFuncType(), "func type not recognised")
	p = createParser("func(blockNumber uint) [32]byte")
	bvmUtils.Assert(t, p.isFuncType(), "second func type not recognised")

}

func TestIsNextAType(t *testing.T) {
	p := createParser(")")
	bvmUtils.Assert(t, !p.isNextAType(), "should not be a type")
	p = createParser(") (int, int), d int) (int, float, int)")
	bvmUtils.Assert(t, !p.isNextAType(), "should not be a type")
}

func TestNotExpVar(t *testing.T) {
	p := createParser("string, string)")
	bvmUtils.Assert(t, !isExplicitVarDeclaration(p), "should not be an expvar")
}

func TestAnnotation(t *testing.T) {
	p := createParser("@Builtin()")
	bvmUtils.Assert(t, isAnnotation(p), "annotation not detected")
}

func TestIsModifier(t *testing.T) {
	//p := createParser("public class OrderBook {}")
	//bvmUtils.Assert(t, isModifier(p), "1 modifier not detected")
	//p := createParser("public ( class OrderBook {} )")
	//bvmUtils.Assert(t, isModifier(p), "2 modifier not detected")
	p := createParser("public ( static ( class OrderBook {} ) )")
	bvmUtils.Assert(t, isModifier(p), "3 modifier not detected")
}

func TestIsNotModifier(t *testing.T) {
	p := createParser("call()")
	bvmUtils.Assert(t, !isModifier(p), "modifier detected")
	p = createParser("create(6, 5)")
	bvmUtils.Assert(t, !isModifier(p), "modifier detected")
	p = createParser(`create("hello", "world")`)
	bvmUtils.Assert(t, !isModifier(p), "modifier detected")
	p = createParser(`assert(now() >= auctionEnd)`)
	bvmUtils.Assert(t, !isModifier(p), "modifier 4 detected")

}
