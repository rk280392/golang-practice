package main

import (
	"fmt"
	"log"

	"github.com/rk280392/golang-practice/my-packages/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)
	fmt.Println(counts)

	for _, line := range lines {
		counts[line]++
	}
	fmt.Println(counts)
	for name, count := range counts {
		fmt.Printf("Votes for %s : %d\n", name, count)
	}
}
