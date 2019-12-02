package dayone

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func executeDayOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalFuel int = 0

	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 10)
		if err != nil {
			fmt.Println("Error occured", err)
		}
		totalFuel += HandleFuel(mass)
	}
	fmt.Println(totalFuel)
}

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

func CalculateFuelNeeded(moduleMass float64) int {
	var fuelNeeded = 0
	var tempFuel float64 = 0
	tempFuel = moduleMass / 3
	tempFuel = math.Floor(tempFuel*100) / 100
	fuelNeeded = int(tempFuel) - 2

	return fuelNeeded
}
