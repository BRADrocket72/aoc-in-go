package main

import (
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Point struct {
	X, Y float64
}

type Shape []Point

type Tree struct {
	xDimension      int
	yDimension      int
	presentCountMap map[int]int
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
	// solve part 1 here
	presentShapes := CreatePresentShapeMap(input)
	trees := CreateTreesWithRules(input)
	doesFitCount := 0
	for _, tree := range trees {
		doesFit := CalculateIfPresentsFitUnderTree(tree, presentShapes)
		if doesFit {
			doesFitCount = doesFitCount + 1
		}
	}
	return doesFitCount
}

func CalculateIfPresentsFitUnderTree(tree Tree, shapes map[int]Shape) bool {

	presentsLeftToPlace := tree.presentCountMap
	totalArea := tree.xDimension * tree.yDimension
	dotsNeeded := 0
	total3x3s := 0
	for i, numberOfThisPresentStillNeededToPlace := range presentsLeftToPlace {
		dotsNeeded = dotsNeeded + (len(shapes[i]) * numberOfThisPresentStillNeededToPlace)
		total3x3s += numberOfThisPresentStillNeededToPlace
	}
	if dotsNeeded > totalArea {
		return false
	} else if (total3x3s * 9) > totalArea {
		return false
	}

	return true

}

func CreateTreesWithRules(input string) []Tree {
	var trees []Tree
	for _, line := range strings.Split(input, "\n") {
		if !strings.Contains(line, "x") {
			continue
		}
		dimensioonThenRuleSplit := strings.Split(line, ":")
		dimension := dimensioonThenRuleSplit[0]
		dimensions := strings.Split(dimension, "x")
		x, _ := strconv.Atoi(dimensions[0])
		y, _ := strconv.Atoi(dimensions[1])
		rules := dimensioonThenRuleSplit[1]
		ruleMap := make(map[int]int)
		for i, rule := range strings.Split(rules, " ") {
			ruleMap[i], _ = strconv.Atoi(rule)
		}
		treeValue := Tree{
			xDimension:      x,
			yDimension:      y,
			presentCountMap: ruleMap,
		}
		trees = append(trees, treeValue)
	}
	return trees
}

func CreatePresentShapeMap(input string) map[int]Shape {
	presentShapes := make(map[int]Shape)
	for _, section := range strings.Split(input, "\n\n") {
		if strings.Contains(section, "x") {
			continue
		}
		presentIndex := 0
		indexThenPresent := strings.Split(section, ":")
		presentIndex, _ = strconv.Atoi(indexThenPresent[0])
		presentShape := indexThenPresent[1]
		present := Shape{}
		for yIndex, presentShapeLine := range strings.Split(presentShape, "\n") {
			if !strings.Contains(presentShapeLine, ".") && !strings.Contains(presentShapeLine, "#") {
				continue
			}
			for xValue, char := range presentShapeLine {
				if char == '#' {
					present = append(present, Point{X: float64(xValue), Y: float64(yIndex - 1)})
				}
			}
		}
		presentShapes[presentIndex] = present
	}
	return presentShapes
}
