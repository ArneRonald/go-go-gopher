package dayone

import (
	"fmt"
	"github.com/ArneRonald/go-go-gopher/util"
	"math"
	"os"
	"strconv"
)

//ExecuteDayOne ...
func ExecuteDayOne() {
	dir, _ := os.Getwd()
	data := util.ReadFileContents(dir + "/2019/DayOne/input.txt")
	var totalFuel int = 0

	for _, line := range data {
		mass, err := strconv.ParseFloat(line, 10)
		if err != nil {
			fmt.Println("Error occured", err)
		}
		totalFuel += HandleFuel(mass)
	}
	fmt.Println(totalFuel)
}

//HandleFuel ...
func HandleFuel(fuel float64) int {
	var totalFuel int = 0
	var tempFuel int = int(fuel)

	for tempFuel > 0 {
		tempFuel = CalculateFuelNeeded(float64(tempFuel))
		if tempFuel > 0 {
			totalFuel += tempFuel
		}
	}

	return totalFuel
}

//CalculateFuelNeeded ...
func CalculateFuelNeeded(moduleMass float64) int {
	var fuelNeeded = 0
	var tempFuel float64 = 0
	tempFuel = moduleMass / 3
	tempFuel = math.Floor(tempFuel*100) / 100
	fuelNeeded = int(tempFuel) - 2

	return fuelNeeded
}
