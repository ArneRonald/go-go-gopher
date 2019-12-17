package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	_2015_dayOne "./2015/DayOne"
	_2019_dayThree "./2019/DayThree"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// directory := readFilesFromDirectory()
	// for {
	// 	fmt.Println("Which day do you want to run?")
	// 	for i, f := range directory {
	// 		fmt.Printf("%v: %v\r\n", i+1, f)
	// 	}
	// 	fmt.Print("->")
	// 	text, _ := reader.ReadString('\n')
	// 	text = strings.Replace(text, "\n", "", -1)
	// 	i, err := strconv.ParseInt(text, 10, 64)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	handleInput(directory[i-1])
	// }

	_2019_dayThree.ExecuteDayThree()

}

func readFilesFromDirectory() []string {
	files, err := ioutil.ReadDir("./src/")
	challenges := []string{}
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.Name() != "adventOfCode.go" && !strings.Contains(f.Name(), "test") {
			challenges = append(challenges, f.Name())
		}
	}
	challenges = append(challenges, "Exit")
	return challenges
}

func handleInput(input string) {
	switch input {
	case "DayOne":
		_2015_dayOne.ExecuteDayOne()
		break
	case "Exit":
		os.Exit(10)
		break
	default:
		fmt.Println("wat")
	}
}
