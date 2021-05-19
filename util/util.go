package util

import (
	"bufio"
	"log"
	"os"
)

func ReadFileContents(path string) (data []string){
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	file.Close()
	return
}
