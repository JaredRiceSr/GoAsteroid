package typing

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestConvertToBits(t *testing.T) {
	bvmUtils.Assert(t, BitsNeeded(0) == 1, "wrong 0")
	bvmUtils.Assert(t, BitsNeeded(1) == 1, "wrong 1")
	bvmUtils.Assert(t, BitsNeeded(2) == 2, "wrong 2")
	bvmUtils.Assert(t, BitsNeeded(10) == 4, "wrong 10")
}

func TestAcceptLiteral(t *testing.T) {
	a := &NumericType{BitSize: 8, Signed: true, Integer: true}
	bvmUtils.AssertNow(t, a.AcceptsLiteral(8, true, true), "1 failed")
	bvmUtils.AssertNow(t, a.AcceptsLiteral(8, true, false), "2 failed")
	bvmUtils.AssertNow(t, a.AcceptsLiteral(7, true, true), "3 failed")
	b := &NumericType{BitSize: 256, Signed: false, Integer: true}
	bvmUtils.AssertNow(t, b.AcceptsLiteral(8, true, false), "4 failed")
}
