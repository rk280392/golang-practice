package main

import (
	"fmt"

	"github.com/rk280392/golang-practice/my-packages/interfaceExample1"
)

func main() {

	// Declare a variable using interface type
	// values of TestType satisfy MyInterface so we can assign this value to a variable with a type MyInterface

	var value interfaceExample1.MyInterface = interfaceExample1.TestType(5)
	value.MethodWithParamater(127.3)
	value.MethodWithoutParameters()
	fmt.Println(value.MethodWithReturnVal())

}
