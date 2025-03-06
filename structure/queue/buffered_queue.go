// fixed size queue with circular buffer
package queue

import (
	"errors"
	"fmt"
	"strings"
)

type CircularQueue[T comparable] struct {
	front int
	rear  int
	cap   int
	size  int
	items []T
}

func NewCircularQueue[T comparable](cap int) Queue[T] {
	return &CircularQueue[T]{
		front: -1,
		rear:  -1,
		cap:   cap,
		size:  0,
		items: make([]T, cap),
	}
}

func (q *CircularQueue[T]) Enqueue(data T) error {
	if q.Size() == q.cap {
		return errors.New("queue is full")
	}
	if q.IsEmpty() {
		q.front = 0
		q.rear = 0
	} else {
		q.rear = (q.rear + 1) % q.cap
	}
	q.items[q.rear] = data
	q.size++
	return nil
}

func (q *CircularQueue[T]) Dequeue() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, errors.New("empty queue")
	}
	result = q.items[q.front]
	if q.front == q.rear {
		q.front, q.rear = -1, -1
	} else {
		q.front = (q.front + 1) % q.cap
	}
	q.size--
	return result, nil
}

func (q *CircularQueue[T]) Front() (T, error) {
	var result T
	if q.IsEmpty() {
		return result, errors.New("lookup on empty queue")
	}
	result = q.items[q.front]
	return result, nil
}

func (q *CircularQueue[T]) Size() int {
	return q.size
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.Size() == 0
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.Size() == q.cap
}

func (q *CircularQueue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	if !q.IsEmpty() {
		cur := q.front
		for i := 0; i < q.size; i++ {
			if sb.Len() > 1 {
				sb.WriteString(", ")
			}
			fmt.Fprintf(&sb, "%v", q.items[cur])
			cur = (cur + 1) % q.cap
		}

	}
	sb.WriteString("]")
	return sb.String()
}

func TestCircularQueue() {
	q := NewCircularQueue[int](5)
	// Test Enqueue
	fmt.Println("Enqueueing elements...")
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	fmt.Println("Queue after enqueues:", q) // [10, 20, 30]
	// Test Front
	front, err := q.Front()
	if err != nil {
		fmt.Println("Front error:", err)
	} else {
		fmt.Println("Front element:", front) // 10
	}
	// Test Dequeue
	fmt.Println("Dequeuing elements...")
	dequeued, err := q.Dequeue()
	if err != nil {
		fmt.Println("Dequeue error:", err)
	} else {
		fmt.Println("Dequeued element:", dequeued) // 10
	}
	fmt.Println("Queue after dequeue:", q) // [20, 30]
	// Test Enqueue again (to test circular behavior)
	fmt.Println("Enqueueing more elements...")
	q.Enqueue(40)
	q.Enqueue(50)
	q.Enqueue(60)                                // This should wrap around if the queue is circular
	fmt.Println("Queue after more enqueues:", q) // [20, 30, 40, 50, 60]
	// Test Size
	fmt.Println("Queue size:", q.Size()) // 5
	// Test IsEmpty
	fmt.Println("Is queue empty?", q.IsEmpty()) // false
	// Test Dequeue until empty
	fmt.Println("Dequeuing until empty...")
	for !q.IsEmpty() {
		val, err := q.Dequeue()
		if err != nil {
			fmt.Println("Dequeue error:", err)
		} else {
			fmt.Println("Dequeued:", val)
		}
	}
	fmt.Println("Queue after dequeuing all elements:", q) // []
	// Test IsEmpty after dequeuing all elements
	fmt.Println("Is queue empty?", q.IsEmpty()) // true
	// Test Front and Dequeue on empty queue
	_, err = q.Front()
	if err != nil {
		fmt.Println("Front error (expected):", err) // Queue is empty
	}
	_, err = q.Dequeue()
	if err != nil {
		fmt.Println("Dequeue error (expected):", err) // Queue is empty
	}
	// Test Enqueue after wrapping around
	fmt.Println("Enqueueing after wrapping around...")
	q.Enqueue(70)
	q.Enqueue(80)
	fmt.Println("Queue after wrapping around:", q) // [70, 80]
}
