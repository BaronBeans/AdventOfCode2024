package days

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y int
	x int
}

type Instruction struct {
	position  Point
	direction Point
}

type Pointer struct {
	position  Point
	direction Point
	actions   map[Instruction]struct{}
	visited   map[Point]struct{}
	loops     bool
	finished  bool
}

func getInputDay6() string {
	filePath := "days/day6.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func parseInputDay6(input string) (Pointer, [][]string) {
	m := [][]string{}
	p := Pointer{actions: make(map[Instruction]struct{}), visited: make(map[Point]struct{}), loops: false, finished: false}
	lines := strings.Split(input, "\n")
	for r, line := range lines {
		l := []string{}
		squares := strings.Split(line, "")
		for c, square := range squares {
			if square == "^" {
				p.direction = Point{y: -1, x: 0}
				p.position = Point{y: r, x: c}
			}
			l = append(l, square)
		}
		m = append(m, l)
	}

	return p, m
}

func getVisited(m [][]string) []Point {
	visited := []Point{}
	for y, row := range m {
		for x := range row {
			if m[y][x] == "X" {
				visited = append(visited, Point{y, x})
			}
		}
	}
	return visited
}

func countVisited(m [][]string) int {
	count := 0
	for y, row := range m {
		for x := range row {
			if m[y][x] == "X" {
				count++
			}
		}
	}
	return count
}

func printPointer(p Pointer) {
	fmt.Println("position", p.position, "direction", p.direction, "visited", len(p.actions))
}

func move(p Pointer, m [][]string) (Pointer, [][]string) {
	r := p.position.y
	c := p.position.x
	dr := p.direction.y
	dc := p.direction.x

	newP := p
	newP.position = Point{y: r + dr, x: c + dc}

	if newP.position.y < 0 || newP.position.y >= len(m) || newP.position.x < 0 || newP.position.x >= len(m[0]) {
		p.finished = true
		return p, m
	}
	if m[newP.position.y][newP.position.x] == "#" {
		p.direction.y = dc
		p.direction.x = -dr
		return p, m
	}
	_, exists := p.actions[Instruction{position: newP.position, direction: newP.direction}]
	if exists {
		p.loops = true
		p.finished = true
		return p, m
	}
	m[newP.position.y][newP.position.x] = "X"
	newP.actions[Instruction{position: newP.position, direction: newP.direction}] = struct{}{}
	newP.visited[newP.position] = struct{}{}
	return newP, m
}

func day6Part1(input string) [][]string {
	p, m := parseInputDay6(input)

	for {
		p, m = move(p, m)
		if p.finished {
			break
		}
	}

	fmt.Println("day 6 part 1", len(p.visited))

	return m
}

func day6Part2(input string, d [][]string) {
	visited := getVisited(d)

	loopsCount := 0

	for _, v := range visited {
		p, m := parseInputDay6(input)
		m[v.y][v.x] = "#"
		for {
			p, m = move(p, m)
			if p.loops {
				loopsCount++
				break
			}
			if p.finished {
				break
			}
		}
	}

	fmt.Println("day 6 part 2", loopsCount)
}

func Day6() {
	input := getInputDay6()
	m := day6Part1(input)
	day6Part2(input, m)
}
