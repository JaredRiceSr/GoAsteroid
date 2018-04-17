package validator

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestSampleClass(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/class.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleContract(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/contracts.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleLoops(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/loops.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleSwitching(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/switching.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleTypes(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/types.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleIterator(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/iterator.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleKV(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/kv.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleAccessRestriction(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/common_patterns/access_restriction.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleRichest(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/common_patterns/richest.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestSampleStateMachine(t *testing.T) {
	_, errs := ValidateFile(NewTestBVM(), nil, "../samples/common_patterns/state_machine.grd")
	bvmUtils.AssertNow(t, len(errs) == 0, errs.Format())
}

func TestEmptyContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/empty.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestConstructorContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/constructors.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestFuncsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/funcs.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestClassesContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/classes.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestInterfacesContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/interfaces.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, fmt.Sprintln(errs))
}

func TestEventsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/events.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestEnumsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/enums.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestTypesContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/types.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestNestedModifiersContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/nested_modifiers.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestCommentsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/comments.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestGroupsContract(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/groups.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}
