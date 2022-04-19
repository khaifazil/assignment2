package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type bookingInfoNode struct {
	car         string
	date        string
	bookingTime int
	userName    string
	pickUp      string
	dropOff     string
	contactInfo int
	remarks     string
	bookingId   string
	prev        *bookingInfoNode
	next        *bookingInfoNode
}

type linkedList struct {
	head *bookingInfoNode
	tail *bookingInfoNode
	size int
}

var bookings = &linkedList{nil, nil, 0}

func init() {
	rand.Seed(time.Now().UnixNano())

	userBst.createUser("khai", "password1")
	userBst.createUser("mary", "password1")
	userBst.createUser("john", "password1")
	userBst.createUser("doug", "password1")

	bookings.makeNewBooking("Car2", "24/11/2022", 200, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "20/04/2022", 300, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/09/2022", 400, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "08/05/2022", 500, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "17/06/2022", 600, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "02/04/2022", 700, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "15/06/2022", 800, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 900, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1000, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1200, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1100, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1300, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1400, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1500, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1600, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1700, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1800, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 1900, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 2000, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 2100, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 2200, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 2300, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "03/06/2022", 2400, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car2", "25/05/1994", 1500, "mary", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
}

func makeRandomBookingId(length int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bookingId := make([]byte, length)
	for i := range bookingId {
		bookingId[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bookingId)
}

func (b *linkedList) makeNewBooking(car string, date string, bookingTime int, userName string, pickUp string, dropOff string, contactInfo int, remarks string) (*bookingInfoNode, error) {
	t := convertTime(bookingTime)
	d := convertDate(date)
	carArr := getCarArr(car)

	if carArr[d][t] != nil {
		return nil, errors.New("there is already a booking at that time and date")
	}

	bookingId := makeRandomBookingId(6)

	newBookingInfoNode := &bookingInfoNode{
		car:         car,
		date:        date,
		bookingTime: bookingTime,
		userName:    userName,
		pickUp:      pickUp,
		dropOff:     dropOff,
		contactInfo: contactInfo,
		remarks:     remarks,
		bookingId:   bookingId,
		next:        nil,
		prev:        nil,
	}
	if b.head == nil {
		b.head = newBookingInfoNode
		b.tail = newBookingInfoNode
	} else {
		//Point tail.next to new
		b.tail.next = newBookingInfoNode
		//Point new.prev to tail
		newBookingInfoNode.prev = b.tail
		//Set tail to new
		b.tail = newBookingInfoNode
	}

	(*carArr)[d][t] = newBookingInfoNode

	userNode, _ := userBst.searchUser(userName)
	userNode.userBookings = append(userNode.userBookings, newBookingInfoNode)
	userNode.userBookings = sortBookingsByDate(userNode.userBookings, len(userNode.userBookings))
	// selectSort(userNode.userBookings, len(userNode.userBookings))

	b.size++
	return newBookingInfoNode, nil
}

// func (b *linkedList) printAllBookings() error {
// 	if b.head == nil {
// 		return errors.New("no bookings")
// 	}
// 	currentNode := b.head
// 	index := 1
// 	fmt.Println("\nBookings:")
// 	fmt.Printf("\nBooking no.%v", index)
// 	fmt.Printf("\nCar: %v", currentNode.car)
// 	fmt.Printf("\nDate: %v", currentNode.date)
// 	fmt.Printf("\nTime: %v", currentNode.bookingTime)
// 	fmt.Printf("\nName: %v", currentNode.userName)
// 	fmt.Printf("\nPickup Address : %v", currentNode.pickUp)
// 	fmt.Printf("\nDropoff Address: %v", currentNode.dropOff)
// 	fmt.Printf("\nContact information: %v", currentNode.contactInfo)
// 	fmt.Printf("\nRemarks: %v", currentNode.remarks)
// 	fmt.Printf("\nBooking ID: %v\n", currentNode.bookingId)
// 	fmt.Println("----------------------------------------------------------------")
// 	for currentNode.next != nil {
// 		currentNode = currentNode.next
// 		index++
// 		fmt.Printf("\nBooking no.%v", index)
// 		fmt.Printf("\nCar: %v", currentNode.car)
// 		fmt.Printf("\nDate: %v", currentNode.date)
// 		fmt.Printf("\nTime: %v", currentNode.bookingTime)
// 		fmt.Printf("\nName: %v", currentNode.userName)
// 		fmt.Printf("\nPickup Address : %v", currentNode.pickUp)
// 		fmt.Printf("\nDropoff Address: %v", currentNode.dropOff)
// 		fmt.Printf("\nContact information: %v", currentNode.contactInfo)
// 		fmt.Printf("\nRemarks: %v", currentNode.remarks)
// 		fmt.Printf("\nBooking ID: %v\n", currentNode.bookingId)
// 		fmt.Println("----------------------------------------------------------------")
// 	}
// 	return nil
// }

func printBookingNode(ptr *bookingInfoNode) error {
	if ptr == nil {
		panic(errors.New("booking not found"))
	}
	fmt.Printf("\nCar: %v", ptr.car)
	fmt.Printf("\nDate: %v", ptr.date)
	fmt.Printf("\nTime: %v", ptr.bookingTime)
	fmt.Printf("\nName: %v", ptr.userName)
	fmt.Printf("\nPickup Address : %v", ptr.pickUp)
	fmt.Printf("\nDropoff Address: %v", ptr.dropOff)
	fmt.Printf("\nContact information: %v", ptr.contactInfo)
	fmt.Printf("\nRemarks: %v", ptr.remarks)
	fmt.Printf("\nBooking ID: %v\n", ptr.bookingId)

	return nil
}
