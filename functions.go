package main

import (
	"errors"
	"time"
)

func convertDate(date string) int {
	// make array of months in accumalated days not including current month's days
	// add date given by user.
	var (
		daysInMonths = [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	)

	userDate, _ := time.Parse("02/01/2006", date)
	month := int(userDate.Month()) - 1
	return daysInMonths[month] + userDate.Day()
}

func convertTime(time int) int {
	if time == 2400 {
		time = 0000
	}
	return time / 100
}

func getCarArr(car string) *[365][24]*bookingInfoNode {
	switch car {
	case "Car1":
		return &car1
	case "Car2":
		return &car2
	case "Car3":
		return &car3
	case "Car4":
		return &car4
	default:
		panic(errors.New("invalid car"))
	}
}

// func updateCarArr(ptr *[365][24]*bookingInfoNode, index1 int, index2 int, address *bookingInfoNode) {
// 	ptr[index1][index2] = address
// }
