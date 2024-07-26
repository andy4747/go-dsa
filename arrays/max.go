package arrays

import "dsa/types"

func Max[T types.Number](arr []T) T {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
