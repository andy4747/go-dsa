package linkedlist

type SNode[T comparable] struct {
	Value T
	Next  *SNode[T]
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

type LinkedList[T comparable] interface {
	// Basic operations
	PushFront(data T)
	PushBack(data T)
	PopFront() error
	PopBack() error
	Insert(int, T) error
	Remove(int, T) error

	// Utility methods
	IsEmpty() bool
	Size() int
	Clear()
	PrintList()

	// Search and access
	Find(data T) int
	ForEach(callback func(data T))
	Reverse()
}
