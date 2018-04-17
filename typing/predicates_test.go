package typing

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestIsAssignableEqualTypes(t *testing.T) {
	a := standards[boolean]
	b := standards[boolean]
	bvmUtils.AssertNow(t, AssignableTo(b, a, false), "equal types should be Assignable")
}

func TestIsAssignableSuperClass(t *testing.T) {
	a := &Class{Name: "OrderBook"}
	b := &Class{Name: "assetType", Supers: []*Class{a}}
	bvmUtils.AssertNow(t, AssignableTo(b, a, false), "super class types should be Assignable")
}

func TestIsAssignableMultipleSuperClass(t *testing.T) {
	a := &Class{Name: "OrderBook"}
	b := &Class{Name: "assetType", Supers: []*Class{a}}
	c := &Class{Name: "Rat", Supers: []*Class{b}}
	bvmUtils.AssertNow(t, AssignableTo(c, a, false), "super class types should be Assignable")
}

func TestIsAssignableParentInterface(t *testing.T) {
	a := &Interface{Name: "OrderBook"}
	b := &Interface{Name: "assetType", Supers: []*Interface{a}}
	bvmUtils.AssertNow(t, AssignableTo(b, a, false), "interface types should be Assignable")
}

func TestIsAssignableClassImplementingInterface(t *testing.T) {
	a := &Interface{Name: "OrderBook"}
	b := &Class{Name: "assetType", Interfaces: []*Interface{a}}
	bvmUtils.AssertNow(t, AssignableTo(b, a, false), "interface types should be Assignable")
}

func TestIsAssignableSuperClassImplementingInterface(t *testing.T) {
	a := &Interface{Name: "OrderBook"}
	b := &Class{Name: "assetType", Interfaces: []*Interface{a}}
	c := &Class{Name: "assetType", Supers: []*Class{b}}
	bvmUtils.AssertNow(t, AssignableTo(c, a, false), "interface types should be Assignable")
}

func TestIsAssignableSuperClassImplementingSuperInterface(t *testing.T) {
	a := &Interface{Name: "OrderBook"}
	b := &Interface{Name: "UserA", Supers: []*Interface{a}}
	c := &Class{Name: "assetType", Interfaces: []*Interface{b}}
	d := &Class{Name: "UserB", Supers: []*Class{c}}
	bvmUtils.AssertNow(t, AssignableTo(d, a, false), "type should be Assignable")
}

func TestIsAssignableClassDoesNotInherit(t *testing.T) {
	c := &Class{Name: "assetType"}
	d := &Class{Name: "UserB"}
	bvmUtils.AssertNow(t, !AssignableTo(d, c, false), "class should not be Assignable")
}

func TestIsAssignableClassFlipped(t *testing.T) {
	d := &Class{Name: "UserB"}
	c := &Class{Name: "assetType", Supers: []*Class{d}}
	bvmUtils.AssertNow(t, !AssignableTo(d, c, true), "class should not be Assignable")
}

func TestIsAssignableClassInterfaceNot(t *testing.T) {
	c := &Class{Name: "assetType"}
	d := &Interface{Name: "UserB"}
	bvmUtils.AssertNow(t, !AssignableTo(d, c, true), "class should not be Assignable")
	bvmUtils.AssertNow(t, !AssignableTo(c, d, true), "interface should not be Assignable")
}

func TestIsAssignableBooleans(t *testing.T) {
	a := Boolean()
	b := Boolean()
	bvmUtils.Assert(t, AssignableTo(a, b, false), "a --> b")
	bvmUtils.Assert(t, AssignableTo(b, a, false), "b --> a")
}

func TestIsAssignableNamedBooleans(t *testing.T) {
	a := &Aliased{Alias: "a", Underlying: Boolean()}
	b := &Aliased{Alias: "b", Underlying: Boolean()}
	bvmUtils.Assert(t, AssignableTo(a, b, false), "a --> b")
	bvmUtils.Assert(t, AssignableTo(b, a, false), "b --> a")
}

func TestIsAssignableFuncParams(t *testing.T) {
	a := NewTuple(Boolean())
	b := NewTuple(&Aliased{Alias: "b", Underlying: Boolean()})
	bvmUtils.Assert(t, AssignableTo(a, b, false), "a --> b")
	bvmUtils.Assert(t, AssignableTo(b, a, false), "b --> a")
}

func TestNumericAssignability(t *testing.T) {
	a := &NumericType{BitSize: 10, Signed: true}
	b := &NumericType{BitSize: 10, Signed: true}
	bvmUtils.Assert(t, AssignableTo(a, b, false), "a --> b")
	// ints --> larger ints
	a = &NumericType{BitSize: 10, Signed: true}
	b = &NumericType{BitSize: 11, Signed: true}
	bvmUtils.Assert(t, AssignableTo(b, a, false), "b --> a")
	bvmUtils.Assert(t, !AssignableTo(a, b, false), "a --> b")
	// uints --> larger uints
	a = &NumericType{BitSize: 10, Signed: false}
	b = &NumericType{BitSize: 11, Signed: false}
	bvmUtils.Assert(t, AssignableTo(b, a, false), "b --> a")
	bvmUtils.Assert(t, !AssignableTo(a, b, false), "a --> b")
	// uints --> larger ints
	a = &NumericType{BitSize: 10, Signed: false}
	b = &NumericType{BitSize: 11, Signed: true}
	bvmUtils.Assert(t, AssignableTo(b, a, false), "b --> a")
	bvmUtils.Assert(t, !AssignableTo(a, b, false), "a --> b")
	// can never go int --> uint
	a = &NumericType{BitSize: 10, Signed: false}
	b = &NumericType{BitSize: 100, Signed: true}
	bvmUtils.Assert(t, !AssignableTo(a, b, false), "a --> b")
}

func TestIsAssignableUnknown(t *testing.T) {
	a := Unknown()
	b := Boolean()
	bvmUtils.Assert(t, !AssignableTo(a, b, false), "a --> b")
	bvmUtils.Assert(t, AssignableTo(a, b, true), "a --> b")
}

func TestIsAssignableEmptyTuples(t *testing.T) {
	a := NewTuple(Unknown())
	b := NewTuple()
	bvmUtils.Assert(t, !AssignableTo(a, b, true), "a --> b")
}
