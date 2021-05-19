package dayseven

import (
	"fmt"
	"github.com/ArneRonald/go-go-gopher/util"
	"os"
)

//ExecuteDaySix ...
func ExecuteDaySix() {
	dir, _ := os.Getwd()
	data := util.ReadFileContents(dir + "/2015/DaySix/input.txt")
	for _, line := range data {
		fmt.Println(line)
	}

}
