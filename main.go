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

func (freqTable FrequencyTable) GenerateNodes() minheap.MinHeap {
	mh := make(minheap.MinHeap, 0, len(freqTable))
	heap.Init(&mh)

	for key, val := range freqTable {
		heap.Push(&mh, &minheap.Node{Char: key, Frequency: val})
	}

	return mh
}

// func main() {
// 	mh := make(minheap.MinHeap, 0)
// 	heap.Init(&mh)

// 	heap.Push(&mh, &minheap.Node{Char: 'a', Frequency: 5})
// 	heap.Push(&mh, &minheap.Node{Char: 'b', Frequency: 1})
// 	heap.Push(&mh, &minheap.Node{Char: 'c', Frequency: 3})

// 	for mh.Len() > 0 {
// 		n := heap.Pop(&mh).(*minheap.Node)
// 		fmt.Print(n)
// 	}
// }

// TODO: Generate huffman tree
// TODO: Encode
// TODO: Decode
