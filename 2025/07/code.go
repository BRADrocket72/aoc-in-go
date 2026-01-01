package main

import (
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
		lines := strings.Split(strings.TrimSpace(string(input)), "\n")
		emptyArray := []int{}
		leftCount := HandleLeftRightSplitPt2(lines, true, emptyArray)
		rightCount := HandleLeftRightSplitPt2(lines, false, emptyArray)
		return leftCount + rightCount

	}
	// solve part 1
	indexesThatAreBeamed := make([]int, 0)
	splitCount := 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {

		chars := strings.Split(s, "")
		for i, char := range chars {
			if char == "S" {
				indexesThatAreBeamed = append(indexesThatAreBeamed, i)
			}
			if char == "^" {
				for nowSplitIndex, beamedIndex := range indexesThatAreBeamed {
					if beamedIndex == i {
						indexesThatAreBeamed = append(indexesThatAreBeamed[:nowSplitIndex], indexesThatAreBeamed[nowSplitIndex+1:]...)
						if i-1 > 0 {
							if !isIndexAlreadySplit(indexesThatAreBeamed, i-1) {
								indexesThatAreBeamed = append(indexesThatAreBeamed, i-1)
							}
						}
						if !isIndexAlreadySplit(indexesThatAreBeamed, i+1) {
							indexesThatAreBeamed = append(indexesThatAreBeamed, i+1)
						}
						splitCount = splitCount + 1
					}

				}
			}
		}
	}
	return splitCount
}

func isIndexAlreadySplit(split []int, indexToCheck int) bool {
	for _, indexThatsBeenSplit := range split {
		if indexThatsBeenSplit == indexToCheck {
			return true
		}
	}
	return false
}

func HandleLeftRightSplitPt2(input []string, isLeft bool, beamedIndexes []int) int {
	splitCount := 0
	if len(input) == 0 {
		return 0
	}
	for _, s := range input {
		nextLines := input[1:]
		chars := strings.Split(s, "")
		indexesThatAreBeamed := beamedIndexes
		for i, char := range chars {
			if char == "S" {
				indexesThatAreBeamed = append(indexesThatAreBeamed, i)
				var leftCount = HandleLeftRightSplitPt2(nextLines, true, indexesThatAreBeamed)
				var rightCount = HandleLeftRightSplitPt2(nextLines, false, indexesThatAreBeamed)
				splitCount = splitCount + leftCount + rightCount
				return splitCount
			}
			if char == "^" {
				for nowSplitIndex, beamedIndex := range indexesThatAreBeamed {
					if beamedIndex == i {
						indexesThatAreBeamed = append(indexesThatAreBeamed[:nowSplitIndex], indexesThatAreBeamed[nowSplitIndex+1:]...)
						if i-1 > 0 {
							if !isIndexAlreadySplit(indexesThatAreBeamed, i-1) && isLeft {
								indexesThatAreBeamed = append(indexesThatAreBeamed, i-1)
								var leftCount = HandleLeftRightSplitPt2(nextLines, true, indexesThatAreBeamed)
								var rightCount = HandleLeftRightSplitPt2(nextLines, false, indexesThatAreBeamed)
								splitCount = splitCount + leftCount + rightCount
							}
						}
						if !isIndexAlreadySplit(indexesThatAreBeamed, i+1) && !isLeft {
							indexesThatAreBeamed = append(indexesThatAreBeamed, i+1)
							var leftCount = HandleLeftRightSplitPt2(nextLines, true, indexesThatAreBeamed)
							var rightCount = HandleLeftRightSplitPt2(nextLines, false, indexesThatAreBeamed)
							splitCount = splitCount + leftCount + rightCount

						}
						//println("totalt", splitCount)
					}

				}
			}
		}
	}
	return splitCount

}
