package main

import (
	"log"
	"os"
	"reflect"
	"testing"
)

func TestGetDesiredLightsFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "FirstLine",
			args: args{"[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}"},
			want: []bool{false, true, true, false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDesiredLightsFromLine(tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDesiredLightsFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

var firstMachineDesiredLights = []bool{false, true, true, false}
var firstMachineFunctions = [][]bool{
	[]bool{false, false, false, true}, //3
	[]bool{false, true, false, true},  //1,3
	[]bool{false, false, true, false}, //2
	[]bool{false, false, true, true},  //2,3
	[]bool{true, false, true, false},  //0,2
	[]bool{true, true, false, false},  //0,1
}

func TestGetFunctionsFromLine(t *testing.T) {
	type args struct {
		line   string
		lights []bool
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "FirstLine",
			args: args{line: "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}", lights: []bool{false, true, true, false}},
			want: [][]bool{
				[]bool{false, false, false, true}, //3
				[]bool{false, true, false, true},  //1,3
				[]bool{false, false, true, false}, //2
				[]bool{false, false, true, true},  //2,3
				[]bool{true, false, true, false},  //0,2
				[]bool{true, true, false, false},  //0,1
			},
		},
		{
			name: "ThirdLine",
			args: args{line: "[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}", lights: []bool{false, true, true, true, false, true}},
			want: [][]bool{
				[]bool{true, true, true, true, true, false},
				[]bool{true, false, false, true, true, false},
				[]bool{true, true, true, false, true, true},
				[]bool{false, true, true, false, false, false},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFunctionsFromLine(tt.args.line, tt.args.lights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFunctionsFromLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateButtonPresses(t *testing.T) {
	type args struct {
		machine LightMachine
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first machine",
			args: args{machine: LightMachine{
				desiredLights: firstMachineDesiredLights,
				functions:     firstMachineFunctions,
			}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateButtonPresses(tt.args.machine); got != tt.want {
				t.Errorf("CalculateButtonPresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part1Sample(t *testing.T) {
	filePath := "input-example.txt" // Replace with your file path

	// Read the entire file into a byte slice
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		// Log the error and exit if the file cannot be read
		log.Fatalf("failed to read file: %v", err)
	}

	// Convert the byte slice to a string
	fileContentString := string(contentBytes)

	value := run(false, fileContentString)
	if value != 7 {
		t.Error(value)
	}
}

func Test_part2Sample(t *testing.T) {
	filePath := "input-example.txt" // Replace with your file path

	// Read the entire file into a byte slice
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		// Log the error and exit if the file cannot be read
		log.Fatalf("failed to read file: %v", err)
	}

	// Convert the byte slice to a string
	fileContentString := string(contentBytes)

	value := run(true, fileContentString)
	if value != 33 {
		t.Error(value)
	}
}

func Test_addFunctions(t *testing.T) {
	type args struct {
		a []bool
		b []bool
	}
	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "os",
			args: args{[]bool{false, false, false, false}, []bool{false, false, false, false}},
			want: []bool{false, false, false, false},
		},
		{
			name: "1s",
			args: args{[]bool{true, true, true, true}, []bool{true, true, true, true}},
			want: []bool{false, false, false, false},
		},
		{
			name: "1s",
			args: args{[]bool{true, true, true, true}, []bool{false, true, false, true}},
			want: []bool{true, false, true, false},
		},
		{
			name: "1s",
			args: args{
				[]bool{false, false, false, true, false, false},
				[]bool{false, true, true, true, true, true}},
			want: []bool{false, true, true, false, true, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addFunctions(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addFunctions() = %v, want %v", got, tt.want)
			}
		})
	}
}
