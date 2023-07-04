package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	if (x1 > x2 && v1 >= v2) || (x2 > x1 && v2 >= v1) || (x1 != x2 && v1 == v2) {
		return "NO"
	}

	for {
		if x1 == x2 {
			return "YES"
		}
		if (x2 > x1 && v2 > v1) || (x1 > x2 && v1 > v2) {
			return "NO"
		}
		x1 += v1
		x2 += v2
	}
}

func main() {
	var x1, v1, x2, v2 int32
	fmt.Println("Enter the starting positions and velocities of the kangaroos:")
	fmt.Scan(&x1, &v1, &x2, &v2)
	result := kangaroo(x1, v1, x2, v2)
	fmt.Println(result)
}
