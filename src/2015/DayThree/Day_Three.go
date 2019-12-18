package dayone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//ExecuteDayThree ...
func ExecuteDayThree() {
	dir, err := os.Getwd()
	file, err := os.Open(dir + "\\2015\\DayThree\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf("Houses visited: %v\n", calculateHousesVisited(scanner.Text()))
		fmt.Printf("Houses visited with RoboSanta: %v\n", calculateRoboSanta(scanner.Text()))
	}
}

type coords struct {
	x int
	y int
}

func calculateRoboSanta(directions string) int {
	splitInput := strings.Split(directions, "")
	startingPos := coords{0, 0}

	santaCoords := []coords{startingPos}
	roboSantaCoords := []coords{startingPos}

	for i, s := range splitInput {
		if i%2 == 0 {
			currentPos := moveHouse(santaCoords[len(santaCoords)-1], s)
			santaCoords = append(santaCoords, currentPos)
		} else {
			currentPos := moveHouse(roboSantaCoords[len(roboSantaCoords)-1], s)
			roboSantaCoords = append(roboSantaCoords, currentPos)
		}
	}

	newSet := unique(append(roboSantaCoords, santaCoords...))

	return len(newSet)
}

func calculateHousesVisited(directions string) int {
	splitInput := strings.Split(directions, "")
	currentPos := coords{0, 0}
	houseCoords := []coords{currentPos}
	for _, s := range splitInput {
		currentPos = moveHouse(currentPos, s)
		houseCoords = append(houseCoords, currentPos)
	}

	uniqueCoords := unique(houseCoords)

	return len(uniqueCoords)
}

func moveHouse(currentPos coords, direction string) coords {
	switch direction {
	case "^":
		currentPos.y++
		break
	case "v":
		currentPos.y--
		break
	case ">":
		currentPos.x++
		break
	case "<":
		currentPos.x--
		break
	}
	return currentPos
}

func unique(intSlice []coords) (list []coords) {
	keys := make(map[coords]bool)
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return
}
