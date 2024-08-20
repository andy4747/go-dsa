package stacks

type Stack[T comparable] struct {
	elements []T
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zeroVal T
		return zeroVal
	}
	element := s.elements[len(s.elements)-1]
	if len(s.elements) > 1 {
		s.elements = s.elements[:len(s.elements)-1]
	} else {
		s.elements = s.elements[0:]
	}
	return element
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
