package daysix

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//ExecuteDaySix ...
func ExecuteDaySix() {
	dir, err := os.Getwd()
	file, err := os.Open(dir + "/2015/DaySix/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	decorations := generateBlankScreen(1000)

	for scanner.Scan() {
		decorations = handleInput(scanner.Text(), decorations)
	}
	// litAmount := calculateLitLights(decorations)
	// fmt.Println("The amount of lit lights are :", litAmount)
	brightness := calculateBrightness(decorations)
	fmt.Println("The brightness of the decorations is", brightness)
}

type coords struct {
	x int
	y int
}

func handleInput(input string, decorations [][]int) [][]int {
	input = strings.ReplaceAll(input, " through ", " ")
	if strings.Contains(input, "turn on") {
		input = strings.ReplaceAll(input, "turn on ", "")
		startCoord, endCoord := buildCoordinate(input)
		decorations = flipSwitches(decorations, "on", startCoord, endCoord)
	}

	if strings.Contains(input, "turn off") {
		input = strings.ReplaceAll(input, "turn off ", "")
		startCoord, endCoord := buildCoordinate(input)
		decorations = flipSwitches(decorations, "off", startCoord, endCoord)
	}

	if strings.Contains(input, "toggle") {
		input = strings.ReplaceAll(input, "toggle ", "")
		startCoord, endCoord := buildCoordinate(input)
		decorations = flipSwitches(decorations, "toggle", startCoord, endCoord)
	}

	return decorations
}

func flipSwitches(decorations [][]int, function string, startCoord coords, endCoord coords) [][]int {
	switch function {
	case "on":
		for row := startCoord.x; row < endCoord.x; row++ {
			for column := startCoord.y; column < endCoord.y; column++ {
				decorations[row][column]++
			}
		}
		break
	case "off":
		for row := startCoord.x; row < endCoord.x; row++ {
			for column := startCoord.y; column < endCoord.y; column++ {
				if decorations[row][column] > 0 {
					decorations[row][column]--
				}
			}
		}
		break
	case "toggle":
		for row := startCoord.x; row < endCoord.x; row++ {
			for column := startCoord.y; column < endCoord.y; column++ {
				// 	switch decorations[row][column] {
				// 	case 0:
				// 		decorations[row][column] = 1
				// 		break
				// 	case 1:
				// 		decorations[row][column] = 0
				// 		break
				// 	}
				decorations[row][column] = decorations[row][column] + 2
			}

		}
		break
	}
	return decorations
}

func buildCoordinate(input string) (startCoord coords, endCoord coords) {
	coordinates := strings.Split(input, " ")

	xStrings := strings.Split(coordinates[0], ",")
	startX, _ := strconv.Atoi(xStrings[0])
	startY, _ := strconv.Atoi(xStrings[1])

	startCoord = coords{x: startX, y: startY}

	yStrings := strings.Split(coordinates[1], ",")
	endX, _ := strconv.Atoi(yStrings[0])
	endY, _ := strconv.Atoi(yStrings[1])

	endCoord = coords{x: endX + 1, y: endY + 1}
	return
}
func calculateBrightness(decorations [][]int) (brightness int) {
	brightness = 0
	for row := 0; row < len(decorations); row++ {
		for column := 0; column < len(decorations[row]); column++ {
			brightness = brightness + decorations[row][column]
		}
	}
	return
}
func calculateLitLights(decorations [][]int) (lit int) {
	lit = 0
	for row := 0; row < len(decorations); row++ {
		for column := 0; column < len(decorations[row]); column++ {
			if decorations[row][column] == 1 {
				lit++
			}
		}
	}
	return
}

func generateBlankScreen(size int) [][]int {
	newBoard := make([][]int, size)
	for row := 0; row < size; row++ {
		newBoard[row] = make([]int, size)
	}
	return newBoard
}

func printPattern(decoration [][]int) {
	for row := 0; row < len(decoration); row++ {
		for column := 0; column < len(decoration[row]); column++ {
			fmt.Printf("%v", decoration[row][column])
		}
		fmt.Printf("\n")
	}
}
