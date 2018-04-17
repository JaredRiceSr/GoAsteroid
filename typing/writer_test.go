package typing

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestWriteMapType(t *testing.T) {
	m := &Map{Key: standards[boolean], Value: standards[boolean]}
	expected := "map[bool]bool"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteArrayType(t *testing.T) {
	m := &Array{Value: standards[unknown], Length: 0, Variable: true}
	expected := "[]unknown"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteTupleTypeEmpty(t *testing.T) {
	m := NewTuple()
	expected := "()"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteTupleTypeSingle(t *testing.T) {
	m := NewTuple(standards[boolean])
	expected := "(bool)"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteTupleTypeMultiple(t *testing.T) {
	m := NewTuple(standards[boolean], standards[unknown])
	expected := "(bool, unknown)"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsEmptyResults(t *testing.T) {
	m := &Func{Params: NewTuple(), Results: NewTuple()}
	expected := "func ()()"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteFuncEmptyParamsSingleResults(t *testing.T) {
	m := &Func{Params: NewTuple(), Results: NewTuple(standards[boolean])}
	expected := "func ()(bool)"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteFuncMultipleParamsMultipleResults(t *testing.T) {
	m := &Func{Params: NewTuple(standards[boolean], standards[unknown]), Results: NewTuple(standards[boolean], standards[unknown])}
	expected := "func (bool, unknown)(bool, unknown)"
	bvmUtils.Assert(t, WriteType(m) == expected, fmt.Sprintf("The wrong type was written %s\n", WriteType(m)))
}

func TestWriteClass(t *testing.T) {

}
