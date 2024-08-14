package minheap

import (
	"testing"
)

var testHeap = MinHeap{
	&Node{Char: 'a', Frequency: 1},
	&Node{Char: 'b', Frequency: 2},
	&Node{Char: 'c', Frequency: 4},
	&Node{Char: 'd', Frequency: 6},
}

func TestLen(t *testing.T) {
	var lenTests = []struct {
		heap MinHeap
		want int
	}{
		{MinHeap{}, 0},
		{testHeap, 4},
	}

	for _, test := range lenTests {
		if got := test.heap.Len(); got != test.want {
			t.Errorf("got: %d\t\twant: %d", got, test.want)
		}
	}
}

func TestLess(t *testing.T) {
	var lessTests = []struct {
		heap       MinHeap
		arg1, arg2 int
		want       bool
	}{
		{testHeap, 1, 3, true},
		{testHeap, 3, 2, false},
	}

	for _, test := range lessTests {
		if got := test.heap.Less(test.arg1, test.arg2); got != test.want {
			t.Errorf("got: %v\twant: %v", got, test.want)
		}
	}
}

func TestSwap(t *testing.T) {
	var swapTests = []struct {
		heap       MinHeap
		arg1, arg2 int
	}{
		{testHeap, 1, 3},
		{testHeap, 3, 2},
	}

	for _, test := range swapTests {
		want := test.heap[test.arg2]
		test.heap.Swap(test.arg1, test.arg2)

		if test.heap[test.arg1].Char != want.Char {
			t.Errorf("got: %c\twant: %c", test.heap[test.arg1].Char, want.Char)
		}
	}
}

func TestPush(t *testing.T) {
	var pushTests = []struct {
		heap MinHeap
		arg1 *Node
	}{
		{testHeap, &Node{Char: 'e'}},
		{testHeap, &Node{Char: 'f'}},
	}

	for _, test := range pushTests {
		oldLen := test.heap.Len()
		want := oldLen + 1

		test.heap.Push(test.arg1)

		if newLen := test.heap.Len(); newLen != want {
			t.Errorf("Len error\ngot: '%d'\twant: '%d'", newLen, want)
		}
		if node := test.heap[len(test.heap)-1]; node.Char != test.arg1.Char && node.index != test.arg1.index {
			t.Errorf("got: '%c' at '%d'\twant: '%c' at '%d'", node.Char, node.index, test.arg1.Char, want)
		}
	}
}

func TestPop(t *testing.T) {
	var popTests = []struct {
		heap MinHeap
		want *Node
	}{
		{testHeap, testHeap[len(testHeap)-1]},
	}

	for _, test := range popTests {
		node := test.heap.Pop()
		if node != test.want {
			t.Errorf("got: '%c'\twanted: '%c'", node.(*Node).Char, test.want.Char)
		}
	}
}

func TestPeekFirst(t *testing.T) {
	var peekFirstTests = []struct {
		heap MinHeap
		want *Node
	}{
		{testHeap, testHeap[0]},
	}

	for _, test := range peekFirstTests {
		if got := test.heap.PeekFirst(); got != test.want {
			t.Errorf("got:%v\twant:%v", got, test.want)
		}
	}
}
