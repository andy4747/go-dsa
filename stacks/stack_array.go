package stacks

type Stack[T comparable] struct {
	elements []T
	len      int
}

func (s *Stack[T]) New() {
	s.elements = make([]T, 0)
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
	s.len += 1
}

func (s *Stack[T]) Pop() *T {
	if s.len == 0 {
		return nil
	}
	element := s.elements[s.len-1]
	if s.len > 1 {
		s.elements = s.elements[:s.len-1]
	} else {
		s.elements = s.elements[0:]
	}
	s.len = len(s.elements)
	return &element
}
