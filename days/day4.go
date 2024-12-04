package days

import (
	"fmt"
	"os"
	"strings"
)

func getInputDay4() string {
	filePath := "days/day4.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func parseInputDay4(input string) [][]string {
	grid := [][]string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		l := []string{}
		chars := strings.Split(line, "")
		for _, char := range chars {
			l = append(l, char)
		}
		grid = append(grid, l)
	}

	return grid
}

func day4Part1(input string) {
	grid := parseInputDay4(input)
	words := 0

	limitR := len(grid)
	for r, row := range grid {
		limitC := len(row)
		for c := range row {
			// right
			if c+3 < limitC && grid[r][c] == "X" && grid[r][c+1] == "M" && grid[r][c+2] == "A" && grid[r][c+3] == "S" {
				fmt.Println("xmas found right", r, c)
				words++
			}
			// left
			if c+3 < limitC && grid[r][c] == "S" && grid[r][c+1] == "A" && grid[r][c+2] == "M" && grid[r][c+3] == "X" {
				fmt.Println("xmas found left", r, c)
				words++
			}
			// down
			if r+3 < limitR && grid[r][c] == "X" && grid[r+1][c] == "M" && grid[r+2][c] == "A" && grid[r+3][c] == "S" {
				fmt.Println("xmas found down", r, c)
				words++
			}
			// up
			if r+3 < limitR && grid[r][c] == "S" && grid[r+1][c] == "A" && grid[r+2][c] == "M" && grid[r+3][c] == "X" {
				fmt.Println("xmas found up", r, c)
				words++
			}

			// down right
			if r+3 < limitR && c+3 < limitC && grid[r][c] == "X" && grid[r+1][c+1] == "M" && grid[r+2][c+2] == "A" && grid[r+3][c+3] == "S" {
				fmt.Println("xmas found down right", r, c)
				words++
			}
			// up left
			if r+3 < limitR && c+3 < limitC && grid[r][c] == "S" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "M" && grid[r+3][c+3] == "X" {
				fmt.Println("xmas found up left", r, c)
				words++
			}
			// up right
			if r-3 >= 0 && c+3 < limitC && grid[r][c] == "X" && grid[r-1][c+1] == "M" && grid[r-2][c+2] == "A" && grid[r-3][c+3] == "S" {
				fmt.Println("xmas found up right", r, c)
				words++
			}
			// down left
			if r-3 >= 0 && c+3 < limitC && grid[r][c] == "S" && grid[r-1][c+1] == "A" && grid[r-2][c+2] == "M" && grid[r-3][c+3] == "X" {
				fmt.Println("xmas found down left", r, c)
				words++
			}
		}
	}

	fmt.Println("day 4 part 1", words)
}

func day4Part2(input string) {
	grid := parseInputDay4(input)
	words := 0

	limitR := len(grid)
	for r, row := range grid {
		limitC := len(row)
		for c := range row {
			if r+2 < limitR && c+2 < limitC && grid[r][c] == "M" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "S" && grid[r+2][c] == "M" && grid[r][c+2] == "S" {
				words++
			}
			if r+2 < limitR && c+2 < limitC && grid[r][c] == "M" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "S" && grid[r+2][c] == "S" && grid[r][c+2] == "M" {
				words++
			}
			if r+2 < limitR && c+2 < limitC && grid[r][c] == "S" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "M" && grid[r+2][c] == "M" && grid[r][c+2] == "S" {
				words++
			}
			if r+2 < limitR && c+2 < limitC && grid[r][c] == "S" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "M" && grid[r+2][c] == "S" && grid[r][c+2] == "M" {
				words++
			}
		}
	}
	fmt.Println("day 4 part 2", words)
}

func Day4() {
	input := getInputDay4()
	day4Part1(input)
	day4Part2(input)
}
