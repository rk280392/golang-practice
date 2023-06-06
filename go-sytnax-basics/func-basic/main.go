package main

import "fmt"

func totalWeight(count int, unitWeight float64) float64 {
	calcWeight := float64(count) * unitWeight
	return calcWeight
}

func main() {
	count := 20
	unitWeight := 0.5
	calcWeight := totalWeight(count, unitWeight)
	fmt.Println(count, "cans weight", calcWeight, "kilograms")
}
