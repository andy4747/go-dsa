package stack

import (
	"errors"
	"fmt"
)

type Stacker[T comparable] interface {
	Push(data T)
	Pop() (T, error)
	Peek() (T, error)
	Size() int
	IsEmpty() bool
	Clear()
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type Stack[T comparable] struct {
	head *Node[T]
	size int
}

func NewStack[T comparable]() Stacker[T] {
	return &Stack[T]{
		head: nil,
		size: 0,
	}
}

func (s *Stack[T]) Push(data T) {
	newNode := &Node[T]{
		Value: data,
		Next:  s.head,
	}
	s.head = newNode
	s.size++
}

func (s *Stack[T]) Pop() (T, error) {
	var result T
	if s.IsEmpty() {
		return result, errors.New("error: pop operation on empty stack")
	}
	result = s.head.Value
	s.head = s.head.Next
	s.size--
	return result, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var result T
	if s.IsEmpty() {
		return result, errors.New("error: peek operation on empty stack")
	}
	result = s.head.Value
	return result, nil
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack[T]) Clear() {
	s.head = nil
	s.size = 0
}

func TestLinkedStack() {
	stack := NewStack[int]()

	// Push some elements
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	// Peek at the top element
	top, _ := stack.Peek()
	fmt.Println("Top element:", top)

	// Pop all elements
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		fmt.Println("Popped:", val)
	}
}
