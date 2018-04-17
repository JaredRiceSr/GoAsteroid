package lexer

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
	"github.com/benchlab/asteroid/token"
)

func checkTokens(t *testing.T, received []token.Token, expected []token.Type) {
	bvmUtils.AssertNow(t, len(received) == len(expected), fmt.Sprintf("wrong num of tokens: a %d / e %d", len(received), len(expected)))
	for index, r := range received {
		bvmUtils.Assert(t, r.Type == expected[index],
			fmt.Sprintf("wrong type %d: %s, expected %d", index, r.Proto.Name, expected[index]))
	}
}
