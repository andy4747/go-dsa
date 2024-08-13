package structure

import (
	"dsa/strings"
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

func (dll *DoublyLinkedList) Delete() *Node {
	if dll.head == nil {
		return nil
	}
	removedNode := dll.tail

	dll.tail = dll.tail.prev
	if dll.tail != nil {
		dll.tail.next = nil
	} else {
		dll.head = nil
	}
	removedNode.prev = nil
	return removedNode
}

func (dll *DoublyLinkedList) String() string {
	var sb strings.Builder
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
	list.Delete()
	listStr := list.String()
	fmt.Println(listStr)
}
