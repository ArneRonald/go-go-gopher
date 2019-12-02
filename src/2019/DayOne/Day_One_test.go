package main

import "testing"

func TestCalculateFuelNeededOne(t *testing.T) {
	var netMass = CalculateFuelNeeded(12)
	if netMass != 2 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}

func TestCalculateFuelNeededTwo(t *testing.T) {
	var netMass = CalculateFuelNeeded(14)
	if netMass != 2 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}

func TestCalculateFuelNeededThree(t *testing.T) {
	var netMass = CalculateFuelNeeded(1969)
	if netMass != 654 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}

func TestCalculateFuelNeededFour(t *testing.T) {
	var netMass = CalculateFuelNeeded(100756)
	if netMass != 33583 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}

func TestHandleFuelOne(t *testing.T) {
	var netMass = HandleFuel(14)
	if netMass != 2 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}

func TestHandleFuelTwo(t *testing.T) {
	var netMass = HandleFuel(1969)
	if netMass != 966 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}

func TestHandleFuelThree(t *testing.T) {
	var netMass = HandleFuel(100756)
	if netMass != 50346 {
		t.Errorf("Incorrect mass calculated: %v", netMass)
	}
}
