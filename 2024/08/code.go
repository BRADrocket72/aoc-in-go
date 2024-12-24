// package main

// import (
// 	"fmt"
// 	"image"
// 	"math"
// 	"strings"

// 	"github.com/jpillora/puzzler/harness/aoc"
// )

// func main() {
// 	aoc.Harness(run)
// }

// // on code change, run will be executed 4 times:
// // 1. with: false (part1), and example input
// // 2. with: true (part2), and example input
// // 3. with: false (part1), and user input
// // 4. with: true (part2), and user input
// // the return value of each run is printed to stdout
// func run(part2 bool, input string) any {
// 	// when you're ready to do part 2, remove this "not implemented" block
// 	if part2 {
// 		return "not implemented"
// 	}
// 	// solve part 1 here
// 	freqs := make(map[rune][]image.Point, 0)
// 	lines := strings.Split(input, "\n")
// 	for y, line := range lines {
// 		for x, char := range line {
// 			if char == '.' {
// 				continue
// 			}
// 			_, exists := freqs[char]
// 			if exists {
// 				new_map := freqs[char]
// 				new_map = append(new_map, image.Pt(x, y))
// 				freqs[char] = new_map
// 			} else {
// 				new_map := make([]image.Point, 0)
// 				new_map = append(new_map, image.Pt(x, y))
// 				freqs[char] = new_map

// 			}
// 		}
// 	}
// 	anti_nodes := make(map[image.Point]bool, 0)
// 	for y, line := range lines {
// 		for x, char := range line {
// 			distancesFromEachFreq := getDistanceFromFreq(freqs, image.Pt(x, y))
// 			indexes_of_doubles := findDoubles(distancesFromEachFreq[char])
// 			for _, index := range indexes_of_doubles {
// 				anti_nodes[freqs[char][index]] = true
// 				if char == '0' {
// 					fmt.Print(distancesFromEachFreq['0'])

// 					fmt.Print(indexes_of_doubles)
// 				}
// 			}
// 		}

// 	}

// 	return len(anti_nodes)
// }

// func getDistanceFromFreq(freqs map[rune][]image.Point, point image.Point) map[rune][]float64 {

// 	ret_map := make(map[rune][]float64, 0)
// 	for freq, points := range freqs {
// 		distances := make([]float64, 0)
// 		for _, point1 := range points {
// 			distance := calcDistance(point, point1)
// 			distances = append(distances, distance)
// 		}
// 		ret_map[freq] = distances
// 	}
// 	return ret_map
// }

// func calcDistance(a image.Point, b image.Point) float64 {
// 	x_diff := a.X - b.X
// 	y_diff := a.Y - b.Y
// 	x_diff_float := math.Pow(float64(x_diff), 2)
// 	y_diff_float := math.Pow(float64(y_diff), 2)
// 	sum := x_diff_float + y_diff_float
// 	c := math.Sqrt(sum)
// 	return c
// }

// func findDoubles(distances []float64) (indexes_of_doubles []int) {
// 	indexes := make([]int, 0)
// 	for i := 0; i < len(distances); i++ {

// 		for j := 1; j < len(distances); j++ {
// 			if i == 0 || j == 0 {
// 				continue
// 			}
// 			if distances[i]*2 == distances[j] {
// 				indexes = append(indexes, j)
// 			}
// 		}
// 	}
// 	return indexes
// }

// //230 too low
