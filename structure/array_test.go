package structure

import (
	"testing"
)

func TestArray(t *testing.T) {
	t.Run("NewArray", func(t *testing.T) {
		arr := NewArray(5)
		if arr.capacity != 5 {
			t.Errorf("Expected capacity 5, got %d", arr.capacity)
		}
		if arr.len != 0 {
			t.Errorf("Expected length 0, got %d", arr.len)
		}
	})

	t.Run("Add and Get", func(t *testing.T) {
		arr := NewArray(3)
		arr.Add(1)
		arr.Add("two")
		arr.Add(3.0)

		if arr.Get(0) != 1 {
			t.Errorf("Expected 1, got %v", arr.Get(0))
		}
		if arr.Get(1) != "two" {
			t.Errorf("Expected 'two', got %v", arr.Get(1))
		}
		if arr.Get(2) != 3.0 {
			t.Errorf("Expected 3.0, got %v", arr.Get(2))
		}
	})

	t.Run("Set", func(t *testing.T) {
		arr := NewArray(3)
		arr.Add(1)
		arr.Set(0, "one")
		if arr.Get(0) != "one" {
			t.Errorf("Expected 'one', got %v", arr.Get(0))
		}
	})

	t.Run("PushBack and PopBack", func(t *testing.T) {
		arr := NewArray(2)
		arr.PushBack(1)
		arr.PushBack(2)
		arr.PushBack(3) // This should trigger a resize

		if arr.capacity != 4 {
			t.Errorf("Expected capacity 4 after resize, got %d", arr.capacity)
		}

		if arr.PopBack() != 3 {
			t.Errorf("Expected PopBack to return 3")
		}
		if arr.len != 2 {
			t.Errorf("Expected length 2 after PopBack, got %d", arr.len)
		}
	})

	t.Run("IsEmpty", func(t *testing.T) {
		arr := NewArray(2)
		if !arr.IsEmpty() {
			t.Errorf("Expected IsEmpty to return true for new array")
		}
		arr.Add(1)
		if arr.IsEmpty() {
			t.Errorf("Expected IsEmpty to return false after adding an element")
		}
	})

	t.Run("Remove", func(t *testing.T) {
		arr := NewArray(3)
		arr.Add(1)
		arr.Add(2)
		arr.Add(3)

		err := arr.Remove(1)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if arr.len != 2 {
			t.Errorf("Expected length 2 after Remove, got %d", arr.len)
		}
		if arr.Get(1) != 3 {
			t.Errorf("Expected 3 at index 1 after Remove, got %v", arr.Get(1))
		}
	})

	t.Run("CheckRange", func(t *testing.T) {
		arr := NewArray(2)
		arr.Add(1)

		if err := arr.CheckRange(0); err != nil {
			t.Errorf("CheckRange(0) should not return an error")
		}
		if err := arr.CheckRange(1); err != nil {
			t.Errorf("CheckRange(1) should not return an error")
		}
		if err := arr.CheckRange(2); err == nil {
			t.Errorf("CheckRange(2) should return an error")
		}
		if err := arr.CheckRange(-1); err == nil {
			t.Errorf("CheckRange(-1) should return an error")
		}
	})
}
