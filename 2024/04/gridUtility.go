package main

import (
	"image"
	"strings"
)

func createGrid(input string) map[image.Point]rune {
	grid := map[image.Point]rune{}
	for y, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for x, r := range s {
			grid[image.Point{x, y}] = r
		}
	}

	return grid
}

func getAdjacentStrings(grid map[image.Point]rune, p image.Point, l int) []string {
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
