package arrays

import (
	"dsa/types"
)

// RemoveDuplicates remove duplicated from the arr
func RemoveDuplicates[T types.Number](arr []T) []T {
	if len(arr) == 0 {
		return arr
	}

	uniqueIndex := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[uniqueIndex] {
			uniqueIndex++
			arr[uniqueIndex] = arr[i]
		}
	}
	return arr[:uniqueIndex+1]
}
