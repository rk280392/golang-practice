package main

import (
	"fmt"
)

func pickingNumbers(a []int32) int32 {
	// Write your code her

	frequencies := make(map[int32]int32)

	// Count the frequencies of each element
	for _, num := range a {
		frequencies[num]++
	}
	fmt.Println(frequencies)
	var maxLength int32 = 0

	// Check the maximum length of subaays
	for num := range frequencies {
		//
		fmt.Println("num", num)
		fmt.Println("num+1", num+1)
		if frequencies[num]+frequencies[num+1] > maxLength {
			maxLength = frequencies[num] + frequencies[num+1]
			fmt.Println("Maxl========================", maxLength)
		}
	}

	return maxLength
}

func main() {
	// Input slice
	a := []int32{14, 18, 17, 10, 9, 20, 4, 13, 19, 19, 8, 15, 15, 17, 6, 5, 15, 12, 18, 2, 18, 7, 20, 8, 2, 8, 11, 2, 16, 2, 12, 9, 3, 6, 9, 9, 13, 7, 4, 6, 19, 7, 2, 4, 3, 4, 14, 3, 4, 9, 17, 9, 4, 20, 10, 16, 12, 1, 16, 4, 15, 15, 9, 13, 6, 3, 8, 4, 7, 14, 16, 18, 20, 11, 20, 14, 20, 12, 15, 4, 5, 10, 10, 20, 11, 18, 5, 20, 13, 4, 18, 1, 14, 3, 20, 19, 14, 2, 5, 13}

	// Call the pickingNumbers function
	result := pickingNumbers(a)

	fmt.Println(result)
}
