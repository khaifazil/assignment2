package main

import (
	"errors"
	"fmt"
)

var carsList = []string{"Car1", "Car2", "Car3", "Car4"}

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
