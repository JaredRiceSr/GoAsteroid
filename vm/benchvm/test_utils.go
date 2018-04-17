package benchvm

import (
	"fmt"
	"testing"

	"github.com/benchlab/bvmUtils"
	"github.com/end-r/vmgen"
)

func checkMnemonics(t *testing.T, is map[byte]*vmgen.Instruction, es []string) {
	fmt.Println("a")
	bvmUtils.AssertNow(t, is != nil, "instructions shouldn't be nil")
	fmt.Println("b")
	bvmUtils.AssertNow(t, len(is) == len(es), "wrong num of instructions")
	fmt.Println("c")
	for index, i := range is {
		fmt.Println("d")
		bvmUtils.AssertNow(t, i != nil, "instruction shouldn't be nil")
		fmt.Println("e")
		bvmUtils.Assert(t, i.Mnemonic == es[index],
			fmt.Sprintf("wrong mnemonic %d: %s, expected %s", index, i.Mnemonic, es[index]))
		fmt.Println("f")
	}
}

func checkStack(t *testing.T, stack *vmgen.Stack, es [][]byte) {
	bvmUtils.AssertNow(t, stack != nil, "stack shouldn't be nil")
	bvmUtils.AssertNow(t, stack.size() == len(es), "wrong stack size")
	for stack.size() > 0 {
		item := stack.Pop()
		expected := es[stack.size()]
		bvmUtils.Assert(t, len(item) == len(expected), fmt.Sprintf("wrong stack item %d length", stack.size()))
		if len(item) != len(expected) {
			continue
		}
		for i, b := range item {
			bvmUtils.Assert(t, b == expected[i],
				fmt.Sprintf("wrong stack item %d: expected %b, got %b", stack.size(), expected[i], b))
		}
	}
}
