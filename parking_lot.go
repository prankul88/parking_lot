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

	temp := 0
	for i := 1; i <= len(parkingLot); i++ {
		if parkingLot[i] == "" {
			parkingLot[i] = carNo + "		" + carColor
			fmt.Println("Allocated slot number:", i)
			temp = 1
			break
		}
	}
	if temp == 0 {
		fmt.Println("Sorry,Parking lot is Full")
	}

}

func leaveCar(ls int) {
	parkingLot[ls] = ""
}

//Status will mention
func status() {

	fmt.Println("Slot No.  Registration No	    Colour")
	for no, details := range parkingLot {
		fmt.Println(no, "          ", details)
	}
}
