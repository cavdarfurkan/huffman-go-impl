package minheap

import (
	"fmt"
)

type Node struct {
	Char        rune
	Frequency   int // priority
	Left, Right *Node
	index       int
}

type MinHeap []*Node

func (mh MinHeap) Len() int {
	return len(mh)
}

func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Frequency < mh[j].Frequency
}

func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
	mh[i].index = i
	mh[j].index = j
}

func (mh *MinHeap) Push(x any) {
	item := x.(*Node)
	item.index = len(*mh)
	*mh = append(*mh, item)
}

func (mh *MinHeap) Pop() any {
	old := *mh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*mh = old[0 : n-1]
	return item
}

func (mh MinHeap) PeekFirst() *Node {
	return mh[0]
}

func (node Node) String() string {
	return fmt.Sprintf("{%c, %d}", node.Char, node.Frequency)
}
