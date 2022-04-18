package main

import (
	"errors"
	"fmt"
)

var carsList = []string{"Car1", "Car2", "Car3", "Car4"}

var car1 [365][24]*bookingInfoNode
var car2 [365][24]*bookingInfoNode
var car3 [365][24]*bookingInfoNode
var car4 [365][24]*bookingInfoNode

func listCars() {
	fmt.Println("\nLimo list")
	fmt.Println("=========")
	
	for _, car := range carsList {
		fmt.Printf("- %v\n", car)
	}
}

func checkCarSelection(car string) error {
	for _, c := range carsList {
		if c == car {
			return nil
		}
	}
	return errors.New("car is not in selection")
}
