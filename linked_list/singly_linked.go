package linkedlist

import "fmt"

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

func (sll *SinglyLinkedList) PrintList() {
	currentNode := sll.head
	fmt.Printf("[ ")
	for currentNode != nil {
		if currentNode.next != nil {
			fmt.Printf("%v, ", currentNode.value)
		} else {
			fmt.Printf("%v ]", currentNode.value)
		}
		currentNode = currentNode.next
	}
	fmt.Println()
}

func SLLMainExample() {
	list := SinglyLinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.Insert(50)
	list.Insert(60)
	list.Insert(70)
	list.Insert(80)
	list.PrintList()
}
