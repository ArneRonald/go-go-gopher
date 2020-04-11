package main

import (
	"fmt"

	twentyFifteenDayFour "github.com/ArneRonald/go-go-gopher/2015/DayFour"
	twentyFifteenDayOne "github.com/ArneRonald/go-go-gopher/2015/DayOne"
	twentyFifteenDayThree "github.com/ArneRonald/go-go-gopher/2015/DayThree"
	twentyFifteenDayTwo "github.com/ArneRonald/go-go-gopher/2015/DayTwo"

	twentyNineteenDayOne "github.com/ArneRonald/go-go-gopher/2019/DayOne"
	twentyNineteenDayThree "github.com/ArneRonald/go-go-gopher/2019/DayThree"
	twentyNineteenDayTwo "github.com/ArneRonald/go-go-gopher/2019/DayTwo"
)

func main() {
	fmt.Println("---Executing 2015 day one solution---")
	twentyFifteenDayOne.ExecuteDayOne()
	fmt.Println("---Finished 2015 day one solution--- \n")

	fmt.Println("---Executing 2015 day two solution---")
	twentyFifteenDayTwo.ExecuteDayTwo()
	fmt.Println("---Finished 2015 day two solution---\n")

	fmt.Println("---Executing 2015 day three solution---")
	twentyFifteenDayThree.ExecuteDayThree()
	fmt.Println("---Finished 2015 day three solution---\n")

	fmt.Println("---Executing 2015 day four solution---")
	twentyFifteenDayFour.ExecuteDayFour()
	fmt.Println("---Finished 2015 day four solution---\n")

	fmt.Println("---Executing 2019 day one solution---")
	twentyNineteenDayOne.ExecuteDayOne()
	fmt.Println("---Finished 2019 day one solution---\n")

	fmt.Println("---Executing 2019 day two solution---")
	twentyNineteenDayTwo.ExecuteDayTwo()
	fmt.Println("---Finished 2019 day two solution---\n")

	fmt.Println("---Executing 2019 day three solution---")
	twentyNineteenDayThree.ExecuteDayThree()
	fmt.Println("---Finished 2019 day three solution---\n")

	fmt.Println("Done!")
}
