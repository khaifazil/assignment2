package main

import "fmt"

func mainMenu() {

	fmt.Println("\nMain Menu")
	fmt.Println("=========")
	fmt.Println("1. View your bookings")
	fmt.Println("2. Make a new booking")
	fmt.Println("3. Add Items")
	fmt.Println("4. Modify Booking")
	fmt.Println("5. Delete Booking")
	fmt.Println("6. Print Current Data")
	fmt.Println("7. Exit Program")
}

func bookingsMenu() {
	fmt.Println("\nBookings Menu")
	fmt.Println("=============")
	fmt.Println("1. View all bookings")
	fmt.Println("2. View bookings by date")
	fmt.Println("3. Back to main menu")
}
