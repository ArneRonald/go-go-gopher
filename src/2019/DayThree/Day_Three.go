package daythree

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type coords struct {
	posX int64
	posY int64
}

//ExecuteDayThree ...
func ExecuteDayThree() {
	dir, err := os.Getwd()
	file, err := os.Open(dir + "\\2019\\DayThree\\input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputArrOne []string
	var inputArrTwo []string
	scanner := bufio.NewScanner(file)
	for i := 0; i < i+1; i++ {
		if !scanner.Scan() {
			break
		}
		if i == 0 {
			inputArrOne = strings.Split(scanner.Text(), ",")
		}
		if i == 1 {
			inputArrTwo = strings.Split(scanner.Text(), ",")
		}
	}
	fmt.Printf("Distance: %v", FindIntersection(inputArrOne, inputArrTwo))
}

//FindIntersection ...
func FindIntersection(inputOne, inputTwo []string) int {
	coordsOne := moveCurrentPos(inputOne)
	coordsTwo := moveCurrentPos(inputTwo)
	var steps = fast1(coordsOne, coordsTwo)
	return int(steps)
}

func fast1(wire0 []coords, wire1 []coords) float64 {
	m0 := make(map[coords]bool)
	var intersect []coords
	for _, item := range wire0 {
		m0[item] = true
	}
	for _, item := range wire1 {
		if _, contained := m0[item]; contained {
			intersect = append(intersect, item)
		}
	}
	return findShortest(intersect)
}

func findShortest(interestions []coords) float64 {
	center := coords{0, 0}
	min := math.MaxFloat64
	for _, v := range interestions {
		dist := manhattan(v, center)
		if min > dist {
			min = dist
		}
	}
	return min
}

func manhattan(pos0 coords, pos1 coords) float64 {
	val := math.Abs(float64(pos0.posX)-float64(pos1.posX)) + math.Abs(float64(pos0.posY)-float64(pos1.posY))
	return val
}

func moveCurrentPos(instructions []string) []coords {
	metaCoords := []coords{}
	metaCoord := coords{0, 0}

	for i := 0; i < len(instructions); i++ {
		var direction = strings.Split(instructions[i], "")[0]
		var distance = strings.Replace(instructions[i], direction, "", 1)
		iDistance, _ := strconv.ParseInt(distance, 10, 16)
		for j := 0; j < int(iDistance); j++ {
			metaCoord = generateCoordinate(direction, metaCoord)
			metaCoords = append(metaCoords, metaCoord)
		}
	}

	return metaCoords
}

func generateCoordinate(direction string, previousCoord coords) coords {
	switch direction {
	case "R":
		previousCoord.posX++
		break
	case "L":
		previousCoord.posX--
		break
	case "U":
		previousCoord.posY++
		break
	case "D":
		previousCoord.posY--
		break
	}
	return previousCoord
}

func calcCulateDistance(inputOne []coords, inputTwo []coords) int {
	return 0
}
