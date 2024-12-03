package linkedlist

import "fmt"

type DoublyLinkedList[T comparable] struct {
	head   *DNode[T]
	tail   *DNode[T]
	length int
}

func NewDoublyLinkedList[T comparable]() List[T] {
	return &DoublyLinkedList[T]{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (dll *DoublyLinkedList[T]) InsertFront(data T) {
	if dll.IsEmpty() {
		newNode := &DNode[T]{Value: data, Next: nil, Prev: nil}
		dll.head = newNode
		dll.tail = newNode
		dll.length++
	} else {
		newNode := &DNode[T]{Value: data, Next: dll.head, Prev: nil}
		dll.head = newNode
		dll.length++
	}
}

func (dll *DoublyLinkedList[T]) InsertBack(data T) {
	if dll.IsEmpty() {
		newNode := &DNode[T]{Value: data, Next: nil, Prev: nil}
		dll.head = newNode
		dll.tail = newNode
		dll.length++
	} else {
		if dll.head == dll.tail {
			newNode := &DNode[T]{Value: data, Next: nil, Prev: dll.head}
			dll.head.Next = newNode
			dll.tail = newNode
			dll.length++
		} else {
			newNode := &DNode[T]{Value: data, Next: nil, Prev: dll.tail}
			dll.tail.Next = newNode
			dll.tail = newNode
			dll.length++
		}
	}
}

func (dll *DoublyLinkedList[T]) DeleteFront() (T, error) {
	var result T
	if dll.IsEmpty() {
		return result, fmt.Errorf("cannot perform delete operation on an empty list")
	} else {
		if dll.head != dll.tail {
			nextNode := dll.head.Next
			nextNode.Prev = nil
			dll.head = nextNode
			dll.length--
		} else {
			dll.head = nil
			dll.tail = nil
			dll.length = 0
		}
	}
	return result, nil
}

func (dll *DoublyLinkedList[T]) DeleteBack() (T, error) {
	var result T
	if dll.IsEmpty() {
		return result, fmt.Errorf("cannot perform delete operation on an empty list")
	} else {
		if dll.head == dll.tail {
			dll.head = nil
			dll.tail = nil
			dll.length = 0
		} else {
			prevNode := dll.tail.Prev
			prevNode.Next = nil
			dll.tail = prevNode
			dll.length--

		}
	}
	return result, nil
}

func (dll *DoublyLinkedList[T]) IsEmpty() bool {
	if dll.head == nil {
		return true
	}
	return false
}

func (dll *DoublyLinkedList[T]) Length() int {
	return dll.length
}

func (dll *DoublyLinkedList[T]) Clear() {
	dll.head = nil
	dll.tail = nil
	dll.length = 0
}

func (dll *DoublyLinkedList[T]) PrintList() {
	for cur := dll.head; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Value)
	}
	fmt.Println()
}

func (dll *DoublyLinkedList[T]) Contains(data T) bool {
	return false
}

func (dll *DoublyLinkedList[T]) PeekFront() (T, error) {
	var result T
	if dll.IsEmpty() {
		return result, fmt.Errorf("cannot peek on empty list")
	}
	result = dll.head.Value
	return result, nil
}

func (dll *DoublyLinkedList[T]) PeekBack() (T, error) {
	var result T
	if dll.IsEmpty() {
		return result, fmt.Errorf("cannot peek on empty list")
	}
	result = dll.tail.Value
	return result, nil
}

func DLLTest() {
	dll := NewDoublyLinkedList[int]()
	dll.InsertFront(10)
	// dll.InsertBack(20)
	// dll.InsertBack(30)
	// dll.InsertFront(05)

	dll.PrintList()
	dll.DeleteBack()
	dll.PrintList()
}
