package validator

import (
	"testing"

	"github.com/benchlab/bvmUtils"
)


func TestParseVotingExample(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/voting.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParseSimpleAuctionExample(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/simple_auction.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParseBlindAuctionExample(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/blind_auction.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParsePurchaseExample(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/purchase.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorBalanceChecker(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/examples/creator_balance_checker.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorBasicIterator(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/examples/basic_iterator.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParseCreatorGreeter(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/examples/greeter.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParseCrowdFunder(t *testing.T) {
	p, errs := ValidateFile(NewTestBVM(), nil, "../samples/tests/solc/crowd_funder.grd")
	bvmUtils.Assert(t, p != nil, "parser should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}

func TestParsePackageDigixDao(t *testing.T) {
	p, errs := ValidatePackage(NewTestBVM(), "../samples/tests/solc/examples/digixdao")
	bvmUtils.Assert(t, p != nil, "ast should not be nil")
	bvmUtils.Assert(t, errs == nil, errs.Format())
}
