package arrays

import (
	"dsa/types"
	"fmt"
)

func ReverseInPlace[T types.Number](arr []T) {
	fmt.Printf("Before: %v\n", arr)
	left := 0
	right := len(arr) - 1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	fmt.Printf("After: %v\n", arr)
}
