package main

import (
	"fmt"
	"strings"
)

func marsExploration(s string) int32 {
	// Write your code here
	var count int32
	noOfSignals := len(s) / 3
	expectedStr := strings.Repeat("SOS", noOfSignals)
	fmt.Println(expectedStr)
	for i := 0; i < len(s) && i < len(expectedStr); i++ {
		if s[i] != expectedStr[i] {
			count++
		}
	}
	return count
}

func main() {
	s := "OOSDSSOSOSWEWSOSOSOSOSOSOSSSSOSOSOSS"
	result := marsExploration(s)
	fmt.Println(result)

}
