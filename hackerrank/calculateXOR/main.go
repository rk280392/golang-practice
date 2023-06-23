package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var s string
	var t string
	fmt.Scan(&s)
	fmt.Scan(&t)

	if len(s) != len(t) {
		fmt.Println("Error")
	}
	var result string
	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			result += "1"
		} else {
			result += "0"
		}
	}
	fmt.Println(result)
}
