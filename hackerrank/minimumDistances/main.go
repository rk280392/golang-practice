package main

import (
	"fmt"
	"math"
)

func minimumDistance(a []int32) int32 {
	frequencyMapIndex := make(map[int32]int32)
	minDistance := math.MaxInt32

	for i, num := range a {
		if j, k := frequencyMapIndex[num]; k {
			distance := i - int(j)
			if minDistance > distance {
				minDistance = distance
			}
		}
		frequencyMapIndex[num] = int32(i)
	}
	if minDistance == math.MaxInt32 {
		return -1
	} else {
		return int32(minDistance)
	}
}

func main() {
	a := []int32{7, 1, 3, 4, 1, 7}
	result := minimumDistance(a)
	fmt.Println(result)
}

/*
Explanation:

1 - Create a map to store indeces
2 - set minDistance to the maxInt32 value
3 - Iterate through array, checks if index is already present, update j and  if it is present k is true
4 - Calculate the diff between index and set to duistance
5 - Check if this distance is smaller than minDistance, if smaller update minDistance
6 - Store current index
7 - Return minDistance


*/
