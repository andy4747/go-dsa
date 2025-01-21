package linkedlist

import "fmt"

type SList[T comparable] struct {
	head *SNode[T]
	size int
}

func NewSList[T comparable]() *SList[T] {
	return &SList[T]{
		head: nil,
		size: 0,
	}
}

func (l *SList[T]) PushFront(data T) {
	newNode := &SNode[T]{
		Value: data,
		Next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}
	newNode.Next = l.head
	l.head = newNode
	l.size++
}

func (l *SList[T]) PushBack(data T) {
	newNode := &SNode[T]{
		Value: data,
		Next:  nil,
	}
	if l.head == nil {
		l.head = newNode
		l.size++
		return
	}
	cur := l.head
	for cur != nil {
		if cur.Next == nil {
			cur.Next = newNode
			break
		}
		cur = cur.Next
	}
	l.size++
}

func (l *SList[T]) PopFront() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}
	next := l.head.Next
	l.head = next
	l.size--
	return nil
}

func (l *SList[T]) PopBack() error {
	if l.head == nil {
		return fmt.Errorf("list is empty")
	}
	cur := l.head
	for cur != nil {
		if cur.Next.Next == nil {
			cur.Next = nil
		}
		cur = cur.Next
	}
	l.size--
	return nil
}

func (l *SList[T]) Insert(index int, data T) error {
	if index >= l.size || index < 0 {
		return fmt.Errorf("index out of bounds")
	}
	if l.head == nil && index != 0 {
		return fmt.Errorf("cannot insert at index %d", index)
	}
	count := 0
	cur := l.head
	for cur != nil {
		if count == index-1 {
			next := cur.Next
			newNode := &SNode[T]{Value: data, Next: next}
			cur.Next = newNode
			l.size++
			return nil
		}
		cur = cur.Next
		count++
	}
	return fmt.Errorf("could not insert at index %d", index)
}

func (l *SList[T]) Remove(index int, data T) error {
	if l.head == nil {
		return fmt.Errorf("cannot remove at index %d", index)
	}
	if index >= l.size || index < 0 {
		return fmt.Errorf("index out of bounds")
	}
	count := 0
	cur := l.head
	for cur != nil {
		if count == index-1 {
			newNext := cur.Next.Next
			cur.Next = newNext
			l.size--
			return nil
		}
		cur = cur.Next
		count++
	}
	return fmt.Errorf("could not remove at index %d", index)
}

func (l *SList[T]) Size() int {
	return l.size
}

func (l *SList[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *SList[T]) Clear() {
	l.head = nil
	l.size = 0
}

func (l *SList[T]) ForEach(callback func(data T)) {
	cur := l.head
	for cur != nil {
		callback(cur.Value)
		cur = cur.Next
	}
}

func (l *SList[T]) PrintList() {
	l.ForEach(func(data T) {
		fmt.Printf("%v ", data)
	})
	fmt.Println()
}

func (l *SList[T]) Find(data T) int {
	cur := l.head
	index := 0
	for cur != nil {
		if cur.Value == data {
			return index
		}
		index++
		cur = cur.Next
	}
	return -1
}

func (l *SList[T]) Reverse() {}

func TestSList() {
	l := NewSList[int]()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)
	l.PushFront(5)
	l.PushFront(1)
	l.PrintList()
}
