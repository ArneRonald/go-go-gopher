package dayseven

import (
	"fmt"
	"github.com/ArneRonald/go-go-gopher/util"
	"os"
)

//ExecuteDaySeven ...
func ExecuteDaySeven() {
	dir, _ := os.Getwd()
	data := util.ReadFileContents(dir + "/2015/DaySeven/input.txt")
	for _, line := range data {
		fmt.Println(line)
	}

}

func ApplyOperations(data []string) map[string]int {
	return nil
}