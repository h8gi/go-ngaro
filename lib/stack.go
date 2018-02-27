package ngaro

import "fmt"

type stack struct {
	data    []Cell
	size    int
	pointer int
}

var (
	ErrStackOverflow  = fmt.Errorf("stack overflow")
	ErrStackUnderflow = fmt.Errorf("stack underflow")
)

func newStack(size int) *stack {
	return &stack{
		data:    make([]Cell, size),
		size:    size,
		pointer: 0,
	}
}

func (s *stack) top() Cell {
	return s.data[s.pointer]
}

func (s *stack) push(elem Cell) error {
	if s.pointer >= s.size {
		return ErrStackOverflow
	}
	s.data[s.pointer] = elem
	s.pointer++
	return nil
}

func (s *stack) pop() (Cell, error) {
	if s.pointer <= 0 {
		return 0, ErrStackUnderflow
	}
	s.pointer--
	elem := s.top()
	return elem, nil
}
