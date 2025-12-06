package main

import (
	"log"
	"os"
	"testing"
)

func Test_parseLineL(t *testing.T) {
	numbers := parseLineToNumberSplice("987")

	if numbers[0] != 9 || numbers[1] != 8 || numbers[2] != 7 {
		t.Error(numbers)
	}
}

func Test_GetTwoDigits(t *testing.T) {
	numbers := parseLineToNumberSplice("987654321111111")
	result := getHighestCombinedNumberInArray(numbers)

	if result != 98 {
		t.Error(result)
	}
}

func Test_GetTwoDigits1(t *testing.T) {
	numbers := parseLineToNumberSplice("811111111111119")
	result := getHighestCombinedNumberInArray(numbers)

	if result != 89 {
		t.Error(result)
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
	if value != 357 {
		t.Error(value)
	}
}

func Test_Highest3Digs(t *testing.T) {

	numbers := parseLineToNumberSplice("987654321111111")
	value := getHighest12CombinedNumberInArray(numbers, 2)
	if value != 987 {
		t.Error(value)
	}
}

func Test_Highest5Digs(t *testing.T) {

	numbers := parseLineToNumberSplice("811111111111119")
	value := getHighest12CombinedNumberInArray(numbers, 11)
	if value != 811111111119 {
		t.Error(value)
	}
}

func Test_Highest12Digs(t *testing.T) {

	numbers := parseLineToNumberSplice("234234234234278")
	value := getHighest12CombinedNumberInArray(numbers, 11)
	if value != 434234234278 {
		t.Error(value)
	}
}
func Test_Highest13DigsSkip1(t *testing.T) {

	numbers := parseLineToNumberSplice("234234")
	value := getHighest12CombinedNumberInArray(numbers, 2)
	if value != 434 {
		t.Error(value)
	}
}

func Test_part2Sample1(t *testing.T) {

	value := run(true, "987654321111111")
	if value != 987654321111 {
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
	if value != 3121910778619 {
		t.Error(value)
	}
}

//6583 too high
//6581 too high
//6000 too low
