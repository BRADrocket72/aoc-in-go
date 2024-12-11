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
		return "not implemented"
	}
	ret := 0
	trail_heads := make([]image.Point, 0)
	x_max := 0
	y_max := 0
	grid := map[image.Point]int{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		y_max = y_max + 1
		for x, r := range s {
			x_max = x_max + 1
			string1 := string(r)
			r_int, _ := strconv.Atoi(string1)
			grid[image.Point{x, y}] = r_int

			if r_int == 0 {
				trail_heads = append(trail_heads, image.Point{x, y})
			}

		}
	}

	for _, trail_head := range trail_heads {
		peaks := make([]image.Point, 0)
		accessed_points := make([]image.Point, 0)
		points_to_traverse := make([]image.Point, 0)

		points_to_traverse = append(points_to_traverse, trail_head)
		fmt.Print("new trail head")

		for len(points_to_traverse) > 0 {
			current_coords := points_to_traverse[0]

			currentPoint := grid[points_to_traverse[0]]
			if current_coords.X > 1 {
				test_coords := image.Point{current_coords.X - 1, current_coords.Y}
				test_point := grid[test_coords]
				if test_point == 0 {
					peaks = append(peaks, test_coords)
				}
				if currentPoint-test_point == 1 || currentPoint-test_point == -1 {
					is_found := false
					for _, travelled_Point := range accessed_points {
						if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
							is_found = true
						}
					}
					for _, travelled_Point := range points_to_traverse {
						if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
							is_found = true
						}
					}
					if !is_found {
						points_to_traverse = append(points_to_traverse, test_coords)
					}
				}
				if current_coords.X < x_max {
					test_coords := image.Point{current_coords.X + 1, current_coords.Y}
					test_point := grid[test_coords]
					if test_point == 0 {
						peaks = append(peaks, test_coords)
					}
					if currentPoint-test_point == 1 || currentPoint-test_point == -1 {
						is_found := false
						for _, travelled_Point := range accessed_points {
							if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
								is_found = true
							}
						}
						for _, travelled_Point := range points_to_traverse {
							if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
								is_found = true
							}
						}
						if !is_found {
							points_to_traverse = append(points_to_traverse, test_coords)
						}
					}
				}

				if current_coords.Y > 1 {
					test_coords := image.Point{current_coords.X, current_coords.Y - 1}
					test_point := grid[test_coords]
					if test_point == 0 {
						peaks = append(peaks, test_coords)
					}
					if currentPoint-test_point == 1 || currentPoint-test_point == -1 {
						is_found := false
						for _, travelled_Point := range accessed_points {
							if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
								is_found = true
							}
						}
						for _, travelled_Point := range points_to_traverse {
							if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
								is_found = true
							}
						}
						if !is_found {
							points_to_traverse = append(points_to_traverse, test_coords)
						}
					}
				}

				if current_coords.Y < y_max {
					test_coords := image.Point{current_coords.X, current_coords.Y + 1}
					test_point := grid[test_coords]
					if test_point == 0 {
						peaks = append(peaks, test_coords)
					}
					if currentPoint-test_point == 1 || currentPoint-test_point == -1 {
						is_found := false
						for _, travelled_Point := range accessed_points {
							if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
								is_found = true
							}
						}
						for _, travelled_Point := range points_to_traverse {
							if travelled_Point.X == test_coords.X && travelled_Point.Y == test_coords.Y {
								is_found = true
							}
						}
						if !is_found {
							points_to_traverse = append(points_to_traverse, test_coords)
						}
					}
				}
			}
			fmt.Printf("%d\n", len(points_to_traverse))
			fmt.Print(len("\n\n"))

			points_to_traverse = points_to_traverse[1:]

			accessed_points = append(accessed_points, current_coords)

		}

		ret = ret + len(peaks)
	}
	// solve part 1 here
	return ret
}
