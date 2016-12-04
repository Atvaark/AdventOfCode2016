package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	rooms, error := openRoomsFile("input.txt")
	if error != nil {
		t.Errorf("unable to read input: %v", error)
		return
	}

	result := part1(rooms)

	const expected = 173787
	if result != expected {
		t.Errorf("invalid result %d expected %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	rooms, error := openRoomsFile("input.txt")
	if error != nil {
		t.Errorf("unable to read input: %v", error)
		return
	}

	result := part2(rooms)

	const expected = 548
	if result != expected {
		t.Errorf("invalid result %d expected %d", result, expected)
	}
}
