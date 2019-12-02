package daytwo

import (
	"fmt"
)

//ExecuteDayTwo ...
func ExecuteDayTwo() {
	var input []int = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 5, 19, 23, 2, 10, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 35, 5, 39, 2, 6, 39, 43, 1, 43, 5, 47, 2, 47, 10, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 1, 71, 5, 75, 1, 13, 75, 79, 1, 6, 79, 83, 2, 83, 13, 87, 1, 87, 6, 91, 1, 10, 91, 95, 1, 95, 9, 99, 2, 99, 13, 103, 1, 103, 6, 107, 2, 107, 6, 111, 1, 111, 2, 115, 1, 115, 13, 0, 99, 2, 0, 14, 0}
	// fmt.Printf("Output: %v\n", strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ConvertOptCode(input))), ","), "[]"))

	fmt.Printf("Calculated: %v\n", findSpesificResult(input, 19690720))
}

//ConvertOptCode ...
func ConvertOptCode(optCode []int) []int {
	var pointer int = 0
	for true {
		if optCode[pointer] == 99 {
			fmt.Printf("Exit Code reached at position: %v\n", pointer)
			break
		}
		firstNumPos := optCode[pointer+1]
		secondNumPos := optCode[pointer+2]
		outputPos := optCode[pointer+3]

		if optCode[pointer] == 1 {
			optCode[outputPos] = add(optCode[firstNumPos], optCode[secondNumPos])
		} else if optCode[pointer] == 2 {
			optCode[outputPos] = mult(optCode[firstNumPos], optCode[secondNumPos])
		}
		pointer += 4
	}
	return optCode
}

func add(val1 int, val2 int) int {
	return val1 + val2
}

func mult(val1 int, val2 int) int {
	return val1 * val2
}

func findSpesificResult(codes []int, result int) int {
	resultVal := 0
	for i := 1; i <= 99; i++ {
		for j := 1; j < 99; j++ {
			var localCopy []int = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 13, 1, 19, 1, 5, 19, 23, 2, 10, 23, 27, 1, 27, 5, 31, 2, 9, 31, 35, 1, 35, 5, 39, 2, 6, 39, 43, 1, 43, 5, 47, 2, 47, 10, 51, 2, 51, 6, 55, 1, 5, 55, 59, 2, 10, 59, 63, 1, 63, 6, 67, 2, 67, 6, 71, 1, 71, 5, 75, 1, 13, 75, 79, 1, 6, 79, 83, 2, 83, 13, 87, 1, 87, 6, 91, 1, 10, 91, 95, 1, 95, 9, 99, 2, 99, 13, 103, 1, 103, 6, 107, 2, 107, 6, 111, 1, 111, 2, 115, 1, 115, 13, 0, 99, 2, 0, 14, 0}
			localCopy[1] = i
			localCopy[2] = j
			resultProgram := ConvertOptCode(localCopy)
			if resultProgram[0] == result {
				resultVal = 100*i + j
				break
			}
		}
	}
	return resultVal
}
