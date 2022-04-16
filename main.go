package main

import (
	"errors"
	"fmt"
)

func showMainMenu() {

	fmt.Println("\nMain Menu")
	fmt.Println("=========")
	fmt.Println("1. View all your bookings")
	fmt.Println("2. Make a new booking")
	fmt.Println("3. Add Items")
	fmt.Println("4. Modify Booking")
	fmt.Println("5. Delete Booking")
	fmt.Println("6. Print Current Data")
	fmt.Println("7. Exit Program")
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Trapped panic:", err)
		}
	}()

	//user interface
	// fmt.Println("Welcome to TRIFECTA LIMO SERVICES")
	// userName:= UserStringInput("Please enter your username: ")
	// UserStringInput("Please enter your password: ")

	// fmt.Printf("Logging in...\n")
	// fmt.Printf("\nWelcome, %v!\n", userName)
	showMainMenu()

	userSelection := UserIntInput("\nWhat would you like to do?")
	if userSelection < 1 || userSelection > 7 {
		panic(errors.New("invalid Selection"))
	}

	switch userSelection {
	case 1:
	case 2:
		// listCars()
		// carSelection := UserStringInput("Enter car selection: ")
		// if err := checkCarSelection(carSelection); err != nil {
		// 	panic(err)
		// }
		// userDate := UserStringInput("Enter date (E.g. dd/mm/yyyy): ")
		// bookingTime := UserIntInput("Enter booking time in 24HR format(E.g., 1300): ")
		// pickUp := UserStringInput("Enter pick up address: ")
		// dropOff := UserStringInput("Enter drop off address: ")
		// contactInfo := UserIntInput("Enter mobile number: ")
		// remarks := UserStringInput("Remarks: ")

		// bookings.makeNewBooking(carSelection, userDate, bookingTime, userName, pickUp, dropOff, contactInfo, remarks)

		bookings.printAllBookings()

	case 3:
	}
}
