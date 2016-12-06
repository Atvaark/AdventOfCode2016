package main

import "testing"

func TestRun(t *testing.T) {
	lines, error := openInput("input.txt")
	if error != nil {
		t.Errorf("unable to read input: %v", error)
		return
	}

	maxResultString, minResultString := run(lines)

	const expectedPart1 = "qoclwvah"
	const expectedPart2 = "ryrgviuv"

	if maxResultString != expectedPart1 {
		t.Errorf("invalid part 1 result %s expected %s", maxResultString, expectedPart1)
		return
	}

	if minResultString != expectedPart2 {
		t.Errorf("invalid part 2 result %s expected %s", minResultString, expectedPart2)
	}
}
