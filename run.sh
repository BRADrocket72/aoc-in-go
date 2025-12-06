#!/bin/bash
set -euf -o pipefail

# functions
function echogrey() {
	echo -e "\033[0;90m$1\033[0m"
}

function template() {
	cat <<EOF
package main

import (
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
	return 42
}
EOF
}

export AOC_SESSION=53616c7465645f5f3cb7b8b5404f4f4357273f5c3e6344b19079852ba78c914f97200a36b4157293a9238c953642f3d3e100fa0ac453e948555de7fd836d349d

# two args YEAR and DAY
YEAR="${1:-}"
DAY="${2:-}"
if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
	echo "Usage: $0 <YEAR> <DAY>"
	exit 1
fi
# pad DAY to 2 digits
DAY=$(printf "%02d" $DAY)
DIR="./$YEAR/$DAY"
# create missing files as needed
if [ ! -d "$DIR" ]; then
	mkdir -p "$DIR"
	echogrey "Created directory $DIR"
fi
if [ ! -f "$DIR/code.go" ]; then
	template >"$DIR/code.go"
	echogrey "Created file code.go"
fi
# go run
cd "$DIR" && go run code.go
