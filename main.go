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
			carNo := subInput[1]
			carColor := subInput[2]
			addCar(carNo, carColor)
		case "status":
			status()
		case "leave":
			leaveSpot := subInput[1]
			ls, _ := strconv.Atoi(leaveSpot)
			leaveCar(ls)
		default:
			fmt.Println("Enter right response")
		}
		fmt.Print("Do you want to continue(yes/no)? ")
		fmt.Scan(&flag)
	}
}
