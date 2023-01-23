package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getFloat(filename string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		number, err := strconv.ParseFloat(scan.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	if scan.Err() != nil {
		return nil, err
	}

	return numbers, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("No arguments passed")
		os.Exit(1)
	}
	numbers, err := getFloat(args[1])
	if err != nil {
		log.Fatal(err)
	}
	var sum float64
	for _, number := range numbers {
		sum += number
	}
	noOfSamples := float64(len(numbers))
	fmt.Printf("Average : %0.2f\n", sum/noOfSamples)

}
