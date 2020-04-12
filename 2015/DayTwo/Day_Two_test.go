package daytwo

import "testing"

func TestScenarioOne(t *testing.T) {
	res := result{0, 0}
	var input string = "2x3x4"
	result := calculateSquareFeet(input, res)
	if result.squareFeet != 58 {
		t.Errorf("Expected 58, Found %v", result)
	}
}

func TestScenarioTwo(t *testing.T) {
	res := result{0, 0}
	var input string = "1x1x10"
	result := calculateSquareFeet(input, res)
	if result.squareFeet != 43 {
		t.Errorf("Expected 42, Found %v", result)
	}
}

func TestScenarioThree(t *testing.T) {
	var input dimentions = dimentions{2, 3, 4}
	result := calculateRibbon(input)
	if result != 34 {
		t.Errorf("Expected 34, Found %v", result)
	}
}

func TestScenarioFour(t *testing.T) {
	var input dimentions = dimentions{1, 1, 10}
	result := calculateRibbon(input)
	if result != 14 {
		t.Errorf("Expected 14, Found %v", result)
	}
}
