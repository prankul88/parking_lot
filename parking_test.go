package main

import (
	"testing"
)

func TestNewLot(t *testing.T) {
	createParkingLot(4)
	if len(parkingLot) != 4 {
		t.Errorf("Expected capacity to be 4 but got %v", len(parkingLot))
	}

	addCar("DL-4C-NE-8184", "white")
	addCar("DL-4C-NE-8185", "white")
	addCar("DL-4C-NE-8186", "white")

	j := 0
	for i := range parkingLot {
		flag := details{carNo: "", carColor: ""}
		if parkingLot[i] == flag {
			j++
		}
	}

	if j != 1 {
		t.Errorf("Expected empty spaces to be 1 but got %v", j)
	}

	leaveCar(2)

	j = 0
	for i := range parkingLot {
		flag := details{carNo: "", carColor: ""}
		if parkingLot[i] == flag {
			j++
		}
	}

	if j != 2 {
		t.Errorf("Expected empty spaces to be 2 but got %v", j)
	}

}
