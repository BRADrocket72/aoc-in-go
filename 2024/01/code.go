package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

		var lines = strings.Split(input, "\n")
		var list_one = [1000]int{}
		var list_two = [1000]int{}
		for i := 0; i < len(lines); i++ {
			var split_line = strings.Split(lines[i], "   ")
			list_one[i], _ = strconv.Atoi(split_line[0])
			list_two[i], _ = strconv.Atoi(split_line[1])
		}

		sort.Sort(sort.IntSlice(list_one[:]))
		sort.Sort(sort.IntSlice(list_two[:]))
		probScore := 0
		for i := 0; i < len(list_one); i++ {
			for j := 0; j < len(list_two); j++ {
				if list_one[i] == list_two[j] {
					probScore += list_one[i]
				}
			}
		}
		return probScore
	}
	// solve part 1 here

	var lines = strings.Split(input, "\n")
	var list_one = [1000]int{}
	var list_two = [1000]int{}
	for i := 0; i < len(lines); i++ {
		var split_line = strings.Split(lines[i], "   ")
		list_one[i], _ = strconv.Atoi(split_line[0])
		list_two[i], _ = strconv.Atoi(split_line[1])
	}

	sort.Sort(sort.IntSlice(list_one[:]))
	sort.Sort(sort.IntSlice(list_two[:]))
	sum := 0
	for i := 0; i < len(list_one); i++ {
		diff := list_one[i] - list_two[i]
		distance := diff
		if distance < 0 {
			distance = -distance
		}
		sum = sum + distance
	}
	return sum
}
