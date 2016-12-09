package main

import "testing"

func TestPart1(t *testing.T) {

	lines, err := openInput("input.txt")
	if err != nil {
		t.Errorf("unable to open input: %v", err)
		return
	}

	lineCount := len(lines)
	if lineCount != 1 {
		t.Errorf("invalid input length %d", lineCount)
		return
	}

	result1 := part1(lines[0])

	const expectedResult = 138735
	if result1 != expectedResult {
		t.Errorf("invalid result %d expected %d", result1, expectedResult)
	}
}

func TestPart2(t *testing.T) {

	lines, err := openInput("input.txt")
	if err != nil {
		t.Errorf("unable to open input: %v", err)
		return
	}

	lineCount := len(lines)
	if lineCount != 1 {
		t.Errorf("invalid input length %d", lineCount)
		return
	}

	result1 := part2(lines[0])

	const expectedResult = 11125026826
	if result1 != expectedResult {
		t.Errorf("invalid result %d expected %d", result1, expectedResult)
	}
}
