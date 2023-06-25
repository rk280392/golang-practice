package main

import (
	"fmt"
)

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	frequencyMap := make(map[int32]int32)
	var maxCount int32
	var commonBird int32

	for _, num := range arr {
		frequencyMap[num]++
		if frequencyMap[num] > maxCount || (frequencyMap[num] == maxCount && num < commonBird) {
			maxCount = frequencyMap[num]
			commonBird = num
		}
	}
	return commonBird
}

func main() {
	var arr []int32 = []int32{1, 4, 4, 4, 5, 3}
	result := migratoryBirds(arr)
	fmt.Println(result)

}
