package main

import (
	"errors"
	"fmt"
	"time"
)

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

func modifyBookingWrapper(booking *bookingInfoNode) {
	modifyBookingMenu()
	if userSelection, err := getSelection(1, 8); err != nil {
		fmt.Println(err)
		main()
	} else {
		switch userSelection {
		case 1:
			listCars()
			carSelection := userStringInput("Enter new car selection: ")
			if err := checkCarSelection(carSelection); err != nil {
				fmt.Println(err)
				modifyBookingWrapper(booking)
			}
			if userInputYN("Confirm the change?") {

				newCarArr := getCarArr(carSelection)
				oldCarArr := getCarArr(booking.car)

				t := convertTime(booking.bookingTime)
				d := convertDate(booking.date)

				if newCarArr[d][t] != nil {
					fmt.Println(fmt.Errorf("%v already has a booking at that time slot", carSelection))
					backToMain()
				}
				(*newCarArr)[d][t] = (*oldCarArr)[d][t]
				(*oldCarArr)[d][t] = nil
				booking.car = carSelection

				fmt.Println("----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()
				backToMain()
			} else {
				main()
			}
		case 2:
			userDate := userStringInput("Enter date (E.g. dd/mm/yyyy): ")
			if date, err := time.Parse(timeFormat, userDate); err != nil {
				fmt.Println(err)
				modifyBookingWrapper(booking)
			} else {
				if date.Before(time.Now()) {
					fmt.Println(errors.New("invalid date given has passed"))
					modifyBookingWrapper(booking)
				} else if date.After(time.Now().AddDate(1, 0, 0)) {
					fmt.Println(errors.New("date given is more than one year away"))
					modifyBookingWrapper(booking)
				}
			}
			carArr := getCarArr(booking.car)
			time := convertTime(booking.bookingTime)
			oldDate := convertDate(booking.date)
			newDate := convertDate(userDate)

			if carArr[newDate][time] != nil {
				fmt.Println(fmt.Errorf("%v already has a booking at that time slot", booking.car))
				modifyBookingWrapper(booking)
			}
			if userInputYN("Confirm the change?") {
				(*carArr)[newDate][time] = (*carArr)[oldDate][time]
				(*carArr)[oldDate][time] = nil
				booking.date = userDate
				currentUser.userBookings = sortBookingsByTime(currentUser.userBookings, len(currentUser.userBookings))
				currentUser.userBookings = sortBookingsByDate(currentUser.userBookings, len(currentUser.userBookings))

				fmt.Println("\n----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()
				backToMain()
			} else {
				main()
			}
		case 3:
			bookingTime := userIntInput("Enter booking time in 24HR format(E.g., 1300): ")
			if bookingTime < 0100 || bookingTime > 2400 || bookingTime%100 != 0 {
				fmt.Println(errors.New("invalid time"))
				modifyBookingWrapper(booking)

			}
			carArr := getCarArr(booking.car)
			oldTime := convertTime(booking.bookingTime)
			newTime := convertTime(bookingTime)
			date := convertDate(booking.date)

			if carArr[date][newTime] != nil {
				fmt.Println(fmt.Errorf("%v already has a booking at that time slot", booking.car))
				modifyBookingWrapper(booking)
			}
			if userInputYN("Confirm the change?") {
				(*carArr)[date][newTime] = (*carArr)[date][oldTime]
				(*carArr)[date][oldTime] = nil
				booking.bookingTime = bookingTime

				currentUser.userBookings = sortBookingsByTime(currentUser.userBookings, len(currentUser.userBookings))
				currentUser.userBookings = sortBookingsByDate(currentUser.userBookings, len(currentUser.userBookings))

				fmt.Println("\n----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()

				backToMain()
			} else {
				main()
			}
		case 4:
			newPickUp := userStringInput("Enter pick up address: ")
			if newPickUp == "" {
				fmt.Println(errors.New("pickup address not specified"))
				modifyBookingWrapper(booking)
				
			}
			if userInputYN("Confirm the change?") {
				booking.pickUp = newPickUp

				fmt.Println("\n----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()

				backToMain()
			} else {
				main()
			}
		case 5:
			newDropOff := userStringInput("Enter drop off address: ")
			if newDropOff == "" {
				fmt.Println(errors.New("dropOff address not specified"))
				modifyBookingWrapper(booking)

			}
			if userInputYN("Confirm the change?") {
				booking.dropOff = newDropOff

				fmt.Println("\n----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()

				backToMain()
			} else {
				main()
			}
		case 6:
			newContactInfo := userIntInput("Enter mobile number: ")
			if newContactInfo == 0 || newContactInfo < 10000000 || newContactInfo > 99999999 {
				fmt.Println(errors.New("invalid mobile number"))
				modifyBookingWrapper(booking)
			}
			if userInputYN("Confirm the change?") {
				booking.contactInfo = newContactInfo

				fmt.Println("\n----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()

				backToMain()
			} else {
				main()
			}
		case 7:
			newRemarks := userStringInput("Remarks: ")
			if userInputYN("Confirm the change?") {
				booking.remarks = newRemarks

				fmt.Println("\n----------------------")
				fmt.Println("Here's your booking after changes: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()

				backToMain()
			} else {
				main()
			}
		case 8:
			main()
		}
	}
}
