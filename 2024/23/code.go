package main

import (
	"fmt"
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
	// solve part 1 here
	lines := strings.Split(input, "\n")
	graph := NewGraph()

	for _, line := range lines {
		connection_string := strings.Split(line, "-")
		if node, ok := graph.nodes[connection_string[0]]; !ok {

			for _, connection := range node {
				if connection != connection_string[1] {
					graph.AddEdge(connection_string[0], connection_string[1])
					graph.AddEdge(connection_string[1], connection_string[0])
				}

			}
		}
	}

	fmt.Print(graph.nodes)
	return 4
}

type Graph struct {
	nodes map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string][]string),
	}
}

func (g *Graph) AddNode(node string) {
	if _, ok := g.nodes[node]; !ok {
		g.nodes[node] = []string{}
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.AddNode(from)
	g.AddNode(to)
	g.nodes[from] = append(g.nodes[from], to)
}

func (g *Graph) Neighbors(node string) []string {
	return g.nodes[node]
}
