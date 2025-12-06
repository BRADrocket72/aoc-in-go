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
		invalid_sum := 0
		for _, numberPair := range strings.Split(strings.TrimSpace(string(input)), ",") {
			a, b := parseNumberPair(numberPair)
			allNumbers := createSpliceOfAllNumbers(a, b)
			mirrors := parseArrayForMirrorsPt2(allNumbers)
			for _, mirror := range mirrors {
				invalid_sum += mirror
			}
		}
		return invalid_sum
	}
	// solve part 1 here
	invalid_sum := 0
	for _, numberPair := range strings.Split(strings.TrimSpace(string(input)), ",") {
		a, b := parseNumberPair(numberPair)
		allNumbers := createSpliceOfAllNumbers(a, b)
		mirrors := parseArrayForMirrors(allNumbers)
		for _, mirror := range mirrors {
			invalid_sum += mirror
		}
	}
	return invalid_sum
}

func parseNumberPair(s string) (int, int) {
	numbersArray := strings.Split(s, "-")
	a, _ := strconv.Atoi(numbersArray[0])
	b, _ := strconv.Atoi(numbersArray[1])
	return a, b
}

func createSpliceOfAllNumbers(a int, b int) []int {
	var allNumbers []int
	for a < (b + 1) {
		allNumbers = append(allNumbers, a)
		a++
	}
	return allNumbers
}

func parseArrayForMirrors(allNumbers []int) []int {
	var mirrorArray []int
	for _, val := range allNumbers {
		intStr := strconv.Itoa(val)
		strLen := len(intStr)
		if (strLen % 2) != 0 {
			continue
		}
		midPoint := strLen / 2
		a := intStr[0:midPoint]
		b := intStr[midPoint:strLen]
		if a == b {
			mirrorArray = append(mirrorArray, val)
			continue
		}

	}
	return mirrorArray
}

func splitIntoGroups(s string, n int) []string {
	var groups []string
	for i := 0; i < len(s); i += n {
		end := i + n
		if end > len(s) {
			end = len(s)
		}
		groups = append(groups, s[i:end])
	}
	return groups
}

func checkGroupsForSameness(groups []string) bool {
	firstGroup := groups[0]
	for _, group := range groups {
		if firstGroup != group {
			return false
		}
	}
	return true
}

func parseArrayForMirrorsPt2(allNumbers []int) []int {
	var mirrorArray []int
	for _, val := range allNumbers {
		intStr := strconv.Itoa(val)
		strLen := len(intStr)
		for i := 1; i < strLen; i++ {
			if (strLen % i) != 0 {
				continue
			}
			groups := splitIntoGroups(intStr, i)

			isGroupSame := checkGroupsForSameness(groups)
			if isGroupSame {
				mirrorArray = append(mirrorArray, val)
				break
			}
		}

	}
	return mirrorArray
}
