package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputDay7() string {
	filePath := "days/day7.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

type Equation struct {
	total int
	parts []int
}

func parseInputDay7(input string) []Equation {
	eqs := []Equation{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		eqParts := strings.Split(line, ": ")
		parts := strings.Split(eqParts[1], " ")
		total, _ := strconv.Atoi(eqParts[0])
		eq := Equation{total: total, parts: []int{}}
		for _, part := range parts {
			num, _ := strconv.Atoi(part)
			eq.parts = append(eq.parts, num)
		}
		eqs = append(eqs, eq)
	}
	return eqs
}

func equationIsValid(total int, parts []int) bool {
	if len(parts) == 1 {
		if parts[0] == total {
			return true
		}
		return false
	}
	last := parts[len(parts)-1]
	if total%last == 0 && equationIsValid(total/last, parts[:len(parts)-1]) {
		return true
	}
	if total > last && equationIsValid(total-last, parts[:len(parts)-1]) {
		return true
	}

	return false
}

func equationIsValid2(total int, parts []int) bool {
	if len(parts) == 1 && parts[0] == total {
		return true
	}
	if len(parts) == 1 && parts[0] != total {
		return false
	}
	if len(parts) == 0 {
		return false
	}
	last := parts[len(parts)-1]
	if total%last == 0 && equationIsValid2(total/last, parts[:len(parts)-1]) {
		return true
	}
	if total > last && equationIsValid2(total-last, parts[:len(parts)-1]) {
		return true
	}
	totalStr := strconv.Itoa(total)
	lastStr := strconv.Itoa(last)

	if len(totalStr) < len(lastStr) {
		return false
	}

	totalWithout := totalStr[:len(totalStr)-len(lastStr)]
	lastFromTotal := totalStr[len(totalStr)-len(lastStr):]

	newTotal, _ := strconv.Atoi(totalWithout)
	if len(totalStr) > len(lastStr) && lastFromTotal == lastStr && equationIsValid2(newTotal, parts[:len(parts)-1]) {
		return true
	}

	return false
}

func day7Part1(input string) {
	count := 0
	eqs := parseInputDay7(input)

	for _, eq := range eqs {
		if equationIsValid(eq.total, eq.parts) {
			count += eq.total
		}
	}
	fmt.Println("day 7 part 1", count)
}

func day7Part2(input string) {
	count := 0
	eqs := parseInputDay7(input)

	for _, eq := range eqs {
		if equationIsValid2(eq.total, eq.parts) {
			count += eq.total
		}
	}
	fmt.Println("day 7 part 2", count)
}

func Day7() {
	input := getInputDay7()
	day7Part1(input)
	day7Part2(input)
}
