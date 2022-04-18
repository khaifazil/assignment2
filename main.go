package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Trapped panic:", err)
		}
	}()

	// user interface
	if currentUser == nil {
		fmt.Println("\nWelcome to TRIFECTA LIMO SERVICES!")
		userName := userRawStringInput("\nPlease enter your username: ")
		password := userRawStringInput("\nPlease enter your password: ")

		if result, err := userBst.validateLogin(userName, password); err != nil {
			fmt.Println(err)
			main()
		} else {
			currentUser = result
		}

		fmt.Printf("Logging in...\n")
		fmt.Printf("\nWelcome, %v!\n", userName)
	}
	mainMenu()

	if userSelection, err := getSelection(1, 7); err != nil {
		fmt.Println(err)
		mainMenu()
	} else {
		switch userSelection {
		case 1:
			bookingsMenu()
			if userSelection, err := getSelection(1, 3); err != nil {
				fmt.Println(err)
				main() //todo: implement for loop
			} else {
				switch userSelection {
				case 1:
					bookingSlice := currentUser.userBookings
					if len(bookingSlice) == 0 {
						fmt.Println()
						fmt.Println(errors.New("user has no bookings"))
						backToMain()
					}
					for i, booking := range bookingSlice {
						fmt.Printf("\nBooking no.: %v", i+1)
						fmt.Println()
						printBookingNode(booking)
						fmt.Println()
						fmt.Println("----------------------")
					}
					backToMain()
				case 2:
					//print by date
				case 3:
					main()
				}
			}
		case 2:
			listCars()
			carSelection := userStringInput("Enter car selection: ")
			if err := checkCarSelection(carSelection); err != nil {
				panic(err)
			}
			userDate := userStringInput("Enter date (E.g. dd/mm/yyyy): ")
			if date, err := time.Parse("02/01/2006", userDate); err != nil {
				panic(err)
			} else {
				if date.Before(time.Now()) {
					panic(errors.New("date given has passed"))
				} else if date.After(time.Now().AddDate(1, 0, 0)) {
					panic(errors.New("date given is more than one year away"))
				}
			}
			bookingTime := userIntInput("Enter booking time in 24HR format(E.g., 1300): ")
			if bookingTime < 0100 || bookingTime > 2400 || bookingTime%100 != 0 {
				panic(errors.New("invalid time"))
			}
			pickUp := userStringInput("Enter pick up address: ")
			if pickUp == "" {
				panic(errors.New("pickup address not specified"))
			}
			dropOff := userStringInput("Enter drop off address: ")
			if dropOff == "" {
				panic(errors.New("dropOff address not specified"))
			}
			contactInfo := userIntInput("Enter mobile number: ")
			if contactInfo == 0 || contactInfo < 10000000 || contactInfo > 99999999 {
				panic(errors.New("invalid mobile number"))
			}
			remarks := userStringInput("Remarks: ")

			userName := currentUser.userName

			if booking, err := bookings.makeNewBooking(carSelection, userDate, bookingTime, userName, pickUp, dropOff, contactInfo, remarks); err != nil {
				panic(err)
			} else {
				fmt.Println("Booking has been made!")
				fmt.Println("----------------------")
				printBookingNode(booking)
			}

		case 3:
		}
	}
}
