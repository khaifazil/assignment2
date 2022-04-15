package main

import (
	"fmt"
)

func showMainMenu(){
		
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
	
	//user interface
	// fmt.Println("Welcome to TRIFECTA LIMO SERVICES")
	// userName:= UserStringInput("Please enter your username: ")
	// UserStringInput("Please enter your password: ")
	
	// fmt.Printf("Logging in...\n")
	// fmt.Printf("\nWelcome, %v!\n", userName)
	showMainMenu()

	userSelection := UserIntInput("\nWhat would you like to do?" )

	switch userSelection {
	case 1:
	case 2:
		userMonth := UserStringInput("Enter month (E.g., Jan) : ")
		userDate := UserIntInput("Enter date: ")
		dateValue, _:= convertDate(userMonth, userDate)
		// carSelection = UserStringInput("Enter car selection: ")
		// bookingTime = UserIntInput("Enter booking time in 24HR format(E.g., 1300): ")
		// pickUp = UserStringInput("Enter pick up address: ")
		// dropOff = UserStringInput("Enter drop off address: ")
		// contactInfo = UserIntInput("Enter mobile number: ")
		// remarks = UserStringInput("Remarks: ")



		// makeNewBooking(userName string, carSelection string, bookingTime int, pickUp string, dropOff string, contactInfo int, remarks string)

	case 3:
	}
}
