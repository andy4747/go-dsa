package linkedlist

type SLL[T comparable] struct {
	Head *Node[T]
	len  int
}

func NewSLL[T comparable]() *SLL[T] {
	return &SLL[T]{}
}

func (l *SLL[T]) AddToHead(value T) {
	newNode := NewNode(value)
	newNode.Next = l.Head
	l.Head = newNode
	l.len++
}

func (l *SLL[T]) LastNode() *Node[T] {
	var lastNode *Node[T]
	for cur := l.Head; cur != nil; cur = cur.Next {
		if cur.Next == nil {
			lastNode = cur
		}
	}
	return lastNode
}

func (l *SLL[T]) Count() int {
	return l.len
}

func (l *SLL[T]) AddToLast(value T) {
	newNode := NewNode(value)
	lastNode := l.LastNode()
	lastNode.Next = newNode
	newNode.Next = nil
	l.len++
}
