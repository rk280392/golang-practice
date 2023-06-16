package main

import (
	"fmt"
)

/*
 * Complete the 'matchingStrings' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY strings
 *  2. STRING_ARRAY queries
 */

func matchingStrings(str []string, queries []string) []int32 {
	// Write your code here
	var countArray []int32
	var count int32
	for _, str1 := range queries {
		for _, str2 := range str {
			if str1 == str2 {
				count++
			}
		}
		countArray = append(countArray, count)
		count = 0
	}
	return countArray
}

func main() {

	str := []string{"aba", "baba", "aba", "xzxb"}
	queries := []string{"aba", "xzxb", "ab"}
	res := matchingStrings(str, queries)
	fmt.Println(res)
}
