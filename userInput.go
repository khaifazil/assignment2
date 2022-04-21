package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var userSelection string
var userSelectionValue int

func userIntInput(question string) int {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	userSelectionValue, _ = strconv.Atoi(userSelection)
	return userSelectionValue
}
func userStringInput(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	userSelectionCapital := strings.Title(strings.ToLower(userSelection))
	return userSelectionCapital
}

func userRawStringInput(question string) string {
	fmt.Println(question)
	userSelection = ""
	fmt.Scanln(&userSelection)
	return userSelection
}

func userInputYN(question string) bool {
	for {
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

func backToMain() {
	fmt.Println("\nPress enter to go back to main menu...")
	fmt.Scanln(&userSelection)
	main()
}

func backToAdminMenu() {
	fmt.Println("\nPress enter to go back to main menu...")
	fmt.Scanln(&userSelection)
	adminMenuWrapper()
}

func getSelection(firstSelection int, lastSelection int) (int, error) {
	userSelection := userIntInput("\nChoose an option: ")
	if userSelection < firstSelection || userSelection > lastSelection {
		return userSelection, errors.New("invalid Selection")
	}
	return userSelection, nil
}
