package main

import (
	"fmt"
	"slices"
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
	return findAllIndexes(input)
}

func findAllIndexes(input string) int {
	startIndexes := make([]int, 0)
	endIndexes := make([]int, 0)

	for point := 0; point < len(input); point++ {
		secondSlice := input[point:]
		startIndex := strings.Index(secondSlice, "mul(")
		endIndex := strings.Index(secondSlice, ")")
		if startIndex > -1 {
			if !slices.Contains(startIndexes, startIndex+point) && startIndex+point > 0 {
				startIndexes = append(startIndexes, startIndex+point)
			}
		}
		if endIndex > -1 {
			if !slices.Contains(endIndexes, endIndex+point) && endIndex+point > 0 {
				endIndexes = append(endIndexes, endIndex+point)
			}
		}

	}
	currentEndIndexIndex := 0
	ret := 0
	for i := 0; i < len(startIndexes); i++ {
		startIndex := startIndexes[i]
		endIndex := endIndexes[currentEndIndexIndex]
		if startIndex > endIndex {
			value, changed := findFirstMatch(endIndexes, startIndex)
			if changed {
				currentEndIndexIndex = value
			}
		}
		if currentEndIndexIndex+1 < len(endIndexes) && startIndex < endIndex {
			testString := input[startIndex : endIndexes[currentEndIndexIndex]+1]
			if testString[0:4] == "mul(" && testString[len(testString)-1] == ')' {
				numbersString := testString[4 : len(testString)-1]
				splitNumbString := strings.Split(numbersString, ",")
				a, a_err := strconv.Atoi(splitNumbString[0])
				b, b_err := strconv.Atoi(splitNumbString[1])
				if a_err == nil && b_err == nil {
					fmt.Print(a, "times", b)
					ret = ret + (a * b)
					currentEndIndexIndex++
				}
			}
		}
	}
	return ret

}

func findFirstMatch(arr []int, target int) (int, bool) {
	for _, num := range arr {
		if num >= target {
			return num, true
		}
	}
	return 0, false
}
