package main

import "fmt"

type car struct {
	name     string
	topSpeed float64
}

func nitroBoost(c *car) { // accept pointer as an argument
	c.topSpeed += 50
}

func main() {
	var mustung car
	mustung.name = "Mustung Cobra"
	mustung.topSpeed = 225
	nitroBoost(&mustung) // pass pointer to update the struct, else the original values won't be updated
	fmt.Println(mustung.name)
	fmt.Println(mustung.topSpeed)
}
