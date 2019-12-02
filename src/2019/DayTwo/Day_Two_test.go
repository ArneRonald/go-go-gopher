package daytwo

import (
	"fmt"
	"strings"
	"testing"
)

func TestOptCodeOne(t *testing.T) {
	var input []int = []int{1, 0, 0, 0, 99}
	outPut := ConvertOptCode(input)
	var expectedOutput string = "[2,0,0,0,99]"
	var stringOutput = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(outPut)), ","), "")

	if stringOutput != expectedOutput {
		t.Errorf("Failed to convert op code, Actual result was: %v. Expected result was %v\n", stringOutput, expectedOutput)
	}
}

func TestOptCodeTwo(t *testing.T) {
	var input []int = []int{2, 3, 0, 3, 99}
	outPut := ConvertOptCode(input)
	var expectedOutput string = "[2,3,0,6,99]"
	var stringOutput = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(outPut)), ","), "")

	if stringOutput != expectedOutput {
		t.Errorf("Failed to convert op code, Actual result was: %v. Expected result was %v\n", stringOutput, expectedOutput)
	}
}

func TestOptCodeThree(t *testing.T) {
	var input []int = []int{2, 4, 4, 5, 99, 0}
	outPut := ConvertOptCode(input)
	var expectedOutput string = "[2,4,4,5,99,9801]"
	var stringOutput = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(outPut)), ","), "")

	if stringOutput != expectedOutput {
		t.Errorf("Failed to convert op code, Actual result was: %v. Expected result was %v\n", stringOutput, expectedOutput)
	}
}

func TestOptCodeFour(t *testing.T) {
	var input []int = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	outPut := ConvertOptCode(input)
	var expectedOutput string = "[30,1,1,4,2,5,6,0,99]"
	var stringOutput = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(outPut)), ","), "")

	if stringOutput != expectedOutput {
		t.Errorf("Failed to convert op code, Actual result was: %v. Expected result was %v\n", stringOutput, expectedOutput)
	}
}
