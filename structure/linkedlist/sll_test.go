package linkedlist

import (
	"reflect"
	"testing"
)

func TestSLL(t *testing.T) {
	list := NewSLL[int]()
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

	list.AddToLast(40)
	list.AddToLast(50)
	list.AddToLast(60)
	list.AddToLast(70)

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

	t.Run("Reverse", func(t *testing.T) {
		want := []int{70, 60, 50, 40, 30, 20, 10}
		got := []int{}
		list.Reverse()
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
