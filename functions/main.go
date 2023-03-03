package main

import (
	"fmt"
	"log"
)

func paintNeeded(height float64, width float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}

	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := height * width
	return area / 10.0, nil

}
func main() {

	amount, err := paintNeeded(4, 2)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(amount)
	}
}
