package main

import (
	"errors"
	"fmt"
	"time"
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
		
		// fmt.Println(car2[1])
		// user interface
		fmt.Println("Welcome to TRIFECTA LIMO SERVICES")
		userName := UserStringInput("Please enter your username: ")
		// password:= UserStringInput("Please enter your password: ")
		
		fmt.Printf("Logging in...\n")
	fmt.Printf("\nWelcome, %v!\n", userName)
	showMainMenu()

	userSelection := UserIntInput("\nWhat would you like to do?")
	if userSelection < 1 || userSelection > 7 {
		panic(errors.New("invalid Selection"))
	}

	switch userSelection {
	case 1:
	case 2:
		listCars()
		carSelection := UserStringInput("Enter car selection: ")
		if err := checkCarSelection(carSelection); err != nil {
			panic(err)
		}
		userDate := UserStringInput("Enter date (E.g. dd/mm/yyyy): ")
		if date, err := time.Parse("02/01/2006", userDate); err != nil {
			panic(err)
		} else {
			if date.Before(time.Now()) {
				panic(errors.New("date given has passed"))
			} else if date.After(time.Now().AddDate(1, 0, 0)) {
				panic(errors.New("date given is more than one year away"))
			}
		}
		bookingTime := UserIntInput("Enter booking time in 24HR format(E.g., 1300): ")
		if bookingTime < 0100 || bookingTime > 2400 || bookingTime%100 != 0 {
			panic(errors.New("invalid time"))
		}
		pickUp := UserStringInput("Enter pick up address: ")
		if pickUp == "" {
			panic(errors.New("pickup address not specified"))
		}
		dropOff := UserStringInput("Enter drop off address: ")
		if dropOff == "" {
			panic(errors.New("dropOff address not specified"))
		}
		contactInfo := UserIntInput("Enter mobile number: ")
		if contactInfo == 0 || contactInfo < 10000000 || contactInfo > 99999999 {
			panic(errors.New("invalid mobile number"))
		}
		remarks := UserStringInput("Remarks: ")

		if err:= bookings.makeNewBooking(carSelection, userDate, bookingTime, userName, pickUp, dropOff, contactInfo, remarks); err != nil {
			panic(err)
		}

		bookings.printAllBookings()

	case 3:
	}
}
