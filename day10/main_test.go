package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	bots, err := openInput("input.txt")
	if err != nil {
		t.Errorf("unable to open input: %v", err)
		return
	}

	part1Result, part2Result := run(bots)

	const expectedPart1Result = 86
	if part1Result != expectedPart1Result {
		t.Errorf("invalid part 1 result %d expected %d", part1Result, expectedPart1Result)
	}

	const expectedPart2Result = 22847
	if part2Result != expectedPart2Result {
		t.Errorf("invalid part 2 result %d expected %d", part2Result, expectedPart2Result)
	}
}
