package maps

import "fmt"

func CountOccurance(arr []int) {
	counts := make(map[int]int)
	for _, item := range arr {
		if _, ok := counts[item]; !ok {
			counts[item] = 1
		} else {
			counts[item] += 1
		}
	}
	fmt.Printf("%+v\n", counts)
}
