package typing

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestAddModifier(t *testing.T) {
	b := Boolean()
	AddModifier(b, "static")
	bvmUtils.AssertNow(t, HasModifier(b, "static"), "doesn't have static")
}
