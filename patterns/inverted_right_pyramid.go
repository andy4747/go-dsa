package patterns

import "fmt"

func PatternInvertedRightPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
