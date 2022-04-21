package main

import (
	"fmt"
)

func loginMenu() {
	fmt.Println("\nLogin Menu")
	fmt.Println("=========")
	fmt.Println("1. Login")
	fmt.Println("2. Signup")
	fmt.Println("3. Admin Login")
	fmt.Println("4. Exit program")
}

func mainMenu() {

	fmt.Println("\nMain Menu")
	fmt.Println("=========")
	fmt.Println("1. View your bookings")
	fmt.Println("2. Make a new booking")
	fmt.Println("3. Modify Booking")
	fmt.Println("4. Delete Booking")
	fmt.Println("5. Exit Program")
}

func listCars() {
	fmt.Println("\nLimo list")
	fmt.Println("=========")

	for _, car := range carsList {
		fmt.Printf("- %v\n", car)
	}
}

func bookingsMenu() {
	fmt.Println("\nBookings Menu")
	fmt.Println("=============")
	fmt.Println("1. View all bookings")
	fmt.Println("2. View bookings by date")
	fmt.Println("3. View booking from bookingID")
	fmt.Println("4. Back to main menu")
}

func modifyBookingMenu() {
	fmt.Println("\nWhat would you like to change?")
	fmt.Println("==============================")
	fmt.Println("1. Car")
	fmt.Println("2. Date")
	fmt.Println("3. Booking Time")
	fmt.Println("4. Pickup address")
	fmt.Println("5. Drop off address")
	fmt.Println("6. Contact Information")
	fmt.Println("7. Remarks")
	fmt.Println("8. Back to main menu")
}

func adminMenu() {
	fmt.Println("\nAdmin Menu")
	fmt.Println("=============")
	fmt.Println("1. View all bookings")
	fmt.Println("2. View All users")
	fmt.Println("3. Delete users")
	fmt.Println("4. Back to main menu")
}
