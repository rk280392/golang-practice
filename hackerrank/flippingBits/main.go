package main

import (
    "fmt"
)


func flippingBits(n int64) int64 {

    // Logic: create a variable with all 1s and run XOR operator to flip the bits
   // hexadecimal value of 11111111 is 0xFFFFFFFF
   mask := int64(0xFFFFFFFF)

    // Flip the bits by performing a bitwise XOR with the bitmask
    flipped := n ^ mask

    return flipped
}

func main() {
	n := int64(1)
        result := flippingBits(n)
        fmt.Println(result)
}

