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
type estimate struct {
	total  int
	values []int
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	//too low 55940032496684
	//too high 106016739583593
	if part2 {
		lines := strings.Split(input, "\n")
		estimates := make([]estimate, 0)
		for _, line := range lines {
			linesOnColon := strings.Split(line, ":")
			total, _ := strconv.Atoi(linesOnColon[0])
			values := make([]int, 0)
			values_strings := strings.Split(linesOnColon[1], " ")
			for _, value := range values_strings {
				value, value_error := strconv.Atoi(value)
				if value_error == nil {
					values = append(values, value)
				}
			}
			estimate := estimate{
				total:  total,
				values: values,
			}
			estimates = append(estimates, estimate)
		}
		ret := 0
		for _, estimate := range estimates {
			if testEstimate(estimate) {
				ret = ret + estimate.total
			}
		}
		return ret
	}
	// solve part 1 here

	lines := strings.Split(input, "\n")
	estimates := make([]estimate, 0)
	for _, line := range lines {
		linesOnColon := strings.Split(line, ":")
		total, _ := strconv.Atoi(linesOnColon[0])
		values := make([]int, 0)
		values_strings := strings.Split(linesOnColon[1], " ")
		for _, value := range values_strings {
			value, value_error := strconv.Atoi(value)
			if value_error == nil {
				values = append(values, value)
			}
		}
		estimate := estimate{
			total:  total,
			values: values,
		}
		estimates = append(estimates, estimate)
	}
	ret := 0
	for _, estimate := range estimates {
		if testEstimate(estimate) {
			ret = ret + estimate.total
		}
	}
	return ret

}

func testEstimate(estimate estimate) bool {
	possible_values := []int{estimate.values[0]}

	for i := 1; i < len(estimate.values); i++ {
		next_possible_values := make([]int, 0)
		for _, possible_value := range possible_values {
			sum := possible_value + estimate.values[i]
			product := possible_value * estimate.values[i]
			concatString := strconv.Itoa(possible_value) + strconv.Itoa(estimate.values[i])
			concat, _ := strconv.Atoi(concatString)

			if sum == estimate.total || product == estimate.total || concat == estimate.total {
				return true
			}
			next_possible_values = append(next_possible_values, sum)
			next_possible_values = append(next_possible_values, product)
			next_possible_values = append(next_possible_values, concat)

		}
		possible_values = next_possible_values
	}
	return false
} //too high 12940396356932
//correct 12940396350192z
