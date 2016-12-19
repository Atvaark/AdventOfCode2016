package main

import "testing"

const elfCount = 3014603

func TestRunPart1(t *testing.T) {
	actual := runPart1(elfCount)
	expected := 1834903
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestRunPart2(t *testing.T) {
	actual := runPart2(elfCount)
	expected := 1420280
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}
