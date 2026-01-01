package main

import (
	"fmt"
	"image"
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
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	points := make([]image.Point, 0)
	for _, line := range strings.Split(input, "\n") {
		lineSplit := strings.Split(line, ",")
		xInt, _ := strconv.Atoi(lineSplit[0])
		yInt, _ := strconv.Atoi(lineSplit[1])
		points = append(points, image.Point{X: xInt, Y: yInt})

	}
	largestArea := float64(0)
	for _, pointA := range points {
		for _, pointB := range points {
			result := CalcArea(pointA, pointB)
			if result > largestArea {
				largestArea = result
			}
		}
	}
	fmt.Printf("%%f output: %f\n", largestArea)

	return largestArea
}

func CalcArea(pointA image.Point, pointB image.Point) float64 {
	xDiff := math.Abs(float64(pointA.X-pointB.X)) + 1
	yDiff := math.Abs(float64(pointA.Y-pointB.Y)) + 1
	result := xDiff * yDiff
	return result
}

// 4529341988 too low
// 476173683 too low
// 571408419 too low
//5714084190
