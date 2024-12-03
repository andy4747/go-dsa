package linkedlist

import "fmt"

type SinglyLinkedList[T comparable] struct {
	head   *Node[T]
	length int
}

func NewSinglyLinkedList[T comparable]() List[T] {
	return &SinglyLinkedList[T]{
		head:   nil,
		length: 0,
	}
}

func (sll *SinglyLinkedList[T]) InsertFront(data T) {
	if sll.IsEmpty() {
		sll.head = &Node[T]{Value: data}
		sll.length++
	} else {
		newNode := &Node[T]{Value: data, Next: sll.head}
		sll.head = newNode
		sll.length++
	}
}

func (sll *SinglyLinkedList[T]) InsertBack(data T) {
	lastNode := sll.lastNode()
	newNode := &Node[T]{Value: data, Next: nil}
	lastNode.Next = newNode
	sll.length++
}

func (sll *SinglyLinkedList[T]) DeleteFront() (T, error) {
	var result T
	if sll.IsEmpty() {
		return result, fmt.Errorf("list is empty. cannot delete from front")
	}
	if sll.head.Next == nil {
		result = sll.head.Value
		sll.head = nil
		sll.length--
		return result, nil
	} else {
		result = sll.head.Value
		newHead := sll.head.Next
		sll.head = newHead
		sll.length--
		return result, nil
	}
}

func (sll *SinglyLinkedList[T]) DeleteBack() (T, error) {
	var result T
	for cur := sll.head; cur != nil; cur = cur.Next {
		if cur.Next.Next == nil {
			result = cur.Next.Value
			cur.Next = nil
			sll.length--
		}
	}
	return result, nil
}

func (sll *SinglyLinkedList[T]) lastNode() *Node[T] {
	var lastNode *Node[T]
	for cur := sll.head; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			lastNode = cur
		}
	}
	return lastNode
}

func (sll *SinglyLinkedList[T]) IsEmpty() bool {
	if sll.head == nil {
		return true
	}
	return false
}

func (sll *SinglyLinkedList[T]) Length() int {
	return sll.length
}

func (sll *SinglyLinkedList[T]) Clear() {
	sll.head = nil
	sll.length = 0
}

func (sll *SinglyLinkedList[T]) PrintList() {
	cur := sll.head
	for cur != nil {
		fmt.Printf("%v ", cur.Value)
		cur = cur.Next
	}
	fmt.Println()
}

func (sll *SinglyLinkedList[T]) Contains(data T) bool {
	for cur := sll.head; cur != nil; cur = cur.Next {
		if cur.Value == data {
			return true
		}
	}
	return false
}

func (sll *SinglyLinkedList[T]) PeekFront() (T, error) {
	var result T
	if sll.IsEmpty() {
		return result, fmt.Errorf("list is empty. no values to peek")
	}
	return sll.head.Value, nil
}

func (sll *SinglyLinkedList[T]) PeekBack() (T, error) {
	var result T
	if sll.IsEmpty() {
		return result, fmt.Errorf("list is empty. no values to peek")
	}
	lastNode := sll.lastNode()
	return lastNode.Value, nil
}

func SLLTest() {
	sll := NewSinglyLinkedList[int]()
	sll.InsertFront(10)
	sll.InsertFront(20)
	sll.InsertFront(30)
	sll.InsertFront(40)
	sll.InsertFront(50)
	sll.InsertFront(60)
	sll.InsertFront(70)
	sll.InsertFront(80)
	sll.InsertBack(01)

	sll.PrintList()

	sll.DeleteFront()
	sll.DeleteFront()

	sll.PrintList()

	sll.DeleteBack()
	sll.DeleteBack()

	sll.PrintList()

	fmt.Println(sll.Contains(20))

	fmt.Println(sll.PeekFront())
	fmt.Println(sll.PeekBack())

	sll.PrintList()
	sll.Clear()
	sll.PrintList()
}
