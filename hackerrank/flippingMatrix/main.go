package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'flippingMatrix' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */

func flippingMatrix(matrix [][]int32) int32 {
	n := len(matrix)
	var topLeftSum int32 = 0

	for i := 0; i < n/2; i++ {
		for j := 0; j < n/2; j++ {
			otherI := n - 1 - i // n-1 because index starts with 0, so example 4-1=3
			otherJ := n - 1 - j

			// loop through the matrix and find the max value between i,j,otheri and otherj combinations.

			maxVal := int(math.Max(math.Max(float64(matrix[i][j]), float64(matrix[otherI][j])), math.Max(float64(matrix[i][otherJ]), float64(matrix[otherI][otherJ]))))
			topLeftSum += int32(maxVal)
		}
	}

	return topLeftSum
}

func main() {
	matrix := [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}

	maxSum := flippingMatrix(matrix)
	fmt.Println("Maximum Sum:", maxSum)
}
