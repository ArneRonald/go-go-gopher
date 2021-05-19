package daytwo

import (
	"fmt"
	"github.com/ArneRonald/go-go-gopher/util"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//ExecuteDayTwo ...
func ExecuteDayTwo() {
	dir,_ := os.Getwd()
	result := result{0, 0}
	data := util.ReadFileContents(dir + "/2015/DayTwo/input.txt")
	for _, line := range data {
		result = calculateSquareFeet(line, result)
	}
	fmt.Println("The total square feet of paper:", result.squareFeet)
	fmt.Println("The total feet of Ribbon:", result.ribbonLength)
}

type dimentions struct {
	x int64
	y int64
	z int64
}

type result struct {
	squareFeet   int64
	ribbonLength int64
}

func calculateSquareFeet(input string, res result) result {
	parsedInput := parseStringToDimentions(input)
	var smallest int64 = 0
	dimX := parsedInput.x * parsedInput.y
	dimY := parsedInput.y * parsedInput.z
	dimZ := parsedInput.z * parsedInput.x

	smallest = dimX
	if dimY < smallest {
		smallest = dimY
	}
	if dimZ < smallest {
		smallest = dimZ
	}

	res.squareFeet += ((dimX * 2) + (dimY * 2) + (dimZ * 2) + smallest)
	res.ribbonLength += calculateRibbon(parsedInput)

	return res
}

func calculateRibbon(input dimentions) (feet int64) {
	ints := []int64{}
	ints = append(ints, input.x)
	ints = append(ints, input.y)
	ints = append(ints, input.z)
	sort.Slice(ints, func(i, j int) bool { return ints[i] < ints[j] })

	feet = (ints[0] * 2) + (ints[1] * 2) + (ints[0] * ints[1] * ints[2])

	return
}

func parseStringToDimentions(inputString string) dimentions {
	splitInput := strings.Split(inputString, "x")
	x, errX := strconv.ParseInt(splitInput[0], 10, 64)
	if errX != nil {
		log.Fatal("Bugger", errX)
	}
	y, errY := strconv.ParseInt(splitInput[1], 10, 64)
	if errY != nil {
		log.Fatal("Bugger", errY)
	}
	z, errZ := strconv.ParseInt(splitInput[2], 10, 64)
	if errZ != nil {
		log.Fatal("Bugger", errZ)
	}

	return dimentions{x, y, z}
}
