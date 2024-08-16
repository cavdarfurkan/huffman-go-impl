package bitarray

import (
	"testing"
)

func createBitArray(size int) *BitArray {
	return NewBitArray(size)
}

func createOnesBitArray(size int) *BitArray {
	ba := createBitArray(size)
	for i := range ba.bits {
		ba.bits[i] = byte(255)
	}
	return ba
}

func TestSetBit(t *testing.T) {
	var tests = []struct {
		name     string
		bits     *BitArray
		setIndex int
		want     *BitArray
		isError  bool
	}{
		{
			name:     "set first bit test",
			bits:     createBitArray(8),
			setIndex: 0,
			want: &BitArray{
				bits: []byte{byte(1)},
				size: 8,
			},
			isError: false,
		},
		{
			name:     "set second bit test",
			bits:     createBitArray(8),
			setIndex: 1,
			want: &BitArray{
				bits: []byte{byte(2)},
				size: 8,
			},
			isError: false,
		},
		{
			name:     "set last bit test",
			bits:     createBitArray(8),
			setIndex: 7,
			want: &BitArray{
				bits: []byte{byte(128)},
				size: 8,
			},
			isError: false,
		},
		{
			name:     "set bit error test",
			bits:     createBitArray(8),
			setIndex: 8,
			want: &BitArray{
				bits: []byte{0},
				size: 8,
			},
			isError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.bits.SetBit(test.setIndex)
			if err != nil && !test.isError {
				t.Errorf("err: %s", err)
			}

			for i := range test.bits.bits {
				if test.bits.bits[i] != test.want.bits[i] {
					t.Errorf("got: %v\t want: %v", test.bits.bits, test.want.bits)
				}
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	var tests = []struct {
		name       string
		bits       *BitArray
		clearIndex int
		want       *BitArray
		isError    bool
	}{
		{
			name:       "clear first bit test",
			bits:       createOnesBitArray(8),
			clearIndex: 0,
			want: &BitArray{
				bits: []byte{byte(254)},
				size: 8,
			},
			isError: false,
		},
		{
			name:       "clear second bit test",
			bits:       createOnesBitArray(8),
			clearIndex: 1,
			want: &BitArray{
				bits: []byte{byte(253)},
				size: 8,
			},
			isError: false,
		},
		{
			name:       "clear last bit test",
			bits:       createOnesBitArray(8),
			clearIndex: 7,
			want: &BitArray{
				bits: []byte{byte(127)},
				size: 8,
			},
			isError: false,
		},
		{
			name:       "clear bit error test",
			bits:       createOnesBitArray(8),
			clearIndex: 8,
			want: &BitArray{
				bits: []byte{255},
				size: 8,
			},
			isError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.bits.ClearBit(test.clearIndex)
			if err != nil && !test.isError {
				t.Errorf("err: %s", err)
			}

			for i := range test.bits.bits {
				if test.bits.bits[i] != test.want.bits[i] {
					t.Errorf("got: %v\t want: %v", test.bits.bits, test.want.bits)
				}
			}
		})
	}
}
