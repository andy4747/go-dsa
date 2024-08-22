package linkedlist

import "fmt"

type SinglyLinkedList[T comparable] struct {
	len  int
	Head *Node[T]
}

func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (l *SinglyLinkedList[T]) AddToHead(item T) {
	newNode := NewNode(item)
	newNode.Next = l.Head
	l.Head = newNode
	l.len++
}

func (l *SinglyLinkedList[T]) IterateList() {
	for cur := l.Head; cur != nil; cur = cur.Next {
		fmt.Println(cur.Value)
	}
}

func (l *SinglyLinkedList[T]) LastNode() *Node[T] {
	var lastNode *Node[T]
	for cur := l.Head; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			lastNode = cur
		}
	}
	return lastNode
}

func (l *SinglyLinkedList[T]) Count() int {
	return l.len
}

func (l *SinglyLinkedList[T]) AddToEnd(item T) {
	newNode := NewNode(item)
	if l.Head == nil {
		l.AddToHead(item)
	}
	lastNode := l.LastNode()
	lastNode.Next = newNode
}
