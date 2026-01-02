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
		cols := make([][]int, len(strings.Split(input, "\n")[0])+1)
		lastLine := ""
		blankColumns := make([]int, 0)

		for _, line := range strings.Split(input, "\n") {
			lastLine = line
			numbers := strings.Split(line, "")

			for i, number := range numbers {
				intVal, err := strconv.Atoi(number)
				if err == nil {
					appendedColumn := append(cols[i], intVal)
					cols[i] = appendedColumn
				}
			}
		}
		for i, col := range cols {
			if len(col) == 0 {
				blankColumns = append(blankColumns, i)
			}
		}
		signIndexes := []int{}
		for i, char := range strings.Split(lastLine, "") {
			if char == "*" || char == "+" {
				signIndexes = append(signIndexes, i)
			}
		}
		retTotal := 0
		for i := 0; i < len(signIndexes); i++ {
			signIndex := signIndexes[i]
			stopIndex := len(cols) + 1
			if i+1 <= len(signIndexes)-1 {
				stopIndex = signIndexes[i+1]
			}
			total := calculateFunction(cols, signIndex, stopIndex, strings.Split(lastLine, "")[signIndex])
			retTotal = retTotal + total
		}
		return retTotal
	}
	//9434894930166 too low
	//9434900007303 too low
	// solve part 1 here
	colsTotal := 0
	re := regexp.MustCompile("\\s*")
	colCount := len(re.Split(strings.Split(input, "\n")[0], -1))
	cols := make([][]int, colCount)

func calculateFunction(cols [][]int, startIndex int, stopIndex int, sign string) int {
	columnsToUse := cols[startIndex : stopIndex-1]
	intsToUse := make([]int, 0)
	for _, columnToUse := range columnsToUse {
		columnString := ""
		for _, intInColumn := range columnToUse {
			intString := strconv.Itoa(intInColumn)
			columnString = columnString + intString
		}
		intValue, _ := strconv.Atoi(columnString)
		intsToUse = append(intsToUse, intValue)
	}

	if sign == "*" {
		total := 1
		for _, intValu := range intsToUse {
			if intValu != 0 {
				total = total * intValu

			}
		}
		return total
	}
	if sign == "+" {
		total := 0
		for _, intValue := range intsToUse {
			total = total + intValue
		}
		return total
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
