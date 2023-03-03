package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rk280392/golang-practice/my-packages/datafile"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("filename is required")
		os.Exit(1)
	}
	lines, err := datafile.GetStrings(args[1])
	if err != nil {
		log.Fatal(err)
	}

	// counting number of times each name appears starts here/

	var names []string
	var counts []int

	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}
		if matched == false {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s : %d\n", name, counts[i])
	}
}
