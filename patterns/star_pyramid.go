/*
Author: Angel Dhakal
File: star_pyramoid.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-7-star-pyramid/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

import "fmt"

func PatternStarPyramid(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}
		for j := 0; j < n-i-1; j++ {
			fmt.Printf(" ")
		}
		fmt.Println()
	}
}
