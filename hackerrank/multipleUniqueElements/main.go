package main

import "fmt"

func findUnique(numbers []int32) []int32 {
	count := make(map[int32]int32)

	for _, num := range numbers {
		count[num]++
	}

	uniqueElements := []int32{}
	for i, k := range count {
		if k == 1 {
			uniqueElements = append(uniqueElements, i)
		}
	}
	return uniqueElements
}

func main() {
	numbers := []int32{2, 4, 5, 2, 4, 6, 7, 8, 5}
	uniqueNumbers := findUnique(numbers)
	fmt.Println(uniqueNumbers)
}

/*

Explanation:

	Make a map of numbers with key and value as the number and the frequency of the number.
	Later check the map if frequency == 1, then it is a unique number. Store it in the slice.
*/
