package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

func diagonalDifference(arr [][]int32) int32 {

	// In a square matrix, the main diagonal consists of elements where the row index is equal to the column index (i == j)
	// In a square matrix, the secondary diagonal consists of elements where the sum of the row index and the column index is equal to matrixSize - 1
	var sumMainDiag int32
	var sumSecondDiag int32
	for i := range arr {
		for j := range arr {
			if i == j {
				sumMainDiag += arr[i][j]
			}
			if i+j == len(arr)-1 {
				sumSecondDiag += arr[i][j]

			}
		}
	}
	return int32(math.Abs(float64(sumMainDiag - sumSecondDiag)))

}

func main() {
	var arr [][]int32
	result := diagonalDifference(arr)

	fmt.Println(result)

}
