package linkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type DNode[T any] struct {
	Value T
	Next  *DNode[T]
	Prev  *DNode[T]
}

type List[T any] interface {
	// Basic operations
	InsertFront(data T)
	InsertBack(data T)
	DeleteFront() (T, error)
	DeleteBack() (T, error)

	// Utility methods
	IsEmpty() bool
	Length() int
	Clear()
	PrintList()

	// Search and access
	Contains(data T) bool
	PeekFront() (T, error)
	PeekBack() (T, error)
}
