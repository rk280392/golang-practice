package datafile

import (
	"bufio"
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
