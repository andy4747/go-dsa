package linkedlist

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

func NewNode[T comparable](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}
