package main

import (
	"fmt"
)

func lonelyinteger(a []int32) int32 {

	var count int32
	for _, num := range a {
		count ^= num
		//XOR operator The XOR operation (^) between two numbers returns 0 if the numbers are the same and 1 if they are different
	}
	return count
}

func main() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	result := lonelyinteger(a)
	fmt.Println(result)
}
