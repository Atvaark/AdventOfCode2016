package main

import (
	"testing"
)

func TestGetRequiredStepsPart1(t *testing.T) {
	grid, err := openInput("input.txt")
	if err != nil {
		t.Errorf ("unable to open input: %v", err)
		return
	}

	const returnToStart = false
	stepsPart1 := getRequiredSteps(grid, returnToStart)

	const expected = 470

	if stepsPart1 != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", stepsPart1, expected)
	}
}


func TestGetRequiredStepsPart2(t *testing.T) {
	grid, err := openInput("input.txt")
	if err != nil {
		t.Errorf ("unable to open input: %v", err)
		return
	}

	const returnToStart = true

	stepsPart2 := getRequiredSteps(grid, returnToStart)
	const expected = 720

	if stepsPart2 != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", stepsPart2, expected)
	}
}
