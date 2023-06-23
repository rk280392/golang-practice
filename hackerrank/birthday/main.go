package main

import (
	"fmt"
)

func birthday(s []int32, d int32, m int32) int32 {
	var count int32
	var n int32 = int32(len(s))

	// Iterate through chocolate bar
	for i := int32(0); i < n-m; i++ {
		var sum int32

		for j := i; j < i+m; j++ {
			sum += s[j]
		}

		if sum == d {
			count++
		}
	}
	return count

}

func main() {

	var s []int32 = []int32{1, 2, 1, 3, 2}
	d := int32(3)
	m := int32(2)

	result := birthday(s, d, m)

	fmt.Println(result)
}
