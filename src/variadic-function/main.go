// variadic function allow to pass n number of arguments to a function. Argument is defind by using ... before the type of last argument. Example:

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func findMax(numbers ...int) int {
	max := math.Inf(-1) // get the least number
	for _, value := range numbers {
		if float64(value) > max {
			max = float64(value)
		}
	}
	return int(max)
}

func inRangeMax(min float64, max float64, numbers ...float64) []float64 {
	// returns numbers between min and max. Discards all other numbers passed to the function except the number between min and max

	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}

	return result
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func main() {
	max := findMax(12, 13, 14, -1, -12, 99)
	fmt.Printf("Max Value is %d\n", max)
	fmt.Println(inRangeMax(5, 15, -1, -2, 6, 7, 8, 9, 16, 19))

	args := os.Args[1:]
	var numbers []float64
	for _, argument := range args {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	fmt.Printf("Average : %0.2f\n", average(numbers...))

}
