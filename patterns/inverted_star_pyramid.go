/*
Author: Angel Dhakal
File: inverted_star_pyramid.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-8-inverted-star-pyramid/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/

package patterns

import "fmt"

func PatternInvertedStarPyramid(n int) {
	for i := n; i > 0; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*i-1; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < n-i; j++ {
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}
