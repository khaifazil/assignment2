package main

import (
	"math/rand"
	"time"
)

type bookingInfo struct {
	bookingTime int
	userName    string
	pickup      string
	dropoff     string
	car         string
	contactInfo int
	remarks     string
	bookingId   string
}

var (

	carSelection string
	bookingTime int
	pickUp 	string
	dropOff string
	contactInfo int
	remarks string
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// func makeNewBooking(userName string, carSelection string, bookingTime int, pickUp string, dropOff string, contactInfo int, remarks string) [365][24]bookingInfo {

// }

func makeRandomBookingId(length int) string {

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bookingId := make([]byte, length)
	for i := range bookingId {
		bookingId[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(bookingId)
}
