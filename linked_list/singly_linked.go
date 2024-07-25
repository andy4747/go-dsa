package linkedlist

import (
	"dsa/mystrings"
	"fmt"
)

type SNode struct {
	value any
	next  *SNode
}

type SinglyLinkedList struct {
	head *SNode
	tail *SNode
}

func (sll *SinglyLinkedList) Insert(value any) {
	newNode := &SNode{value: value}
	if sll.head == nil { // if the list is empty set head and tail to same node
		sll.head = newNode
		sll.tail = newNode
	} else { // else list has items set the new node after the current tail node and update the new node as the tail node
		sll.tail.next = newNode // set the new node after the current tail node
		sll.tail = newNode      // update the new node as the tail node
	}
}

func (dll *SinglyLinkedList) String() string {
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

func SLLMain() {
	list := SinglyLinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.Insert(50)
	list.Insert(60)
	list.Insert(70)
	list.Insert(80)
	listStr := list.String()
	fmt.Println(listStr)
}
