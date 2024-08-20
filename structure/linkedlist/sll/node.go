package sll

type Node[T comparable] struct {
	value any
	next  *Node[T]
}
