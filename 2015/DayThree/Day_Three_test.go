package daythree

import "testing"

func TestScenarioOne(t *testing.T) {
	result := calculateHousesVisited(">")
	if result != 2 {
		t.Errorf("Expected 2, found %v", result)
	}
}

func TestScenarioTwo(t *testing.T) {
	result := calculateHousesVisited("^>v<")
	if result != 4 {
		t.Errorf("Expected 4, found %v", result)
	}
}

func TestScenarioThree(t *testing.T) {
	result := calculateHousesVisited("^v^v^v^v^v")
	if result != 2 {
		t.Errorf("Expected 2, found %v", result)
	}
}

func TestScenarioFour(t *testing.T) {
	result := calculateRoboSanta("^v")
	if result != 3 {
		t.Errorf("Expected 3, found %v", result)
	}
}

func TestScenarioFive(t *testing.T) {
	result := calculateRoboSanta("^>v<")
	if result != 3 {
		t.Errorf("Expected 3, found %v", result)
	}
}
func TestScenarioSix(t *testing.T) {
	result := calculateRoboSanta("^v^v^v^v^v")
	if result != 11 {
		t.Errorf("Expected 11, found %v", result)
	}
}
