package main

import (
	"fmt"
	"sort"
)

func twoArrays(k int32, A []int32, B []int32) string {
	// Sort A

	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})

	// sort B in reverse

	sort.Slice(B, func(i, j int) bool {
		return B[i] > B[j]
	})

	// once sorted check the sum of pairs and return "NO" if sum is less than k, else return "YES"

	for i := range A {
		if A[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var t int32
	fmt.Scan(&t)

	for i := 0; int32(i) < t; i++ {
		var n, k int32
		fmt.Scan(&n, &k)

		arrA := make([]int32, n)
		arrB := make([]int32, n)

		for j := 0; int32(j) < n; j++ {
			fmt.Scan(&arrA[j])
		}

		for j := 0; int32(j) < n; j++ {
			fmt.Scan(&arrB[j])
		}

		result := twoArrays(k, arrA, arrB)
		fmt.Println(result)
	}
}
