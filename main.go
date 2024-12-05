package main

import (
	"aoc2024/days"
	"aoc2024/util"
	"flag"
	"fmt"
	"os"
)

func main() {
	init := flag.Bool("i", false, "initialise new day")
	day := flag.Int("d", 0, "provide the day to use")
	flag.Parse()

	// error if there are no flags
	if *day == 0 && *init == false {
		fmt.Println("Please provide at least one flag")
		os.Exit(0)
	}

	// main logic
	if *init == true {
		util.Init(*day)
		return
	} else {
		runOne(day)
	}
}

func runOne(day *int) {
	str := fmt.Sprintf("day%d", *day)
	funcMap[str].(func())()
}

var funcMap = map[string]interface{}{
	"day1": days.Day1,
	"day2": days.Day2,
	"day3": days.Day3,
	"day4": days.Day4,
	"day5": days.Day5,
	//insert here
}
