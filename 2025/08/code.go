package main

import (
	"math"
	"sort"
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
		junctionBoxes := CreateJunctionBoxes(input)
		circuits := makeCircuits(junctionBoxes)
		junctionBoxes = SetNearestNeighbor(junctionBoxes)
		mappedPairs := make([]Pair, 0)
		result := 0
		isTest := len(circuits) < 100
		if isTest {
			for {
				junctionBoxes, circuits, mappedPairs, result = MergeClosestJunctionBoxesPt2(junctionBoxes, circuits, mappedPairs)
				if result > 0 {
					return result
				}
			}
		} else {
			for {
				junctionBoxes, circuits, mappedPairs, result = MergeClosestJunctionBoxesPt2(junctionBoxes, circuits, mappedPairs)
				if result > 0 {
					return result
				}
			}
		}

	}
	// solve part 1 here

	junctionBoxes := CreateJunctionBoxes(input)
	circuits := makeCircuits(junctionBoxes)
	junctionBoxes = SetNearestNeighbor(junctionBoxes)
	mappedPairs := make([]Pair, 0)
	isTest := len(circuits) < 100
	if isTest {
		for i := 0; i < 10; i++ {
			junctionBoxes, circuits, mappedPairs = MergeClosestJunctionBoxes(junctionBoxes, circuits, mappedPairs)
		}

	} else {
		for i := 0; i < 1000; i++ {
			junctionBoxes, circuits, mappedPairs = MergeClosestJunctionBoxes(junctionBoxes, circuits, mappedPairs)
		}
	}

	return findThreeLargestCircuits(circuits)

}

func findThreeLargestCircuits(circuits map[int][]JunctionBox) int {
	// 1. Create a slice to hold all the circuit sizes
	sizes := make([]int, 0, len(circuits))
	for _, boxes := range circuits {
		sizes = append(sizes, len(boxes))
	}

	// 2. Sort the sizes in descending order
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	// 3. Sum the first three elements (handling cases with fewer than 3 circuits)
	total := 1
	limit := 3
	if len(sizes) < limit {
		limit = len(sizes)
	}

	for i := 0; i < limit; i++ {
		total = total * sizes[i]
	}

	return total
}

type Pair struct {
	to   int
	from int
}

func MergeClosestJunctionBoxes(junctionBoxes map[int]JunctionBox, circuits map[int][]JunctionBox, mappedPair []Pair) (map[int]JunctionBox, map[int][]JunctionBox, []Pair) {
	closestDistance := math.MaxFloat64
	closestPointA := junctionBoxes[0]
	closestPointB := junctionBoxes[1]
	for _, box := range junctionBoxes {
		if box.closestNeighborDistance < closestDistance {
			closestDistance = box.closestNeighborDistance
			closestPointA = box
			closestPointB = junctionBoxes[box.closestNeighborId]
		}
	}
	mappedPair = append(mappedPair, Pair{
		to:   closestPointA.id,
		from: closestPointB.id,
	})

	circuitIdBoxAIsIn := math.MaxInt
	circuitIdBoxBIsIn := math.MaxInt

	for circuitId, circuit := range circuits {
		for _, boxInCircuit := range circuit {
			if boxInCircuit.id == closestPointA.id {
				circuitIdBoxAIsIn = circuitId
			}
			if boxInCircuit.id == closestPointB.id {
				circuitIdBoxBIsIn = circuitId
			}
		}
	}
	if circuitIdBoxAIsIn == circuitIdBoxBIsIn {
		junctionBoxes = SetNearestNeighborForPoint(closestPointA, junctionBoxes, mappedPair)
		junctionBoxes = SetNearestNeighborForPoint(closestPointB, junctionBoxes, mappedPair)
		return junctionBoxes, circuits, mappedPair
	}

	newCircuit := append(circuits[circuitIdBoxAIsIn], circuits[circuitIdBoxBIsIn]...)
	circuits[circuitIdBoxAIsIn] = newCircuit
	delete(circuits, circuitIdBoxBIsIn)
	junctionBoxes = SetNearestNeighborForPoint(closestPointA, junctionBoxes, mappedPair)
	junctionBoxes = SetNearestNeighborForPoint(closestPointB, junctionBoxes, mappedPair)
	return junctionBoxes, circuits, mappedPair
	//junctionBoxes = SetNearestNeighborIgnoreCircuit(junctionBoxes, circuits)
	//return junctionBoxes, circuits, 0
}

