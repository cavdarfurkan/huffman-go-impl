package main

import (
	"reflect"
	"testing"

	"github.com/cavdarfurkan/huffman-go-impl/minheap"
	"github.com/vishalkuo/bimap"
)

func TestFreqTable(t *testing.T) {
	var freqTableTests = []struct {
		input    string
		expected FrequencyTable
	}{
		{"deneme", FrequencyTable{'e': 3, 'd': 1, 'n': 1, 'm': 1}},
		{"this is an example of a huffman tree", FrequencyTable{' ': 7, 'a': 4, 'e': 4, 'f': 3, 'h': 2, 'i': 2, 'm': 2, 'n': 2, 's': 2, 't': 2, 'l': 1, 'o': 1, 'p': 1, 'r': 1, 'u': 1, 'x': 1}},
	}

	for _, test := range freqTableTests {
		if got := GenerateFrequencyTable(test.input); !reflect.DeepEqual(got, test.expected) {
			t.Errorf("For input '%s'\ngot: %v\nwanted: %v", test.input, got, test.expected)
		}
	}
}

func TestGenerateNodes(t *testing.T) {
	var generateNodesTests = []struct {
		freqTable FrequencyTable
		want      minheap.MinHeap
	}{
		{
			GenerateFrequencyTable("deneme"),
			minheap.MinHeap{
				&minheap.Node{Char: 'e', Frequency: 3},
				&minheap.Node{Char: 'd', Frequency: 1},
				&minheap.Node{Char: 'n', Frequency: 1},
				&minheap.Node{Char: 'm', Frequency: 1},
			},
		},
		{
			GenerateFrequencyTable("this is an example of a huffman tree"),
			minheap.MinHeap{
				&minheap.Node{Char: ' ', Frequency: 7},
				&minheap.Node{Char: 'a', Frequency: 4},
				&minheap.Node{Char: 'e', Frequency: 4},
				&minheap.Node{Char: 'f', Frequency: 3},
				&minheap.Node{Char: 'h', Frequency: 2},
				&minheap.Node{Char: 'i', Frequency: 2},
				&minheap.Node{Char: 'm', Frequency: 2},
				&minheap.Node{Char: 'n', Frequency: 2},
				&minheap.Node{Char: 's', Frequency: 2},
				&minheap.Node{Char: 't', Frequency: 2},
				&minheap.Node{Char: 'l', Frequency: 1},
				&minheap.Node{Char: 'o', Frequency: 1},
				&minheap.Node{Char: 'p', Frequency: 1},
				&minheap.Node{Char: 'r', Frequency: 1},
				&minheap.Node{Char: 'u', Frequency: 1},
				&minheap.Node{Char: 'x', Frequency: 1},
			},
		},
	}

	for _, test := range generateNodesTests {
		nodes := test.freqTable.generateNodes()
		nodesMap := make(map[rune]int, nodes.Len())

		// Fill the nodesMap
		for _, val := range *nodes {
			nodesMap[val.Char] = val.Frequency
		}

		if nodes.Len() == test.want.Len() {
			for _, val := range test.want {
				gotFreq, ok := nodesMap[val.Char]
				if !ok {
					t.Errorf("got: %v\t want: %v", nodes, test.want)
				}
				if gotFreq != val.Frequency {
					t.Errorf("Incorrect frequency error\nFor: '%c'\ngot: %d\twant: %d", val.Char, gotFreq, val.Frequency)
				}
			}
		} else {
			t.Errorf("Length error\ngot: '%d'\twanted: '%d'", nodes.Len(), test.want.Len())
		}
	}
}

func TestBuildHuffmanTree(t *testing.T) {
	var tests = []struct {
		name           string
		frequencyTable FrequencyTable
		want           *minheap.MinHeap
	}{
		{
			name:           "test 1",
			frequencyTable: GenerateFrequencyTable("aaabbc"),
			want: &minheap.MinHeap{
				&minheap.Node{
					Frequency: 6,
					Left: &minheap.Node{
						Char:      'a',
						Frequency: 3,
					},
					Right: &minheap.Node{
						Frequency: 3,
						Left: &minheap.Node{
							Char:      'b',
							Frequency: 2,
						},
						Right: &minheap.Node{
							Char:      'c',
							Frequency: 1,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.frequencyTable.BuildHuffmanTree()

			// Unbelievable
			if (got.Len() != test.want.Len()) ||
				(got.PeekFirst().Frequency != test.want.PeekFirst().Frequency) ||
				(got.PeekFirst().Left.Char != test.want.PeekFirst().Left.Char ||
					got.PeekFirst().Left.Frequency != test.want.PeekFirst().Left.Frequency) ||
				(got.PeekFirst().Right.Frequency != test.want.PeekFirst().Right.Frequency) ||
				(got.PeekFirst().Right.Left.Char != test.want.PeekFirst().Right.Left.Char ||
					got.PeekFirst().Right.Left.Frequency != test.want.PeekFirst().Right.Left.Frequency) ||
				(got.PeekFirst().Right.Right.Char != test.want.PeekFirst().Right.Right.Char ||
					got.PeekFirst().Right.Right.Frequency != test.want.PeekFirst().Right.Right.Frequency) {
				t.Errorf("got: %v,%v,%v,%v,%v\twant:%v,%v,%v,%v,%v",
					got.PeekFirst(), got.PeekFirst().Left, got.PeekFirst().Right, got.PeekFirst().Right.Left, got.PeekFirst().Right.Right,
					test.want.PeekFirst(), test.want.PeekFirst().Left, test.want.PeekFirst().Right, test.want.PeekFirst().Right.Left, test.want.PeekFirst().Right.Right)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	var tests = []struct {
		name, input, want string
	}{
		{
			name:  "encode test 1",
			input: "aaabbc",
			want:  "000101011",
		},
		{
			name:  "encode empty string",
			input: "",
			want:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := Encode(test.input)
			if got.EncodedValue != test.want {
				t.Errorf("got: %v\twant: %s", got.EncodedValue, test.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	var tests = []struct {
		name, want string
		input      EncodedString
	}{
		{
			name: "decode test 1",
			input: EncodedString{
				EncodedValue: "000101011",
				Codes: func() *bimap.BiMap[rune, string] {
					val, _ := generateCodes(*GenerateFrequencyTable("aaabbc").BuildHuffmanTree())
					return val
				}(),
			},
			want: "aaabbc",
		},
		{
			name: "decode empty string",
			input: EncodedString{
				EncodedValue: "",
				Codes: func() *bimap.BiMap[rune, string] {
					val, _ := generateCodes(*GenerateFrequencyTable("").BuildHuffmanTree())
					return val
				}(),
			},
			want: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Decode(test.input)
			if got != test.want {
				t.Errorf("got: %s\twant: %s", got, test.want)
			}
		})
	}
}
