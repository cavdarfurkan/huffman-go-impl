package bitarray

import (
	"fmt"
)

type BitArray struct {
	bits []byte
	size int
}

func NewBitArray(size int) *BitArray {
	return &BitArray{
		bits: make([]byte, (size+7)/8),
		size: size,
	}
}

func (bt *BitArray) SetBit(index int) error {
	if index < 0 || index >= bt.size {
		return fmt.Errorf("index (%d) out of range of (%d)", index, bt.size-1)
	}

	bt.bits[index/8] |= (1 << index)
	return nil
}

func (bt *BitArray) ClearBit(index int) error {
	if index < 0 || index >= bt.size {
		return fmt.Errorf("index (%d) out of range of (%d)", index, bt.size-1)
	}

	var mask byte = ^(1 << index)
	bt.bits[index/8] &= mask
	return nil
}
