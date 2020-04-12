package dayfive

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//ExecuteDayFive ...
func ExecuteDayFive() {
	dir, err := os.Getwd()
	file, err := os.Open(dir + "/2015/DayFive/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	amount := 0

	for scanner.Scan() {
		if determineNaughtyOrNice(scanner.Text()) {
			amount++
		}
	}
	fmt.Println("The amount of Nice names on the list were:", amount)
}

func determineNaughtyOrNice(name string) (nice bool) {
	nice = false
	if containsThreeVowels(name) && containsOneLetterThatAppearsTwiceInARow(name) && !containRestrictedWords(name) {
		nice = true
	}
	return
}

func containsThreeVowels(name string) bool {
	vowels := map[string]int{"a": 0, "e": 0, "i": 0, "o": 0, "u": 0}
	for _, char := range strings.Split(name, "") {
		if char == "a" || char == "e" || char == "i" || char == "o" || char == "u" {
			vowels[char]++
		}
	}
	vowelCount := 0
	for _, amount := range vowels {
		vowelCount = vowelCount + amount
	}
	if vowelCount >= 3 {
		return true
	}
	return false
}

func containsOneLetterThatAppearsTwiceInARow(name string) bool {
	previousChar := ""
	for _, char := range strings.Split(name, "") {
		if char == previousChar {
			return true
		}
		previousChar = char
	}
	return false
}

func containRestrictedWords(name string) bool {
	restrictedWords := []string{"ab", "cd", "pq", "xy"}
	for _, restrictedWord := range restrictedWords {
		if strings.Contains(name, restrictedWord) {
			return true
		}
	}
	return false
}
