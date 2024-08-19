/*
Author: Angel Dhakal
File: half_diamond_star.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-10-half-diamond-star-pattern/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/

package patterns

import "fmt"

func PatternHaldDiamondStar(n int) {
	//top triangle
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	//bottom triangle
	for i := 0; i < n-1; i++ {
		for j := n - i - 1; j > 0; j-- {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
