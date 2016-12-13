package main

import "testing"

func TestIsWall1(t *testing.T) {
	actual := isWall(1, 0, 10)
	const expected = true
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestIsWall2(t *testing.T) {
	actual := isWall(0, 0, 10)
	const expected = false
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestIsWall3(t *testing.T) {
	actual := isWall(1, 3, 10)
	const expected = true
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}
}

func TestRun(t *testing.T) {
	start := &pos{1, 1}
	end := &pos{31, 39}
	seed := testSeed

	actualPart1Result, actualPart2Result := run(start, end, seed)
	const expectedPart1 = 90
	if actualPart1Result != expectedPart1 {
		t.Errorf("invalid part 1 result. got '%v' expected '%v'", actualPart1Result, expectedPart1)
	}

	const expectedPart2 = 135
	if actualPart2Result != expectedPart2 {
		t.Errorf("invalid part 2 result. got '%v' expected '%v'", actualPart2Result, expectedPart2)
	}
}