func MergeClosestJunctionBoxesPt2(junctionBoxes map[int]JunctionBox, circuits map[int][]JunctionBox, mappedPair []Pair) (map[int]JunctionBox, map[int][]JunctionBox, []Pair, int) {
	result := 0
	closestDistance := math.MaxFloat64
	closestPointA := junctionBoxes[0]
	closestPointB := junctionBoxes[1]
	for _, box := range junctionBoxes {
		if box.closestNeighborDistance < closestDistance {
			closestDistance = box.closestNeighborDistance
			closestPointA = box
			closestPointB = junctionBoxes[box.closestNeighborId]
		}
	}
	mappedPair = append(mappedPair, Pair{
		to:   closestPointA.id,
		from: closestPointB.id,
	})

	circuitIdBoxAIsIn := math.MaxInt
	circuitIdBoxBIsIn := math.MaxInt

	for circuitId, circuit := range circuits {
		for _, boxInCircuit := range circuit {
			if boxInCircuit.id == closestPointA.id {
				circuitIdBoxAIsIn = circuitId
			}
			if boxInCircuit.id == closestPointB.id {
				circuitIdBoxBIsIn = circuitId
			}
		}
	}
	if circuitIdBoxAIsIn == circuitIdBoxBIsIn {
		junctionBoxes = SetNearestNeighborForPoint(closestPointA, junctionBoxes, mappedPair)
		junctionBoxes = SetNearestNeighborForPoint(closestPointB, junctionBoxes, mappedPair)
		return junctionBoxes, circuits, mappedPair, 0
	}
	if len(circuits) == 2 {
		result = closestPointA.x * closestPointB.x

	}

	newCircuit := append(circuits[circuitIdBoxAIsIn], circuits[circuitIdBoxBIsIn]...)
	circuits[circuitIdBoxAIsIn] = newCircuit
	delete(circuits, circuitIdBoxBIsIn)
	junctionBoxes = SetNearestNeighborForPoint(closestPointA, junctionBoxes, mappedPair)
	junctionBoxes = SetNearestNeighborForPoint(closestPointB, junctionBoxes, mappedPair)
	return junctionBoxes, circuits, mappedPair, result
	//junctionBoxes = SetNearestNeighborIgnoreCircuit(junctionBoxes, circuits)
	//return junctionBoxes, circuits, 0
}

func makeCircuits(junctionBoxes map[int]JunctionBox) map[int][]JunctionBox {
	circuits := make(map[int][]JunctionBox)

	for _, box := range junctionBoxes {
		circuits[box.id] = append(circuits[box.id], box)
	}

	return circuits
}

func SetNearestNeighbor(junctionBoxes map[int]JunctionBox) map[int]JunctionBox {
	for _, junctionBoxA := range junctionBoxes {
		shortestDistance := math.MaxFloat64
		nearestNeighborId := -1
		for _, junctionBoxB := range junctionBoxes {
			aToBDistance := CalcDistanceBetweenPoints(junctionBoxA, junctionBoxB)
			if aToBDistance < shortestDistance {
				shortestDistance = aToBDistance
				nearestNeighborId = junctionBoxB.id
			}
			junctionBoxA.closestNeighborDistance = shortestDistance
			junctionBoxA.closestNeighborId = nearestNeighborId
			junctionBoxes[junctionBoxA.id] = junctionBoxA
		}
	}
	return junctionBoxes
}

