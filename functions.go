package main

import (
	"errors"
	"time"
)

var timeFormat = "02/01/2006"

func convertDate(date string) int {
	// make array of months in accumalated days not including current month's days
	// add date given by user.
	var (
		daysInMonths = [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	)

	userDate, _ := time.Parse(timeFormat, date)
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

func selectSort(arr []*bookingInfoNode, n int) {
	for last := n - 1; last >= 1; last-- {
		largest := indexOfLargest(arr, last+1)
		swapPtr(&arr[largest], &arr[last])
	}
}

func indexOfLargest(arr []*bookingInfoNode, n int) int {
	largestIndex := 0
	for i := 1; i < n; i++ {
		checkAtIndex, _ := time.Parse(timeFormat, arr[i].date)
		currentLargest, _ := time.Parse(timeFormat, arr[largestIndex].date)
		if checkAtIndex.After(currentLargest) {
			largestIndex = i
		}
	}
	return largestIndex
}

func swapPtr(x **bookingInfoNode, y **bookingInfoNode) {
	temp := *x
	*x = *y
	*y = temp
}

func getTimeFromParse(d time.Time, _ error) time.Time {
	return d
}

func sortBookingsByDate(arr []*bookingInfoNode, n int) []*bookingInfoNode {
	for i := 1; i < n; i++ {
		data := arr[i]
		last := i
		dataDate, _ := time.Parse(timeFormat, data.date)
		for (last > 0) && (getTimeFromParse(time.Parse(timeFormat, arr[last-1].date)).After(dataDate)) {
			arr[last] = arr[last-1]
			last--
		}

		arr[last] = data
	}
	return arr
}

func sortBookingsByTime(arr []*bookingInfoNode, n int) []*bookingInfoNode {
	for i := 1; i < n; i++ {
		data := arr[i]
		last := i
		
		for (last > 0) && (arr[last-1].bookingTime > data.bookingTime) {
			arr[last] = arr[last-1]
			last--
		}

		arr[last] = data
	}
	return arr
}

// func updateCarArr(ptr *[365][24]*bookingInfoNode, index1 int, index2 int, address *bookingInfoNode) {
// 	ptr[index1][index2] = address
// }

func checkDate(date string) error {
	parsedDate, err := time.Parse(timeFormat, date)
	if err != nil {
		return err
	} else if parsedDate.Before(time.Now()) {
		return errors.New("date given has passed")
	}
	return nil
}

func binarySearch(arr []*bookingInfoNode, target string)int {
	first := 0
	last := len(arr) - 1

	for first <= last {
		mid := (first+last)/2
		if arr[mid].date == target {
			return mid
		} else {
			if target < arr[mid].date {
				last = mid -1
			}else {
				first = mid + 1
			}
		}

	}
	return -1
}

func lookForDuplicates(arr []*bookingInfoNode, n int, target string) []*bookingInfoNode {
	temp := []*bookingInfoNode{}
	for i := n; i >= 0; i-- {
		if target != arr[i].date {
			break
		}else {
			temp = append(temp, arr[i])
		}
	}
	for i := n+1; i < len(arr); i++ {
		if target != arr[i].date {
			break
		}else {
			temp = append(temp, arr[i])
		}
	}
	return temp
}

func searchBookingByDate(arr []*bookingInfoNode, date string) ([]*bookingInfoNode, error){
	n := binarySearch(arr, date)
	if n == -1 {
		return nil, errors.New("no bookings found at that date")
	}
	return lookForDuplicates(arr, n, date), nil
}