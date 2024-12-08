package main

import (
	"fmt"
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
			is_line_invalid, _ := testCaseAgainstAllRules(rules, case_numbers)
			if !is_line_invalid {

				ordered_list := orderList(rules, case_numbers)
				ret = ret + getMidPoint(ordered_list)
			}
		}
		return ret
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
	fmt.Print(case_line)
	midIndex := len(case_line) / 2
	midPoint := case_line[midIndex]
	fmt.Print(midPoint)

	return midPoint
}

// func orderList(rules [][]int, case_line []int) []int {
// 	internal_relevant_rules := make([][]int, 0)

// 	orderList := make([]int, len(case_line))
// 	for i := 0; i < len(case_line); i++ {

// 		for test_index, test_number := range case_line[i:] {
// 			number_must_be_after_something := false
// 			number_must_be_before_something := false
// 			for _, rule := range internal_relevant_rules {
// 				//if number matches the high number
// 				next_revelevant_rules := make([][]int, 0)
// 				if rule[1] == test_number {
// 					next_revelevant_rules = append(next_revelevant_rules, rule)
// 					for remainingNumber := range case_line[test_index:] {
// 						if rule[0] == remainingNumber {
// 							number_must_be_after_something = true
// 						}
// 					}
// 				}
// 				if rule[0] == test_number {
// 					next_revelevant_rules = append(next_revelevant_rules, rule)
// 					for remainingNumber := range case_line[test_index:] {
// 						if rule[1] == remainingNumber {
// 							number_must_be_before_something = true
// 						}
// 					}
// 				}
// 				internal_relevant_rules = next_revelevant_rules
// 			}

// 			// if !number_must_be_before_something {
// 			// 	orderList[len(case_line)-1] = test_number
// 			// }
// 		}
// 		if !number_must_be_after_something {
// 			orderList[i] = test_number
// 		}
// 	}

// 	return orderList
// }

func orderList(rules [][]int, case_line []int) []int {
	n := len(case_line)
	for i := 0; i < n-1; i++ {
		for j := 1; j < n; j++ {
			if searchRules(case_line[i], case_line[j], rules) {
				a := case_line[i]
				b := case_line[j]
				case_line[j] = a
				case_line[i] = b
			}
		}
	}
	return case_line
	// for (int i = 0; i < n - 1; i++) {
	//     swapped = false;
	//     for (int j = 0; j < n - i - 1; j++) {
	//         if (arr[j] > arr[j + 1]) {
	//             swap(arr[j], arr[j + 1]);
	//             swapped = true;
	//         }
	//     }

	//     // If no two elements were swapped, then break
	//     if (!swapped)
	//         break;
	// }
}

func isABeforeB(a int, b int, rules [][]int) bool {
	for _, rule := range rules {
		if rule[1] == a && rule[0] == b {
			return false
		} else if rule[0] == a && rule[1] == b {
			return true
		}
	}
	return true
}

func searchRules(a int, b int, rules [][]int) bool {
	if a == b {
		return true
	}
	found := false
	for _, rule := range rules {
		if rule[0] == a {
			found = searchRules(rule[1], b, rules)
		}
	}
	return found
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
