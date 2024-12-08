package main

import (
	"fmt"
	"image"
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
		grid := createGrid(input)

		part2 := 0
		for p := range grid {
			part2 += strings.Count("AMAMASASAMAMAS", strings.Join(getAdjacentStrings(grid, p, 2)[4:], ""))
		}
		return part2
	}
	// solve part 1 here
	grid := map[image.Point]rune{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}

	adj := func(p image.Point, l int) []string {
		delta := []image.Point{
			{0, -1}, {1, 0}, {0, 1}, {-1, 0},
			{-1, -1}, {1, -1}, {1, 1}, {-1, 1},
		}

		words := make([]string, len(delta))
		for i, d := range delta {
			for n := range l {
				words[i] += string(grid[p.Add(d.Mul(n))])
			}
		}
		return words
	}

	part1 := 0
	for p := range grid {
		part1 += strings.Count(strings.Join(adj(p, 4), " "), "XMAS")
	}
	return part1
}

func testForward(lines []string) int {
	sum := 0
	for lineIndex, line := range lines {
		for charIndex := range line {
			if charIndex+4 < len(line) {
				testString := line[charIndex : charIndex+4]
				if testString == "XMAS" || testString == "SAMX" {
					sum++
					if line[charIndex] == 'X' {
						fmt.Printf("(%s,%s)", strconv.Itoa(lineIndex), strconv.Itoa(charIndex))
					}
				}
			}
		}
	}
	return sum
}

func testDownwardLines(lines []string) int {
	sum := 0
	for lineIndex, line := range lines {
		for charIndex := range line {
			if lineIndex+3 < len(line) {
				testString := string(lines[lineIndex][charIndex]) + string(lines[lineIndex+1][charIndex]) + string(lines[lineIndex+2][charIndex]) + string(lines[lineIndex+3][charIndex])

				if testString == "SAMX" || testString == "XMAS" {
					sum++
					if lines[lineIndex][charIndex] == 'X' {
						fmt.Printf("(%s,%s)", strconv.Itoa(lineIndex), strconv.Itoa(charIndex))
					}
				}
			}
		}
	}
	return sum

}

func testDiaganolDown(lines []string) int {
	sum := 0
	for lineIndex, line := range lines {
		for charIndex := range line {
			if lineIndex+3 < len(line) && charIndex+3 < len(line) {
				testString := string(lines[lineIndex][charIndex]) + string(lines[lineIndex+1][charIndex+1]) + string(lines[lineIndex+2][charIndex+2]) + string(lines[lineIndex+3][charIndex+3])

				if testString == "SAMX" || testString == "XMAS" {
					sum++
					if lines[lineIndex][charIndex] == 'X' {
						fmt.Printf("(%s,%s)", strconv.Itoa(lineIndex), strconv.Itoa(charIndex))
					}
				}
			}
		}
	}
	return sum
}

func testDiaganolUp(lines []string) int {
	sum := 0

	for lineIndex, line := range lines {
		for charIndex := range line {
			if lineIndex-4 > -2 && charIndex+3 < len(line) {
				testString := string(lines[lineIndex][charIndex]) + string(lines[lineIndex-1][charIndex+1]) + string(lines[lineIndex-2][charIndex+2]) + string(lines[lineIndex-3][charIndex+3])

				if testString == "SAMX" || testString == "XMAS" {
					sum++
					if lines[lineIndex][charIndex] == 'X' {
						fmt.Printf("(%s,%s)", strconv.Itoa(lineIndex), strconv.Itoa(charIndex))
					}

				}
			}
		}
	}
	return sum
}

// func testDiaganolUp(cols [][]string, lines []string) int {
// 	sum := 0
// 	for lineIndex, col := range cols {
// 		for charIndex := range col {
// 			if charIndex+4 < len(col) && lineIndex+4 < len(lines) {
// 				testString := cols[lineIndex][charIndex] + cols[lineIndex+1][charIndex+1] + cols[lineIndex+2][charIndex+2] + cols[lineIndex+3][charIndex+3]
// 				if testString == "XMAS" {
// 					fmt.Print(lineIndex)
// 					fmt.Print(cols[lineIndex])

// 					sum++
// 				}
// 				if testString == "SAMX" {
// 					fmt.Print(lineIndex)
// 					fmt.Print(cols[lineIndex])

// 					sum++
// 				}
// 			}
// 		}

// 	}
// 	return sum
// }
