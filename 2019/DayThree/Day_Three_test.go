package daythree

import (
	"testing"
)

func TestCalculateDistanceExample(t *testing.T) {
	var testInputOne []string = []string{"R8", "U5", "L5", "D3"}
	var testInputTwo []string = []string{"U7", "R6", "D4", "L4"}

	distance := FindIntersection(testInputOne, testInputTwo)
	if distance != 6 {
		t.Errorf("Incorrect! Expected: 6, Actual result: %v\n", distance)
	}
}

func TestCalculateDistanceOne(t *testing.T) {
	var testInputOne []string = []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"}
	var testInputTwo []string = []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"}

	distance := FindIntersection(testInputOne, testInputTwo)
	if distance != 159 {
		t.Errorf("Incorrect! Expected: 159, Actual result: %v\n", distance)
	}
}

func TestCalculateDistanceTwo(t *testing.T) {
	var testInputOne []string = []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"}
	var testInputTwo []string = []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"}

	distance := FindIntersection(testInputOne, testInputTwo)
	if distance != 135 {
		t.Errorf("Incorrect! Expected: 135, Actual result: %v\n", distance)
	}
}
