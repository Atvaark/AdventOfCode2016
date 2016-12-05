package main

import "testing"

const input = "abbhdwsy"

func TestPart1(t *testing.T) {
	actual := part1(input)
	const expected = "801b56a7"
	if actual != expected {
		t.Errorf("invalid result. got %s expected %s", actual, expected)
	}
}

func TestPart2(t *testing.T) {
	actual := part2(input)
	const expected = "424a0197"
	if actual != expected {
		t.Errorf("invalid result. got %s expected %s", actual, expected)
	}
}
