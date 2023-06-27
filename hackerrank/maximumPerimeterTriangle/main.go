package main

import (
	"fmt"
	"sort"
)

/*
A non-degenerate triangle is a triangle that satisfies the triangle inequality rule, which states that the sum of the lengths of
any two sides of a triangle must be greater than the length of the third side. In other words, a non-degenerate triangle is a
triangle that has a positive area and is not collapsed into a line segment or a single point.

For example, a triangle with side lengths [3, 4, 5] is non-degenerate because it satisfies the triangle inequality:
3 + 4 > 5, 4 + 5 > 3, and 5 + 3 > 4. This triangle has a positive area and is not degenerate.

On the other hand, a triangle with side lengths [1, 2, 3] is degenerate because it does not satisfy the triangle inequality:
1 + 2 = 3, which means two sides are equal in length to the third side. This triangle is collapsed into a line segment and does not have a positive area.

*/

func maximumPerimeterTriangle(sticks []int32) []int32 {
	sort.Slice(sticks, func(i, j int) bool {
		return sticks[i] > sticks[j]

	})
	for i := 0; i <= len(sticks)-3; i++ {
		if sticks[i]+sticks[i+1] > sticks[i+2] && sticks[i]+sticks[i+2] > sticks[i+1] && sticks[i+2]+sticks[i+1] > sticks[i] {
			return []int32{sticks[i+2], sticks[i+1], sticks[i]}
		}
	}

	return []int32{-1}

}

func main() {

	sticks := []int32{1, 1, 1, 2, 3, 5}
	result := maximumPerimeterTriangle(sticks)
	fmt.Println(result)
}
