package evm

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestBytesRequired(t *testing.T) {
	bvmUtils.Assert(t, bytesRequired(1) == 1, fmt.Sprintf("wrong 1: %d", bytesRequired(1)))
	bvmUtils.Assert(t, bytesRequired(8) == 1, fmt.Sprintf("wrong 8: %d", bytesRequired(8)))
	bvmUtils.Assert(t, bytesRequired(257) == 2, fmt.Sprintf("wrong 257: %d", bytesRequired(257)))
}
