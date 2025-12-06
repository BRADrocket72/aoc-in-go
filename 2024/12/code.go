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
	// solve part 1 here
	init_value := 50
	result_value := init_value
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		result_value += parseLine(s)
		if result_value > 50 {
			result_value = result_value - 50
		}
		if result_value < 0 {
			result_value = 50 + result_value
		}
		println(result_value)
	}
	return result_value
}

func parseLine(line string) int {
	r_int, _ := strconv.Atoi(s[1:])
	if s[0] == 'L' {
		return 0 - r_int
	} else {
		return r_int
	}
}
