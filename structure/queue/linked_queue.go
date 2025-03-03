// linked list queue
package queue

import (
	"errors"
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type LinkedQueue[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewLinkedQueue[T comparable]() Queue[T] {
	return &LinkedQueue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *LinkedQueue[T]) Enqueue(data T) {
	newNode := &Node[T]{
		Value: data,
		Next:  nil,
	}
	if q.IsEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.Next = newNode
		q.tail = newNode
	}
	q.size++
}

func (q *LinkedQueue[T]) Dequeue() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, errors.New("dequeue on empty queue")
	}
	result = q.head.Value
	q.head = q.head.Next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return result, nil
}

func (q *LinkedQueue[T]) Front() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, errors.New("lookup on empty queue")
	}
	result = q.head.Value
	return result, nil
}

func (q *LinkedQueue[T]) Size() int {
	return q.size
}

func (q *LinkedQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *LinkedQueue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	cur := q.head
	for cur != nil {
		if sb.Len() > 1 {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%v", cur.Value)
		cur = cur.Next
	}
	sb.WriteString("]")
	return sb.String()
}

func TestLinkedQueue() {
	q := NewLinkedQueue[int]()
	// Enqueue elements
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println("Queue:", q) // [10, 20, 30]
	// Front element
	front, _ := q.Front()
	fmt.Println("Front:", front) // 10
	// Dequeue elements
	dequeued, _ := q.Dequeue()
	fmt.Println("Dequeued:", dequeued) // 10
	fmt.Println("Queue:", q)           // [20, 30]
	// Enqueue more
	q.Enqueue(40)
	q.Enqueue(50)
	fmt.Println("Queue:", q) // [20, 30, 40, 50]
	// Multiple dequeues
	q.Dequeue()
	q.Dequeue()
	fmt.Println("Queue:", q) // [40, 50]
	// Verify size
	fmt.Println("Size:", q.Size()) // 2
	// Empty check
	fmt.Println("Empty?", q.IsEmpty()) // false
	// Dequeue remaining
	q.Dequeue()
	q.Dequeue()
	fmt.Println("Empty?", q.IsEmpty()) // true
}
