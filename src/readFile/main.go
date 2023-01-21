package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloat(filename string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(filename)
	if err != nil {
		return numbers, err
	}

	scan := bufio.NewScanner(file)
	i := 0
	for scan.Scan() {
		numbers[i], err = strconv.ParseFloat(scan.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}
	err = file.Close()
	if err != nil {
		return numbers, err

	}
	if scan.Err() != nil {
		return numbers, err
	}
	return numbers, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		os.Exit(1)
	}
	fmt.Println(args[1])

	numbers, err := getFloat(args[1])
	if err != nil {
		log.Fatal(err)
	}

	var sum float64 = 0
	for _, numbers := range numbers {
		sum += numbers
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
