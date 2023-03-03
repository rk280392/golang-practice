package main

import "fmt"

func double(number *int) *int {
	*number *= 2
	fmt.Println(*number)
	return number
}

func main() {
	myint := 4
	value := double(&myint)
	fmt.Println(value)
	fmt.Println(&value)
	fmt.Println(*value)

}
