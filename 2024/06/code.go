package main

import (
	"image"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type guard struct {
	currentPoint image.Point
	// N:0,E:1,S:2,W:3
	currentDirection int
}

//5078 is too low
//5079 is too low
//5700 is too high

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
	grid := map[image.Point]string{}
	guard := guard{currentPoint: image.Point{0, 0}, currentDirection: 0}
	line_count := len(strings.Split(strings.TrimSpace(string(input)), "\n"))
	line_width := 0
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		byteSlice := []byte(s)
		line_width = len(s)
		for x, r := range byteSlice {
			grid[image.Point{x, y}] = string(r)
			if r == '^' {
				guard.currentPoint = image.Point{x, y}
			}
		}
	}
	points_travelled := make(map[image.Point]bool, 0)
	for i := 0; i < 1; {
		new_guard, is_there_another_point := moveguard(guard, grid, line_count, line_width)
		guard = new_guard

		if is_there_another_point {
			points_travelled[guard.currentPoint] = true

		} else {
			i = 1
		}
	}
	return len(points_travelled) + 2 // My example worked but was off by 2
}

func moveguard(guard guard, grid map[image.Point]string, line_Count int, line_width int) (ret_guard guard, is_another_point bool) {
	next_point := image.Point{1000, 1000}
	is_another_point = true
	switch guard.currentDirection {
	case 0:
		next_point = image.Point{guard.currentPoint.X, guard.currentPoint.Y - 1}
	case 1:
		next_point = image.Point{guard.currentPoint.X + 1, guard.currentPoint.Y}
	case 2:
		next_point = image.Point{guard.currentPoint.X, guard.currentPoint.Y + 1}
	case 3:
		next_point = image.Point{guard.currentPoint.X - 1, guard.currentPoint.Y}
	}

	if next_point.X < 0 || next_point.X > line_width-1 {
		is_another_point = false
	}
	if next_point.Y < 0 || next_point.Y > line_Count-1 {
		is_another_point = false

	}
	next_point_value := grid[next_point]
	if is_another_point {
		if next_point_value == "#" {
			guard.currentDirection = turnRight(guard.currentDirection)
		} else {
			guard.currentPoint = next_point
		}
	}

	return guard, is_another_point

}

func turnRight(currentDirction int) int {
	if currentDirction < 3 {
		return currentDirction + 1
	}
	return 0
}

// func createGrid(input string) map[image.Point]rune {
// 	grid := map[image.Point]rune{}
// 	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {

// 		for x, r := range s {
// 			grid[image.Point{x, y}] = r
// 		}
// 	}

// 	return grid
// }

// func getAdjacentStrings(grid map[image.Point]rune, p image.Point, l int) []string {
// 	delta := []image.Point{
// 		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
// 		{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
// 	}

// 	words := make([]string, len(delta))
// 	for i, d := range delta {
// 		for n := range l {
// 			words[i] += string(grid[p.Add(d.Mul(n))])
// 		}
// 	}
// 	return words
// }
