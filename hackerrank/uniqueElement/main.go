package main

import (
	"fmt"
)

func lonelyinteger(a []int32) int32 {

	var count int32
	for _, num := range a {
		count = count ^ num
		fmt.Println(count)
	}
	return count
}

func main() {
	a := []int32{1, 2, 3, 4, 3, 2, 1}
	result := lonelyinteger(a)
	fmt.Println(result)
}

/* 
Explanation:

1 = 00000001
2 = 00000010
3 = 00000011
4 = 00000100

NOTE:- XOR operator The XOR operation (^) between two numbers returns 0 if the bits are the same and 1 if they are different

loop for 1 results in count = 1
loop for 2 --> XOR of 1 and 2 (count) = 00000011 = 3
loop for 3 --> XOR of count(3) and 3 = 00000011 ^ 00000011 = 00000000 = 0
loop for 4 --> XOR of count(0) and 4 = 00000000 ^ 00000100 = 00000100 = 4
loop for 3 --> XOR of count(4) and 3 = 00000100 ^ 00000011 = 00000111 = 7
loop for 2 --> XOR of count(7) and 2 = 00000111 ^ 00000010 = 00000101 = 5
loop for 1 --> XOR of count(5) and 1 = 00000101 ^ 00000001 = 00000100 = 4

Answer is 4

This operation is reversible, meaning that if we XOR the result (4) with one of the operands (5 or 1), we will obtain the other operand. So the last result should be the unique number. 
This method only works if other numbers work exactly twice.  If there are multiple unique numbers or an odd number of occurrences for any number, the XOR approach alone cannot find the unique number(s) correctly


