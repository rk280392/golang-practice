package main

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {
	
	inputTime := s

	// this needs to have 0 in start of each hour,min,sec as in 12 hour format 3:4:5 is also valid. Valid examples "3:04PM","03:04PM","3:04 PM","03:04 PM"
	// invalid examples 07:15:05. Adding 1 forces for a 2 digit number.
	layout := "03:04:05PM" 

	parsedTime, err := time.Parse(layout, inputTime)
	if err != nil {
		return "Failed to parse time:"
	}

	// Check if it's midnight (12:00 AM) and convert to military time
	militaryTime := parsedTime.Format("15:04:05")
	return militaryTime
}

func main() {
	timeStr := "07:00:00PM"
	fmt.Println(timeConversion(timeStr))

}
