package stack

import "errors"

type Stack []interface{}

func (s Stack) Size() int {
	return len(s)
}

func (s Stack) IsEmpty() bool {
	return s.Size() <= 0
}

func (s Stack) Peek() (interface{}, error) {
	n := s.Size()
	if n <= 0 {
		return nil, errors.New("there is no element in the stack")
	}

	return s[n-1], nil
}

func (s *Stack) Push(n interface{}) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (interface{}, error) {
	n := s.Size()
	if n <= 0 {
		return nil, errors.New("there is no element in the stack")
	}

	old := *s
	item := old[n-1]
	*s = old[:n-1]

	return item, nil
}
