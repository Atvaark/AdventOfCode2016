package main

import "testing"

const inputRow = ".^.^..^......^^^^^...^^^...^...^....^^.^...^.^^^^....^...^^.^^^...^^^^.^^.^.^^..^.^^^..^^^^^^.^^^..^"

func TestRunPart1(t *testing.T) {
	const part1RowCount = 40
	actual := run(inputRow, part1RowCount)
	const expected = 1987
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestRunPart2(t *testing.T) {
	const part2RowCount = 400000
	actual := run(inputRow, part2RowCount)
	const expected = 19984714
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}
