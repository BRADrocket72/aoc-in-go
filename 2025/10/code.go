package main

import (
	"regexp"
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
		machines := make([]JoltageMachine, 0)
		for _, line := range strings.Split(input, "\n") {
			desiredJoltage := GetJoltagesFromLine(line)
			desiredLights := GetDesiredLightsFromLine(line)
			functions := GetFunctionsFromLine(line, desiredLights)
			machine := JoltageMachine{
				desiredJoltage: desiredJoltage,
				functions:      functions,
			}
			machines = append(machines, machine)
		}
		totalPresses := 0
		for _, machine := range machines {
			presses := CalculateButtonPressesPt2(machine)
			totalPresses += presses
		}
		return totalPresses
	}
	// solve part 1 here
	machines := make([]LightMachine, 0)
	for _, line := range strings.Split(input, "\n") {
		desiredLights := GetDesiredLightsFromLine(line)
		functions := GetFunctionsFromLine(line, desiredLights)
		machine := LightMachine{
			desiredLights: desiredLights,
			functions:     functions,
		}
		machines = append(machines, machine)
	}
	totalPresses := 0
	for _, machine := range machines {
		presses := CalculateButtonPresses(machine)
		totalPresses += presses
	}
	return totalPresses
}

func CalculateButtonPressesPt2(machine JoltageMachine) int {
	turnCount := 1
	lastStepStates := make([][]int, 0)
	for _, function := range machine.functions {
		emptyStep := make([]int, len(machine.desiredJoltage))
		nextStep := addJoltage(emptyStep, function)
		lastStepStates = append(lastStepStates, nextStep)
	}
	for {
		for _, lastStepState := range lastStepStates {
			found, _ := checkJoltage(lastStepState, machine.desiredJoltage)
			if found {
				return turnCount
			}
		}

		nextStepStates := make([][]int, 0)

		for _, lastStepState := range lastStepStates {
			for _, function := range machine.functions {
				nextStep := addJoltage(lastStepState, function)
				hasStepAlready := false
				isGoalExceeded := false
				for _, nextStepState := range nextStepStates {
					foundStep := checkIntArraysForEquality(nextStepState, nextStep)
					if foundStep {
						foundStepIsgoal, goalExceeded := checkJoltage(nextStep, machine.desiredJoltage)
						hasStepAlready = true
						isGoalExceeded = goalExceeded
						if !foundStepIsgoal {
							turnCount++
							return turnCount
						}
					}
				}
				if !hasStepAlready || isGoalExceeded {
					nextStepStates = append(nextStepStates, nextStep)

				}
			}
		}
		turnCount++
		lastStepStates = nextStepStates
	}
}

func checkJoltage(a, goal []int) (found bool, exceeds bool) {
	isFound := true
	for i, joltageVal := range a {
		if joltageVal > goal[i] {
			exceeds = true
		}
		if joltageVal != goal[i] {
			isFound = false
		}
	}
	return isFound, exceeds
}

func addJoltage(state []int, function []bool) []int {
	stateCopy := state
	for i, value := range function {
		if value {
			stateCopy[i] = stateCopy[i] + 1
		}

	}
	return stateCopy
}

func CalculateButtonPresses(machine LightMachine) int {
	turnCount := 1
	lastStepFunctions := machine.functions
	for {
		if turnCount > 10 {
			panic("too many turns")
		}
		for _, lastStepFunction := range lastStepFunctions {
			found := checkArraysForEquality(lastStepFunction, machine.desiredLights)
			if found {
				return turnCount
			}
		}
		nextStepFunctions := make([][]bool, 0)
		for _, lastStepFunction := range lastStepFunctions {
			for _, function := range machine.functions {
				nextStep := addFunctions(lastStepFunction, function)
				hasStepAlready := false
				for _, nextStepFunction := range nextStepFunctions {
					foundStep := checkArraysForEquality(nextStepFunction, nextStep)
					if foundStep {
						foundStepIsgoal := checkArraysForEquality(nextStep, machine.desiredLights)
						hasStepAlready = true
						if foundStepIsgoal {
							turnCount++
							return turnCount
						}
					}
				}
				if !hasStepAlready {
					nextStepFunctions = append(nextStepFunctions, nextStep)

				}
			}
		}
		turnCount++
		lastStepFunctions = nextStepFunctions
	}
}

func addFunctions(a, b []bool) []bool {
	result := make([]bool, len(a))

	for i := 0; i < len(a); i++ {
		if a[i] && b[i] {
			result[i] = false
		} else if a[i] || b[i] {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
}

func checkArraysForEquality(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func checkIntArraysForEquality(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type LightMachine struct {
	desiredLights []bool
	functions     [][]bool
}
type JoltageMachine struct {
	desiredJoltage []int
	functions      [][]bool
}

func GetFunctionsFromLinePt2(line string, lights []bool) [][]bool {
	functions := make([][]bool, 0)
	// The regex pattern explained:
	// \\( : Matches a literal opening parenthesis.
	// ([^)]+) : This is the capturing group (group 1).
	//          [^)]+ means match one or more characters that are NOT a closing parenthesis.
	// \\) : Matches a literal closing parenthesis.
	re := regexp.MustCompile(`\(([^)]+)\)`)

	// Find all matches and submatches. The second argument, -1, means find all occurrences.
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		functionResult := make([]bool, len(lights))
		stringSplit := strings.Split(match[1], ",")
		for _, s := range stringSplit {
			intValue, _ := strconv.Atoi(s)
			functionResult[intValue] = true
		}
		functions = append(functions, functionResult)
	}
	return functions
}
func GetFunctionsFromLine(line string, lights []bool) [][]bool {
	functions := make([][]bool, 0)
	// The regex pattern explained:
	// \\( : Matches a literal opening parenthesis.
	// ([^)]+) : This is the capturing group (group 1).
	//          [^)]+ means match one or more characters that are NOT a closing parenthesis.
	// \\) : Matches a literal closing parenthesis.
	re := regexp.MustCompile(`\(([^)]+)\)`)

	// Find all matches and submatches. The second argument, -1, means find all occurrences.
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		functionResult := make([]bool, len(lights))
		stringSplit := strings.Split(match[1], ",")
		for _, s := range stringSplit {
			intValue, _ := strconv.Atoi(s)
			functionResult[intValue] = true
		}
		functions = append(functions, functionResult)
	}
	return functions
}

func GetJoltagesFromLine(line string) []int {
	// The regex pattern explained:
	// \\( : Matches a literal opening parenthesis.
	// ([^)]+) : This is the capturing group (group 1).
	//          [^)]+ means match one or more characters that are NOT a closing parenthesis.
	// \\) : Matches a literal closing parenthesis.
	re := regexp.MustCompile("\\{(.*?)\\}")

	// Find all matches and submatches. The second argument, -1, means find all occurrences.
	matches := re.FindAllStringSubmatch(line, -1)
	joltages := make([]int, 0)

	for _, match := range matches {
		stringSplit := strings.Split(match[1], ",")
		for _, s := range stringSplit {
			intValue, _ := strconv.Atoi(s)
			joltages = append(joltages, intValue)
		}
	}
	return joltages
}

func GetDesiredLightsFromLine(line string) []bool {
	allData := strings.Split(line, "]")
	lights := allData[0][1:]
	lightArray := strings.Split(lights, "")
	desiredLight := make([]bool, len(lightArray))
	for i, codedLight := range lightArray {
		if codedLight == "." {
			desiredLight[i] = false
		} else if codedLight == "#" {
			desiredLight[i] = true
		}
	}
	return desiredLight
}
