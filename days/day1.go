package days

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputDay1() string {
	filePath := "days/day1.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func parseInputDay1(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	left := []int{}
	right := []int{}
	for _, line := range lines {
		parts := strings.Split(line, "   ")
		l, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error converting to int", err)
		}
		r, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting to int", err)
		}
		left = append(left, l)
		right = append(right, r)
	}
	return left, right
}

func day1Part1(input string) {
	left, right := parseInputDay1(input)

	sort.Ints(left)
	sort.Ints(right)

	answer := 0
	for i, l := range left {
		diff := l - right[i]
		if diff < 0 {
			diff = -diff
		}
		answer += diff
	}

	fmt.Println("day 1 part 1", answer)
}

func countOccurrences(slice []int, target int) int {
	count := 0
	for _, num := range slice {
		if num == target {
			count++
		}
	}
	return count
}

func day1Part2(input string) {
	left, right := parseInputDay1(input)

	answer := 0
	for _, l := range left {
		number := countOccurrences(right, l)
		answer += number * l
	}

	fmt.Println("day 1 part 2", answer)
}

func Day1() {
	input := getInputDay1()
	day1Part1(input)
	day1Part2(input)
}
