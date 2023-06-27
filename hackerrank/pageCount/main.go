package main

import (
	"fmt"
)

/*
 * Complete the 'pageCount' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER p
 */

func pageCount(n int32, p int32) int32 {
	totalPages := n / 2
	targetPage := p / 2

	if p == 1 || p == n {
		return 0
	}
	fromFront := targetPage
	fromBack := totalPages - targetPage
	// fmt.Println(fromBack)

	// Check for odd total pages.
	if n%2 == 0 {
		fromBack += n % 2
	}

	if fromFront < fromBack {
		return fromFront
	}
	return fromBack

}

func main() {
	n := int32(6)
	p := int32(2)

	result := pageCount(n, p)

	fmt.Println(result)
}
