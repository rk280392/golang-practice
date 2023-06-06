package main

import (
	"fmt"
	"log"

	"github.com/rk280392/golang-practice/my-packages/calender"
)

func main() {
	event := calender.Event{}
	err := event.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetYear(1992)
	if err != nil {
		log.Fatal(err)
	}

	err = event.SetTitle("My Birthday")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(event.Title())
	fmt.Println(event.Day())
	fmt.Println(event.Month())
	fmt.Println(event.Year())

	// also possible to call this way
	fmt.Println(event.Date.Month())
}
