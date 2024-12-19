package main

import (
	"fmt"
	"strconv"
)

func main() {
	run(false, "")
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
	init := []int{125, 17}
	rocks, _ := splitRocks(init, 1)
	fmt.Print("\n\n :Dfoliasjdfjaslkdjf:     ")

	fmt.Print(len(rocks))
	return len(rocks)

}

func splitRocks(rocks []int, depth int) (split_rocks []int, is_complete bool) {

	returnRocks := make([]int, 0)
	if depth >= 25 {
		return rocks, true
	}
	fmt.Printf("%d: %d", depth, len(rocks))
	to_split := []int{rock}

	for _, rock := range rocks {
		to_split := []int{rock}

		split := false
		num_string := strconv.Itoa(rock)

		if rock == 0 {
			rock = 1
			split = true
			to_split = append(to_split, rock)
		}
		if !split && len(num_string)%2 == 0 {
			midPoint := len(num_string) / 2
			rock_1 := num_string[0:midPoint]
			rock_2 := num_string[midPoint:]
			rock_1_int, err1 := strconv.Atoi(rock_1)
			rock_2_int, err2 := strconv.Atoi(rock_2)
			if err1 != nil {
				fmt.Print(err1)
			}
			if err2 != nil {
				fmt.Print(err2)
			}
			to_split := []int{rock_1_int, rock_2_int}
			split = true

		}
		if !split {
			rock = rock * 2024
			to_split := []int{rock}

		}
	}
	return splitRocks(to_split, depth+1)

}

// func splitRock(rock int, depth int) (split_rocks []int, is_complete bool) {

// 	returnRocks := make([]int,0)
// 	if depth == 25 {
// 		return returnRocks, true
// 	}

// 	split := false
// 	num_string := strconv.Itoa(rock)

// 	if rock == 0 {
// 		rock = 1
// 		split = true
// 		results,is_done := splitRock(rock, depth+1)
// 		if is_done{
// 			return result,true
// 		}
// 		for _,result : range results{
// 			splitRock()
// 		}
// 		returnRocks = append(returnRocks, result)
// 	}
// 	if !split && len(num_string)%2 == 0 {
// 		midPoint := len(num_string) / 2
// 		rock_1 := num_string[0:midPoint]
// 		rock_2 := num_string[midPoint:]
// 		rock_1_int := strconv.Atoi(rock_1)
// 		rock_2_int := strconv.Atoi(rock_1)

// 		split = true
// 		splitRock(rock_1_int, depth+1)
// 		splitRock(rock_2_int, depth+1)
// 		return
// 	}
// 	if !split {
// 		rock = rock * 2024
// 		splitRock(rock, depth+1)
// 	}
// }
