package linkedlist

import "fmt"

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewList[T comparable]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *List[T]) PushFront(data T) error {
	newNode := &Node[T]{Value: data}
	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return nil
	}
	newNode.Next = l.head
	l.head.Prev = newNode
	l.head = newNode
	l.size++
	return nil
}

func (l *List[T]) PushBack(data T) error {
	newNode := &Node[T]{Value: data}
	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return nil
	}
	l.tail.Next = newNode
	newNode.Prev = l.tail
	l.tail = newNode
	l.size++
	return nil
}

func (l *List[T]) PopBack() (T, error) {
	var result T
	if l.IsEmpty() {
		return result, fmt.Errorf("pop operation on empty list")
	}
	if l.head == l.tail {
		result = l.head.Value
		l.head = nil
		l.tail = nil
		l.size = 0
		return result, nil
	}
	result = l.tail.Value
	newTail := l.tail.Prev
	newTail.Next = nil
	l.tail = newTail
	l.size--
	return result, nil
}

func (l *List[T]) PopFront() (T, error) {
	var result T
	if l.IsEmpty() {
		return result, fmt.Errorf("pop operation on empty list")
	}
	if l.head == l.tail {
		result = l.head.Value
		l.head = nil
		l.tail = nil
		l.size = 0
		return result, nil
	}
	result = l.head.Value
	newHead := l.head.Next
	newHead.Prev = nil
	l.head = newHead
	l.size--
	return result, nil
}

func (l *List[T]) Insert(index int, data T) error {
	if l.IsEmpty() || index == 0 {
		return l.PushFront(data)
	}
	if index == l.Size() {
		return l.PushBack(data)
	}
}

func (l *List[T]) Remove(index int, data *T) error {
}

func (l *List[T]) Size() int {
}

func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[T]) Clear() {
}

func (l *List[T]) PrintList() {
}

func (l *List[T]) Find(data T) int {
}

func (l *List[T]) ForEach(callback func(data T)) {
}

func (l *List[T]) TraverseFront(index int, callback func(data T)) {
}

func (l *List[T]) TraverseBack(index int, callback func(data T)) {
}

func (l *List[T]) Reverse() {

}

func TestList() {
}
