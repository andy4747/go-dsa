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
	newNode := &Node[T]{
		Value: data,
		Prev:  nil,
		Next:  l.head,
	}
	if l.IsEmpty() {
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return nil
	}
	l.head.Prev = newNode
	l.head = newNode
	l.size++
	return nil
}

func (l *List[T]) PushBack(data T) error {
	newNode := &Node[T]{
		Value: data,
		Prev:  l.tail,
		Next:  nil,
	}
	if l.IsEmpty() {
		newNode.Prev = nil
		l.head = newNode
		l.tail = newNode
		l.size = 1
		return nil
	}
	cur := l.head
	for cur != nil {
		if cur.Next == nil {
			cur.Next = newNode
			l.tail = newNode
			l.size++
			return nil
		}
		cur = cur.Next
	}
	return fmt.Errorf("error in \"PushBack\" operation")
}

func (l *List[T]) PopBack() (T, error) {
	var result T
	if l.head == l.tail {
		result = l.head.Value
		l.head = nil
		l.tail = nil
		l.size--
		return result, nil
	}
	cur := l.head
	for cur != nil {
		if cur.Next.Next == nil {
			result = cur.Next.Value
			cur.Next = nil
			l.tail = cur
			l.size--
			return result, nil
		}
	}
	return result, fmt.Errorf("error in \"PopBack\" operation")
}

func (l *List[T]) PopFront() (T, error) {
	var result T
	if l.head == l.tail {
		result = l.head.Value
		l.head = nil
		l.tail = nil
		l.size--
		return result, nil
	}
	result = l.head.Value
	newHead := l.head.Next
	l.head = newHead
	l.size--
	return result, nil
}

func (l *List[T]) Insert(index int, data T) error {
	if index < 0 || index >= l.Size() {
		return fmt.Errorf("index out of bounds")
	}
	if l.head == nil {
		return fmt.Errorf("insert error, index out of bounds")
	}
	if index == 0 {
		return l.PushFront(data)
	}
	if index == l.Size()-1 {
		return l.PushBack(data)
	}
	count := 0
	cur := l.head
	for cur != nil {
		if count == index-1 {
			newNode := &Node[T]{Value: data, Next: cur.Next}
			cur.Next = newNode
			l.size++
			return nil
		}
		count++
		cur = cur.Next
	}
	return nil
}

func (l *List[T]) Remove(index int, data *T) error {
	if index < 0 || index >= l.Size() {
		return fmt.Errorf("index out of bounds")
	}
	if l.head == nil {
		return fmt.Errorf("remove error, index out of bounds")
	}
	if index == 0 {
		_, err := l.PopFront()
		return err
	}
	if index == l.Size()-1 {
		_, err := l.PopBack()
		return err
	}
	count := 0
	cur := l.head
	for cur != nil {
		if count == index-1 {
			newNext := cur.Next.Next
			newNext.Prev = cur
			*data = cur.Next.Value
			cur.Next = newNext
			l.size--
			return nil
		}
		count++
		cur = cur.Next
	}
	return fmt.Errorf("remove error")
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l *List[T]) PrintList() {
	l.ForEach(func(data T) {
		fmt.Printf("%v ", data)
	})
	fmt.Println()
}

func (l *List[T]) Find(data T) int {
	if l.IsEmpty() {
		return -1
	}
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

func (l *List[T]) ForEach(callback func(data T)) {
	cur := l.head
	for cur != nil {
		callback(cur.Value)
		cur = cur.Next
	}
}

func (l *List[T]) TraverseFront(index int, callback func(data T)) {
	count := 0
	cur := l.head
	for cur != nil {
		if count >= index {
			callback(cur.Value)
		}
		count++
		cur = cur.Next
	}
}

func (l *List[T]) TraverseBack(index int, callback func(data T)) {
	count := l.Size() - 1
	cur := l.tail
	for cur != nil {
		if count <= index {
			callback(cur.Value)
		}
		count--
		cur = cur.Prev
	}
}

func (l *List[T]) Reverse() {
	cur := l.head
	var prev *Node[T] = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		cur.Prev = next
		prev = cur
		cur = next
	}
	l.tail = l.head
	l.head = prev
}

func TestList() {
	l := NewList[int]()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)
	l.PushBack(40)
	l.PushFront(5)
	l.PushFront(1)
	l.PrintList()

	var remVal int
	l.Remove(1, &remVal)
	fmt.Printf("Removed Value is: %d\n", remVal)
	l.PrintList()
	l.Reverse()
	fmt.Println("After Reverse")
	l.PrintList()

	fmt.Println("Test Traverse Back")
	l.TraverseBack(4, func(data int) {
		fmt.Printf("%d ", data)
	})
	fmt.Println()
	fmt.Println("Test Traverse Front")
	l.TraverseFront(0, func(data int) {
		fmt.Printf("%d ", data)
	})
	fmt.Println()
}
