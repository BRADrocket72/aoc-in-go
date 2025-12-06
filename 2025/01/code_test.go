package main

import (
	"log"
	"os"
	"testing"
)

func Test_parseLineL(t *testing.T) {
	line := "L21"
	testValue, _ := parseLine(line)

	if testValue != -21 {
		t.Error("parse L21 should be -21")
	}
}

func Test_parseLineR(t *testing.T) {
	line := "R21"
	testValue, _ := parseLine(line)

	if testValue != 21 {
		t.Error("parse L21 should be 21")
	}
}

func Test_parseLineMod10(t *testing.T) {
	line := "R10"
	testValue, _ := parseLine(line)

	if testValue != 10 {
		t.Error("parse R10 should be 10")
	}
}

func Test_parseLineMod100(t *testing.T) {
	line := "R100"
	testValue, _ := parseLine(line)

	if testValue != 0 {
		t.Error("parse R101 should be 0")
	}
}

func Test_parseLineModNeg100(t *testing.T) {
	line := "L100"
	testValue, _ := parseLine(line)

	if testValue != 0 {
		t.Error("parse L100 should be 0")
	}
}

func Test_parseLineMod105(t *testing.T) {
	line := "R105"
	testValue, _ := parseLine(line)

	if testValue != 5 {
		t.Error("parse R105 should be 5")
	}
}

func Test_overflowNormal(t *testing.T) {
	testValue, _ := handleOverflowAndUnderflow(21)

	if testValue != 21 {
		t.Error("overflow L21 should be 21")
	}
}

func Test_overflowHigh(t *testing.T) {
	testValue, _ := handleOverflowAndUnderflow(21)

	if testValue != 21 {
		t.Error("overflow L21 should be 21")
	}
}

func Test_overflowHigh100(t *testing.T) {
	testValue, _ := handleOverflowAndUnderflow(100)

	if testValue != 0 {
		t.Error("overflow 100 should be 0", testValue)
	}
}

func Test_overflowHigh236(t *testing.T) {
	testValue, _ := handleOverflowAndUnderflow(100)

	if testValue != 0 {
		t.Error("overflow 100 should be 0", testValue)
	}
}

func Test_overflowLow(t *testing.T) {
	testValue, _ := handleOverflowAndUnderflow(-79)
	if testValue != 21 {
		t.Error("undeflow -29 should be 21", testValue)
	}
}

func Test_overflowLow1(t *testing.T) {
	testValue, _ := handleOverflowAndUnderflow(-1)
	if testValue != 99 {
		t.Error("undeflow -1 should be 99", testValue)
	}
}

func Test_applyLine(t *testing.T) {
	testValue := handleLineApplyToNumber(50, "L68")
	if testValue != 82 {
		t.Error(testValue)
	}
}

func Test_applyLine2(t *testing.T) {
	testValue := handleLineApplyToNumber(82, "L30")
	if testValue != 52 {
		t.Error(testValue)
	}
}

func Test_applyLine3(t *testing.T) {
	testValue := handleLineApplyToNumber(52, "R48")
	if testValue != 0 {
		t.Error(testValue)
	}
}

func Test_applyLine8(t *testing.T) {
	testValue := handleLineApplyToNumber(99, "L99")
	if testValue != 0 {
		t.Error(testValue)
	}
}
func Test_applyLine9(t *testing.T) {
	testValue := handleLineApplyToNumber(14, "L82")
	if testValue != 32 {
		t.Error(testValue)
	}
}
func Test_applyLine10(t *testing.T) {
	testValue := handleLineApplyToNumber(99, "L99")

	if testValue != 0 {
		t.Error(testValue)
	}
}

func Test_applyLineSeries(t *testing.T) {
	testValue := handleLineApplyToNumber(95, "R60")
	testValue = handleLineApplyToNumber(testValue, "L55")
	testValue = handleLineApplyToNumber(testValue, "L1")
	testValue = handleLineApplyToNumber(testValue, "L99")
	if testValue != 0 {
		t.Error(testValue)
	}
}

func Test_applyLineSeries2(t *testing.T) {
	testValue := handleLineApplyToNumber(50, "L68")
	testValue = handleLineApplyToNumber(testValue, "L30")
	testValue = handleLineApplyToNumber(testValue, "R48")
	if testValue != 0 {
		t.Error(testValue)
	}
}

//1145 pt. 1

// Part 2
func Test_applyLineSeries3(t *testing.T) {

	testValue, timesPassed := handleLineApplyToNumberPt2(50, "L68", 0)
	if timesPassed != 1 {
		t.Error(testValue)
	}
}

func Test_applyLineSeriesall(t *testing.T) {

	testValue, timesPassed := handleLineApplyToNumberPt2(50, "L68", 0)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "L30", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "R48", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "L5", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "R60", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "L55", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "L1", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "L99", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "R14", timesPassed)
	testValue, timesPassed = handleLineApplyToNumberPt2(testValue, "L82", timesPassed)
	if timesPassed != 6 {
		t.Error(timesPassed)
	}
}

func Test_applyLineSeriesShouldFail(t *testing.T) {

	_, timesPassed := handleLineApplyToNumberPt2(0, "L5", 0)
	if timesPassed != 0 {
		t.Error(timesPassed)
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
	if value != 6 {
		t.Error(value)
	}
}

func Test_part2User(t *testing.T) {
	filePath := "input-user.txt" // Replace with your file path

	// Read the entire file into a byte slice
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		// Log the error and exit if the file cannot be read
		log.Fatalf("failed to read file: %v", err)
	}

	// Convert the byte slice to a string
	fileContentString := string(contentBytes)

	value := run(true, fileContentString)
	if value != 6 {
		t.Error(value)
	}
}

//6583 too high
//6581 too high
//6000 too low
