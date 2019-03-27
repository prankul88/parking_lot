package main

import "fmt"

var parkingLot = make(map[int]string)

func createParkingLot(c int) {
	//parkingLot = make(map[int]string)
	fmt.Println("Created a parking lot with", c, "slots")
	for i := 1; i <= c; i++ {
		parkingLot[i] = ""
	}

}

func addCar(carNo string, carColor string) {

	for i := 1; i <= len(parkingLot); i++ {
		if parkingLot[i] == "" {
			parkingLot[i] = carNo + " " + carColor
			fmt.Println("Allocated slot number:", i)
		}
	}
}

//Status will mention
func status(c int) {

	fmt.Println("Capap")
}
