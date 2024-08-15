package stack

import "testing"

func TestSize(t *testing.T) {
	var tests = []struct {
		name string
		s    Stack
		want int
	}{
		{
			name: "empty stack test",
			s:    Stack{},
			want: 0,
		},
		{
			name: "stack test",
			s:    Stack{1, 2, 3},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.s.Size(); got != test.want {
				t.Errorf("got: %v\twant: %v", got, test.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	var tests = []struct {
		name string
		s    Stack
		want bool
	}{
		{
			name: "empty stack test",
			s:    Stack{},
			want: true,
		},
		{
			name: "stack test",
			s:    Stack{1, 2, 3},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.s.IsEmpty(); got != test.want {
				t.Errorf("got: %v\twant: %v", got, test.want)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	var tests = []struct {
		name string
		s    Stack
		want interface{}
	}{
		{
			name: "empty stack test",
			s:    Stack{},
			want: nil,
		},
		{
			name: "stack test",
			s:    Stack{1, 2, 3},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			if got, _ := test.s.Peek(); got != test.want {
				t.Errorf("got: %v\twant: %v", got, test.want)
			}
		})
	}
}

func TestPush(t *testing.T) {
	var tests = []struct {
		name string
		s    *Stack
		item interface{}
	}{
		{
			name: "empty stack test",
			s:    &Stack{},
			item: 999,
		},
		{
			name: "stack test",
			s:    &Stack{1, 2, 3},
			item: 888,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.s.Push(test.item)

			if got, _ := test.s.Peek(); got != test.item {
				t.Errorf("got: %v\twant: %v", got, test.item)
			}
		})
	}
}

func TestPop(t *testing.T) {
	var tests = []struct {
		name string
		s    Stack
		want interface{}
	}{
		{
			name: "empty stack test",
			s:    Stack{},
			want: nil,
		},
		{
			name: "stack test",
			s:    Stack{1, 2, 3},
			want: 3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got, _ := test.s.Pop(); got != test.want {
				t.Errorf("got: %v\twant: %v", got, test.want)
			}
		})
	}
}
