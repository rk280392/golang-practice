package main

import (
	"fmt"
)

/*
 * Complete the 'rotateLeft' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER d
 *  2. INTEGER_ARRAY arr
 */

func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here

	n := len(arr)
	d = d % int32(n) // cases where d > n (4 % 5 = 4)

	fmt.Println(d)

	// Create a new rotated array
	rotatedArray := make([]int32, n)
	for i := 0; i < n; i++ {
		newIndex := (i + int(d)) % n
		rotatedArray[i] = arr[newIndex]
	}

	return rotatedArray
}

func main() {
	d := int32(4)
	var arr []int32 = []int32{1, 2, 3, 4, 5}
	result := rotateLeft(d, arr)
	fmt.Println(result)
}
