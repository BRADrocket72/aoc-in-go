package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

const MaxNumber = 99

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
	if part2 {
		result_value := 50
		times_past_zero := 0
		for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
			result_value, times_past_zero = handleLineApplyToNumberPt2(result_value, s, times_past_zero)
		}
		return times_past_zero
	}
	// solve part 1 here
	init_value := 50
	result_value := init_value
	times_hit_zero := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		result_value = handleLineApplyToNumber(result_value, s)
		if result_value == 0 {
			times_hit_zero++
		}
	}
	return times_hit_zero
}

func parseLine(line string) (int, int) {
	rInt, _ := strconv.Atoi(line[1:])
	passedZero := int(math.Floor(float64(rInt / (MaxNumber + 1))))
	amountToMove := rInt % (MaxNumber + 1)
	if line[0] == 'L' {
		return 0 - amountToMove, passedZero
	} else {
		return amountToMove, passedZero
	}

}

func handleOverflowAndUnderflow(value int) (int, int) {
	passedZero := 0
	if value == 0 {
		return value, passedZero
	}
	if value >= 0 && value <= MaxNumber {
		return value, passedZero
	}
	if value > MaxNumber {
		for value >= MaxNumber {
			if value < 0 {
				println("i am below 0")
			}
			value = value - (MaxNumber + 1)
			passedZero++
		}
	}

	if value < 0 {
		for value <= 0 {
			if value > MaxNumber {
				println("i am greater than max")
			}
			value = value + (MaxNumber + 1)
			passedZero++
		}
	}
	return value, passedZero
}

func handleOverflowAndUnderflowPt2(initValue int, value int) (int, int) {
	combinedValue := initValue + value
	if combinedValue < MaxNumber && combinedValue > 0 {
		return combinedValue, 0
	}
	if combinedValue < 0 {
		combinedValue = 100 + combinedValue
		if initValue != 0 {
			return combinedValue, 1
		}
		return combinedValue, 0
	}
	if combinedValue == 0 {
		return combinedValue, 1
	}
	if combinedValue == MaxNumber {
		return 99, 0
	}
	combinedValue = combinedValue - 100
	if initValue != 0 {
		return combinedValue, 1
	}
	return combinedValue, 0

}

func handleLineApplyToNumber(initNumber int, line string) int {
	movement, _ := parseLine(line)
	initNumber = initNumber + movement
	newValue, _ := handleOverflowAndUnderflow(initNumber)
	return newValue
}

func handleLineApplyToNumberPt2(initNumber int, line string, linesPassedSoFar int) (int, int) {
	movement, passedZero := parseLine(line)
	newValue, passedAgain := handleOverflowAndUnderflowPt2(initNumber, movement)
	return newValue, passedAgain + passedZero + linesPassedSoFar
}
