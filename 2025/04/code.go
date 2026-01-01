package main

import (
	"image"
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
		grid, maxX, maxY := fillGrid(input)
		totalResults := 0
		for range 100 {
			results := 0
			results, grid = checkSurroundingPoints(grid, maxX, maxY)
			totalResults += results
			if results == 0 {
				return totalResults
			}
		}
		return 11111
	}
	// solve part 1 here
	grid, maxX, maxY := fillGrid(input)
	results, _ := checkSurroundingPoints(grid, maxX, maxY)
	return results
}

func fillGrid(input string) (map[image.Point]string, int, int) {
	var grid = make(map[image.Point]string)
	maxX := 0
	maxY := 0
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			grid[image.Point{X: x, Y: y}] = string(char)
			maxX = x
		}
		maxY = y
	}
	return grid, maxX, maxY
}

func checkSurroundingPoints(grid map[image.Point]string, maxX int, maxY int) (int, map[image.Point]string) {
	newMap := make(map[image.Point]string)
	blockedLines := 0
	for y := 0; y <= maxY+1; y++ {
		for x := 0; x <= maxX+1; x++ {
			point := image.Point{X: x, Y: y}
			atSigns := 0
			//Going clockwize from top left
			if point.X-1 >= 0 && point.Y-1 >= 0 {
				if grid[image.Pt(point.X-1, point.Y-1)] == "@" {
					atSigns++
				}
			}
			if point.Y-1 >= 0 {
				if grid[image.Pt(point.X, point.Y-1)] == "@" {
					atSigns++
				}
			}
			if point.X+1 <= maxX && point.Y-1 >= 0 {
				if grid[image.Pt(point.X+1, point.Y-1)] == "@" {
					atSigns++
				}
			}
			if point.X+1 <= maxX {
				if grid[image.Pt(point.X+1, point.Y)] == "@" {
					atSigns++
				}
			}
			if point.X+1 <= maxX && point.Y+1 <= maxY {
				if grid[image.Pt(point.X+1, point.Y+1)] == "@" {
					atSigns++
				}
			}
			if point.Y+1 <= maxY {
				if grid[image.Pt(point.X, point.Y+1)] == "@" {
					atSigns++
				}
			}
			if point.X-1 <= maxX && point.Y+1 <= maxY {
				if grid[image.Pt(point.X-1, point.Y+1)] == "@" {
					atSigns++
				}
			}
			if point.X-1 <= maxX {
				if grid[image.Pt(point.X-1, point.Y)] == "@" {
					atSigns++
				}
			}
			newMap[point] = grid[point]
			if atSigns < 4 && grid[image.Pt(point.X, point.Y)] == "@" {
				blockedLines++
				newMap[point] = "."
			}
		}
	}
	return blockedLines, newMap
}
