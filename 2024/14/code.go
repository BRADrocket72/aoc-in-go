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

type bot struct {
	position image.Point
	velocity image.Point
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		bots := make([]bot, 0)
		// solve part 1 here
		lines := strings.Split(input, "\n")
		for _, line := range lines {
			splitLine := strings.Split(line, " v=")
			position_string := strings.Split(splitLine[0], "p=")[1]
			x_and_y := strings.Split(position_string, ",")
			x, x_err := strconv.Atoi(x_and_y[0])
			y, y_err := strconv.Atoi(x_and_y[1])
			if x_err != nil || y_err != nil {
				fmt.Print(x_err)
				fmt.Print(y_err)

			}

			velocity_string := strings.Split(line, "v=")[1]
			v_xand_v_y := strings.Split(velocity_string, ",")
			v_x, _ := strconv.Atoi(v_xand_v_y[0])
			v_y, _ := strconv.Atoi(v_xand_v_y[1])

			bot := bot{
				position: image.Point{X: x, Y: y},
				velocity: image.Point{X: v_x, Y: v_y},
			}
			bots = append(bots, bot)
		}
		is_test := false
		if bots[0].position.X == 0 {
			is_test = true
		}

		x_max_inx := 100
		y_max_inx := 102
		if is_test {
			x_max_inx = 10
			y_max_inx = 6
		}
		seconds := 0
		for i := range 500000000000 {
			for i, bot := range bots {
				bot = moveBot(bot, x_max_inx, y_max_inx)
				bots[i] = bot
			}
			is_triangle_started := false
			levels := 1

			for _, bot1 := range bots {
				rigt := false
				left := false
				for _, bot2 := range bots {
					if bot2.position.X == bot1.position.X+levels && bot2.position.Y == bot1.position.Y+levels {
						rigt = true
					}
					if bot2.position.X == bot1.position.X-levels && bot2.position.Y == bot1.position.Y-levels {
						left = true
					}
				}
				if rigt && left {
					levels++
				}

			}
			if levels == 20 {
				seconds = i
			}
			if is_triangle_started {
				fmt.Print(i)
			}
		}
		return seconds
	}

	bots := make([]bot, 0)
	// solve part 1 here
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		splitLine := strings.Split(line, " v=")
		position_string := strings.Split(splitLine[0], "p=")[1]
		x_and_y := strings.Split(position_string, ",")
		x, _ := strconv.Atoi(x_and_y[0])
		y, _ := strconv.Atoi(x_and_y[1])

		velocity_string := strings.Split(line, "v=")[1]
		v_xand_v_y := strings.Split(velocity_string, ",")
		v_x, _ := strconv.Atoi(v_xand_v_y[0])
		v_y, _ := strconv.Atoi(v_xand_v_y[1])

		bot := bot{
			position: image.Point{X: x, Y: y},
			velocity: image.Point{X: v_x, Y: v_y},
		}
		bots = append(bots, bot)
	}
	is_test := false
	if bots[0].position.X == 0 {
		is_test = true
	}

	x_max_inx := 100
	y_max_inx := 102
	if is_test {
		x_max_inx = 10
		y_max_inx = 6
	}
	for i := range 500 {
		for i, bot := range bots {
			bot = moveBot(bot, x_max_inx, y_max_inx)
			bots[i] = bot
		}
		unique_points := make(map[image.Point]bool)
		are_all_unique := true
		for _, bot := range bots {
			if unique_points[bot.position] {
				are_all_unique = false
			} else {
				unique_points[bot.position] = true
			}
		}
		if are_all_unique {
			fmt.Print(i)
		}
	}
	ret := judgeQuandrant(bots, x_max_inx, y_max_inx)
	return ret

}

func moveBot(bot bot, x_max int, y_max int) (ret_bot bot) {
	bot.position.X = moveX(bot, x_max)
	bot.position.Y = moveY(bot, y_max)

	return bot

}

func moveX(bot bot, x_max int) int {
	new_pos := bot.position.X + bot.velocity.X
	if new_pos < 0 {
		new_pos = x_max + new_pos + 1
	}
	if new_pos > x_max {
		new_pos = new_pos - x_max - 1
	}
	return new_pos
}

func moveY(bot bot, y_max int) int {
	new_pos := bot.position.Y + bot.velocity.Y
	if new_pos < 0 {
		new_pos = y_max + new_pos + 1
	}
	if new_pos > y_max {
		new_pos = new_pos - y_max - 1
	}
	return new_pos

}

func judgeQuandrant(bots []bot, x_max int, y_max int) int {
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	mid_x := (x_max + 1) / 2
	mid_y := (y_max + 1) / 2

	for _, bot := range bots {
		if bot.position.X < mid_x && bot.position.Y < mid_y {
			q1++
		}
		if bot.position.X > mid_x && bot.position.Y < mid_y {
			q2++
		}
		if bot.position.X < mid_x && bot.position.Y > mid_y {
			q3++
		}
		if bot.position.X > mid_x && bot.position.Y > mid_y {
			q4++
		}
	}
	return q1 * q2 * q3 * q4

}
