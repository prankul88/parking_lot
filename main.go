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

	if flag == "yes" {

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
		case "exit":
			flag = "no"
			break
		default:
			fmt.Print("Enter right response")
		}
	}
}
