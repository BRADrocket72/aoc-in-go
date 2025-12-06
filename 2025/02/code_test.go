package main

import (
	"log"
	"os"
	"testing"
)

func Test_parseLineL(t *testing.T) {
	number_pair := "11-22"
	a, b := parseNumberPair(number_pair)

	if a != 11 || b != 22 {
		t.Error("parse L21 should be -21")
	}
}

func Test_parseLineSizeOfRange(t *testing.T) {
	allNumbers := createSpliceOfAllNumbers(11, 22)
	if allNumbers[0] != 11 || allNumbers[len(allNumbers)-1] != 22 {
		t.Error("list all from 11 to 22", allNumbers)
	}
}

func Test_parseMirror(t *testing.T) {
	allNumbers := createSpliceOfAllNumbers(11, 22)
	mirrors := parseArrayForMirrors(allNumbers)
	if len(mirrors) != 2 {
		t.Error("list all from 11 to 22", allNumbers)
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
	if value != 1227775554 {
		t.Error(value)
	}
}

func Test_parseMirrorPart2(t *testing.T) {
	allNumbers := createSpliceOfAllNumbers(2121212118, 2121212124)
	mirrors := parseArrayForMirrorsPt2(allNumbers)
	if mirrors[0] != 2121212121 {
		t.Error("list all from 11 to 22", allNumbers)
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
	if value != 4174379265 {
		t.Error(value)
	}
}
