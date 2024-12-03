package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputDay2() string {
	filePath := "days/day2.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func parseInputDay2(input string) [][]int {
	reports := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		report := []int{}
		parts := strings.Split(line, " ")
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Println("Error converting to int", err)
			}
			report = append(report, num)
		}
		reports = append(reports, report)
	}

	return reports
}

func checkDirection(report []int) int {
	forward := false
	backward := false
	same := false
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			forward = true
		} else if report[i] > report[i+1] {
			backward = true
		} else {
			same = true
		}
	}

	if forward == true && backward == false && same == false {
		return 1
	}
	if forward == false && backward == true && same == false {
		return -1
	}
	if forward == false && backward == false && same == true {
		return 0
	}
	return 0
}

func checkSafety(report []int) bool {
	safe := true
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if diff < 0 {
			diff = -diff
		}

		if diff > 3 {
			safe = false
		}
	}

	return safe
}

func checkSafety2(report []int) bool {
	for i := range report {
		newSlice := make([]int, 0, len(report)-1)
		for j := range report {
			if i != j {
				newSlice = append(newSlice, report[j])
			}
		}
		dir := checkDirection(newSlice)
		safety := checkSafety(newSlice)

		if dir != 0 && safety == true {
			return true
		}
	}

	return false
}

func day2Part1(input string) {
	reports := parseInputDay2(input)
	safeReports := 0
	for _, report := range reports {
		dir := checkDirection(report)
		if dir != 0 {
			if checkSafety(report) {
				safeReports++
			}
		}
	}

	fmt.Println("day 2 part 1", safeReports)
}

func day2Part2(input string) {
	reports := parseInputDay2(input)
	safeReports := 0
	for _, report := range reports {
		if checkSafety2(report) {
			safeReports++
		}
	}

	fmt.Println("day 2 part 2", safeReports)
}

func Day2() {
	input := getInputDay2()
	day2Part1(input)
	day2Part2(input)
}
