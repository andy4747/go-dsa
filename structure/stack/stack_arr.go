package stack

import "fmt"

type StackArr[T comparable] struct {
	elements []T
	size     int
}

func NewStackArr[T comparable]() *StackArr[T] {
	return &StackArr[T]{
		elements: make([]T, 0),
		size:     0,
	}
}

func (s *StackArr[T]) Push(data T) {
	s.elements = append(s.elements[:s.size], data)
	s.size++
}
func (s *StackArr[T]) Pop() (T, error) {
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

func (s *StackArr[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty stack")
	}
	return s.elements[s.Size()-1], nil
}

func (s *StackArr[T]) Size() int {
	return s.size
}

func (s *StackArr[T]) IsEmpty() bool {
	return s.Size() == 0
}

func (s *StackArr[T]) Print() {
	for i := 0; i < s.Size(); i++ {
		fmt.Printf("%+v\n", s.elements[i])
	}
}

func TestStack() {
	s := NewStackArr[int]()
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
