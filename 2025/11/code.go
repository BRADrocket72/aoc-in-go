package main

import (
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Device struct {
	Name    string
	Outputs []string
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
		allDevices, _ := CreateAllDevicesPt2(input)
		emptyPath := make([]string, 0)
		emptyPath = append(emptyPath, "svr")
		_, totalCount := navigateFromDevicePt2(allDevices["svr"], allDevices, emptyPath, 0)
		return totalCount
	}
	// solve part 1 here

	allDevices, youIndex := CreateAllDevices(input)
	totalCount := navigateFromDevice(allDevices[youIndex], allDevices, 0)
	return totalCount
}

func navigateFromDevice(device Device, allDevices []Device, currentOutCount int) int {
	for _, output := range device.Outputs {
		if output == "out" {
			currentOutCount++
			return currentOutCount
		}
		for _, possibleOutPutDevice := range allDevices {
			if possibleOutPutDevice.Name == output {
				currentOutCount = navigateFromDevice(possibleOutPutDevice, allDevices, currentOutCount)
			}
		}
	}
	return currentOutCount
}

func CreateAllDevices(input string) ([]Device, int) {
	allDevices := []Device{}
	youIndex := 0
	for _, line := range strings.Split(input, "\n") {
		linesplit := strings.Split(line, ":")
		deviceName := linesplit[0]
		connections := strings.Split(linesplit[1], " ")
		filteredConnections := make([]string, 0)
		for _, connection := range connections {
			if len(connection) == 3 {
				filteredConnections = append(filteredConnections, connection)
			}
		}
		device := Device{Name: deviceName, Outputs: filteredConnections}
		if deviceName == "you" {
			youIndex = len(allDevices)
		}
		allDevices = append(allDevices, device)
	}
	return allDevices, youIndex
}

func navigateFromDevicePt2(device Device, allDevices map[string]Device, pathSoFar []string, currentOutCount int) ([]string, int) {
	pathBeforeThisStop := pathSoFar
	if device.Name == "out" {
		hitDac := false
		hitFft := false
		for _, path := range pathSoFar {
			if path == "dac" {
				hitDac = true
			}
			if path == "fft" {
				hitFft = true
			}
		}
		if hitDac && hitFft {
			currentOutCount++
			return pathSoFar, currentOutCount
		} else {
			return make([]string, 0), currentOutCount
		}
	}
	for _, pathsVisited := range pathSoFar {
		if pathsVisited == device.Name {
			return make([]string, 0), currentOutCount
		}
	}
	for _, output := range device.Outputs {
		pathBeforeThisStop = append(pathBeforeThisStop, output)
		for key, possibleOutPutDevice := range allDevices {
			if key == output {
				_, currentPathOutCount := navigateFromDevicePt2(possibleOutPutDevice, allDevices, pathBeforeThisStop, currentOutCount)
				if currentPathOutCount > 0 {
					currentOutCount = currentOutCount + 1
				}
			}
		}
		if output == "out" {
			outDevice := Device{Name: output, Outputs: []string{}}
			_, currentPathOutCount := navigateFromDevicePt2(outDevice, allDevices, pathBeforeThisStop, currentOutCount)
			if currentPathOutCount > 0 {
				currentOutCount = currentOutCount + 1
			}
			return make([]string, 0), currentOutCount
		}

	}
	return pathSoFar, currentOutCount
}

func CreateAllDevicesPt2(input string) (map[string]Device, int) {
	devices := make(map[string]Device)
	svrIndex := 0
	for _, line := range strings.Split(input, "\n") {
		linesplit := strings.Split(line, ":")
		deviceName := linesplit[0]
		connections := strings.Split(linesplit[1], " ")
		filteredConnections := make([]string, 0)
		for _, connection := range connections {
			if len(connection) == 3 {
				filteredConnections = append(filteredConnections, connection)
			}
		}
		device := Device{Name: deviceName, Outputs: filteredConnections}
		if deviceName == "svr" {
			svrIndex = len(devices)
		}
		devices[deviceName] = device
	}
	return devices, svrIndex
}
