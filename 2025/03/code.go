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
		totalJoltage := 0
		for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
			numbers := parseLineToNumberSplice(s)
			highestInLine := getHighest12CombinedNumberInArray(numbers, 11)
			totalJoltage += highestInLine
		}
		return totalJoltage
	}
	// solve part 1 here
	totalJoltage := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		numbers := parseLineToNumberSplice(s)
		highestInLine := getHighestCombinedNumberInArray(numbers)
		totalJoltage += highestInLine
	}
	return totalJoltage
}

func parseLineToNumberSplice(line string) []int {
	result := make([]int, len(line))
	for i, c := range line {
		result[i], _ = strconv.Atoi(string(c))
	}
	return result
}

func getHighestNumberInArray(array []int) (highest int, highestIndex int) {
	for i, c := range array {
		if c > highest {
			highest = c
			highestIndex = i
		}
	}
	return highest, highestIndex
}

func getHighestCombinedNumberInArray(array []int) (result int) {
	highestFirstDigitSoFar := 0
	for i, c := range array {
		subResult := 0
		if c < highestFirstDigitSoFar {
			continue
		}
		subArray := array[i+1:]
		highestSecondDigit, _ := getHighestNumberInArray(subArray)
		subResult = c*10 + highestSecondDigit
		if subResult > result {
			result = subResult
		}
	}
	return result
}

func getHighest12CombinedNumberInArray(array []int, extraDigitsNeeded int) (result int) {
	highestFirstDigitSoFar := 0
	for i, c := range array {
		subResult := 0
		if c < highestFirstDigitSoFar {
			continue
		}
		if i+extraDigitsNeeded > len(array) {
			break
		}
		highTwelveArray := make([]int, extraDigitsNeeded+1)
		highTwelveArray[0] = c
		lastIndexUsed := len(array)
		localHighTotal := 0
		for j := extraDigitsNeeded; j > 0; j-- {
			if i+j > lastIndexUsed {
				continue
			}

			subArray := array[i+j : lastIndexUsed]
			highestDigitInSubstring, highestIndex := getHighestNumberInArray(subArray)
			lastIndexUsed = i + j + highestIndex
			testArray := highTwelveArray
			testArray[j] = highestDigitInSubstring
			localResult, _ := parseArrayIntoNumber(testArray)
			if localResult > localHighTotal {
				localHighTotal = localResult
				highTwelveArray[j] = highestDigitInSubstring
			}
		}
		subResult, _ = parseArrayIntoNumber(highTwelveArray)
		if subResult > result {
			result = subResult
		}
	}
	return result
}

func parseArrayIntoNumber(array []int) (int, error) {
	resultString := ""
	for _, c := range array {
		str := strconv.Itoa(c)
		resultString += str
	}
	return strconv.Atoi(resultString)
}
