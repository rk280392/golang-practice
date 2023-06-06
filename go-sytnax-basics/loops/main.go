package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ask the user for a series of numbers and add them together
// It should stop only when the user types “n” in response to the “Add more?” prompt.
// (Notice that the “Y” in “Y/n” is capitalized, indicating that it’s the default.
// If the user hits Enter without typing anything, or indeed if they enter anything at all besides “n”,
// the program should ask for another number to add.)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var sum float64
	var more string = "yes"

	for more != "n" && more != "N" {
		fmt.Print("Enter a number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		if _, err := strconv.Atoi(input); err != nil {
			log.Fatal("An integer is expected")
		}

		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}

		sum += number

		fmt.Print("Add more? [Y/n] ")
		more, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		more = strings.TrimSpace(more)
	}
	fmt.Println("Sum is ", sum)
}
