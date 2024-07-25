package linkedlist

import (
	"dsa/mystrings"
	"fmt"
)

type Node struct {
	prev  *Node
	value any
	next  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) Insert(value any) {
	newNode := &Node{value: value}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) String() string {
	var sb mystrings.Builder
	sb.WriteString("[ ")
	currentNode := dll.head
	for currentNode != nil {
		sb.WriteString(fmt.Sprintf("%v", currentNode.value))
		if currentNode.next != nil {
			sb.WriteString(", ")
		}
		currentNode = currentNode.next
	}
	sb.WriteString(" ]")
	return sb.String()
}

func DLLMain() {
	list := DoublyLinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	listStr := list.String()
	fmt.Println(listStr)
}
