package main

import (
	"fmt"
	"sort"
)

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	maxA := a[len(a)-1]
	minB := b[0]
	var count int32
	isFactorial := func(x int32, arr []int32) bool {
		for _, num := range arr {
			if x%num != 0 {
				return false
			}
		}
		return true
	}

	isDivisible := func(x int32, arr []int32) bool {
		for _, num := range arr {
			if num%x != 0 {
				return false
			}
		}
		return true
	}
	for x := maxA; x <= minB; x += maxA {
		if isFactorial(x, a) && isDivisible(x, b) {
			fmt.Println(x)
			count++
		}
	}
	return count
}

func main() {
	var arr []int32 = []int32{4, 2}
	var brr []int32 = []int32{16, 32, 96}
	total := getTotalX(arr, brr)
	fmt.Println(total)
}
