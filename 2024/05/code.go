package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	ret := 0
	// solve part 1 here
	rulesLinesSplit := strings.Split(input, "\n\n")
	rulesLines := strings.Split(rulesLinesSplit[0], "\n")
	rules := make([][]int, 0)
	for _, rules_line := range rulesLines {
		splitLine := strings.Split(rules_line, "|")
		a, a_error := strconv.Atoi(splitLine[0])
		b, b_error := strconv.Atoi(splitLine[1])
		if a_error == nil && b_error == nil {
			rule := []int{a, b}
			rules = append(rules, rule)
		}
	}
	cases_lines := strings.Split(rulesLinesSplit[1], "\n")
	for _, case_line := range cases_lines {
		numbers := strings.Split(case_line, ",")
		case_numbers := make([]int, 0)
		for _, number := range numbers {
			a, a_error := strconv.Atoi(number)
			if a_error == nil {
				case_numbers = append(case_numbers, a)
			}
		}
		is_line_invalid, midPoint := testCaseAgainstAllRules(rules, case_numbers)
		if is_line_invalid {
			ret = ret + midPoint
		}
	}
	return ret

}

func testCaseAgainstAllRules(rules [][]int, case_line []int) (bool, int) {
	is_line_invalid := true
	for test_index, test_number := range case_line {
		for _, rule := range rules {
			//if number matches the high number
			if rule[1] == test_number {
				for i := test_index; i < len(case_line); i++ {
					if case_line[i] == rule[0] {
						is_line_invalid = false
					}
				}
			}
		}
	}
	midIndex := len(case_line) / 2
	midPoint := case_line[midIndex]

	return is_line_invalid, midPoint
}

func getMidPoint(case_line []int) int {
	midIndex := len(case_line) / 2
	midPoint := case_line[midIndex]
	return midPoint
}

func reorderLine(rules [][]int, case_line []int) []int {
	return_line = make([]int, 0)
	for test_index, test_number := range case_line {
		must_be_before := make([]int, 0)
		must_be_after := make([]int, 0)

		for _, rule := range rules {
			//if number matches the high number
			if rule[1] == test_number {
				must_be_after = append(must_be_after, rule[0])
			}
			if rule[0] == test_number {
				must_be_after = append(must_be_after, rule[1])
			}
		}

	}
}

func testLineValidity(rules [][]int, case_line []int) bool {
	is_line_invalid := true

	for test_index, test_number := range case_line {
		for _, rule := range rules {
			//if number matches the high number
			if rule[1] == test_number {
				for i := test_index; i < len(case_line); i++ {
					if case_line[i] == rule[0] {
						is_line_invalid = false
					}
				}
			}
		}
	}
	return is_line_invalid
}
