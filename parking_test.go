package main

import (
	"testing"
)

func TestNewLot(t *testing.T) {
	createParkingLot(4)
	if len(parkingLot) != 4 {
		t.Errorf("Expected capacity to be 4 but got %v", len(parkingLot))
	}
}
