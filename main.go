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

	if userSelection, err := getSelection(1, 5); err != nil {
		fmt.Println(err)
		mainMenu()
	} else {
		switch userSelection {
		case 1:
			bookingsMenu()
			if userSelection, err := getSelection(1, 4); err != nil {
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
					userDate := userStringInput("Enter date (E.g. dd/mm/yyyy): ")
					if err = checkDate(userDate); err != nil {
						panic(err)
					}
					bookingsByDate, err := searchBookingByDate(currentUser.userBookings, userDate)
					if err != nil {
						panic(err)
					}
					bookingsByDate = sortBookingsByTime(bookingsByDate, len(bookingsByDate))
					for i, booking := range bookingsByDate {
						fmt.Printf("\nBooking no.: %v\n", i+1)
						fmt.Println()
						printBookingNode(booking)
						fmt.Println()
						fmt.Println("----------------------")
					}
					backToMain()
				case 3:
					//print by booking number feature
					// collect user booking number
					userBookingId := userRawStringInput("Enter booking ID")
					//retrive from current user array the booking
					if booking, _, err := recursiveSeqSearchId(len(currentUser.userBookings), 0, currentUser.userBookings, userBookingId); err != nil {
						panic(err)
					} else {
						//print booking
						fmt.Println("----------------------")
						fmt.Println("Here's your booking: ")
						printBookingNode(booking)
						fmt.Println()
						fmt.Println("----------------------")
					}
					backToMain()
				case 4:
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
			if date, err := time.Parse(timeFormat, userDate); err != nil {
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
		case 3: //modify bookings
			fmt.Println("\nModifying Booking")
			fmt.Println("=================")
			userBookingId := userRawStringInput("Enter booking ID: ")
			if booking, _, err := recursiveSeqSearchId(len(currentUser.userBookings), 0, currentUser.userBookings, userBookingId); err != nil {
				panic(err)
			} else {
				fmt.Println("----------------------")
				fmt.Println("Here's your booking: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()
				modifyBookingMenu()
				if userSelection, error := getSelection(1, 8); error != nil {
					fmt.Println(err)
					main()
				} else {
					switch userSelection {
					case 1: //car
						listCars()
						carSelection := userStringInput("Enter new car selection: ")
						if err := checkCarSelection(carSelection); err != nil {
							panic(err)
						}
						if userInputYN("Confirm the change?"){
							//get new & old car from user userInput
							newCarArr := getCarArr(carSelection)
							oldCarArr := getCarArr(booking.car)
							// convert booking time and date
							t := convertTime(booking.bookingTime)
							d := convertDate(booking.date)
							// check if new car slot is full
							if newCarArr[d][t] != nil {
								fmt.Println(fmt.Errorf("%v already has a booking at that time slot", carSelection))
								backToMain()
							}
							// transfer old booking to new arraycar
							(*newCarArr)[d][t] = (*oldCarArr)[d][t]
							// delete old booking
							(*oldCarArr)[d][t] = nil
							// change car at node
							booking.car = carSelection
							// print changed booking
							fmt.Println("----------------------")
							fmt.Println("Here's your booking after changes: ")
							printBookingNode(booking)
							fmt.Println()
							fmt.Println("----------------------")
							fmt.Println()
							backToMain()
						}else {
							main()
						}
					case 2: //date
						userDate := userStringInput("Enter date (E.g. dd/mm/yyyy): ")
						if date, err := time.Parse(timeFormat, userDate); err != nil {
							panic(err)
						} else {
							if date.Before(time.Now()) {
								panic(errors.New("date given has passed"))
							} else if date.After(time.Now().AddDate(1, 0, 0)) {
								panic(errors.New("date given is more than one year away"))
							}
						}
						carArr := getCarArr(booking.car)
						time := convertTime(booking.bookingTime)
						oldDate := convertDate(booking.date)
						newDate := convertDate(userDate)

						if carArr[newDate][time] != nil {
							fmt.Println(fmt.Errorf("%v already has a booking at that time slot", booking.car))
								backToMain()
						}
						if userInputYN("Confirm the change?"){
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
						}else{
							main()
						}
					case 3: // booking time
						bookingTime := userIntInput("Enter booking time in 24HR format(E.g., 1300): ")
						if bookingTime < 0100 || bookingTime > 2400 || bookingTime%100 != 0 {
							panic(errors.New("invalid time"))
						}
						carArr := getCarArr(booking.car)
						oldTime := convertTime(booking.bookingTime)
						newTime := convertTime(bookingTime)
						date := convertDate(booking.date)

						if carArr[date][newTime] != nil {
							fmt.Println(fmt.Errorf("%v already has a booking at that time slot", booking.car))
								backToMain()
						}
						if userInputYN("Confirm the change?"){
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
						}else{
							main()
						}
					case 4: //Pickup address
						pickUp := userStringInput("Enter pick up address: ")
						if pickUp == "" {
							panic(errors.New("pickup address not specified"))
						}
					case 5: //DropOff
						dropOff := userStringInput("Enter drop off address: ")
						if dropOff == "" {
							panic(errors.New("dropOff address not specified"))
						}
					case 6: // contact
						contactInfo := userIntInput("Enter mobile number: ")
						if contactInfo == 0 || contactInfo < 10000000 || contactInfo > 99999999 {
							panic(errors.New("invalid mobile number"))
						}
					case 7: // Remarks
						// remarks := userStringInput("Remarks: ")
					case 8: // back to main
						main()
					}
				}
			}
		case 4: //delete bookings
			userBookingId := userRawStringInput("Enter booking ID")
			if booking, index, err := recursiveSeqSearchId(len(currentUser.userBookings), 0, currentUser.userBookings, userBookingId); err != nil {
				panic(err)
			} else {

				fmt.Println("----------------------")
				fmt.Println("Here's your booking: ")
				printBookingNode(booking)
				fmt.Println()
				fmt.Println("----------------------")
				fmt.Println()
				if userInputYN("Delete booking?") {
					err := bookings.deleteBooking(booking, index)
					if err != nil {
						panic(err)
					}
					fmt.Println("Booking has been deleted!")
					backToMain()
				} else {
					main()
				}
			}
		case 5:
			//exit program
			if userInputYN("Are you sure you want to exit?") {
				fmt.Println("Goodbye!")
			} else {
				main()
			}
		}
	}
}
