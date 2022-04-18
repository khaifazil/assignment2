package main

import (
	"fmt"
	"strconv"
	"strings"
)

var userSelection string
var userSelectionValue int

func UserIntInput(question string) int {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	userSelectionValue, _ = strconv.Atoi(userSelection)
	return userSelectionValue
}
func UserStringInput(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	userSelectionCapital := strings.Title(strings.ToLower(userSelection))
	return userSelectionCapital
}

func UserInputYN(question string) bool{
	for{
		fmt.Println(question, "(y/n): ")
		fmt.Scanln(&userSelection)
		if userSelection == "y" {
			return true
		} else if userSelection == "n" {
			return false
		} else {
			fmt.Println("Invalid input. Please reply with 'y' or 'n'.")
		}
	}
}
