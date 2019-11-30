package DayOne

import "testing"

func TestGetFloorScenarioOne(t *testing.T) {
	firstResult := getFloor("(())")
	if firstResult != 0 {
		t.Errorf("Incorrect, Result was %v", firstResult)
	}

	secondResult := getFloor("()()")
	if secondResult != 0 {
		t.Errorf("Incorrect, Result was %v", secondResult)
	}
}

func TestGetFloorScenarioTwo(t *testing.T) {
	firstResult := getFloor("(((")
	if firstResult != 3 {
		t.Errorf("Incorrect, Result was %v", firstResult)
	}

	secondResult := getFloor("(()(()(")
	if secondResult != 3 {
		t.Errorf("Incorrect, Result was %v", secondResult)
	}
}

func TestGetFloorScenarioThree(t *testing.T) {
	firstResult := getFloor("))(((((")
	if firstResult != 3 {
		t.Errorf("Incorrect, Result was %v", firstResult)
	}
}

func TestGetFloorScenarioFour(t *testing.T) {
	firstResult := getFloor("())")
	if firstResult != -1 {
		t.Errorf("Incorrect, Result was %v", firstResult)
	}

	secondResult := getFloor("))(")
	if secondResult != -1 {
		t.Errorf("Incorrect, Result was %v", secondResult)
	}
}

func TestGetFloorScenarioFive(t *testing.T) {
	firstResult := getFloor(")))")
	if firstResult != -3 {
		t.Errorf("Incorrect, Result was %v", firstResult)
	}

	secondResult := getFloor(")())())")
	if secondResult != -3 {
		t.Errorf("Incorrect, Result was %v", secondResult)
	}
}


