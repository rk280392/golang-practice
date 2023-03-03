package main

import "fmt"

type Centimeters float64
type Meters float64
type Kilometers float64

// We can define new types using defined types. Each defined type is based on the underlying type.
// Methods are similar to function, we pass a receiver parameter to methods.
// Calling a function only recieves a copy of the value defined in the function, use pointers to change the actual value.
// Use capital first letter to export the function, methods, type, variables, etc.

func (m *Meters) ToCentimeters() Centimeters {
	return Centimeters(*m * 100)
}

func (k *Kilometers) ToCentimeters() Centimeters {
	return Centimeters(*k * 100000)
}

func (c *Centimeters) ToMeters() Meters {
	return Meters(*c * 0.01)
}
func main() {
	fmt.Println("Starting")
	scale := Meters(10)
	fmt.Printf("%0.3f meters equals %0.3f centimeters\n", scale, scale.ToCentimeters())
	scale1 := Kilometers(10)
	fmt.Printf("%0.3f kilometers equal %0.3f centimeters\n", scale1, scale1.ToCentimeters())
	scale2 := Centimeters(100)
	fmt.Printf("%0.3f centimeters equal %0.3f meters\n", scale2, scale2.ToMeters())
}
