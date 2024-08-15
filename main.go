package main

import (
	"container/heap"

	"github.com/cavdarfurkan/huffman-go-impl/minheap"
)

type FrequencyTable map[rune]int

func GenerateFrequencyTable(str string) FrequencyTable {
	freqTable := make(FrequencyTable, 0)
	for _, val := range str {
		freqTable[val] += 1
	}
	return freqTable
}

func (freqTable FrequencyTable) generateNodes() *minheap.MinHeap {
	mh := make(minheap.MinHeap, 0, len(freqTable))
	heap.Init(&mh)

	for key, val := range freqTable {
		heap.Push(&mh, &minheap.Node{Char: key, Frequency: val})
	}

	return &mh
}

func (freqTable FrequencyTable) BuildHuffmanTree() *minheap.MinHeap {
	mh := freqTable.generateNodes()
	heap.Init(mh)

	for mh.Len() > 1 {
		node1 := heap.Pop(mh).(*minheap.Node)
		node2 := heap.Pop(mh).(*minheap.Node)

		var left, right *minheap.Node
		if node1.Frequency < node2.Frequency {
			left = node2
			right = node1
		} else {
			left = node1
			right = node2
		}

		newNode := &minheap.Node{
			Frequency: node1.Frequency + node2.Frequency,
			Left:      left,
			Right:     right,
		}

		heap.Push(mh, newNode)
	}

	return mh
}

func Encode(input string) string {
	panic("Implement encode function")
}

func Decode(input string) string {
	panic("Implement decode function")
}

// https://cgi.luddy.indiana.edu/~yye/c343-2019/huffman.php