func SetNearestNeighborForPoint(boxToUpdate JunctionBox, junctionBoxes map[int]JunctionBox, mappedPair []Pair) map[int]JunctionBox {
	for _, junctionBoxA := range junctionBoxes {
		if junctionBoxA.id != boxToUpdate.id {
			continue
		}
		shortestDistance := math.MaxFloat64
		nearestNeighborId := -1
		for _, junctionBoxB := range junctionBoxes {
			isAlreadyPaired := false
			for _, pair := range mappedPair {
				if pair.to == junctionBoxA.id && pair.from == junctionBoxB.id {
					isAlreadyPaired = true
				} else if pair.to == junctionBoxB.id && pair.from == junctionBoxA.id {
					isAlreadyPaired = true
				}
			}
			if isAlreadyPaired {
				continue
			}
			aToBDistance := CalcDistanceBetweenPoints(junctionBoxA, junctionBoxB)
			if aToBDistance < shortestDistance {
				shortestDistance = aToBDistance
				nearestNeighborId = junctionBoxB.id
			}
		}
		boxToUpdate.closestNeighborDistance = shortestDistance
		boxToUpdate.closestNeighborId = nearestNeighborId
	}
	for currentJunctionBoxIndex, junctionBoxA := range junctionBoxes {
		if junctionBoxA.id == boxToUpdate.id {
			junctionBoxes[currentJunctionBoxIndex] = boxToUpdate
		}
	}
	return junctionBoxes
}

//func SetNearestNeighborIgnoreCircuit(junctionBoxes map[int]junctionBox, circuits [][]junctionBox) map[int]junctionBox {
//	for _, junctionBoxA := range junctionBoxes {
//		shortestDistance := math.MaxFloat64
//		nearestNeighborId := math.MinInt
//		boxesInCircuit := make([]int, 0)
//		for _, circuit := range circuits {
//			junctionAInThisCircuit := false
//
//			for _, boxInCircuit := range circuit {
//				for _, boxInCircuitB := range circuit {
//					if boxInCircuitB.id == junctionBoxA.id {
//						junctionAInThisCircuit = true
//					}
//				}
//				if junctionAInThisCircuit {
//					boxesInCircuit = append(boxesInCircuit, boxInCircuit.id)
//				}
//
//			}
//		}
//		for _, junctionBoxB := range junctionBoxes {
//			skip := false
//			for _, boxToSkip := range boxesInCircuit {
//				if boxToSkip == junctionBoxB.id {
//					skip = true
//				}
//			}
//			if skip {
//				continue
//			}
//			aToBDistance := CalcDistanceBetweenPoints(junctionBoxA, junctionBoxB)
//			if aToBDistance < shortestDistance {
//				shortestDistance = aToBDistance
//				nearestNeighborId = junctionBoxB.id
//			}
//			junctionBoxA.closestNeighborDistance = shortestDistance
//			junctionBoxA.closestNeighborId = nearestNeighborId
//			junctionBoxes[junctionBoxA.id] = junctionBoxA
//		}
//	}
//	return junctionBoxes
//}

func CreateJunctionBoxes(input string) map[int]JunctionBox {
	junctionBoxes := make(map[int]JunctionBox)

	for pointId, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		coords := strings.Split(line, ",")
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		newPoint := JunctionBox{
			id:                      pointId,
			x:                       x,
			y:                       y,
			z:                       z,
			closestNeighborId:       math.MaxInt,
			closestNeighborDistance: math.MaxInt,
		}

		junctionBoxes[pointId] = newPoint
	}
	return junctionBoxes
}

func CalcDistanceBetweenPoints(boxA, boxB JunctionBox) float64 {
	if boxA.id == boxB.id {
		return math.MaxFloat64
	}

	dx := float64(boxA.x - boxB.x)
	dy := float64(boxA.y - boxB.y)
	dz := float64(boxA.z - boxB.z)

	sumOfSquares := dx*dx + dy*dy + dz*dz

	return math.Sqrt(sumOfSquares)
}

type JunctionBox struct {
	id                      int
	x                       int
	y                       int
	z                       int
	closestNeighborId       int
	closestNeighborDistance float64
}
