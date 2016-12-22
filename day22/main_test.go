package main

import "testing"

func TestGetViablePairCount(t *testing.T) {
	nodes, err := openInput("input.txt")
	if err != nil {
		t.Errorf("unable to open input: %v", err)
		return
	}

	actual := getViablePairCount(nodes)
	const expected = 1038
	if actual != expected {
		t.Errorf("invalid result. got '%v' expected '%v'", actual, expected)
	}

	// part 2 result: 252 (calculated by hand)
}
