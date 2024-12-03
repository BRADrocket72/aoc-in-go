package main

import (
	"fmt"
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
	var lines = strings.Split(input, "\n")
	var ret = 0

	for i := 0; i < len(lines); i++ {
		var is_safe = true
		var split_line = strings.Split(lines[i], " ")
		var first, _ = strconv.Atoi(split_line[0])
		var second, _ = strconv.Atoi(split_line[1])
		var is_increasing = first < second

		for j := 0; j < len(split_line); j++ {
			if j > 0 {
				var a, aerr = strconv.Atoi(split_line[j-1])
				var b, berr = strconv.Atoi(split_line[j])
				if aerr == nil && berr == nil {
					if is_increasing {
						if a > b || a == b || ((b - a) >= 4) {
							is_safe = false
							fmt.Print(split_line)
							fmt.Print(a)
							fmt.Print(b)
							fmt.Print(is_increasing)
							fmt.Print(split_line[0] <= split_line[1])

						}
					} else if !is_increasing {
						if a < b || a == b || ((a - b) >= 4) {
							is_safe = false
							fmt.Print(split_line)
							fmt.Print(a)
							fmt.Print(b)

						}
					}
				}
			}
		}
		if is_safe {
			ret += 1
		}
	}
	return ret

}
