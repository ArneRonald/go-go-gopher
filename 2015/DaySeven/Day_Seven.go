package dayseven

import (
	"bufio"
	"log"
	"os"
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

	for scanner.Scan() {
		scanner.Text()
	}

}
