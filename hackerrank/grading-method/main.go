package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i := 0; i < len(grades); i++ {
		if grades[i] > 33 {
			nextNumber := (grades[i]/5 + 1) * 5
			if (nextNumber - grades[i]) < 3 {
				grades[i] = nextNumber
			}
		}
	}
	return grades

}

func main() {
	grades := []int32{73, 67, 38, 45, 29, 91}
	updatedGrades := gradingStudents(grades)
	fmt.Println(updatedGrades)
}
