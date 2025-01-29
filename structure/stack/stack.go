package stack

import "fmt"

type Stack[T comparable] struct {
	elements []T
	size     int
}

func NewStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
		size:     0,
	}
}

func (s *Stack[T]) Push(data T) {
	s.elements = append(s.elements[:s.size], data)
	s.size++
}
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty stack")
	}
	result := s.elements[s.Size()-1]
	if s.Size() > 1 {
		s.elements = s.elements[:s.Size()-1]
	} else {
		s.elements = s.elements[0:]
	}
	s.size--
	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty stack")
	}
	return s.elements[s.Size()-1], nil
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *Stack[T]) Print() {
	for i := 0; i < s.Size(); i++ {
		fmt.Printf("%+v\n", s.elements[i])
	}
}

func TestStack() {
	s := NewStack[int]()
	fmt.Printf("is stack empty? : %v\n", s.IsEmpty())
	s.Push(0)
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Push(40)
	popValue, _ := s.Pop()
	fmt.Printf("pop value: %d\n", popValue)
	s.Print()
	fmt.Printf("size of stack: %d\n", s.Size())
	peekValue, _ := s.Peek()
	fmt.Printf("peek value: %d\n", peekValue)
	fmt.Printf("is stack empty? : %v\n", s.IsEmpty())
}
