/*
Author: Angel Dhakal
File: diamond_star_pyramid.go
Date: 2024-08-18
Problem: https://takeuforward.org/pattern/pattern-9-diamond-star-pattern/
Time Complexity: O(n^2)
Space Complexity: O(1)
*/
package patterns

func PatternDiamondStarPyramid(n int) {
	PatternStarPyramid(n)
	PatternInvertedStarPyramid(n)
}
