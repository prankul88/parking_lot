package main

import "fmt"

var parkingLot = make(map[int]details)

type details struct {
	carNo    string
	carColor string
}

func createParkingLot(c int) {
	//parkingLot = make(map[int]string)
	fmt.Println("Created a parking lot with", c, "slots")
	for i := 1; i <= c; i++ {
		parkingLot[i] = details{carNo: "", carColor: ""}
	}
}

func addCar(No string, Color string) {
	temp := 0
	flag := details{carNo: "", carColor: ""}
	for i := 1; i <= len(parkingLot); i++ {

		if parkingLot[i] == flag {

			carDetails := details{
				carNo:    No,
				carColor: Color,
			}
			parkingLot[i] = carDetails
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
	parkingLot[ls] = details{carNo: "", carColor: ""}
	fmt.Println("Slot number", ls, "is empty")
}

//Status will mention
func status() {

	fmt.Println("Slot No.  Registration No	 Colour")
	for no, detail := range parkingLot {
		fmt.Println(no, "       ", detail.carNo, "  ", detail.carColor)
	}
}

func registrationNumbersForCarsWithColour(c string) {
	flag := 0
	for _, detail := range parkingLot {
		if detail.carColor == c {
			fmt.Print(detail.carNo, ",")
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Println("Not Found")
	}
}

func slotNumbersForCarsWithColour(c string) {
	flag := 0
	for no, detail := range parkingLot {
		if detail.carColor == c {
			fmt.Print(no, ",")
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Println("Not Found")
	}
}

func slotNumberForRegistrationNumber(n string) {
	flag := 0
	for no, detail := range parkingLot {
		if detail.carNo == n {
			fmt.Print(no, ",")
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Println("Not Found")
	}
}
