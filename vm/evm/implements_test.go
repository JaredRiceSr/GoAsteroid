package evm

import (
	"testing"

	"github.com/benchlab/asteroid/validator"

	"github.com/benchlab/bvmUtils"
)

func TestImplements(t *testing.T) {
	var v validator.BVM
	var e AsteroidEVM
	v = e
	bvmUtils.Assert(t, v.BooleanName() == e.BooleanName(), "this doesn't matter")
}
