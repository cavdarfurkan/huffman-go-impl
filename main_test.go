package main

import (
	"reflect"
	"testing"
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
