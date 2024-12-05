package days

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	orders  [][]int
	updates [][]int
}

type Validity struct {
	valid   [][]int
	invalid [][]int
}

func getInputDay5() string {
	filePath := "days/day5.input"
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
	}

	return strings.TrimRight(string(file), "\n")
}

func parseInputDay5(input string) Input {
	parts := strings.Split(input, "\n\n")
	orders := [][]int{}
	for _, o := range strings.Split(parts[0], "\n") {
		p := strings.Split(o, "|")
		a, _ := strconv.Atoi(p[0])
		b, _ := strconv.Atoi(p[1])

		orders = append(orders, []int{a, b})
	}
	rules := [][]int{}
	for _, r := range strings.Split(parts[1], "\n") {
		rule := []int{}
		p := strings.Split(r, ",")
		for i := range p {
			page, _ := strconv.Atoi(p[i])
			rule = append(rule, page)
		}
		rules = append(rules, rule)
	}

	return Input{orders, rules}
}

func checkValidity(a, b int, input Input) bool {
	for _, o := range input.orders {
		if o[0] == b && o[1] == a {
			return false
		}
	}
	return true
}

func getPairs(input []int) [][]int {
	pairs := [][]int{}

	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			pairs = append(pairs, []int{input[i], input[j]})
		}
	}
	return pairs
}

func getValidUpdates(input Input) Validity {
	valid := [][]int{}
	invalid := [][]int{}

	for _, update := range input.updates {
		v := true
		pairs := getPairs(update)

		for _, p := range pairs {
			if checkValidity(p[0], p[1], input) == false {
				v = false
			}
		}

		if v {
			valid = append(valid, update)
		} else {
			invalid = append(invalid, update)
		}
	}

	return Validity{valid, invalid}
}

func day5Part1(input string) {
	ans := 0
	instructions := parseInputDay5(input)
	validity := getValidUpdates(instructions)
	for _, v := range validity.valid {
		ans += v[len(v)/2]
	}
	fmt.Println("day 5 part 1", ans)
}

func swapItems(arr []int, a, b int) []int {
	newArr := []int{}
    newArr = append(newArr, arr[:a]...)
	newArr = append(newArr, arr[b])
    newArr = append(newArr, arr[a+1:b]...)
	newArr = append(newArr, arr[a])
    newArr = append(newArr, arr[b+1:]...)

	return newArr
}

func day5Part2(input string) {
	instructions := parseInputDay5(input)
	validity := getValidUpdates(instructions)
    ans :=  0
	for _, item := range validity.invalid {
		changed := true
		for changed {
			changed = false
			for i := 0; i < len(item)-1; i++ {
				for _, o := range instructions.orders {
					if o[0] == item[i+1] && o[1] == item[i] {
						item = swapItems(item, i, i+1)
						changed = true
						break
					}
				}
				if changed {
					break
				}
			}
		}

        ans += item[len(item)/2]
	}
	fmt.Println("day 5 part 2", ans)
}

func Day5() {
	input := getInputDay5()
	day5Part1(input)
	day5Part2(input)
}
