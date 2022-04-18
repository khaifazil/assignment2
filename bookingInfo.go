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


	bookings.makeNewBooking("Car2", "03/06/2022", 0100, "khai", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car3", "25/05/1994", 1400, "john", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
	bookings.makeNewBooking("Car1", "25/05/1994", 1500, "mary", "dfasdfas", "sadfsdf", 98196006, "kdslfj")
}

func makeRandomBookingId(length int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bookingId := make([]byte, length)
	for i := range bookingId {
		bookingId[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bookingId)
}

func (b *linkedList) makeNewBooking(car string, date string, bookingTime int, userName string, pickUp string, dropOff string, contactInfo int, remarks string) error {
	t, _ := convertTime(bookingTime)
	d := convertDate(date)
	carArr := getCarArr(car)

	if carArr[d][t] != nil {
		return errors.New("there is already a booking at that time and date")
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
	// fmt.Println(carArr[d][t])
	b.size++
	return nil
}

func (b *linkedList) printAllBookings() error {
	if b.head == nil {
		return errors.New("no bookings")
	}
	currentNode := b.head
	index := 1
	fmt.Println("\nBookings:")
	fmt.Printf("\nBooking no.%v", index)
	fmt.Printf("\nCar: %v", currentNode.car)
	fmt.Printf("\nDate: %v", currentNode.date)
	fmt.Printf("\nTime: %v", currentNode.bookingTime)
	fmt.Printf("\nName: %v", currentNode.userName)
	fmt.Printf("\nPickup Address : %v", currentNode.pickUp)
	fmt.Printf("\nDropoff Address: %v", currentNode.dropOff)
	fmt.Printf("\nContact information: %v", currentNode.contactInfo)
	fmt.Printf("\nRemarks: %v", currentNode.remarks)
	fmt.Printf("\nBooking ID: %v\n", currentNode.bookingId)
	fmt.Println("----------------------------------------------------------------")
	for currentNode.next != nil {
		currentNode = currentNode.next
		index++
		fmt.Printf("\nBooking no.%v", index)
		fmt.Printf("\nCar: %v", currentNode.car)
		fmt.Printf("\nDate: %v", currentNode.date)
		fmt.Printf("\nTime: %v", currentNode.bookingTime)
		fmt.Printf("\nName: %v", currentNode.userName)
		fmt.Printf("\nPickup Address : %v", currentNode.pickUp)
		fmt.Printf("\nDropoff Address: %v", currentNode.dropOff)
		fmt.Printf("\nContact information: %v", currentNode.contactInfo)
		fmt.Printf("\nRemarks: %v", currentNode.remarks)
		fmt.Printf("\nBooking ID: %v\n", currentNode.bookingId)
		fmt.Println("----------------------------------------------------------------")
	}
	return nil
}
