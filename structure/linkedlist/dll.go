package linkedlist

import "fmt"

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewList[T comparable]() *List[T] {
	return &List[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *List[T]) PushFront(data T) error {
	newNode := &Node[T]{Value: data}
	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return nil
	}
	newNode.Next = l.head
	l.head.Prev = newNode
	l.head = newNode
	l.size++
	return nil
}

func (l *List[T]) PushBack(data T) error {
	newNode := &Node[T]{Value: data}
	if l.head == nil && l.tail == nil {
		l.head = newNode
		l.tail = newNode
		l.size++
		return nil
	}
	l.tail.Next = newNode
	newNode.Prev = l.tail
	l.tail = newNode
	l.size++
	return nil
}

func (l *List[T]) PopBack() (T, error) {
	var result T
	if l.IsEmpty() {
		return result, fmt.Errorf("pop operation on empty list")
	}
	if l.head == l.tail {
		result = l.head.Value
		l.head = nil
		l.tail = nil
		l.size = 0
		return result, nil
	}
	result = l.tail.Value
	newTail := l.tail.Prev
	newTail.Next = nil
	l.tail = newTail
	l.size--
	return result, nil
}

func (l *List[T]) PopFront() (T, error) {
	var result T
	if l.IsEmpty() {
		return result, fmt.Errorf("pop operation on empty list")
	}
	if l.head == l.tail {
		result = l.head.Value
		l.head = nil
		l.tail = nil
		l.size = 0
		return result, nil
	}
	result = l.head.Value
	newHead := l.head.Next
	newHead.Prev = nil
	l.head = newHead
	l.size--
	return result, nil
}

func (l *List[T]) Insert(index int, data T) error {
	if l.IsEmpty() && index > 0 {
		return fmt.Errorf("error: index out of bound")
	}
	if index == 0 {
		return l.PushFront(data)
	}
	if index == l.Size() {
		return l.PushBack(data)
	}
	newNode := &Node[T]{Value: data}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	newNode.Next = cur.Next
	newNode.Prev = cur
	cur.Next.Prev = newNode
	cur.Next = newNode
	l.size++
	return nil
}

func (l *List[T]) Remove(index int, data *T) (T, error) {
	var result T
	if l.IsEmpty() {
		return result, fmt.Errorf("error: index out of bounds")
	}
	if index == 0 {
		return l.PopFront()
	}
	if index == l.size-1 {
		return l.PopBack()
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	result = cur.Value
	cur.Next.Prev = cur.Prev
	cur.Prev.Next = cur.Next
	l.size--
	return result, nil
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) IsEmpty() bool {
	return l.size == 0
}

func (l *List[T]) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *List[T]) PrintList() {
	cur := l.head
	fmt.Printf("[")
	for cur != nil {
		if cur.Next == nil {
			fmt.Printf("%v", cur.Value)
		} else {
			fmt.Printf("%v ", cur.Value)
		}
		cur = cur.Next
	}
	fmt.Printf("]\n")
}

// returns -1 if not found
func (l *List[T]) Find(data T) int {
	cur := l.head
	count := 0
	for cur != nil {
		if cur.Value == data {
			return count
		}
		count++
		cur = cur.Next
	}
	return -1
}

func (l *List[T]) ForEach(callback func(data T)) {
	cur := l.head
	for cur != nil {
		callback(cur.Value)
		cur = cur.Next
	}
}

func (l *List[T]) TraverseFront(index int, callback func(data T)) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("error: index out of bounds")
	}
	cur := l.head
	for i := 0; i < index; i++ {
		if cur == nil {
			return fmt.Errorf("error: unexpected nil node")
		}
		cur = cur.Next
	}
	for cur != nil {
		callback(cur.Value)
		cur = cur.Next
	}
	return nil
}

func (l *List[T]) TraverseBack(index int, callback func(data T)) error {
	if index < 0 || index >= l.size {
		return fmt.Errorf("error: index out of bounds")
	}
	cur := l.head
	for i := 0; i < index; i++ {
		if cur == nil {
			return fmt.Errorf("error: unexpected nil node")
		}
		cur = cur.Next
	}
	for cur != nil {
		callback(cur.Value)
		cur = cur.Prev
	}
	return nil
}

func (l *List[T]) Reverse() {
	if l.head == nil || l.head.Next == nil {
		//empty or single node, nothing to reverse
		return
	}
	cur := l.head
	var prev *Node[T] = nil
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		cur.Prev = next
		prev = cur
		cur = next
	}
	prev = l.head
	l.head = l.tail
	l.tail = prev
}

func CreateCycle(l *List[int]) {
	l.Clear()
	nodes := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150}
	start_index := 8
	for _, v := range nodes {
		l.PushBack(v)
	}
	cur := l.head
	for i := 0; i < start_index; i++ {
		cur = cur.Next
	}
	l.tail.Next = cur
}

