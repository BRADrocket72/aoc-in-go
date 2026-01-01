package main

import (
	"fmt"
	"math"
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
		colsTotal := 0
		re := regexp.MustCompile("\\s*")
		colCount := len(re.Split(strings.Split(input, "\n")[0], -1))
		cols := make([][]int, colCount)

		for _, line := range strings.Split(input, "\n") {
			numbers := strings.Fields(line)
			for i, number := range numbers {
				intVal, err := strconv.Atoi(number)
				if err == nil {
					cols[i] = append(cols[i], intVal)
				}
				if err != nil {
					if number == "*" {
						colsTotal = colsTotal + multiplySlicePt2(cols[i])
					} else if number == "+" {
						colsTotal = colsTotal + sumSlicePt2(cols[i])
					}
				}
			}
		}
		return colsTotal
	}
	// solve part 1 here
	colsTotal := 0
	re := regexp.MustCompile("\\s*")
	colCount := len(re.Split(strings.Split(input, "\n")[0], -1))
	cols := make([][]int, colCount)

	for _, line := range strings.Split(input, "\n") {
		numbers := strings.Fields(line)
		for i, number := range numbers {
			intVal, err := strconv.Atoi(number)
			if err == nil {
				cols[i] = append(cols[i], intVal)
			}
			if err != nil {
				if number == "*" {
					colsTotal = colsTotal + multiplySlice(cols[i])
				} else if number == "+" {
					colsTotal = colsTotal + sumSlice(cols[i])
				}
			}
		}
	}
	return colsTotal
}

func sumSlice(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func multiplySlice(numbers []int) int {
	total := 1
	for _, number := range numbers {
		total = total * number
	}
	return total
}

func sumSlicePt2(numbers []int) int {
	sum := 0
	maxDigit := math.Pow10(len(numbers) - 1)
	addedVal := make([]int, 0)
	//312
	for colInd, number := range numbers {
		place := maxDigit / math.Pow10(colInd)
		//does not handle single digits well=rconv.Itoa(number)
		numSplitToDigits := strings.Split(str, "")
		for _, numSplit := range numSplitToDigits {
			var digit, _ = strconv.Atoi(numSplit)
			placeInt := int(place)
			sum += placeInt * digit
			addedVal = append(addedVal, (placeInt * digit))
		}
		fmt.Println(addedVal)
	}
	return sum
}

func multiplySlicePt2(numbers []int) int {
	total := 1
	for _, number := range numbers {
		total = total * number
	}
	return total
}
