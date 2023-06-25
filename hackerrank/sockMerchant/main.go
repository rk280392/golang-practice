package main

import "fmt"

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */

func sockMerchant(n int32, ar []int32) int32 {
	// Write your code here
	sockMap := make(map[int32]int32)

	for _, num := range ar {
		sockMap[num]++
	}
	var count int32 = 0
	for _, v := range sockMap {
		n := v / 2
		count += n
	}
	return count

}

func main() {
	n := int32(10)
	var ar []int32 = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20, 10}

	result := sockMerchant(n, ar)
	fmt.Println(result)

}
