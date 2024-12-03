package days

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInputDay3() string {
	filePath := "days/day3.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func parseInputDay3(input string) []string {
	reg := regexp.MustCompile("mul\\(\\d+,\\d+\\)")
	matches := reg.FindAllString(input, -1)

	return matches
}

func parseInputDay3Part2(input string) []string {
	reg := regexp.MustCompile("mul\\(\\d+,\\d+\\)|do\\(\\)|don't\\(\\)")
	matches := reg.FindAllString(input, -1)

	validMatches := []string{}
	enabled := true
	for _, act := range matches {
		if act == "do()" {
			enabled = true
            continue
		}
		if act == "don't()" {
			enabled = false
            continue
		}
		if enabled {
			validMatches = append(validMatches, act)
		}
	}

    return validMatches
}

func performCalc(input string) int {
	reg := regexp.MustCompile("\\d+")

	matches := reg.FindAllString(input, 2)

	a, err := strconv.Atoi(matches[0])
	if err != nil {
		fmt.Println("error converting string to number")
	}
	b, err := strconv.Atoi(matches[1])
	if err != nil {
		fmt.Println("error converting string to number")
	}

	return a * b
}

func day3Part1(input string) {
	matches := parseInputDay3(input)
	answer := 0
	for _, match := range matches {
		a := performCalc(match)
		answer += a
	}
	fmt.Println("day 3 part 1", answer)
}

func day3Part2(input string) {
    validMatches := parseInputDay3Part2(input)
    answer :=0
    for _, match := range validMatches {
        a := performCalc(match)
        answer += a 
    }
	fmt.Println("day 3 part 2", answer)
}

func Day3() {
	input := getInputDay3()
	day3Part1(input)
	day3Part2(input)
}
