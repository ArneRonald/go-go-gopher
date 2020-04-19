package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	twentyFifteenDayFive "github.com/ArneRonald/go-go-gopher/2015/DayFive"
	twentyFifteenDayFour "github.com/ArneRonald/go-go-gopher/2015/DayFour"
	twentyFifteenDayOne "github.com/ArneRonald/go-go-gopher/2015/DayOne"
	twentyFifteenDaySeven "github.com/ArneRonald/go-go-gopher/2015/DaySeven"
	twentyFifteenDaySix "github.com/ArneRonald/go-go-gopher/2015/DaySix"
	twentyFifteenDayThree "github.com/ArneRonald/go-go-gopher/2015/DayThree"
	twentyFifteenDayTwo "github.com/ArneRonald/go-go-gopher/2015/DayTwo"

	twentyNineteenDayOne "github.com/ArneRonald/go-go-gopher/2019/DayOne"
	twentyNineteenDayThree "github.com/ArneRonald/go-go-gopher/2019/DayThree"
	twentyNineteenDayTwo "github.com/ArneRonald/go-go-gopher/2019/DayTwo"
)

func main() {
	currentDir := "./"
	reader := bufio.NewReader(os.Stdin)
	rootDir := true
	for {
		if currentDir == "./" || currentDir == "." {
			rootDir = true
		} else {
			rootDir = false
		}
		challenges := readFilesFromDirectory(currentDir, rootDir)
		if len(challenges) < 3 {
			currentDir = makeMenuSelection(currentDir)
			challenges = readFilesFromDirectory(currentDir, rootDir)
		}
		for i, challenge := range challenges {
			if i == 0 {
				continue
			}
			fmt.Printf("%v: %v\n", i, challenge)
		}
		fmt.Printf("%v: %v\n", 0, challenges[0])
		fmt.Print("->")

		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		selectedOption, err := strconv.Atoi(input)
		if err != nil {
			fmt.Errorf(err.Error())
		}
		if selectedOption == 0 {
			break
		}

		currentDir = handleUserInput(currentDir, selectedOption, challenges)
	}
	fmt.Println("Thanks for playing")
}

func handleUserInput(currentDir string, selectedOption int, menu []string) string {
	if menu[selectedOption] == "Previous" {
		menuArr := strings.Split(currentDir, "/")
		currentDir = ""
		for i := 0; i < len(menuArr)-1; i++ {
			if menuArr[i] == "." {
				currentDir = currentDir + "."
			} else {
				currentDir = currentDir + "/" + menuArr[i]
			}
		}
		return currentDir
	}
	if currentDir != "./" {
		return currentDir + "/" + menu[selectedOption]

	}
	return currentDir + menu[selectedOption]

}

func readFilesFromDirectory(dir string, rootDir bool) []string {
	challenges := []string{}
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	challenges = append(challenges, "Exit")

	for _, file := range files {
		if file.IsDir() && file.Name() != ".git" {
			challenges = append(challenges, file.Name())
		}
	}

	if !rootDir {
		challenges = append(challenges, "Previous")
	}

	return challenges
}

func clearScreen() {
	print("\033[H\033[2J")
}

func makeMenuSelection(choice string) string {
	switch choice {
	case "./2015/DayOne":
		fmt.Println("---Executing 2015 day one solution---")
		twentyFifteenDayOne.ExecuteDayOne()
		fmt.Println("---Finished 2015 day one solution---")
	case "./2015/DayTwo":
		fmt.Println("---Executing 2015 day two solution---")
		twentyFifteenDayTwo.ExecuteDayTwo()
		fmt.Println("---Finished 2015 day two solution---")
	case "./2015/DayThree":
		fmt.Println("---Executing 2015 day three solution---")
		twentyFifteenDayThree.ExecuteDayThree()
		fmt.Println("---Finished 2015 day three solution---")
	case "./2015/DayFour":
		fmt.Println("---Executing 2015 day four solution---")
		twentyFifteenDayFour.ExecuteDayFour()
		fmt.Println("---Finished 2015 day four solution---")
	case "./2015/DayFive":
		fmt.Println("---Executing 2015 day five solution---")
		twentyFifteenDayFive.ExecuteDayFive()
		fmt.Println("---Finished 2015 day five solution---")
	case "./2015/DaySix":
		fmt.Println("---Executing 2015 day six solution---")
		twentyFifteenDaySix.ExecuteDaySix()
		fmt.Println("---Finished 2015 day six solution---")
	case "./2015/DaySeven":
		fmt.Println("---Executing 2015 day seven solution---")
		twentyFifteenDaySeven.ExecuteDaySix()
		fmt.Println("---Finished 2015 day seven solution---")
	case "./2019/DayOne":
		fmt.Println("---Executing 2019 day one solution---")
		twentyNineteenDayOne.ExecuteDayOne()
		fmt.Println("---Finished 2019 day one solution---")
	case "./2019/DayTwo":
		fmt.Println("---Executing 2019 day two solution---")
		twentyNineteenDayTwo.ExecuteDayTwo()
		fmt.Println("---Finished 2019 day two solution---")
	case "./2019/Daythree":
		fmt.Println("---Executing 2019 day three solution---")
		twentyNineteenDayThree.ExecuteDayThree()
		fmt.Println("---Finished 2019 day three solution---")
	}
	menuArr := strings.Split(choice, "/")
	currentDir := ""
	for i := 0; i < len(menuArr)-1; i++ {
		if menuArr[i] == "." {
			currentDir = currentDir + "."
		} else {
			currentDir = currentDir + "/" + menuArr[i]
		}
	}
	return currentDir
}
