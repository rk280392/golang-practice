package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rk280392/golang-practice/hackerrank/combineCamelCase"
	"github.com/rk280392/golang-practice/hackerrank/splitCamelCase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
S;M;plasticCup()
C;V;mobile phone
C;C;coffee machine
S;C;LargeSoftwareBook
C;M;white sheet of paper
S;V;pictureFrame
*/

func processInput(input string) string {
	splittedString := strings.Split(input, ";")
	operation := splittedString[0]
	valueType := splittedString[1]
	words := strings.TrimSpace(splittedString[2])

	if operation == "S" {
		if valueType == "M" {
			words = words[:len(words)-2]
			return splitCamelCase.SplitCamelCase(words)
		}
		return splitCamelCase.SplitCamelCase(words)

	} else if operation == "C" {
		if valueType == "C" {
			parts := strings.Split(words, " ")
			for i, _ := range parts {
				parts[i] = cases.Title(language.Und).String(parts[i])
			}
			fmt.Println(strings.Join(parts, ""))
			return strings.Join(parts, "")
		} else if valueType == "M" {
			return combineCamelCase.CombineCamelCase(words) + "()"
		} else {
			return combineCamelCase.CombineCamelCase(words)
		}
	}
	return ""
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	for scanner.Scan() {
		line := scanner.Text()
		result := processInput(line)
		fmt.Println(result)
	}
}
