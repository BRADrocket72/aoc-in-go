package main

import (
	"math"
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

type RangeObject struct {
	Min, Max int
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		rangeList, _ := CreateRnageListAndIds(input)
		mergedRanges := mergeRanges(rangeList)
		totalValue := 0
		for _, mergedRange := range mergedRanges {
			totalValue = totalValue + mergedRange.Max - mergedRange.Min
		}
		return totalValue
		//399163709260641 too high
	}
	// solve part 1 here
	rangeList, ingredientIds := CreateRnageListAndIds(input)
	freshIdList := make([]int, 0)

	for _, ingredientId := range ingredientIds {
		for _, rangeObject := range rangeList {
			if ingredientId > rangeObject.Min && ingredientId <= rangeObject.Max {
				freshIdList = append(freshIdList, ingredientId)
			}
		}
	}
	uniqueIds := UniqueSliceElements(freshIdList)
	return len(uniqueIds)
}

func mergeRanges(ranges []RangeObject) []RangeObject {

	currentRanges := ranges
	for range len(currentRanges) {
		newRangeObjects := make([]RangeObject, 0)
		for _, rangeObject := range currentRanges {
			isAdded := false
			for newRangeIndex, newRangeObject := range newRangeObjects {
				aMaxIsGreater := rangeObject.Max >= newRangeObject.Max
				aMinIsGreater := rangeObject.Min >= newRangeObject.Min
				AMinIsLesser := rangeObject.Min <= newRangeObject.Min
				AMaxIsLesser := rangeObject.Max <= newRangeObject.Max
				if aMinIsGreater && AMinIsLesser {
					
				}
				if aMaxIsGreater {
					if rangeObject.Min >= newRangeObject.Max {
						continue
					}
					newExpandedRange := RangeObject{int(math.Min(float64(rangeObject.Min), float64(newRangeObject.Min))), int(math.Max(float64(rangeObject.Max), float64(newRangeObject.Max)))}
					newRangeObjects[newRangeIndex] = newExpandedRange
				}

			}
			if !isAdded {
				newRangeObjects = append(newRangeObjects, rangeObject)
			}
		}
		currentRanges = newRangeObjects
	}

	return currentRanges
}

func CreateRnageListAndIds(input string) ([]RangeObject, []int) {
	var rangeList []RangeObject
	var ingredientIds []int
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		if strings.Contains(s, "-") {
			var splitString = strings.Split(strings.TrimSpace(string(s)), "-")
			minValue, _ := strconv.Atoi(splitString[0])
			maxValue, _ := strconv.Atoi(splitString[1])
			rangeList = append(rangeList, RangeObject{minValue, maxValue})
		} else {
			if len(s) > 0 {
				{
					idInt, _ := strconv.Atoi(s)
					ingredientIds = append(ingredientIds, idInt)
				}
			}
		}
	}
	return rangeList, ingredientIds
}

func UniqueSliceElements(inputSlice []int) []int {
	uniqueSlice := make([]int, 0, len(inputSlice)) // Pre-allocate capacity
	seen := make(map[int]bool, len(inputSlice))    // Use a map as a set

	for _, element := range inputSlice {
		if !seen[element] {
			uniqueSlice = append(uniqueSlice, element)
			seen[element] = true
		}
	}

	return uniqueSlice
}

func expandRanges(rangeList []RangeObject) []RangeObject {
	expandedRanges := make([]RangeObject, 0)
	expandedRanges = append(expandedRanges, rangeList[0])
	for _, rangeObject := range rangeList {
		added := false

		for expandedRangeIndex, expandedRange := range expandedRanges {
			//check for no overlap //too low than too high
			if rangeObject.Min == expandedRange.Min && rangeObject.Max == expandedRange.Max {
				added = true
			} else if rangeObject.Min <= expandedRange.Min && rangeObject.Max >= expandedRange.Min {
				//do nothing
			} else if rangeObject.Min >= expandedRange.Max && rangeObject.Max >= expandedRange.Max {
				//do nothing
			} else if rangeObject.Min > expandedRange.Min && rangeObject.Max > expandedRange.Max {
				//Add to upper bound
				if !added {
					added = true
					expandedRanges[expandedRangeIndex] = RangeObject{expandedRange.Min, rangeObject.Max}
				}

			} else if rangeObject.Min < expandedRange.Min && rangeObject.Max < expandedRange.Max {
				//Add to lower bound
				if !added {
					added = true
					expandedRanges[expandedRangeIndex] = RangeObject{rangeObject.Min, expandedRange.Max}
				}
			}
		}
		if !added {
			expandedRanges = append(expandedRanges, rangeObject)
		}

	}
	return expandedRanges
}
