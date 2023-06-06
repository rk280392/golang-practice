package interfaceExample1

import "fmt"

// Declare interface type
type MyInterface interface {

	// A type satisfies the interface if it has all the three methods

	MethodWithoutParameters()
	MethodWithParamater(float64)
	MethodWithReturnVal() string
}

type TestType int // Declare a type and satisfy the interface by creating all the three methods.

// Method 1
func (t TestType) MethodWithoutParameters() {
	fmt.Println("Method without parameter")
}

// Method2

func (t TestType) MethodWithParamater(f float64) {
	fmt.Println("Method with parameter", f)
}

// method3

func (t TestType) MethodWithReturnVal() string {
	return "Method Without Return Value"
}

// Extra method not present in Interface
// A type still satisfies an interface even if it has methods that are not part of interface
func (t TestType) MethodNotInInterface() {
	fmt.Println("Method is not present in Interface")
}

// If a type has al/ the methods declared in an interface then it can be used anywhere that interface is required, with no further decalarations.
