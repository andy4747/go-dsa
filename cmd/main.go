package main

import (
	"dsa/arrays"
	"fmt"
)

func main() {
	// linkedlist.DLLMain()
	// fmt.Println(arrays.Max([]int{1, 2, 3, 4, 6, 7, 8, 10, 12, 17, 24}))
	// arrays.ReverseInPlace([]int{1, 2, 3, 4, 6, 7, 8, 10, 12, 17, 24})
	arr := []int{1, 1, 2, 2, 3, 4, 4, 5}
	arr = arrays.RemoveDuplicates(arr)
	fmt.Println(arr)
}