func (l *List[T]) CycleExists() bool {
	fast := l.head
	slow := l.head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func (l *List[T]) CycleStartNode() *Node[T] {
	fast := l.head
	slow := l.head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = l.head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}

func TestDLL() {
	// Create a new list of integers.
	list := NewList[int]()
	// Test IsEmpty and Size on an empty list.
	if !list.IsEmpty() {
		fmt.Println("Test Failed: Expected list to be empty initially.")
	}
	if list.Size() != 0 {
		fmt.Printf("Test Failed: Expected size 0 initially, got %d\n", list.Size())
	}
	// Test PushBack and PushFront.
	if err := list.PushBack(10); err != nil {
		fmt.Println("Test Failed: PushBack error:", err)
	}
	if err := list.PushFront(5); err != nil {
		fmt.Println("Test Failed: PushFront error:", err)
	}
	if err := list.PushBack(15); err != nil {
		fmt.Println("Test Failed: PushBack error:", err)
	}
	// Expected list: 5, 10, 15.
	if list.Size() != 3 {
		fmt.Printf("Test Failed: Expected size 3 after pushes, got %d\n", list.Size())
	}
	fmt.Print("List after Push operations: ")
	list.PrintList()
	// Test Find.
	idx := list.Find(10)
	if idx != 1 {
		fmt.Printf("Test Failed: Expected value 10 at index 1, got index %d\n", idx)
	}
	// Test Insert: Insert 7 at index 1.
	if err := list.Insert(1, 7); err != nil {
		fmt.Println("Test Failed: Insert error:", err)
	}
	// Expected list: 5, 7, 10, 15.
	if list.Size() != 4 {
		fmt.Printf("Test Failed: Expected size 4 after insert, got %d\n", list.Size())
	}
	fmt.Print("List after Insert: ")
	list.PrintList()
	// Test Remove: Remove element at index 2 (should remove 10).
	// The Remove method expects a pointer argument (perhaps to store extra info).
	var dummy int
	removed, err := list.Remove(2, &dummy)
	if err != nil {
		fmt.Println("Test Failed: Remove error:", err)
	} else if removed != 10 {
		fmt.Printf("Test Failed: Expected removed value 10, got %v\n", removed)
	}
	if list.Size() != 3 {
		fmt.Printf("Test Failed: Expected size 3 after removal, got %d\n", list.Size())
	}
	fmt.Print("List after Remove: ")
	list.PrintList()
	// Test PopFront.
	front, err := list.PopFront()
	if err != nil {
		fmt.Println("Test Failed: PopFront error:", err)
	} else if front != 5 {
		fmt.Printf("Test Failed: Expected PopFront value 5, got %v\n", front)
	}
	// Test PopBack.
	back, err := list.PopBack()
	if err != nil {
		fmt.Println("Test Failed: PopBack error:", err)
	} else if back != 15 {
		fmt.Printf("Test Failed: Expected PopBack value 15, got %v\n", back)
	}
	if list.Size() != 1 {
		fmt.Printf("Test Failed: Expected size 1 after pops, got %d\n", list.Size())
	}
	fmt.Println("Before ForEach")
	list.PrintList()
	// Test ForEach: Should iterate over remaining elements.
	fmt.Print("ForEach output: ")
	list.ForEach(func(data int) {
		fmt.Printf("%v ", data)
	})
	fmt.Println()
	// Test TraverseFront.
	fmt.Print("TraverseFront output: ")
	if err = list.TraverseFront(0, func(data int) {
		fmt.Printf("%v ", data)
	}); err != nil {
		fmt.Println("Test Failed: TraverseFront error:", err)
	}
	fmt.Println()
	// Add another element for testing TraverseBack.
	if err = list.PushBack(20); err != nil {
		fmt.Println("Test Failed: PushBack error:", err)
	}
	// Now list should have: 7, 20.
	fmt.Print("List before TraverseBack: ")
	list.PrintList()
	fmt.Print("TraverseBack output: ")
	if err = list.TraverseBack(list.Size()-1, func(data int) {
		fmt.Printf("%v ", data)
	}); err != nil {
		fmt.Println("Test Failed: TraverseBack error:", err)
	}
	fmt.Println()
	// Test Clear.
	list.Clear()
	if !list.IsEmpty() {
		fmt.Println("Test Failed: Expected list to be empty after Clear.")
	} else {
		fmt.Println("List cleared successfully.")
	}
	// --- Test Reverse ---
	fmt.Println("Testing Reverse:")
	// Clear the list and add multiple elements.
	list.Clear()
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		if err := list.PushBack(v); err != nil {
			fmt.Println("Test Failed: PushBack error while adding", v, ":", err)
		}
	}
	fmt.Print("List before Reverse: ")
	list.PrintList()
	// Reverse the list.
	list.Reverse()
	fmt.Print("List after Reverse: ")
	list.PrintList()
	CreateCycle(list)
	cycleExists := list.CycleExists()
	if cycleExists {
		fmt.Println("Test Passed: Cycle exists in the list")
	} else {
		fmt.Println("Test Failed: cycle not found in the list")
	}
	cycleNode := list.CycleStartNode()
	if cycleNode == nil {
		fmt.Println("Test Failed: cycle start node not found")
	} else {
		fmt.Printf("Test Passed: cycle node value: %v\n", cycleNode.Value)
	}
	// --- End Reverse test ---
	// Finally, test Clear.
	list.Clear()
	if !list.IsEmpty() {
		fmt.Println("Test Failed: Expected list to be empty after Clear.")
	} else {
		fmt.Println("List cleared successfully.")
	}
	fmt.Println("All tests completed.")
}
