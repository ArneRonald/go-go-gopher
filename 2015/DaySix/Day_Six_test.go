package daysix

import (
	"testing"
)

func TestLitLightsOne(t *testing.T) {
	emptyBoard := generateBlankScreen(1000)
	newBoard := handleInput("turn on 0,0 through 999,999", emptyBoard)
	lit := calculateLitLights(newBoard)
	if lit != 1000000 {
		t.Errorf("Was expecting 1000000 but returned %v", lit)
	}
}

func TestLitLightsTest(t *testing.T) {
	emptyBoard := generateBlankScreen(10)
	newBoard := handleInput("turn on 0,0 through 5,5", emptyBoard)
	lit := calculateLitLights(newBoard)
	if lit != 36 {
		t.Errorf("Was expecting 36 but returned %v", lit)
	}
}

func TestLitLightsTwo(t *testing.T) {
	emptyBoard := generateBlankScreen(1000)
	newBoard := handleInput("toggle 0,0 through 999,0", emptyBoard)
	lit := calculateLitLights(newBoard)
	if lit != 1000 {
		t.Errorf("Was expecting 1000 but returned %v", lit)
	}
}

func TestLitLightsThree(t *testing.T) {
	emptyBoard := generateBlankScreen(1000)

	for row := 0; row < len(emptyBoard); row++ {
		for column := 0; column < len(emptyBoard[row]); column++ {
			emptyBoard[row][column] = 1
		}
	}

	newBoard := handleInput("turn off 499,499 through 500,500", emptyBoard)
	lit := calculateLitLights(newBoard)
	if lit != 999996 {
		t.Errorf("Was expecting 4 but returned %v", lit)
	}
}

func TestLitLightsScenarioTwoOne(t *testing.T) {
	emptyBoard := generateBlankScreen(1000)
	newBoard := handleInput("turn on 0,0 through 0,0", emptyBoard)
	lit := calculateBrightness(newBoard)
	if lit != 1 {
		t.Errorf("Was expecting 1 but returned %v", lit)
	}
}
func TestLitLightsScenarioTwoTwo(t *testing.T) {
	emptyBoard := generateBlankScreen(1000)
	newBoard := handleInput("toggle 0,0 through 999,999", emptyBoard)
	lit := calculateBrightness(newBoard)
	if lit != 2000000 {
		t.Errorf("Was expecting 2000000 but returned %v", lit)
	}
}
