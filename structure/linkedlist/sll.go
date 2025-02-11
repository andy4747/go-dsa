package linkedlist

import (
	"fmt"
)

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

func (l *SList[T]) Remove(index int, data *T) error {
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
			*data = cur.Next.Value
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

func (l *SList[T]) Reverse() {
	cur := l.head
	var pv *SNode[T] = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pv
		pv = cur
		cur = next
	}
	l.head = pv
}

func (l *SList[T]) FloydCycleDetection() bool {
	if l.head == nil || l.head.Next == nil {
		return false
	}
	slow := l.head
	fast := l.head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func (l *SList[T]) FloydCycleDetectionNode() *SNode[T] {
	fast := l.head
	slow := l.head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = l.head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func (l *SList[T]) MiddleNode() *SNode[T] {
	if l.head == nil {
		return nil
	}
	slow := l.head
	fast := l.head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func (l *SList[T]) NthNodeFromBack(n int) *SNode[T] {
	if l.head == nil {
		return nil
	}
	first := l.head
	last := l.head
	count := 0
	for count != n && last != nil {
		last = last.Next
		count++
	}
	for first != nil && last != nil {
		first = first.Next
		last = last.Next
	}
	return first
}

func (list *SList[T]) CreateLoop(index int) {
	count := 0
	cur := list.head
	var idxNode *SNode[T] = nil
	for cur != nil {
		if count == index {
			idxNode = cur
		}
		if cur.Next == nil {
			cur.Next = idxNode
			break
		}
		count++
		cur = cur.Next
	}
	fmt.Printf("The loop is created from node at index %d to tail node\n", index)
}

func InsertIntItems(l *SList[int]) {
	l.Clear()
	items := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140}
	for _, v := range items {
		l.PushBack(v)
	}
}

func TestSList() {
	l := NewSList[int]()
	// l.PushBack(10)
	// l.PushBack(20)
	// l.PushBack(30)
	// l.PushBack(40)
	// l.PushFront(5)
	// l.PushFront(1)
	// var remVal int
	// l.Remove(1, &remVal)
	// fmt.Printf("Removed Value is: %d\n", remVal)
	// l.PrintList()
	// l.Reverse()
	// l.PrintList()
	InsertIntItems(l)
	l.PrintList()
	// l.Reverse()
	// l.PrintList()
	// l.CreateLoop(8)
	// fmt.Printf("Cycle in loop: %v\n", l.FloydCycleDetection())
	// loopNode := l.FloydCycleDetectionNode()
	// fmt.Printf("%+v\n", loopNode)
	fmt.Printf("%+v\n", l.MiddleNode())
	fmt.Printf("%+v\n", l.NthNodeFromBack(5))
}
