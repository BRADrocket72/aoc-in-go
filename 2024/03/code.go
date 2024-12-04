package main

import (
	"fmt"
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
	findAllIndexes(input)
	return 42
}

func findAllIndexes(input string) {
	slicedInput := input
	for point := 0; point < len(input); point++ {
		startIndex := strings.Index(input, "mul(")
		fmt.Print(startIndex)
		if startIndex == -1 || len(slicedInput) == 0 {
			point = len(input)
		}
		if startIndex > -1 && len(slicedInput) > 0 {
			slicedInput = slicedInput[startIndex:]
			var endIndex = strings.Index(input, ")")

			if endIndex == -1 {
				point = len(input)
			}
			if endIndex > -1 && endIndex < len(slicedInput) {
				digits := slicedInput[:endIndex]
				point = endIndex
				fmt.Print(digits)

			}

		}

	}

}
