package main

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Pattern struct {
	characters       []string
	towelsUsedToMake [][][]string
}

type Towel []string

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
	lines := strings.Split(input, "\n")
	towelsLine := lines[0]
	towelsStrings := strings.Split(towelsLine, ", ")
	towels := make([]Towel, 0)
	for _, towel := range towelsStrings {
		towels = append(towels, strings.Split(towel, ""))
	}

	patterns := lines[2:]
	for _, pattern := range patterns {
		findAllCombos(towels, pattern)

	}
	return 42
}

func findAllCombos(towels []Towel, pattern string) {
	for i, _ := range pattern {
		runeValue, _ := utf8.DecodeRuneInString(pattern[i:])
		for j, towel := range towels {
			if i == 0 && j == 0 {
				fmt.Print(towel)
				fmt.Print(runeValue)

			}
		}
	}

}
