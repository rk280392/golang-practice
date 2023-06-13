package main

import (
	"fmt"
	"strconv"
)

func countBits(num uint32) int32 {
	s2 := strconv.FormatInt(int64(num), 2)
	var val int32
	for i := 0; i < len(s2); i++ {
		if s2[i] == '1' {
			val = val + 1
		}
	}
	return val
}

func main() {

	var num uint32 = 127
	result := countBits(num)
	fmt.Printf("%d\n", result)

}
