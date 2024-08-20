package stacks

import (
	"slices"
	"testing"
)

func cleanup[T comparable](stk *Stack[T]) {
	stk.elements = make([]T, 0)
}

func TestStackArray(t *testing.T) {
	tests := []struct {
		list []int
		size int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{10, 20, 30, 40, 50}, 5},
		{[]int{100}, 1},        // Test case for a single element stack
		{[]int{-1, -2, -3}, 3}, // Test case with negative numbers
	} //push
	t.Run("Push", func(t *testing.T) {
		skt := NewStack[int]()
		for i, test := range tests {
			for _, v := range test.list {
				skt.Push(v)
			}
			if slices.Equal(skt.elements, test.list) {
				t.Logf("Test Passed #%d\n", i)
			} else {
				t.Errorf("Test Failed #%d. Want %#v, Got %#v", i, test.list, skt.elements)
			}
			cleanup(skt)
		}
	})

	//test size
	t.Run("Stack Size", func(t *testing.T) {
		skt := NewStack[int]()
		for i, test := range tests {
			for _, v := range test.list {
				skt.Push(v)
			}
			if test.size == skt.Size() {
				t.Logf("Test Passed #%d\n", i)
			} else {
				t.Errorf("Test Failed #%d\n", i)
			}
			cleanup(skt)
		}
	})

	//test pop
	t.Run("Pop", func(t *testing.T) {
		skt := NewStack[int]()
		for i, test := range tests {
			for _, v := range test.list {
				skt.Push(v)
			}
			popped := skt.Pop()
			//comparing popped element with the last element in test arr
			if popped == test.list[len(test.list)-1] {
				t.Logf("Test Passed #%d\n", i)
			} else {
				t.Errorf("Test Failed #%d\n", i)
			}
			cleanup(skt)
		}
	})
}
