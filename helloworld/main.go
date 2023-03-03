package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./hello-world <argument>\n")
		os.Exit(1)
	}
	fmt.Printf("Hello world\nArguments: %v\nFirst Argument: %v\nAll Arguments: %v\n", args, args[1], args[1:]) // handling args
}
