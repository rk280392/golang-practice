package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'pangrams' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func pangrams(s string) string {
	// Write your code here
	lettrs := "abcdefghijklmnopqrstuvwxyz"
	slower := strings.ToLower(s)
	for _, str := range lettrs {
		if !strings.ContainsRune(slower, str) {
			return "not pangram"
		}
	}
	return "pangram"

}

func main() {
	s := "We promptly judged antique ivory buckles for the next prize"

	result := pangrams(s)

	fmt.Println(result)

}
