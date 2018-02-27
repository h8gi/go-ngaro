package ngaro

import "fmt"

type stack struct {
	data    []int
	pointer int
}

var (
	ErrStackOverflow  = fmt.Errorf("stack overflow")
	ErrStackUnderflow = fmt.Errorf("stack underflow")
)

func newStack(size int) *stack {
	return &stack{
		data:    make([]int, size),
		pointer: 0,
	}
}

func (s *stack) top() int {
	return s.data[s.pointer]
}

func (s *stack) push(elem int) error {
	if s.pointer >= len(s.data) {
		return ErrStackOverflow
	}
	s.data[s.pointer] = elem
	s.pointer++
	return nil
}

func (s *stack) pop() (int, error) {
	if s.pointer <= 0 {
		return 0, ErrStackUnderflow
	}
	s.pointer--
	elem := s.top()
	return elem, nil
}
