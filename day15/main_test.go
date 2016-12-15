package main

import "testing"

func TestRun(t *testing.T) {
	discs := openInput()

	actual1, actual2 := run(discs)
	const expected1 = 16824
	if actual1 != expected1 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual1, expected1)
	}

	const expected2 = 3543984
	if actual2 != expected2 {
		t.Errorf("invalid result. got '%v' expected '%v'", actual2, expected2)
	}

}
