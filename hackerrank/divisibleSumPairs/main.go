package main

import (
	"fmt"
)

/*
 * Complete the 'divisibleSumPairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER_ARRAY ar
 */

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	// Write your code here
	pairs := 0
	for i, num1 := range ar {
		for _, num2 := range ar[i+1:] {
			if (num1+num2)%k == 0 {
				pairs += 1
			}
		}
	}
	return int32(pairs)
}

func main() {
	n := int32(8)
	k := int32(5)
	ar := []int32{1, 5, 5, 7, 0, 3, 4, 6}
	result := divisibleSumPairs(n, k, ar)
	fmt.Println(result)

}
