package main

import (
	"fmt"
	"regexp"
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
		pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`)
		dos := regexp.MustCompile(`do()`) // Match one or more digits

		donts := regexp.MustCompile(`don't()`)

		matches := pattern.FindAllString(input, -1) // -1 means find all matches
		bytes := []byte(input)
		matchesIndexes := pattern.FindAllIndex(bytes, -1) // -1 means find all matches
		dosIndexes := dos.FindAllIndex(bytes, -1)         // -1 means find all matches
		dontsIndexes := donts.FindAllIndex(bytes, -1)     // -1 means find all matches

		fmt.Print(dosIndexes)
		var isDo = true
		for indexOfMatch, match := range matches {
			currentPointIndex := matchesIndexes[indexOfMatch][0]
			lastDo := 0
			lastDont := 1000000000000
			for _, doIndex := range dosIndexes {
				if doIndex[1] < currentPointIndex {
					lastDo = doIndex[0]
				}
			}
			for _, dontIndex := range dontsIndexes {
				if dontIndex[1] < currentPointIndex {
					lastDont = dontIndex[0]
				}
			}
			if lastDo > lastDont || lastDont == 1000000000000 {
				isDo = true
			} else {
				isDo = false
			}
			numStrings := match[4 : len(match)-1]
			splitStrings := strings.Split(numStrings, ",")
			a, a_err := strconv.Atoi(splitStrings[0])
			b, b_err := strconv.Atoi(splitStrings[1])
			if a_err == nil && b_err == nil && isDo {
				ret = ret + (a * b)
			}

		}
		return ret
	}
	ret := 0
	pattern := regexp.MustCompile(`mul\([0-9]+,[0-9]+\)`) // Match one or more digits

	matches := pattern.FindAllString(input, -1) // -1 means find all matches

	for _, match := range matches {
		numStrings := match[4 : len(match)-1]
		splitStrings := strings.Split(numStrings, ",")
		a, a_err := strconv.Atoi(splitStrings[0])
		b, b_err := strconv.Atoi(splitStrings[1])
		if a_err == nil && b_err == nil {
			ret = ret + (a * b)
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
