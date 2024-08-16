package main

import (
	"container/heap"
	"errors"
	"fmt"
	"strings"

	"github.com/cavdarfurkan/huffman-go-impl/minheap"
	"github.com/cavdarfurkan/huffman-go-impl/stack"
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

type nodeCodes struct {
	node  *minheap.Node
	value interface{}
	code  string
}

// CodeMap is an alias for map[rune]string
type CodeMap map[rune]string

func generateCodes(mh minheap.MinHeap) (CodeMap, error) {
	if mh.Len() != 1 {
		return nil, errors.New("heap's size must be 1 to be able to generate codes")
	}

	stack := make(stack.Stack, 0)
	result := make(CodeMap, 0)

	root := mh.PeekFirst()
	stack.Push(nodeCodes{
		node:  root,
		value: root.Char,
		code:  "",
	})

	for !stack.IsEmpty() {
		nodeCode, err := stack.Pop()
		if err != nil {
			fmt.Printf("error: %s", err)
		}

		if val, ok := nodeCode.(nodeCodes); ok {
			if val.node.Char != 0 {
				result[val.node.Char] = val.code
			}

			if val.node.Right != nil {
				stack.Push(nodeCodes{
					node:  val.node.Right,
					value: val.node.Right.Char,
					code:  val.code + "1",
				})
			}
			if val.node.Left != nil {
				stack.Push(nodeCodes{
					node:  val.node.Left,
					value: val.node.Left.Char,
					code:  val.code + "0",
				})
			}
		}
	}

	return result, nil
}

type EncodedString struct {
	EncodedValue string
	Codes        CodeMap
}

func Encode(input string) (EncodedString, error) {
	huff := GenerateFrequencyTable(input).BuildHuffmanTree()
	heap.Init(huff)

	codes, err := generateCodes(*huff)
	if err != nil {
		return EncodedString{}, err
	}

	var encodedString strings.Builder

	for _, v := range input {
		code := codes[v]
		encodedString.WriteString(code)
	}

	return EncodedString{
		EncodedValue: encodedString.String(),
		Codes:        codes,
	}, nil
}

func Decode(input EncodedString) string {
	panic("Implement decode function")
}

func main() {
	val, err := Encode("")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("val.EncodedValue: %v\n", val.EncodedValue)
}

// https://cgi.luddy.indiana.edu/~yye/c343-2019/huffman.php
