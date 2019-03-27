package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("This in an interactive shell.")
	flag := "yes"

	for flag == "yes" {

		fmt.Print("\n\nEnter the input (only strings valid) :")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()

		subInput := strings.Split(input, " ")

		switch subInput[0] {
		case "create_parking_lot":
			capacity := subInput[1]
			c, _ := strconv.Atoi(capacity)
			createParkingLot(c)
		case "park":
			No := subInput[1]
			Color := subInput[2]
			addCar(No, Color)
		case "status":
			status()
		case "leave":
			leaveSpot := subInput[1]
			ls, _ := strconv.Atoi(leaveSpot)
			leaveCar(ls)
		case "registration_numbers_for_cars_with_colour":
			Color := subInput[1]
			registrationNumbersForCarsWithColour(Color)
		case "slot_numbers_for_cars_with_colour":
			Color := subInput[1]
			slotNumbersForCarsWithColour(Color)
		case "slot_number_for_registration_number":
			No := subInput[1]
			slotNumberForRegistrationNumber(No)
		default:
			fmt.Println("Enter right response")
		}
		fmt.Print("Do you want to continue(yes/no)? ")
		fmt.Scan(&flag)
	}
}
