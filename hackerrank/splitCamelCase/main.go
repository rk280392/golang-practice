package main

import (
	"fmt"
	"strings"
	"unicode"
)

func splitCamelCase(words string) string {
	var currentWord string
	previousCharIsLower := false
	var result []string

	for _, ch := range words {
		//fmt.Println("first loop", string(ch))
		if unicode.IsUpper(ch) {
			//	fmt.Println("second loop", string(ch))
			//	fmt.Println("current word is", currentWord)
			if currentWord != "" {
				result = append(result, currentWord)
			}
			currentWord = strings.ToLower(string(ch))
			//	fmt.Println("current word in outside if", currentWord)
		} else {
			if previousCharIsLower {
				currentWord += string(ch)
			} else {
				// first uppercase letter will be handled here
				currentWord += string(ch)
			}
			previousCharIsLower = unicode.IsLower(ch)
		}
	}
	// for last word, as there won't be any uppercase letters
	if currentWord != "" {
		result = append(result, currentWord)
	}
	return strings.Join(result, " ")

}

func main() {
	input := "LargeSoftwareBook"
	output := splitCamelCase(input)
	fmt.Println(output)
}
