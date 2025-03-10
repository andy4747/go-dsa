package queue

import "fmt"

type Queue[T comparable] interface {
	Enqueue(data T) error
	Dequeue() (T, error)
	Front() (T, error)
	Size() int
	IsEmpty() bool
	fmt.Stringer
}
