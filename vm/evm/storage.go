package evm

import (
	"encoding/binary"

	"github.com/benchlab/bvmCreate"
)

type machineStorageBlock struct {
	name   string
	size   uint
	offset uint
	slot   uint
}

type machineMemoryBlock struct {
	size   uint
	offset uint
}

const (
	wordSize = uint(256)
)

func (s machineStorageBlock) retrieve() (code bvmCreate.Bytecode) {
	if s.size > wordSize {
		// over at least 2 slots
		first := wordSize - s.offset
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, first))

		remaining := s.size - first
		slot := s.slot
		for remaining >= wordSize {
			// get whole slot
			slot += 1
			code.Concat(getByteSectionOfSlot(slot, 0, wordSize))
			remaining -= wordSize
		}
		if remaining > 0 {
			// get first remaining bits from next slot
			code.Concat(getByteSectionOfSlot(slot+1, 0, remaining))
		}
	} else if s.offset+s.size > wordSize {
		// over 2 slots
		// get last wordSize - s.offset bits from first
		first := wordSize - s.offset
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, first))
		// get first s.size - (wordSize - s.offset) bits from second
		code.Concat(getByteSectionOfSlot(s.slot+1, 0, s.size-first))
	} else {
		// all within 1 slot
		code.Concat(getByteSectionOfSlot(s.slot, s.offset, s.size))
	}
	return code
}

func getByteSectionOfSlot(slot, start, size uint) (code bvmCreate.Bytecode) {
	code.Add("PUSH", uintAsBytes(slot)...)
	code.Add("SLOAD")
	mask := ((1 << size) - 1) << start
	code.Add("PUSH", uintAsBytes(uint(mask))...)
	code.Add("AND")
	return code
}

func uintAsBytes(a uint) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(a))
	return bs
}

func (m machineMemoryBlock) retrieve() (code bvmCreate.Bytecode) {
	code.Concat(push(uintAsBytes(m.offset)))
	/*for size := m.size; size > wordSize; size -= wordSize {
		loc := m.offset + (m.size - size)
		code.Add("PUSH", uintAsBytes(loc)...)
		code.Add("MLOAD")
	}*/
	return code
}

func (s machineStorageBlock) store() (code bvmCreate.Bytecode) {
	code.Concat(push(EncodeName(s.name)))
	code.Add("SSTORE")
	return code
}

func (m machineMemoryBlock) store() (code bvmCreate.Bytecode) {
	free := []byte{0x40}
	code.Add("PUSH", free...)
	code.Add("PUSH", uintAsBytes(m.offset)...)
	code.Add("MSTORE")
	return code
}

func (evm *AsteroidEVM) freeMachineMemory(name string) {
	m, ok := evm.machineMemory[name]
	if ok {
		if evm.freedMachineMemory == nil {
			evm.freedMachineMemory = make([]*machineMemoryBlock, 0)
		}
		evm.freedMachineMemory = append(evm.freedMachineMemory, m)
		evm.machineMemory[name] = nil
	}
}

func (evm *AsteroidEVM) allocateMachineMemory(name string, size uint) {
	if evm.machineMemory == nil {
		evm.machineMemory = make(map[string]*machineMemoryBlock)
	}
	// try to use previously reclaimed machineMemory
	if evm.freedMachineMemory != nil {
		for i, m := range evm.freedMachineMemory {
			if m.size >= size {
				// we can use this block
				evm.machineMemory[name] = m
				// remove it from the freed list
				// TODO: check remove function
				evm.freedMachineMemory = append(evm.freedMachineMemory[i:], evm.freedMachineMemory[:i]...)
				return
			}
		}
	}

	block := machineMemoryBlock{
		size:   size,
		offset: evm.machineMemoryCursor,
	}
	evm.machineMemoryCursor += size
	evm.machineMemory[name] = &block
}

func (evm *AsteroidEVM) allocateStorageToMachine(name string, size uint) {
	// TODO: check whether there's a way to reduce machineStorage using some weird bin packing algo
	// with a modified heuristic to reduce the cost of extracting variables using bitshifts
	// maybe?
	if evm.machineStorage == nil {
		evm.machineStorage = make(map[string]*machineStorageBlock)
	}
	block := machineStorageBlock{
		size:   size,
		offset: evm.lastOffset,
		slot:   evm.lastSlot,
	}
	for size > wordSize {
		size -= wordSize
		evm.lastSlot += 1
	}
	evm.lastOffset += size
	evm.machineStorage[name] = &block
}

func (evm *AsteroidEVM) lookupMachineStorage(name string) *machineStorageBlock {
	if evm.machineStorage == nil {
		return nil
	}
	return evm.machineStorage[name]
}

func (evm *AsteroidEVM) lookupMachineMemory(name string) *machineMemoryBlock {
	if evm.machineMemory == nil {
		return nil
	}
	return evm.machineMemory[name]
}
