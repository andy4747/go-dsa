package patterns

import "fmt"

func PatternRightAngledNumberPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 1; j <= i+1; j++ {
			fmt.Printf("%d", j)
		}
		fmt.Println()
	}
}
