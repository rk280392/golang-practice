package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	for i := 0; i < len(grades); i++ {
		if grades[i] > 34 {
			nextNumber := (grades[i]/5 + 1) * 5
			if (nextNumber - grades[i]) < 3 {
				grades[i] = nextNumber
			}
		}
	}
	return grades

}

func main() {
	grades := []int32{86,30,0,16,51,53,42,48,22,69,12,27,34,24,95,16,32,22,52,56,71,95}
	updatedGrades := gradingStudents(grades)
	fmt.Println(updatedGrades)
}
