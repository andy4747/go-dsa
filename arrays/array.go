package arrays

import (
	"fmt"
)

type Data interface{}

/*
Design Dynamic Array (Resizable Array)
A Dynamic Array struct implementation, such as an ArrayList in Java or a vector in C++.

Array struct support the following operations:

  - int get(int i) will return the element at index i. Assume that index i is valid.

  - void set(int i, int n) will set the element at index i to n. Assume that index i is valid.

  - void pushback(int n) will push the element n to the end of the array.

  - int popback() will pop and return the element at the end of the array. Assume that the array is non-empty.

  - void resize() will double the capacity of the array.

when pushback(int n) is called but the array is full, the arr is resized first.
*/
type Array struct {
	len      int
	capacity int
	arr      []Data
}

// NewArray will initialize an empty array with a capacity of capacity, where capacity > 0.
func NewArray(capacity int) *Array {
	if capacity <= 0 {
		panic("capacity must be greater than zero")
	}
	return &Array{
		len:      0,
		capacity: capacity,
		arr:      make([]Data, capacity),
	}
}

// Get T will return the element at index i.
func (a *Array) Get(i int) Data {
	if i < 0 || i >= a.len {
		return nil
	}
	return a.arr[i]
}

func (a *Array) Add(v Data) {
	if a.len == a.capacity {
		a.Resize()
	}
	a.arr[a.len] = v
	a.len++
}

// Set will set the element at index i to v.
func (a *Array) Set(i int, data Data) {
	if i < 0 || i >= a.len {
		return
	}
	a.arr[i] = data
}

// func (a *Array[T]) PushBack(v T) {
// 	for i := 0; i< a.len; i++ {
// 		if a.arr[i] == v {
// 			a.arr[i], a.arr[a.len-1] = a.arr[a.len-1], a.arr[i]
// 		}
// 	}
// }

// PopBack will pop and return the element at the end of the array.
func (a *Array) PopBack() Data {
	item := a.arr[a.len-1]
	a.arr[a.len-1] = nil
	a.len--
	return item
}

// Resize will double the capacity of the array
func (a *Array) Resize() {
	newArr := make([]Data, a.capacity*2)
	newArr = append(newArr, a.arr...)
	a.arr = newArr
}

// CheckRange checks the index i exists
func (a *Array) CheckRange(i int) error {
	if i < 0 || i > +a.len {
		return fmt.Errorf("index out of bound")
	}
	return nil
}

// IsEmpty checks array is empty or not
func (a *Array) IsEmpty() bool {
	return a.len == 0
}

func (a *Array) Remove(i int) error {
	err := a.CheckRange(i)
	if err != nil {
		return err
	}
	copy(a.arr[i:], a.arr[i+1:a.len])
	a.arr[a.len-1] = nil
	a.len--
	return nil
}
