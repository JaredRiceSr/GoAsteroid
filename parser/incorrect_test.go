package parser

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)

func TestParseClassTerminating(t *testing.T) {
	_, errs := ParseString(`class`)
	bvmUtils.AssertNow(t, len(errs) == 3, errs.Format())
}

func TestParseClassNoIdentifier(t *testing.T) {
	_, errs := ParseString(`class {}`)
	bvmUtils.AssertNow(t, len(errs) > 1, errs.Format())
}

func TestParseClassNoBraces(t *testing.T) {
	_, errs := ParseString(`class OrderBook`)
	bvmUtils.AssertNow(t, len(errs) == 2, errs.Format())
}

func TestParseClassNoOpenBrace(t *testing.T) {
	_, errs := ParseString(`class OrderBook }`)
	bvmUtils.AssertNow(t, len(errs) > 1, errs.Format())
}

func TestParseClassNoCloseBrace(t *testing.T) {
	_, errs := ParseString(`class OrderBook {`)
	bvmUtils.AssertNow(t, len(errs) == 1, errs.Format())
}

func TestParseVarNoVarStatement(t *testing.T) {
	_, errs := ParseString(`name string`)
	bvmUtils.AssertNow(t, len(errs) > 0, errs.Format())
}
