package main

import (
	"fmt"
	"log"

	"github.com/rk280392/golang-practice/my-packages/geo"
)

func main() {
	location := geo.Landmark{}

	err := location.SetName("my home")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
