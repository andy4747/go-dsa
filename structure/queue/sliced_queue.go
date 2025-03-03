// slice queue
package queue

import (
	"errors"
	"fmt"
	"strings"
)

type SliceQueue[T comparable] struct {
	items []T
	head  int
}

func NewSliceQueue[T comparable]() Queue[T] {
	return &SliceQueue[T]{
		items: make([]T, 0),
		head:  0,
	}
}

func (q *SliceQueue[T]) Enqueue(data T) {
	q.items = append(q.items, data)
}

func (q *SliceQueue[T]) Dequeue() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, errors.New("dequeue on empty queue")
	}
	result = q.items[q.head]
	q.head++
	if q.head > len(q.items)/2 {
		q.items = q.items[q.head:]
		q.head = 0
	}
	return result, nil
}

func (q *SliceQueue[T]) Front() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, errors.New("lookup on empty queue")
	}
	result = q.items[q.head]
	return result, nil
}

func (q *SliceQueue[T]) Size() int {
	return len(q.items) - q.head
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *SliceQueue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i := q.head; i < len(q.items); i++ {
		if i > q.head {
			sb.WriteString(", ")
		}
		fmt.Fprintf(&sb, "%v", q.items[i])
	}
	sb.WriteString("]")
	return sb.String()
}

func TestSliceQueue() {

	q := NewSliceQueue[int]()

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
	fmt.Printf("String Method: %v\n", q.String())

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
