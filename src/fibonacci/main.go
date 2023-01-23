// 0, 1, 1, 2, 3, 5, 8.... n=(n-1)+(n-2)
// This script checks if the given number is a fibonacci number

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkFibo(number int) bool {
	n1, n2 := 0, 1
	result := false

	// 0 and 1 are fibonacci numbers
	if number == 0 || number == 1 {
		result = true
		return result
	}

	// check each number from 1 to nth. If number == Nth, it is a fibonacci number
	for i := 1; i <= number; i++ {
		nth := n1 + n2
		n1 = n2
		n2 = nth
		if number == nth {
			result = true
			return result
		}
	}

	return result

}
func main() {

	// Read number from input
	value := bufio.NewReader(os.Stdin)
	fmt.Print("check number is fibo? ")
	input, err := value.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)   // trimspace from the input string
	number, err := strconv.Atoi(input) // convert string to int

	result := checkFibo(number)
	if result == true {
		fmt.Printf("%d is a fibonacci number \n", number)
	} else {
		fmt.Printf("%d is not a fibonacci number \n", number)
	}

}
