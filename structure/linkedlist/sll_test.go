package linkedlist

import (
	"reflect"
	"testing"
)

func TestSLL(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.AddToHead(30)
	list.AddToHead(20)
	t.Run("AddToHead", func(t *testing.T) {
		want := []int{20, 30}
		got := []int{}
		for cur := list.Head; cur != nil; cur = cur.Next {
			got = append(got, cur.Value)
		}
		if reflect.DeepEqual(want, got) {
			t.Logf("Test Passed. Want %#v, Got %#v", want, got)
		} else {
			t.Errorf("Test Failed. Want %#v, Got %#v", want, got)
		}
	})
	list.AddToHead(10)
	t.Run("Count", func(t *testing.T) {
		wantCount := len([]int{10, 20, 30})
		count := list.Count()
		if count == wantCount {
			t.Logf("Test Passed. Want Count: %v, Got Count: %v", wantCount, count)
		} else {
			t.Errorf("Test Failed. Want Count: %v, Got Count: %v", wantCount, count)
		}
	})
	t.Run("LastNode", func(t *testing.T) {
		currentLastNodeValue := 30
		lastNode := list.LastNode()
		if lastNode.Next != nil {
			t.Errorf("Test Failed. Next is not nil. Not a last node.")
		}
		if lastNode.Value == currentLastNodeValue {
			t.Logf("Test Passed. Want Node With Value: %v, Got Node With Value: %v", currentLastNodeValue, lastNode.Value)
		} else {
			t.Errorf("Test Failed. Want Node With Value: %v, Got Node With Value: %v", currentLastNodeValue, lastNode.Value)
		}
	})
	list.AddToEnd(40)
	list.AddToEnd(50)
	list.AddToEnd(60)
	list.AddToEnd(70)
	t.Run("AddToEnd", func(t *testing.T) {
		want := []int{10, 20, 30, 40, 50, 60, 70}
		got := []int{}
		for cur := list.Head; cur != nil; cur = cur.Next {
			got = append(got, cur.Value)
		}
		if reflect.DeepEqual(want, got) {
			t.Logf("Test Passed. Want %#v, Got %#v", want, got)
		} else {
			t.Errorf("Test Failed. Want %#v, Got %#v", want, got)
		}
	})
}
